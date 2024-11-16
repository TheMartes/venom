package main

import (
	"fmt"
	"os"
	"os/exec"
)

func requiredCheck(cmds []string) {
	for _, cmd := range cmds {
		_, err := exec.LookPath(cmd)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			fmt.Println("Venom needs these executable to be able to run\n")
			os.Exit(1)
		}
	}
}
