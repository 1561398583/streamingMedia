package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main()  {
	cmd := exec.Command("sh", "testShell.sh")
	bs, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(bs)
}
