package goo

import (
	"bytes"
	"fmt"
)

func locatePanic(stack []byte) string {
	lines := bytes.Split(stack, []byte("\n"))
	if len(lines) < 1 {
		return string(stack)
	}

	lines = lines[1:]
	var located bool
	var index int
	for i, line := range lines {
		if bytes.Contains(line, []byte("runtime/panic.go:770 +0x124")) {
			located = true
			index = i
			break
		}
	}

	if located && len(lines)-1 > index+2 {
		return fmt.Sprintf("caller: %s, line: %s", lines[index+1], bytes.TrimSpace(lines[index+2]))
	}

	return string(stack)
}
