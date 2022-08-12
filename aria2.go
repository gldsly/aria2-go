package aria2go

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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
func (a Aria2Client) SendRequest(body []byte) (result []byte, err error) {
	serverAddr := fmt.Sprintf("http://%s:%s/jsonrpc", a.Addr, a.Port)
	request, err := http.NewRequest("POST", serverAddr, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	request.Header.Set("ContentType", DEFAULT_CONTENT_TYPE)
	request.Header.Set("Accept-Charset", "utf-8")

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	resultJsonData, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	result = resultJsonData

	return
}

func (a Aria2Client) Download(uri string) (gid string, err error) {
	downloadRequest, _, err := NewRequestWithToken(a.Token).AddUri([]string{uri}, nil).Create()
	if err != nil {
		return "", err
	}

	requestResult, err := a.SendRequest(downloadRequest)
	if err != nil {
		return "", err
	}

	resp := &Response{}
	err = json.Unmarshal(requestResult, &resp)
	if err != nil {
		return "", err
	}

	if resp.Error != nil {
		return "", fmt.Errorf("code: %d message: %s", resp.Error.Code, resp.Error.Message)
	}
	return resp.Result, nil
}

func (a Aria2Client) DownloadWithLocalTorrent(filePath string) (gid string, err error) {
	downloadRequest, _, err := NewRequestWithToken(a.Token).AddTorrent(filePath, nil).Create()
	if err != nil {
		return "", err
	}

	requestResult, err := a.SendRequest(downloadRequest)
	if err != nil {
		return "", err
	}

	resp := &Response{}
	err = json.Unmarshal(requestResult, &resp)
	if err != nil {
		return "", err
	}

	if resp.Error != nil {
		return "", fmt.Errorf("code: %d message: %s", resp.Error.Code, resp.Error.Message)
	}
	return resp.Result, nil
}

func (a Aria2Client) QueryTaskStatus(gid string) (status *TaskStatusData, err error) {
	request, _, err := NewRequestWithToken(a.Token).TellStatus(gid).Create()
	if err != nil {
		return nil, err
	}
	requestResult, err := a.SendRequest(request)
	if err != nil {
		return nil, err
	}
	resp := &TellStatusResponse{}
	if err := json.Unmarshal(requestResult, &resp); err != nil {
		return nil, err
	}

	if resp.Error != nil {
		return nil, err
	}
	return resp.Result, nil
}

