package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func main() {
	out, err := exec.Command("ip", "a").Output()
	if err != nil {
		log.Fatal(err)
	}
	var s []string = strings.Split(string(out), " ")
	for i := 0; i < len(s); i++ {
		if s[i] == "inet" {
			fmt.Println(s[i+1])
		}
	}
}
