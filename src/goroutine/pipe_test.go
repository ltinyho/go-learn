package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
	"testing"
)

func TestPipe(t *testing.T) {
	cmd1 := exec.Command("ps", "aux")
	cmd2 := exec.Command("grep", "apipe")
	stdout1, err := cmd1.StdoutPipe()
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := cmd1.Start(); err != nil {
		fmt.Printf("Error: Start %v\n", err)
	}
	outputBuf1 := bufio.NewReader(stdout1)
	stdin2, err := cmd2.StdinPipe()
	if err != nil {
		fmt.Printf("Error:stdin :%v\n", err)
		return
	}
	outputBuf1.WriteTo(stdin2)
	var outputBuf2 bytes.Buffer
	cmd2.Stdout = &outputBuf2
	if err := cmd2.Start(); err != nil {
		fmt.Printf("Error: startup:%v\n", err)
		return
	}
	err = stdin2.Close()
	if err != nil {
		fmt.Printf("Error: stdin2 close:%v\n", err)
		return
	}
	if err := cmd2.Wait(); err != nil {
		fmt.Printf("Error: wait %v\n", err)
		return
	}
	fmt.Printf("%s\n", outputBuf2.Bytes())
}

func TestPipeFunc(t *testing.T) {
	cmd0 := exec.Command("echo", "-n", "My first command from golang.\nadsf ")
	stdout0, err := cmd0.StdoutPipe()
	if err != nil {
		fmt.Printf("Error: Can not obtain the stdout pipe for command No.0: %s\n", err)
		return
	}
	if err := cmd0.Start(); err != nil {
		fmt.Printf("Error: command No.0 can not be startup: %s\n", err)
	}
	outputBufo := bufio.NewReader(stdout0)
	output0, _, err := outputBufo.ReadLine()
	if err != nil {
		fmt.Printf("Error: Can not read data from the pipe: %s\n", err)
		return
	}
	fmt.Printf("%s\n", string(output0))
	output0, _, _ = outputBufo.ReadLine()
	fmt.Printf("%s\n", string(output0))
}
