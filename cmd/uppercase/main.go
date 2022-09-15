package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		panic(fmt.Errorf("[Err]: expected 2 args received %d", len(os.Args)))
	}
	input := os.Args[1]
	output := strings.ToUpper(input)
	pipe := os.NewFile(uintptr(3), "pipe")
	err := json.NewEncoder(pipe).Encode(output)
	if err != nil {
		panic(err)
	}
}
