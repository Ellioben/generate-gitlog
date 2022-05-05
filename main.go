/*
@auther: xieyongbo@bytedance.com
@date: 2022/4/29
@comment
*/
package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os/exec"
	"strings"
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

	out, err := git("log", "--pretty=format:\"%B%H\"")
	if err != nil {
		return
	}

	tags := strings.Split(out, "\n")
	records := make([]record, 0, len(tags))
	for _, t := range tags {
		segs := strings.Split(t, ";")
		if len(segs) < 3 || len(segs[0]) == 0 {
			continue
		}
		version := segs[0]
		date := strings.ReplaceAll(segs[1], "-", "/")
		records = append(records, record{Version: version, Date: date})
	}
	fmt.Println(records)

	//gitOut()
}

func git(args ...string) (string, error) {
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
	return string(out), err
}

func gitOut(args ...string) error {
	cmd := exec.Command("git", "log", "--pretty=format:\"%B%H\"")
	//cmd := exec.Command("git", args...)

	stdout, _ := cmd.StdoutPipe()
	cmd.Start()
	reader := bufio.NewReader(stdout)
	var i int
	for {
		i++
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if err != nil || io.EOF == err {
			break
		}
		fmt.Printf("line:%v -- %v\n", i, line)
	}

	cmd.Wait()
	return nil
}
