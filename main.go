/*
@auther: xieyongbo@bytedance.com
@date: 2022/4/29
@comment
*/
package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os/exec"
)

var (
	source_path = flag.String("source", "", "--source_path")
	output      = flag.String("output", "CHANGELOG.md", "--output CHANGELOG.md")
	fetch       = flag.Bool("fetch", false, "--fetch")
)

type record struct {
	Version  string
	Date     string
	Features []string
	BugFixes []string
}

func main() {
	flag.Parse()

	git("log", "log", "--pretty=format:\"%B%H\"")
}

func git(dir string, args ...string) (string, error) {
	var stdin, stdout, stderr bytes.Buffer

	cmd := exec.Command("git", args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	cmd.Stdin = &stdin
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	fmt.Printf(">>out:\n%s\nerr:\n%s\n", outStr, errStr)
	fmt.Printf("combined out:\n%s\n", string(out))

	return "xxx", err
}