// Golang program to get current date and time
package main

import (
	"fmt"
	"time"
)

func main() {

	// get date and time and store in variable on left
	currentTime := time.Now()

	// printing the date and time in string format
	fmt.Println("Current Time: ", currentTime.String())
}
