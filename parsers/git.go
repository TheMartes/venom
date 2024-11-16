package parsers

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/themartes/venom/dictionary"
)

func ParseGit(cmds []string) {
	for _, c := range cmds {
		var cmd *exec.Cmd

		cmd, err := dictionary.GetGitCommand(c)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}

		// we don't care about output
		_, err = cmd.CombinedOutput()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}

		fmt.Println(fmt.Sprintf("Applied: git %s", c))
	}

	fmt.Println("Git Configured")
}
