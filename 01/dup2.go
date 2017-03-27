package main

import (
	"os"
	"io/ioutil"
	"fmt"
	"strings"
)

func main() {
	counter := make(map[string][]string)
	for _, fileName := range os.Args[1:] {
		data, err := ioutil.ReadFile(fileName)
		if err != nil {
			fmt.Println(err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counter[line] = append(counter[line], fileName)
		}
	}
	for k, v := range counter {
		fmt.Println(k, v)
	}
}
