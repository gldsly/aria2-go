package aria2go

import (
	"encoding/json"
	"fmt"
	"testing"
)

var client *Aria2Client

func TestMain(m *testing.M) {
	client = NewAria2Client("thanks")
	m.Run()
}

func TestTellWaiting(t *testing.T) {
	request, replayID, err := NewRequestWithToken(client.Token).TellWaiting(0, 20).Create()
	if err != nil {
		t.Error(err.Error())
		return
	}
	result, err := client.SendRequest(request)
	if err != nil {
		t.Error(err.Error())
		return
	}
	resp := &TellTaskListResponse{}
	if err := json.Unmarshal(result, &resp); err != nil {
		t.Error(err.Error())
		return
	}

	if resp.Error != nil {
		t.Error(fmt.Errorf("code: %d  message: %s", resp.Error.Code, resp.Error.Message))
		return
	}
	fmt.Println("replay id: ", replayID)
	fmt.Println("等待开始的任务: ")
	for _, task := range resp.Result {
		fmt.Printf("gid: %s 已下载: %s 总大小: %s 下载速度: %s 状态: %s 连接数: %s\n",
			task.Gid, task.CompletedLength, task.TotalLength, task.DownloadSpeed,
			task.Status, task.Connections)
		for _, file := range task.Files {
			fmt.Printf("\t 索引: %s 文件信息: %s 已下载: %s 总大小: %s \n", file.Index,
				file.Path, file.CompletedLength, file.Length)
		}
	}
}

func TestTellStopped(t *testing.T) {
	request, replayID, err := NewRequestWithToken(client.Token).TellStopped(0, 20).Create()
	if err != nil {
		t.Error(err.Error())
		return
	}
	result, err := client.SendRequest(request)
	if err != nil {
		t.Error(err.Error())
		return
	}
	resp := &TellTaskListResponse{}
	if err := json.Unmarshal(result, &resp); err != nil {
		t.Error(err.Error())
		return
	}

	if resp.Error != nil {
		t.Error(fmt.Errorf("code: %d  message: %s", resp.Error.Code, resp.Error.Message))
		return
	}

	fmt.Println("replay id: ", replayID)
	fmt.Println("停止的任务: ")
	for _, task := range resp.Result {
		fmt.Printf("gid: %s 已下载: %s 总大小: %s 下载速度: %s 状态: %s 连接数: %s\n",
			task.Gid, task.CompletedLength, task.TotalLength, task.DownloadSpeed,
			task.Status, task.Connections)
		for _, file := range task.Files {
			fmt.Printf("\t 索引: %s 文件信息: %s 已下载: %s 总大小: %s \n", file.Index,
				file.Path, file.CompletedLength, file.Length)
		}
	}
}

func TestGetUris(t *testing.T) {
	request, replayID, err := NewRequestWithToken(client.Token).GetUris("20bf613078e59fff").Create()
	if err != nil {
		t.Error(err.Error())
		return
	}
	result, err := client.SendRequest(request)
	if err != nil {
		t.Error(err.Error())
		return
	}
	resp := &GetUrisResponse{}
	if err := json.Unmarshal(result, &resp); err != nil {
		t.Error(err.Error())
		return
	}

	if resp.Error != nil {
		t.Error(fmt.Errorf("code: %d  message: %s", resp.Error.Code, resp.Error.Message))
		return
	}

	fmt.Println(replayID, resp.Result)
}

func TestTellStatus(t *testing.T) {
	request, replayID, err := NewRequestWithToken(client.Token).TellStatus("4c2e5dabceed5065").Create()
	if err != nil {
		t.Error(err.Error())
		return
	}
	result, err := client.SendRequest(request)
	if err != nil {
		t.Error(err.Error())
		return
	}
	resp := &TellStatusResponse{}
	if err := json.Unmarshal(result, &resp); err != nil {
		t.Error(err.Error())
		return
	}

	if resp.Error != nil {
		t.Error(fmt.Errorf("code: %d  message: %s", resp.Error.Code, resp.Error.Message))
		return
	}

	fmt.Println("replay id: ", replayID)
	fmt.Printf("gid: %s 已下载: %s 总大小: %s 下载速度: %s 状态: %s 连接数: %s\n",
		resp.Result.Gid, resp.Result.CompletedLength, resp.Result.TotalLength, resp.Result.DownloadSpeed,
		resp.Result.Status, resp.Result.Connections)
	for _, file := range resp.Result.Files {
		fmt.Printf("\t 索引: %s 文件信息: %s 已下载: %s 总大小: %s \n", file.Index,
			file.Path, file.CompletedLength, file.Length)
	}
}

func TestTellActive(t *testing.T) {
	// 1686b089f8a2e41f
	request, replayID, err := NewRequestWithToken(client.Token).TellActive().Create()
	if err != nil {
		t.Error(err.Error())
		return
	}
	result, err := client.SendRequest(request)
	if err != nil {
		t.Error(err.Error())
		return
	}
	resp := &TellTaskListResponse{}
	if err := json.Unmarshal(result, &resp); err != nil {
		t.Error(err.Error())
		return
	}

	if resp.Error != nil {
		t.Error(fmt.Errorf("code: %d  message: %s", resp.Error.Code, resp.Error.Message))
		return
	}
	fmt.Println("replay id: ", replayID)
	fmt.Println("进行中的任务: ")
	for _, task := range resp.Result {
		fmt.Printf("gid: %s 已下载: %s 总大小: %s 下载速度: %s 状态: %s 连接数: %s\n",
			task.Gid, task.CompletedLength, task.TotalLength, task.DownloadSpeed,
			task.Status, task.Connections)
		for _, file := range task.Files {
			fmt.Printf("\t 索引: %s 文件信息: %s 已下载: %s 总大小: %s \n", file.Index,
				file.Path, file.CompletedLength, file.Length)
		}
	}
}

