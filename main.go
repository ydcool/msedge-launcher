package main

import (
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
		log.Fatalf("environment variable %s not found", pf86err)
	}
	edgePath := filepath.Join(pf86, DefaultEdgePath)
	args := append(os.Args[1:], EdgeFlag)
	cmd := exec.Command(edgePath, args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("failed excute edge command: %s %v: %v", edgePath, args, err)
	}
	log.Print(string(out))
}
