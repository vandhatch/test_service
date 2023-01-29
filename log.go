package test_service

import "fmt"

func Write(message any) {
	log(message)
}

func log(message any) {
	fmt.Println(message)
}
