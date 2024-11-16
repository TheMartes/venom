package parsers

import (
	"fmt"
	"os"
	"os/exec"
)

func ParseGit(cmds []string) {
	for _, c := range cmds {
		var cmd *exec.Cmd

		// TODO: move to some sort of dictionary??
		switch c {
		case "store-credentials":
			cmd = exec.Command("git", "config", "--global", "credential.helper", "store")
		}

		// we don't care about output
		_, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}

		fmt.Println(fmt.Sprintf("Applied: git %s", c))
	}

	fmt.Println("Git Configured")
}
