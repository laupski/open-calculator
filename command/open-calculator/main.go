package command

import (
	"flag"
	"fmt"
	"github.com/laupski/open-calculator/api"
	"os"
	"strconv"

	"github.com/laupski/open-calculator/internal/arithmetic"
)

var commands = []string{
	"evaluate",
	"tokenize",
	"postfix",
	"api",
}

func Run(args []string) int {
	var help bool
	flag.BoolVar(&help, "h", false, "print help")
	flag.Parse()

	a := flag.Args()
	if help || len(a) == 0 || len(args) == 0{
		fmt.Println("Usage: open-calculator <command> <argument>")
		os.Exit(0)
	}

	check := checkCommands(args[0])
	if !check {
		printValidCommands()
		os.Exit(0)
	}

	switch args[0] {
	case "api":
		if len(args) == 2 {
			port,err := strconv.ParseUint(args[1], 5,64)
			if err != nil {
				api.StartApi(int(port))
			} else {
				fmt.Printf("Invalid port specified: %v\n", args[1])
			}
		} else {
			api.StartApi(api.DefaultPort)
		}
	case "evaluate":
		if len(args) == 2 {
			evaluate(args[1])
		} else {
			fmt.Println("Invalid amount of arguments found for evaluation")
		}
	case "postfix":
		if len(args) == 2 {
			postfix(args[1])
		} else {
			fmt.Println("Invalid amount of arguments found for postfix")
		}
	case "tokenize":
		if len(args) == 2 {
			tokenize(args[1])
		} else {
			fmt.Println("Invalid amount of arguments found for tokenize")
		}
	default:
		fmt.Println("Usage: open-calculator <command> <argument>")
		os.Exit(0)
	}

	return 0
}

func checkCommands(arg string) bool {
	for _, cmd := range commands {
		if cmd == arg {
			return true
		}
	}

	return false
}

func evaluate(expression string) {
	fmt.Printf("Running evaluation on: %v\n", expression)
	tokenList, err := arithmetic.Tokenize(expression)
	if err != nil {
		fmt.Printf("Error in evaluation: %v\n", err)

	}
	output := arithmetic.ToTokenQueue(tokenList)
	sy := arithmetic.NewShuntingYard(output)
	postfixQueue := sy.InfixToPostFix()
	answer, err := arithmetic.Evaluate(postfixQueue)
	if err != nil {
		fmt.Printf("Error in evaluation: %v\n", err)

	} else {
		fmt.Printf("Evaluated to: %v\n", answer)
	}
}

func tokenize(expression string) {
	fmt.Printf("Running evaluation on: %v\n", expression)
	tokenList, err := arithmetic.Tokenize(expression)
	if err != nil {
		fmt.Printf("Error in evaluation: %v\n", err)

	} else {
		fmt.Printf("Evaluated to: %v\n", tokenList)
	}
}

func postfix(expression string) {
	fmt.Printf("Running evaluation on: %v\n", expression)
	tokenList, err := arithmetic.Tokenize(expression)
	output := arithmetic.ToTokenQueue(tokenList)
	sy := arithmetic.NewShuntingYard(output)
	postfixQueue := sy.InfixToPostFix()
	if err != nil {
		fmt.Printf("Error in evaluation: %v\n", err)

	} else {
		fmt.Printf("Evaluated to: %v\n", postfixQueue)
	}
}

func printValidCommands() {
	fmt.Println("List of valid commands:")
	for _, c := range commands {
		fmt.Printf("\t%v\n", c)
	}
}