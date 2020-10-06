// +build windows

package main

import (
	"fmt"
	"github.com/ydcool/msedge-launcher/pkg"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

const DefaultEdgePath = "Microsoft/Edge/Application/msedge.exe"
const EdgeFlag = "--disable-features=RendererCodeIntegrity"

func main() {
	pf86, pf86err := os.Getenv("PROGRAMFILES(x86)"), "%programfiles(x86)%"
	if pf86 == "" {
		_ = pkg.MessageBoxPlain("ErrorLaunchEdge", fmt.Sprintf("environment variable %s not found", pf86err))
		return
	}

	edgePath := filepath.Join(pf86, DefaultEdgePath)
	_, err := os.Stat(edgePath)
	if err != nil {
		_ = pkg.MessageBoxPlain("ErrorLaunchEdge", fmt.Sprintf("Failed stat msedge location, Error:\n %v", err))
		return
	}

	args := append(os.Args[1:], EdgeFlag)
	cmd := exec.Command(edgePath, args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		_ = pkg.MessageBoxPlain("ErrorLaunchEdge", fmt.Sprintf("failed excute edge command:\n %s %v \nError:\n %v", edgePath, args, err))
		return
	}
	log.Print(string(out))
}
