package logger

import "fmt"

// Debug determinate to print a message
var Debug bool

// CounterLog track the amount of calls of the functionLog
var CounterLog int

// Log check for Debug and display the passed data
func Log(functionName string, args ...interface{}) {

	if !Debug {
		return
	}

	const separator = "========================================"

	CounterLog++
	fmt.Printf("%v\n#%v Excution in the function: %v\n", separator, CounterLog, functionName)
	if len(args) > 0 {
		fmt.Println("Running the following commands:")
	}
	for v := range args {
		fmt.Printf("%+v\n", v)
	}
	fmt.Println(separator)
}
