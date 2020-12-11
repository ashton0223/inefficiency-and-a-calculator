// Interpreter for the inefficiency programming language aimed at producing very large
// source code.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type runtime struct {
	vars        map[string]string
	conditional string
}

func (r *runtime) initRun() {
	r.vars = make(map[string]string, 100)
}

func (r *runtime) runLine(line string) string {
	// Break up instruction
	parts := strings.SplitN(line, ":", 2)
	if strings.Contains(parts[1], ":") {
		firstArg := ""
		ran := ""
		test := strings.SplitN(parts[1], ":", 2)
		if strings.Contains(test[0], ",") {
			args := strings.SplitN(parts[1], ",", 2)
			firstArg = args[0] + ","
			ran = r.runLine(args[1])
		} else {
			ran = r.runLine(parts[1])
		}
		parts[1] = firstArg + ran
	}

	// Check for conditional statements
	for {
		if string(parts[0][0]) == " " {
			if r.conditional == "true" {
				parts[0] = parts[0][1:]
			} else {
				return ""
			}
		} else {
			break
		}
	}

	// Check for variables
	for name, value := range r.vars {
		if strings.Contains(parts[1], "*"+name+"*") {
			parts[1] = strings.ReplaceAll(parts[1], "*"+name+"*", value)
		}
	}

	// Run instruction
	switch parts[0] {
	case "print":
		fmt.Println(parts[1])
		break
	case "cmp":
		inputs := strings.Split(parts[1], ",")
		if inputs[0] == inputs[1] {
			return "true"
		} else {
			return "false"
		}
	case "if":
		r.conditional = parts[1]
		break
	case "set":
		inputs := strings.Split(parts[1], ",")
		r.vars[inputs[0]] = inputs[1]
		break
	case "input":
		fmt.Print(parts[1])
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		return text
	}
	return ""
}

func main() {
	writeCode()
	filename := os.Args[1]

	program, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer program.Close()

	r := runtime{}
	r.initRun()

	scanner := bufio.NewScanner(program)
	for scanner.Scan() {
		line := scanner.Text()
		r.runLine(line)
	}
}

func writeCode() {
	file, err := os.Create("calculator.inef")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.WriteString("set:input1,input:Enter the first number --> \n")
	_, err = file.WriteString("set:input2,input:Enter the second number --> \n")

	for i := 0; i < 10000; i++ {
		for j := 0; j < 10000; j++ {
			_, err = file.WriteString(fmt.Sprintf("set:val1,cmp:*input1*,%d\n", i))
			_, err = file.WriteString(fmt.Sprintf("set:val2,cmp:*input2*,%d\n", j))
			_, err = file.WriteString("if:cmp:*val1*,true\n")
			_, err = file.WriteString("    if:cmp:*val2*,true\n")
			_, err = file.WriteString(fmt.Sprintf("    print:%d + %d = %d\n", i, j, i+j))
		}
	}
}
