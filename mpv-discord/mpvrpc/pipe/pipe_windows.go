// +build windows

package pipe

import (
	"fmt"
	"net"
	"path/filepath"
	"time"

	npipe "gopkg.in/natefinch/npipe.v2"
)

func GetPipeSocket(path string) (net.Conn, error) {
	path = filepath.FromSlash(path) // convert it into a Windows-compatible path
	return npipe.DialTimeout(fmt.Sprintf(`\\.\pipe\%s`, path), time.Second*5)
}
