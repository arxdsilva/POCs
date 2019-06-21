package main

import (
	"errors"
	"fmt"
)

func main() {
	var err error
	var str string
	f := func() error {
		err = errors.New("some error")
		str = "some string"
		return nil
	}
	fmt.Println(f(), err, str)
}
