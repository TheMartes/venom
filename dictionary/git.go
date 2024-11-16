package dictionary

import (
	"errors"
	"fmt"
	"os/exec"
)

var git = map[string]*exec.Cmd{
	"store-credentials": exec.Command("git", "config", "--global", "credential.helper", "store"),
}

func GetGitCommand(key string) (*exec.Cmd, error) {
	output := git[key]

	if output == nil {
		return nil, errors.New(fmt.Sprintf("error: command \"%s\" doest not exists in dictionary", key))
	}
	return output, nil
}
