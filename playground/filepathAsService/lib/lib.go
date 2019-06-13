package lib

import (
	"bufio"
	"fmt"
	"os"
)

func ImportHello() {
	f, err := os.Open("./tmp/tmp.html")
	fmt.Println("file err: ", err)
	s := bufio.NewScanner(f)
	for s.Scan() {
		fmt.Println(s.Text())
	}
	// fmt.Println("file err: ", f.Read())
	fmt.Println(os.Getwd())
}
