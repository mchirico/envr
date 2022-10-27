package pb

import (
	"os/exec"
)

var (
	pasteCmdArgs = "pbpaste"
	copyCmdArgs  = "pbcopy"
)

func paste() *exec.Cmd {
	return exec.Command(pasteCmdArgs)
}

func copy() *exec.Cmd {
	return exec.Command(copyCmdArgs)
}

func Read() (string, error) {
	pasteCmd := paste()
	out, err := pasteCmd.Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func Write(text string) error {
	copyCmd := copy()
	in, err := copyCmd.StdinPipe()
	if err != nil {
		return err
	}

	if err := copyCmd.Start(); err != nil {
		return err
	}
	if _, err := in.Write([]byte(text)); err != nil {
		return err
	}
	if err := in.Close(); err != nil {
		return err
	}
	return copyCmd.Wait()
}
