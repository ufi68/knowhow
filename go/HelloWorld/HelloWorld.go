package main
import (
	"fmt"
	"runtime"
)

func main() {
    fmt.Println("Hello World")
	fmt.Println("Number of CPUs:", runtime.NumCPU())
}
