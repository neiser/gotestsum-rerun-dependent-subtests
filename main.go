package main

import (
	"errors"
	"fmt"
	"os"
)

const harryIsThereFile = "harryIsThere"

func sayHello(name string) error {
	if len(name) == 0 {
		return nil
	} else if name == "Voldemort" {
		_, err := os.ReadFile(harryIsThereFile)
		if err == nil {
			fmt.Println("Not afraid of Voldemort!")
			return os.Remove(harryIsThereFile)
		}
		err = os.WriteFile(harryIsThereFile, []byte{}, 0644)
		if err != nil {
			return err
		}
		return errors.New("afraid of Voldemort, cannot say hello")
	}
	fmt.Printf("Hello, %s\n", name)
	return nil
}
