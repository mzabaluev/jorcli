// Package jcli - jormungander jcli helpers.
package jcli

import (
	"bytes"
	"os/exec"
)

// execStd ...
func execStd(stdin []byte, name string, arg ...string) ([]byte, error) {
	var (
		cmd    *exec.Cmd
		stdout bytes.Buffer
		stderr bytes.Buffer
	)
	cmd = exec.Command(name, arg...)
	// TODO: check also exec.CommandContext(ctx context.Context, name string, arg ...string)

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if stdin != nil /* && len(stdin) > 0 */ {
		cmd.Stdin = bytes.NewBuffer(stdin)
	}

	if err := cmd.Run(); err != nil {
		return stderr.Bytes(), err
	}
	return stdout.Bytes(), nil
}

// JormungandrVersion - get jormungandr version.
// Provided just for convenience.
func JormungandrVersion() ([]byte, error) {
	return execStd(nil, "jormungandr", "--version")
}

// Version - get jcli version.
func Version() ([]byte, error) {
	return execStd(nil, "jcli", "--version")
}
