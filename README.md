# aria2-go

使用 jsonrpc 和 aria2 通信的 sdk

完整实现 文档中 jsonrpc 的 method 请求构造

预置常用的操作方法

下载文件
```go
package main

import (
	"encoding/json"
	"fmt"

	aria2go "github.com/gldsly/aria2-go"
)

var client *aria2go.Aria2Client

// Download 下载文件
func Download() {
	gid, err := client.Download("https://dl.google.com/go/go1.18.4.linux-amd64.tar.gz")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(gid)
}

// GetTaskStatus 查询任务状态
func GetTaskStatus() {
	status, err := client.QueryTaskStatus("fc9b15cfe91ee80b")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(status)
}

// GetStoppedTask 构建请求发送到 aria2 中查询
func GetStoppedTask() {
	request, replayID, err := aria2go.NewRequestWithToken(client.Token).TellStopped(0, 10).Create()

	if err != nil {
		fmt.Println(err)
		return
	}

	result, err := client.SendRequest(request)
	if err != nil {
		fmt.Println(err)
		return
	}

	resp := &aria2go.TellTaskListResponse{}
	if err := json.Unmarshal(result, &resp); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("replay id: %s \nresult: %#v\n",replayID, resp.Result[0])
}

func main() {
	client = aria2go.NewAria2Client("thanks")
	GetStoppedTask()
}
```