package main

import (
	"fmt"
	"runtime"
)

/*
switch statement provides a multi-way execution.
An expression or type specifier is compared to the cases inside the switch to determine which branch to execute.
Unlike in other languages such as C, Java, or PHP, each case is terminated by an implicit break;
therefore, we do not have to write it explicitly.
The default statement can be used for a branch that is executed, when no other cases fit.
The default statement is optional.
*/
func main() {
	/*The GOOS environment variable is the
	running program's operating system target: one of darwin, freebsd, linux, and so on.
	Based on the value of the variable, we print the OS version.*/
	os := runtime.GOOS
	switch os {
	case "windows":
		fmt.Println("Windows")
	case "darwin":
		fmt.Println("MAC operating system")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s.\n", os)
	}
}
