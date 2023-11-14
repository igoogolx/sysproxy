package sysproxy

import (
	_ "embed"
	"fmt"
	"os/exec"
	"syscall"

	"github.com/getlantern/byteexec"
)

// Note this is a universal binary that runs on amd64 and arm64
//
//go:embed binaries/darwin/sysproxy
var sysproxy []byte

func ensureElevatedOnDarwin(be *byteexec.Exec, prompt string, iconFullPath string) (err error) {
	var s syscall.Stat_t
	// we just checked its existence, not bother checking specific error again
	if err = syscall.Stat(be.Filename, &s); err != nil {
		return fmt.Errorf("error starting helper tool %s: %v", be.Filename, err)
	}

	return nil
}

func detach(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setpgid: true,
	}
}
