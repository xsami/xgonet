package logger

import "fmt"

// Debug determinate to print a message
var Debug bool

// Log check for Debug and display the passed data
func Log(functionName string, args ...interface{}) {

	if !Debug {
		return
	}

	fmt.Printf("Excution in the function: %v\n", functionName)
	if len(args) > 0 {
		fmt.Println("Running the following commands:")
	}
	for v := range args {
		fmt.Printf("%+v\n", v)
	}
}
