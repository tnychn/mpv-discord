// +build windows

package pipe

import (
	"net"
	"time"

	npipe "gopkg.in/natefinch/npipe.v2"
)

func GetPipeSocket() (net.Conn, error) {
	return npipe.DialTimeout(`\\.\pipe\discord-ipc-0`, time.Second*5)
}
