//go:build windows && go1.19
// +build windows,go1.19

package tzlocal

import (
	"os/exec"
	"strings"
	"syscall"
)

// localTZfromTzutil executes command `tzutil /g` to get the name of the time zone Windows is configured to use.
func localTZfromTzutil() (string, error) {
	cmd := exec.Command("tzutil", "/g")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow: true,
	}
	data, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(data)), nil
}
