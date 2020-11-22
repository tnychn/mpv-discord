// +build !windows

package pipe

import (
	"fmt"
	"net"
	"time"
)

func GetPipeSocket(pid string) (net.Conn, error) {
	return net.DialTimeout("unix", fmt.Sprintf("/tmp/mpv-discord-%s", pid), time.Second*5)
}
