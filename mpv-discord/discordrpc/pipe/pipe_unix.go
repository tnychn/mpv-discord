// +build !windows

package pipe

import (
	"net"
	"os"
	"path/filepath"
	"time"
)

func GetPipeSocket() (net.Conn, error) {
	tempdir := os.Getenv("XDG_RUNTIME_DIR")
	if tempdir == "" {
		tempdir = os.TempDir()
	}
	
	path := filepath.Join(tempdir, "discord-ipc-0")
	

	// If discord is installed as a snap, it has a special tmp folder for isolation
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		path = filepath.Join(tempdir, "snap.discord/discord-ipc-0")
	}
	//TODO: Add support for other install methods that place discord's rpc socket somewhere else

	return net.DialTimeout("unix", path, time.Second*5)
}
