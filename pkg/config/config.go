package config

import (
	"bytes"
	"os/exec"
)

func GetModuleName() string {
	cmd := exec.Command("go", "list", "-m")
	var out bytes.Buffer
	cmd.Stdout = &out

	return out.String()
}
