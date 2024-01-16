package exec

import (
	"fmt"
	"os/exec"
	"strings"
)

func SafeExec(name string, arg ...string) (output string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()

	raw, err := exec.Command(name, arg...).CombinedOutput()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(raw)), nil
}
