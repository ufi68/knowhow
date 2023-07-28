package main

import (
	"fmt"
	"log"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/v3/mem"
)

func main() {
	v, _ := mem.VirtualMemory()

	// almost every return value is a struct
	fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)

	// convert to JSON. String() is also implemented
	fmt.Println(v)

	// show the number of CPU's on the machine
	fmt.Println("Number of CPUs:", runtime.NumCPU())

	// save the date and time in a variable
	var t time.Time
	t = time.Now()
	// print the date and time
	fmt.Println("Date and Time:", t)
	// print the date
	fmt.Println("YearDate:     ", t.YearDay())
	// print the time
	fmt.Println("UTC Year:     ", t.UTC().Year())

	// Async function to show environment infos
	showEnvironment()

	// Configure log output
	log.SetPrefix("INFO: ")
	log.SetFlags(log.Ldate | log.Ltime | log.LUTC)
	log.Println("Starting...")
	log.Println("OS:", runtime.GOOS)
	log.SetPrefix("FATAL: ")
	log.Fatalln("Fatal error")
	log.SetPrefix("PANIC: ")
	log.Panicln("Panic error")
	log.SetPrefix("INFO: ")
	log.Println("Ending...")

}

// Function to show environment infos
func showEnvironment() {
	fmt.Println("OS:", runtime.GOOS)
	fmt.Println("Arch:", runtime.GOARCH)
	fmt.Println("Version:", runtime.Version())
	fmt.Println("Compiler:", runtime.Compiler)
	fmt.Println("NumGoroutine:", runtime.NumGoroutine())
	fmt.Println("NumCPU:", runtime.NumCPU())
	fmt.Println("NumCgoCall:", runtime.NumCgoCall())
}
