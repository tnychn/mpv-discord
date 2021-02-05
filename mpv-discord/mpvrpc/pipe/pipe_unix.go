// +build !windows

package pipe

import (
	"fmt"
	"net"
	"time"
)

func GetPipeSocket(pipe_path string) (net.Conn, error) {

	return net.DialTimeout("unix", fmt.Sprintf(pipe_path), time.Second*5)

}
