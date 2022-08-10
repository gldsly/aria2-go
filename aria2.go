package aria2go

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Aria2Client struct {
	Token string
	Addr  string
	Port  string
}

func NewAria2Client(token string, serverAddrPort ...string) *Aria2Client {
	token = strings.TrimSpace(token)
	var addr string
	var port string
	if len(serverAddrPort) != 2 {
		addr = DEFAULT_ARIA2_ADDR
		port = DEFAULT_ARIA2_PORT
	} else {
		addr = serverAddrPort[0]
		port = serverAddrPort[1]
	}
	log.Printf("[aria2go] server addr: http://%s:%s/jsonrpc", DEFAULT_ARIA2_ADDR, DEFAULT_ARIA2_PORT)
	return &Aria2Client{
		Token: token,
		Addr:  addr,
		Port:  port,
	}
}

// SendRequest 发送请求
func (a Aria2Client) SendRequest(body []byte) (result *Response, err error) {
	serverAddr := fmt.Sprintf("http://%s:%s/jsonrpc", a.Addr, a.Port)
	request, err := http.NewRequest("POST", serverAddr, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	request.Header.Set("ContentType", DEFAULT_CONTENT_TYPE)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	resultJsonData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resultJsonData, &result)
	if err != nil {
		return nil, err
	}
	return
}
