package discordrpc

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"net"
	"time"

	"github.com/tnychn/mpv-discord/discordrpc/payloads"
	"github.com/tnychn/mpv-discord/discordrpc/pipe"
)

type ClientError struct {
	Code    int
	Message string
}

func (err *ClientError) Error() string { return err.Message }

type Client struct {
	cid    string
	socket net.Conn
}

func NewClient(cid string) *Client {
	return &Client{cid: cid}
}

func (c *Client) read() error {
	// read 1024 bytes data from socket
	data := make([]byte, 1024)
	if _, err := c.socket.Read(data); err != nil {
		return err
	}
	// parse first 8 bytes (header)
	var header struct {
		OPCode int32
		Length int32
	}
	if err := binary.Read(bytes.NewReader(data[:8]), binary.LittleEndian, &header); err != nil {
		return err
	}
	// parse remaining bytes (payload)
	var payload struct {
		// - case 1
		Code    int    `json:"code"`
		Message string `json:"message"`
		// - case 2
		Evt  string `json:"evt"`
		Data struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		} `json:"data"`
	}
	if err := json.Unmarshal(bytes.Trim(data[8:], "\x00"), &payload); err != nil {
		return err
	}
	// handle error (if any)
	if (payload.Code != 0 && payload.Message != "") || payload.Evt == "ERROR" {
		return &ClientError{payload.Code, payload.Message}
	}
	return nil
}

func (c *Client) send(opcode int, payload payloads.Payload) error {
	d := time.Now().Add(time.Second * 3)
	if err := c.socket.SetDeadline(d); err != nil {
		return err
	}
	// encode data into JSON format
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	//  form the payload -> [header(opcode,length)][data]
	buffer := new(bytes.Buffer)
	if err = binary.Write(buffer, binary.LittleEndian, int32(opcode)); err != nil {
		return err
	}
	if err = binary.Write(buffer, binary.LittleEndian, int32(len(data))); err != nil {
		return err
	}
	if _, err = buffer.Write(data); err != nil {
		return err
	}
	// send out the payload
	if _, err = c.socket.Write(buffer.Bytes()); err != nil {
		return err
	}
	// wait for response and read it
	return c.read()
}

func (c *Client) Open() (err error) {
	if c.socket, err = pipe.GetPipeSocket(); err != nil {
		return
	}
	return c.send(0, payloads.Handshake{V: "1", ClientID: c.cid})
}

func (c *Client) Close() error {
	defer func() {
		c.socket = nil
	}()
	return c.socket.Close()
}

func (c *Client) IsClosed() bool { return c.socket == nil }
