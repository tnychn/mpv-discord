// +build windows

package pipe

import (
	"fmt"
	"net"
	"time"

	npipe "gopkg.in/natefinch/npipe.v2"
)

func GetPipeSocket(pid string) (net.Conn, error) {
	return npipe.DialTimeout(fmt.Sprintf(`\\.\pipe\tmp\mpv-discord-%s`, pid), time.Second*5)
}
