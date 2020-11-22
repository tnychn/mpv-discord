// +build !windows

package pipe

import (
	"net"
	"os"
	"path/filepath"
	"time"
)

func GetPipeSocket() (net.Conn, error) {
	path := os.Getenv("XDG_RUNTIME_DIR")
	if path == "" {
		path = os.TempDir()
	}
	path = filepath.Join(path, "discord-ipc-0")
	return net.DialTimeout("unix", path, time.Second*5)
}
