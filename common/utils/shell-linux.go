//go:build !windows
// +build !windows

package utils

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"os/exec"
	"syscall"
)

type Result struct {
	output string
	err    error
}

// ExecShell 执行shell命令，可设置执行超时时间
func ExecShell(ctx context.Context, command string) (string, error) {

	cmd := exec.Command("/bin/bash", "-c", command)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setpgid: true,
	}
	resultChan := make(chan Result)
	go func() {
		output, err := cmd.CombinedOutput()
		resultChan <- Result{string(output), err}
	}()
	select {
	case <-ctx.Done():
		if cmd.Process.Pid > 0 {
			syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
		}
		return "", status.Error(codes.DeadlineExceeded, "Deadline exceeded")
	case result := <-resultChan:
		return result.output, result.err
	}
}
