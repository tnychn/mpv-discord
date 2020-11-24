package mpvrpc

import (
	"bufio"
	"encoding/json"
	"errors"
	"io"
	"net"
	"sync"

	"github.com/tnychn/mpv-discord/mpvrpc/pipe"
)

type Client struct {
	pid    string
	reqid  int
	socket net.Conn

	mutex    *sync.Mutex
	qchan    chan struct{}
	requests map[int]*request
	// subscribers map[string]func()
}

func NewClient(pid string) *Client {
	client := &Client{pid: pid, reqid: 1}
	client.mutex = new(sync.Mutex)
	client.qchan = make(chan struct{})
	client.requests = make(map[int]*request)
	// client.subscribers = make(map[string]func())
	return client
}

func (c *Client) Open() (err error) {
	c.socket, err = pipe.GetPipeSocket(c.pid)
	go c.readloop()
	return
}

func (c *Client) readloop() {
	reader := bufio.NewReader(c.socket)
	for {
		select {
		case <-c.qchan:
			break
		default:
			// # in case the client is closed already
			if c.socket == nil {
				return
			}
			// read data from socket
			data, err := reader.ReadBytes('\n')
			if err != nil {
				if errors.Is(err, io.EOF) {
					break
				}
				continue
			}
			// unmarshal received data
			var res response
			if err = json.Unmarshal(data, &res); err != nil {
				continue
			}
			// handle response
			c.mutex.Lock()
			if res.Event != "" {
				// // FIXME: not working
				// if cb, ok := c.subscribers[res.Event]; ok {
				// 	cb()
				// }
			} else if req, ok := c.requests[res.RequestID]; ok {
				delete(c.requests, res.RequestID)
				req.reschan <- &res
			}
			c.mutex.Unlock()
		}
	}
}

func (c *Client) write(req *request) (*request, error) {
	defer func() {
		c.reqid += 1
	}()
	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	if _, err = c.socket.Write(append(data, '\n')); err != nil {
		return nil, err
	}
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.requests[req.RequestID] = req
	return req, nil
}

// func (c *Client) Subscribe(event string, callback func()) {
// 	c.mutex.Lock()
// 	defer c.mutex.Unlock()
// 	c.subscribers[event] = callback
// }

func (c *Client) Call(cmd string, args ...interface{}) (interface{}, error) {
	req, err := c.write(newRequest(c.reqid, cmd, args...))
	if err != nil {
		return nil, err
	}
	return req.Response()
}

func (c *Client) GetProperty(key string) (interface{}, error) {
	return c.Call("get_property", key)
}

func (c *Client) GetPropertyString(key string) (string, error) {
	value, err := c.Call("get_property_string", key)
	if err != nil {
		return "", err
	}
	return value.(string), nil
}

func (c *Client) Close() error {
	defer func() {
		c.socket = nil
	}()
	c.qchan <- struct{}{}
	return c.socket.Close()
}

func (c *Client) IsClosed() bool { return c.socket == nil }
