// +build windows

package pipe

import (
	"fmt"
	"net"
	"time"

	npipe "gopkg.in/natefinch/npipe.v2"
)

func GetPipeSocket(pipe_path string) (net.Conn, error) {
	return npipe.DialTimeout(fmt.Sprintf(pipe_path), time.Second*5)

}
