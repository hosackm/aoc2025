# aoc2025

This repository contains my solutions for Advent of Code 2025. It includes a command line tool named aoc that can scaffold new puzzle days and run solutions.

## Building

The project includes a Makefile with several useful targets. To build the tool:

```bash
make build
```


This compiles the CLI binary into ./bin/aoc.

You can also run it directly without building:


## Scaffolding a new day

To create the folder structure and template files for a specific day, run:

```bash
aoc scaffold 6
```

This generates the scaffolding for day 6.

## Running solutions

After implementing a day's solution, run it with:

```bash
aoc run 6
```

This executes the code for day 6.

## Inputs

You must manually download your Advent of Code input for each day and place it in:

`inputs/DD.txt`

Where `DD` is a two-digit number representing the day.

For example, day 6 would use:

`inputs/06.txt`

## Development workflow

During development you can use the watch target to automatically rebuild and run a day whenever its files change:

```bash
make watch
```

Run tests with:

```bash
make test
```

Run benchmarks with:

```bash
make bench
```

Clean the build directory with:

```bash
make clean
```

## Project structure

Solutions are located under `internal/days`, one folder per day, created by the scaffolding tool.

The CLI lives under `cmd/aoc`.

