package commands

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

const sourceTmpl = `package {{ printf "day%02d" .Day }}

type Day{{ printf "%02d" .Day }} struct{}

func (d {{ printf "Day%02d" .Day }}) Part1(input string) (string, error) { return "", nil }

func (d {{ printf "Day%02d" .Day }}) Part2(input string) (string, error) { return "", nil }
`
const testTmpl = `package {{ printf "day%02d" .Day }}

import "testing"

func TestDay{{ printf "%02d" .Day }}(t *testing.T) {}
`

func HandleScaffold(day int, filename string) error {
	if day < 1 || day > 25 {
		return fmt.Errorf("day must be between 1 through 25")
	}

	dayStr := fmt.Sprintf("day%02d", day)
	dir := filepath.Join("internal", "days", dayStr)
	filenames := []string{
		filepath.Join(dir, dayStr+".go"),
		filepath.Join(dir, dayStr+"_test.go"),
	}

	if _, err := os.Stat(dir); errors.Is(err, os.ErrNotExist) {
		err = os.Mkdir(dir, 0o777)
		if err != nil {
			return err
		}
	}

	for _, f := range filenames {
		_, err := os.Stat(f)
		if err == nil || errors.Is(err, os.ErrExist) {
			return fmt.Errorf("file %s already exists", f)
		}
	}

	templates := make([]*template.Template, len(filenames))
	for i, tmpl := range []string{sourceTmpl, testTmpl} {
		t, err := template.New("").Parse(tmpl)
		if err != nil {
			return err
		}

		templates[i] = t

		f, err := os.OpenFile(filenames[i], os.O_CREATE|os.O_WRONLY, 0o777)
		if err != nil {
			return err
		}

		err = t.Execute(f, struct{ Day int }{day})
		if err != nil {
			return err
		}
	}

	return nil
}
