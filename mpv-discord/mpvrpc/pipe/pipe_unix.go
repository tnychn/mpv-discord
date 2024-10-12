// +build !windows

package pipe

import (
	"net"
	"time"
)

func GetPipeSocket(path string) (net.Conn, error) {
	return net.DialTimeout("unix", path, time.Second*5)
}
