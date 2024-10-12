package mpvrpc

import "errors"

type response struct {
	Error     string      `json:"error"`
	Data      interface{} `json:"data"`
	Event     string      `json:"event"`
	RequestID int         `json:"request_id"`
}

type request struct {
	reschan chan *response

	Command   []interface{} `json:"command"`
	RequestID int           `json:"request_id"`
}

func (req *request) Response() (interface{}, error) {
	res := <-req.reschan
	if res.Error != "success" {
		return nil, errors.New(res.Error)
	}
	return res.Data, nil
}

func newRequest(reqid int, cmd string, args ...interface{}) *request {
	return &request{
		reschan:   make(chan *response, 1),
		Command:   append([]interface{}{cmd}, args...),
		RequestID: reqid,
	}
}
