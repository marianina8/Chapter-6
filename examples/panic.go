package examples

import (
	"bytes"
	"fmt"
	"os/exec"
	"path/filepath"
)

func Panic() {
	cmd := exec.Command(filepath.Join("bin", "panic"))
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return
	}
}
