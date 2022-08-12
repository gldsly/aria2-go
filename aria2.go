package aria2go

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Aria2Client struct {
	Token string
	Addr  string
	Port  string
}

type Aria2ClientOption func(*Aria2Client)

func ClientSetAddr(addr string) Aria2ClientOption {
	return func(client *Aria2Client) {
		client.Addr = addr
	}
}

func ClientSetPort(port string) Aria2ClientOption {
	return func(client *Aria2Client) {
		client.Port = port
	}
}

func NewAria2Client(token string, opt ...Aria2ClientOption) *Aria2Client {
	token = strings.TrimSpace(token)
	client := &Aria2Client{Token: token, Addr: DEFAULT_ARIA2_ADDR, Port: DEFAULT_ARIA2_PORT}

	for _, obj := range opt {
		obj(client)
	}

	return client
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
		return nil, fmt.Errorf("code: %d message: %s", resp.Error.Code, resp.Error.Message)
	}
	return resp.Result, nil
}

func (a Aria2Client) QueryWaitingTask(offset int, limit int) (tasks []*TaskStatusData, err error) {
	request, _, err := NewRequestWithToken(a.Token).TellWaiting(offset, limit).Create()
	if err != nil {
		return nil, err
	}
	requestResult, err := a.SendRequest(request)
	if err != nil {
		return nil, err
	}
	resp := &TellTaskListResponse{}
	if err := json.Unmarshal(requestResult, &resp); err != nil {
		return nil, err
	}

	if resp.Error != nil {
		return nil, fmt.Errorf("code: %d message: %s", resp.Error.Code, resp.Error.Message)
	}
	return resp.Result, nil
}

func (a Aria2Client) QueryStoppedTask(offset int, limit int) (tasks []*TaskStatusData, err error) {
	request, _, err := NewRequestWithToken(a.Token).TellStopped(offset, limit).Create()
	if err != nil {
		return nil, err
	}
	requestResult, err := a.SendRequest(request)
	if err != nil {
		return nil, err
	}
	resp := &TellTaskListResponse{}
	if err := json.Unmarshal(requestResult, &resp); err != nil {
		return nil, err
	}

	if resp.Error != nil {
		return nil, fmt.Errorf("code: %d message: %s", resp.Error.Code, resp.Error.Message)
	}
	return resp.Result, nil
}

func (a Aria2Client) QueryDownloadingTask() (tasks []*TaskStatusData, err error) {
	request, _, err := NewRequestWithToken(a.Token).TellActive().Create()
	if err != nil {
		return nil, err
	}
	requestResult, err := a.SendRequest(request)
	if err != nil {
		return nil, err
	}
	resp := &TellTaskListResponse{}
	if err := json.Unmarshal(requestResult, &resp); err != nil {
		return nil, err
	}

	if resp.Error != nil {
		return nil, fmt.Errorf("code: %d message: %s", resp.Error.Code, resp.Error.Message)
	}
	return resp.Result, nil
}

func (a Aria2Client) QueryNotDownloadingTask() (tasks []*TaskStatusData, err error) {
	offset := 0
	limit := 50
	count := 1
	waitingTasks := make([]*TaskStatusData, 0)
	stoppedTasks := make([]*TaskStatusData, 0)

	for {
		waitingReq := NewRequestWithToken(a.Token).TellWaiting(offset, limit)
		stoppedReq := NewRequestWithToken(a.Token).TellStopped(offset, limit)
		request, _, err := NewRequest().MultiCall(waitingReq, stoppedReq).Create()
		if err != nil {
			return nil, err
		}
		result, err := a.SendRequest(request)
		if err != nil {
			return nil, err
		}
		resp := &QueryNotDownloadingTaskResponse{}
		err = json.Unmarshal(result, &resp)
		if err != nil {
			return nil, err
		}
		if resp.Error != nil {
			return nil, fmt.Errorf("code: %d message: %s", resp.Error.Code, resp.Error.Message)
		}

		waitingTaskRes := resp.Result[0][0]
		stoppedTaskRes := resp.Result[1][0]

		if len(waitingTaskRes) == 0 && len(stoppedTaskRes) == 0 {
			break
		}
		waitingTasks = append(waitingTasks, waitingTaskRes...)
		stoppedTasks = append(stoppedTasks, stoppedTaskRes...)
		count += 1
		offset = count * limit
	}

	allTasks := make([]*TaskStatusData, 0, len(waitingTasks)+len(stoppedTasks))
	allTasks = append(allTasks, waitingTasks...)
	allTasks = append(allTasks, stoppedTasks...)

	return allTasks, nil
}

func (a Aria2Client) Pause(gid string) error {
	request, _, err := NewRequestWithToken(a.Token).Pause(gid, false).Create()
	if err != nil {
		return err
	}
	requestResult, err := a.SendRequest(request)
	if err != nil {
		return err
	}
	resp := &Response{}
	if err := json.Unmarshal(requestResult, &resp); err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf("code: %d message: %s", resp.Error.Code, resp.Error.Message)
	}
	return nil
}

func (a Aria2Client) Unpause(gid string) error {
	request, _, err := NewRequestWithToken(a.Token).Unpause(gid).Create()
	if err != nil {
		return err
	}
	requestResult, err := a.SendRequest(request)
	if err != nil {
		return err
	}
	resp := &Response{}
	if err := json.Unmarshal(requestResult, &resp); err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf("code: %d message: %s", resp.Error.Code, resp.Error.Message)
	}
	return nil
}

func (a Aria2Client) PauseAll(gid string) error {
	request, _, err := NewRequestWithToken(a.Token).PauseAll(false).Create()
	if err != nil {
		return err
	}
	requestResult, err := a.SendRequest(request)
	if err != nil {
		return err
	}
	resp := &Response{}
	if err := json.Unmarshal(requestResult, &resp); err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf("code: %d message: %s", resp.Error.Code, resp.Error.Message)
	}
	return nil
}

func (a Aria2Client) UnpauseAll(gid string) error {
	request, _, err := NewRequestWithToken(a.Token).UnpauseAll().Create()
	if err != nil {
		return err
	}
	requestResult, err := a.SendRequest(request)
	if err != nil {
		return err
	}
	resp := &Response{}
	if err := json.Unmarshal(requestResult, &resp); err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf("code: %d message: %s", resp.Error.Code, resp.Error.Message)
	}
	return nil
}

func (a Aria2Client) RemoveTask(gid string) error {
	request, _, err := NewRequestWithToken(a.Token).RemoveDownloadResult(gid).Create()
	if err != nil {
		return err
	}
	requestResult, err := a.SendRequest(request)
	if err != nil {
		return err
	}
	resp := &Response{}
	if err := json.Unmarshal(requestResult, &resp); err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf("code: %d message: %s", resp.Error.Code, resp.Error.Message)
	}
	return nil
}

func (a Aria2Client) RemoveAllTask() error {
	request, _, err := NewRequestWithToken(a.Token).PurgeDownloadResult().Create()
	if err != nil {
		return err
	}
	requestResult, err := a.SendRequest(request)
	if err != nil {
		return err
	}
	resp := &Response{}
	if err := json.Unmarshal(requestResult, &resp); err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf("code: %d message: %s", resp.Error.Code, resp.Error.Message)
	}
	return nil
}
