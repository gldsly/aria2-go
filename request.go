package aria2go

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"reflect"

	"github.com/google/uuid"
)

// BasicRequestBody 基础请求结构体
type BasicRequestBody struct {
	JSONRPC   string        `json:"jsonrpc"`
	ReplayID  string        `json:"id"`
	Method    string        `json:"method"`
	Params    []interface{} `json:"params"`
	errorInfo error
}

// NewRequest 创建请求
func NewRequest() *BasicRequestBody {
	return &BasicRequestBody{
		JSONRPC:  DEFAULT_JSONRPC_VERSION,
		ReplayID: "",
		Method:   "",
		Params:   []interface{}{},
	}
}

// NewRequestWithToken 创建请求(Token)
func NewRequestWithToken(token string) *BasicRequestBody {
	return &BasicRequestBody{
		JSONRPC:  DEFAULT_JSONRPC_VERSION,
		ReplayID: "",
		Method:   "",
		Params: []interface{}{
			"token:" + token,
		},
	}
}

// SetToken 设置访问令牌
func (b *BasicRequestBody) SetToken(token string) *BasicRequestBody {
	if b.errorInfo != nil {
		return b
	}

	b.Params = append(b.Params, "token:"+token)
	return b
}

// Create 创建请求数据
// 所有的请求链式调用末尾必须调用此函数来生成请求体数据
func (b *BasicRequestBody) Create() (result []byte, replayID string, err error) {
	if b.errorInfo != nil {
		return nil, "", b.errorInfo
	}
	replayID = uuid.New().String()
	b.ReplayID = replayID
	result, err = json.Marshal(b)
	return
}

// addParamsOption 添加 option 数据到 params 中
func (b *BasicRequestBody) addParamsOption(option *Option) {
	if option != nil {
		availableOption := make(map[string]string)

		v := reflect.ValueOf(*option)
		t := reflect.TypeOf(*option)
		totalFieldNum := v.NumField()
		for i := 0; i < totalFieldNum; i++ {
			key := t.Field(i).Tag.Get("json")
			value := v.Field(i).Interface().(string)

			if value != "" && key != "" {
				availableOption[key] = value
			}
		}

		b.Params = append(b.Params, availableOption)
	}
}

// AddUri 下载文件请求
func (b *BasicRequestBody) AddUri(downloadSourceUri []string, option *Option) *BasicRequestBody {
	if b.errorInfo != nil {
		return b
	}

	if len(downloadSourceUri) < 1 {
		b.errorInfo = errors.New("download source uri is required")
		return b
	}

	b.Method = "aria2.addUri"
	b.Params = append(b.Params, downloadSourceUri)
	b.addParamsOption(option)

	return b
}

// AddTorrent 添加本地 bt 文件创建下载任务
func (b *BasicRequestBody) AddTorrent(torrentFilePath string, option *Option) *BasicRequestBody {
	if b.errorInfo != nil {
		return b
	}
	torrentFile, err := os.Open(torrentFilePath)
	if err != nil {
		b.errorInfo = err
		return b
	}
	fileContent, err := ioutil.ReadAll(torrentFile)
	if err != nil {
		b.errorInfo = err
		return b
	}
	err = torrentFile.Close()
	if err != nil {
		b.errorInfo = err
		return b
	}
	torrent := base64.StdEncoding.EncodeToString(fileContent)

	b.Method = "aria2.addTorrent"
	b.Params = append(b.Params, torrent)
	b.addParamsOption(option)

	return b
}

// Remove 删除下载记录
// 如果是正在进行的下载任务,则会先停止在删除
// 返回删除的任务 gid
// 如果 force 为 true 则会直接删除.不会执行其他操作,例如联系 BitTorrent trackers 取消下载
func (b *BasicRequestBody) Remove(gid string, force bool) *BasicRequestBody {
	if b.errorInfo != nil {
		return b
	}
	if force {
		b.Method = "aria2.forceRemove"
	} else {
		b.Method = "aria2.remove"
	}
	b.Params = append(b.Params, gid)
	return b
}

// Pause 暂停任务下载
// 修改任务的状态为 paused
// 返回结果为被暂停的任务 gid
func (b *BasicRequestBody) Pause(gid string, force bool) *BasicRequestBody {
	if b.errorInfo != nil {
		return b
	}
	if force {
		b.Method = "aria2.forcePause"
	} else {
		b.Method = "aria2.pause"
	}
	b.Params = append(b.Params, gid)
	return b
}

// PauseAll 暂停所有下载
// 成功操作返回结果为 ok
func (b *BasicRequestBody) PauseAll(force bool) *BasicRequestBody {
	if b.errorInfo != nil {
		return b
	}
	if force {
		b.Method = "aria2.forcePauseAll"
	} else {
		b.Method = "aria2.pauseAll"
	}
	return b
}

// Unpause 取消任务暂停
// 修改任务状态从 paused -> waiting
// 操作成功返回: 任务 gid
func (b *BasicRequestBody) Unpause(gid string) *BasicRequestBody {
	if b.errorInfo != nil {
		return b
	}
	b.Method = "aria2.unpause"
	b.Params = append(b.Params, gid)
	return b
}

// UnpauseAll 取消所有任务暂停
// 操作成功返回: ok
func (b *BasicRequestBody) UnpauseAll() *BasicRequestBody {
	if b.errorInfo != nil {
		return b
	}
	b.Method = "aria2.unpauseAll"
	return b
}

// TellStatus 查询任务状态
// keys 可以指定返回字段,可指定的字段名参考 TaskStatusData
func (b *BasicRequestBody) TellStatus(gid string, keys ...string) *BasicRequestBody {
	if b.errorInfo != nil {
		return b
	}

	// TODO: 等待实现
	b.Method = "aria2.tellStatus"
	b.Params = append(b.Params, gid)
	if len(keys) > 0 {
		b.Params = append(b.Params, keys)
	}

	return b
}

// TellActive 查询所有正在进行中的任务
// keys 可以指定返回字段,可指定的字段名参考 TaskStatusData
func (b *BasicRequestBody) TellActive(keys ...string) *BasicRequestBody {
	if b.errorInfo != nil {
		return b
	}
	b.Method = "aria2.tellActive"
	if len(keys) > 0 {
		b.Params = append(b.Params, keys)
	}
	return b
}

// Shutdown 关闭 aria2
func (b *BasicRequestBody) Shutdown() *BasicRequestBody {
	if b.errorInfo != nil {
		return b
	}

	b.Method = "aria2.shutdown"
	return b
}