func TestRemove(t *testing.T) {
	// 1686b089f8a2e41f
	request, replayID, err := NewRequestWithToken(client.Token).Remove("c363ab43eea0182d", false).Create()
	if err != nil {
		t.Error(err.Error())
		return
	}
	result, err := client.SendRequest(request)
	if err != nil {
		t.Error(err.Error())
		return
	}

	resp := &Response{}
	if err := json.Unmarshal(result, &resp); err != nil {
		t.Error(err.Error())
		return
	}

	if resp.Error != nil {
		t.Error(fmt.Errorf("code: %d  message: %s", resp.Error.Code, resp.Error.Message))
		return
	}

	fmt.Printf("replay id: %s download result: %v\n", replayID, string(result))
}

func TestPause(t *testing.T) {
	// 1686b089f8a2e41f
	request, replayID, err := NewRequestWithToken(client.Token).Pause("1686b089f8a2e41f", false).Create()
	if err != nil {
		t.Error(err.Error())
		return
	}
	result, err := client.SendRequest(request)
	if err != nil {
		t.Error(err.Error())
		return
	}

	resp := &Response{}
	if err := json.Unmarshal(result, &resp); err != nil {
		t.Error(err.Error())
		return
	}

	if resp.Error != nil {
		t.Error(fmt.Errorf("code: %d  message: %s", resp.Error.Code, resp.Error.Message))
		return
	}

	fmt.Printf("replay id: %s download result: %v\n", replayID, result)
}

func TestPauseAll(t *testing.T) {
	request, replayID, err := NewRequestWithToken(client.Token).PauseAll(false).Create()
	if err != nil {
		t.Error(err.Error())
		return
	}
	result, err := client.SendRequest(request)
	if err != nil {
		t.Error(err.Error())
		return
	}

	resp := &Response{}
	if err := json.Unmarshal(result, &resp); err != nil {
		t.Error(err.Error())
		return
	}

	if resp.Error != nil {
		t.Error(fmt.Errorf("code: %d  message: %s", resp.Error.Code, resp.Error.Message))
		return
	}

	fmt.Printf("replay id: %s download result: %v\n", replayID, string(result))
}

func TestUnpauseAll(t *testing.T) {
	request, replayID, err := NewRequestWithToken(client.Token).UnpauseAll().Create()
	if err != nil {
		t.Error(err.Error())
		return
	}
	result, err := client.SendRequest(request)
	if err != nil {
		t.Error(err.Error())
		return
	}

	resp := &Response{}
	if err := json.Unmarshal(result, &resp); err != nil {
		t.Error(err.Error())
		return
	}

	if resp.Error != nil {
		t.Error(fmt.Errorf("code: %d  message: %s", resp.Error.Code, resp.Error.Message))
		return
	}

	fmt.Printf("replay id: %s download result: %v\n", replayID, string(result))
}

func TestTorrentDownload(t *testing.T) {
	testTorrentFile := "/Users/yw/Downloads/123.torrent"
	request, id, err := NewRequest().SetToken(client.Token).AddTorrent(testTorrentFile, nil).Create()
	if err != nil {
		t.Error(err.Error())
		return
	}

	result, err := client.SendRequest(request)
	if err != nil {
		t.Error(err.Error())
		return
	}

	resp := &Response{}
	if err := json.Unmarshal(result, &resp); err != nil {
		t.Error(err.Error())
		return
	}

	if resp.Error != nil {
		t.Error(fmt.Errorf("code: %d  message: %s", resp.Error.Code, resp.Error.Message))
		return
	}

	fmt.Printf("replay id: %s download result: %v", id, result)
}

func TestDownload(t *testing.T) {
	// downloadFileUri := "https://dl.google.com/go/go1.18.4.linux-amd64.tar.gz"
	downloadFileUri := "magnet:?xt=urn:btih:a9ff280d1418f98f9ad016e3a20acee1beb2de46&dn=zh-cn_windows_11_consumer_editions_version_21h2_updated_july_2022_x64_dvd_50ad4acf.iso&xl=5858816000"
	downloadRequest, id, err := NewRequest().SetToken(client.Token).AddUri([]string{downloadFileUri}, &Option{
		Out: "bbb.tar.gz",
	}).Create()
	if err != nil {
		t.Error(err.Error())
		return
	}

	result, err := client.SendRequest(downloadRequest)
	if err != nil {
		t.Error(err.Error())
		return
	}

	fmt.Printf("replay id: %s result: %s\n", id, string(result))
}

func TestShutdown(t *testing.T) {
	shutdownRequest, id, err := NewRequestWithToken(client.Token).Shutdown().Create()
	if err != nil {
		t.Error(err.Error())
		return
	}

	result, err := client.SendRequest(shutdownRequest)
	if err != nil {
		t.Error(err.Error())
		return
	}

	resp := &Response{}
	if err := json.Unmarshal(result, &resp); err != nil {
		t.Error(err.Error())
		return
	}

	if resp.Error != nil {
		t.Error(fmt.Errorf("code: %d  message: %s", resp.Error.Code, resp.Error.Message))
		return
	}

	fmt.Printf("replay id: %s, response: %s\n", id, string(result))
}
