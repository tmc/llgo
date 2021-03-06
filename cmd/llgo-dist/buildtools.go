// Copyright 2013 The llgo Authors.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const llgoBuildPath = "github.com/axw/llgo/cmd/llgo-build"

var llgobuildbin string

func buildLlgoTools() error {
	log.Println("Building llgo-build")

	var err error
	if llgobuildbin, err = findCommand(llgoBuildPath); err != nil {
		return err
	}
	if _, err = os.Stat(llgobuildbin); err == nil {
		err = os.Remove(llgobuildbin)
		if err != nil {
			return fmt.Errorf("Failed to remove llgo-build: %s", err)
		}
	}

	// We set default values in the llgo-build binary.
	ldflags := []string{
		fmt.Sprintf("-X main.llgobin %q", llgobin),
		fmt.Sprintf("-X main.llvmbindir %q", llvmbindir),
		fmt.Sprintf("-X main.defaulttriple %q", triple),
	}
	if triple == "pnacl" {
		pnaclClangFlag := fmt.Sprintf("-X main.defaultclang %q", pnaclClang)
		ldflags = append(ldflags, pnaclClangFlag)
	}

	args := []string{"install", "-ldflags", strings.Join(ldflags, " "), llgoBuildPath}
	output, err := command("go", args...).CombinedOutput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", string(output))
		return err
	}

	log.Printf("Built %s", llgobuildbin)
	return nil
}
