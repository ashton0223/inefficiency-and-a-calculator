# inefficiency-and-a-calculator
Creating a programming language in order to run a calculator.

## Be careful running this!
It will create a 12 gigabyte script in the directory in which this is run.

## How to run:
Run `go run inefficiency.go calculator.inef` in the directory `inefficiency.go` is located.

### Why I created a new language for this:
After seeing [AceLewis](https://github.com/AceLewis)'s [calculator](https://github.com/AceLewis/my_first_calculator.py), I wanted to see how far I could push it. To be specific, I wanted to see if I could create a version of it that was over a gigabyte in size. This was not possible in Go, C++, or Python, due to the amount of memory usage. Therefore, the only solution was to create a new programming language that only reads a single line at a time. Inefficiency can run scripts of any length, due to only requiring the current line of code to be in memory. This allows me to generate and run a script 12 gigabytes in size.
