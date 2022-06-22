package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

func main() {
	fmt.Println("Hello")
	for i := 0; i < 20; i++ {
		fmt.Println("This is i:", i)
	}
	logrus.Info("Hello from logrues, wink")
}
