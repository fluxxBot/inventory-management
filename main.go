package main

import (
	"fmt"
	"git.jfrog.info/kanishkg/inventory-management/application"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic: ", r)
		}
	}()
	application.Application()
}
