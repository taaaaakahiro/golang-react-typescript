package practice

import (
	"fmt"
	"time"
)

func Display() {
	fmt.Println("test", time.Now())
} 

func MathTest(a int, b int) int {
	return a * b 
}