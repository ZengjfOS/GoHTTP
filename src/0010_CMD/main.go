package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	fmt.Println("Hello, World")

	cmd := exec.Command("cmd", "/c", "dir", "/?")
	data, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("output: %s", data)

	cmd = exec.Command("cmd", "/c", "adb devices")
	data, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("output: %s", data)

	cmd = exec.Command("cmd", "/c", "adb shell")
	data, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("output: %s", data)
}
