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

// RequestBody 基础请求结构体
type RequestBody struct {
	JSONRPC   string        `json:"jsonrpc"`
	ReplayID  string        `json:"id"`
	Method    string        `json:"method"`
	Params    []interface{} `json:"params"`
	errorInfo error
}

// NewRequest 创建请求
func NewRequest() *RequestBody {
	return &RequestBody{
		JSONRPC:  DEFAULT_JSONRPC_VERSION,
		ReplayID: "",
		Method:   "",
		Params:   []interface{}{},
	}
}

// NewRequestWithToken 创建请求(Token)
func NewRequestWithToken(token string) *RequestBody {
	return &RequestBody{
		JSONRPC:  DEFAULT_JSONRPC_VERSION,
		ReplayID: "",
		Method:   "",
		Params: []interface{}{
			"token:" + token,
		},
	}
}

// SetToken 设置访问令牌
func (r *RequestBody) SetToken(token string) *RequestBody {
	if r.errorInfo != nil {
		return r
	}

	r.Params = append(r.Params, "token:"+token)
	return r
}

// Create 创建请求数据
// 所有的请求链式调用末尾必须调用此函数来生成请求体数据
func (r *RequestBody) Create() (result []byte, replayID string, err error) {
	if r.errorInfo != nil {
		return nil, "", r.errorInfo
	}
	replayID = uuid.New().String()
	r.ReplayID = replayID
	result, err = json.Marshal(r)
	return
}

// addParamsOption 添加 option 数据到 params 中
func (r *RequestBody) addParamsOption(option *Option) {
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

		r.Params = append(r.Params, availableOption)
	}
}

// AddUri 下载文件请求
func (r *RequestBody) AddUri(downloadSourceUri []string, option *Option) *RequestBody {
	if r.errorInfo != nil {
		return r
	}

	if len(downloadSourceUri) < 1 {
		r.errorInfo = errors.New("download source uri is required")
		return r
	}

	r.Method = "aria2.addUri"
	r.Params = append(r.Params, downloadSourceUri)
	r.addParamsOption(option)

	return r
}

// AddTorrent 添加本地 bt 文件创建下载任务
func (r *RequestBody) AddTorrent(torrentFilePath string, option *Option) *RequestBody {
	if r.errorInfo != nil {
		return r
	}
	torrentFile, err := os.Open(torrentFilePath)
	if err != nil {
		r.errorInfo = err
		return r
	}
	fileContent, err := ioutil.ReadAll(torrentFile)
	if err != nil {
		r.errorInfo = err
		return r
	}
	err = torrentFile.Close()
	if err != nil {
		r.errorInfo = err
		return r
	}
	torrent := base64.StdEncoding.EncodeToString(fileContent)

	r.Method = "aria2.addTorrent"
	r.Params = append(r.Params, torrent)
	r.addParamsOption(option)

	return r
}

// Remove 删除下载记录
// 如果 force 为 true 则会直接删除.不会执行其他操作,例如联系 BitTorrent trackers 取消下载
func (r *RequestBody) Remove(gid string, force bool) *RequestBody {
	if r.errorInfo != nil {
		return r
	}
	if force {
		r.Method = "aria2.forceRemove"
	} else {
		r.Method = "aria2.remove"
	}
	r.Params = append(r.Params, gid)
	return r
}

// Pause 暂停任务下载
// 修改任务的状态为 paused
func (r *RequestBody) Pause(gid string, force bool) *RequestBody {
	if r.errorInfo != nil {
		return r
	}
	if force {
		r.Method = "aria2.forcePause"
	} else {
		r.Method = "aria2.pause"
	}
	r.Params = append(r.Params, gid)
	return r
}

// PauseAll 暂停所有下载
func (r *RequestBody) PauseAll(force bool) *RequestBody {
	if r.errorInfo != nil {
		return r
	}
	if force {
		r.Method = "aria2.forcePauseAll"
	} else {
		r.Method = "aria2.pauseAll"
	}
	return r
}

// Unpause 取消任务暂停
// 修改任务状态从 paused -> waiting
func (r *RequestBody) Unpause(gid string) *RequestBody {
	if r.errorInfo != nil {
		return r
	}
	r.Method = "aria2.unpause"
	r.Params = append(r.Params, gid)
	return r
}

// UnpauseAll 取消所有任务暂停
func (r *RequestBody) UnpauseAll() *RequestBody {
	if r.errorInfo != nil {
		return r
	}
	r.Method = "aria2.unpauseAll"
	return r
}

// TellStatus 查询任务状态
// keys 可以指定返回字段,可指定的字段名参考 TaskStatusData
func (r *RequestBody) TellStatus(gid string, keys ...string) *RequestBody {
	if r.errorInfo != nil {
		return r
	}

	// TODO: 等待实现
	r.Method = "aria2.tellStatus"
	r.Params = append(r.Params, gid)
	if len(keys) > 0 {
		r.Params = append(r.Params, keys)
	}

	return r
}

// TellActive 查询所有正在进行中的任务
// keys 可以指定返回字段,可指定的字段名参考 TaskStatusData
func (r *RequestBody) TellActive(keys ...string) *RequestBody {
	if r.errorInfo != nil {
		return r
	}
	r.Method = "aria2.tellActive"
	if len(keys) > 0 {
		r.Params = append(r.Params, keys)
	}
	return r
}

// TellWaiting 查询等待执行的任务
// offset 设置偏移量 limit 限制每次显示多少
func (r *RequestBody) TellWaiting(offset, limit int, keys ...string) *RequestBody {
	r.Method = "aria2.tellWaiting"
	r.Params = append(r.Params, offset)
	r.Params = append(r.Params, limit)
	return r
}

// TellStopped 查询已经完成或者停止的任务
func (r *RequestBody) TellStopped(offset, limit int, keys ...string) *RequestBody {
	r.Method = "aria2.tellStopped"
	r.Params = append(r.Params, offset)
	r.Params = append(r.Params, limit)
	return r
}

// Shutdown 关闭 aria2
func (r *RequestBody) Shutdown() *RequestBody {
	if r.errorInfo != nil {
		return r
	}

	r.Method = "aria2.shutdown"
	return r
}

// GetUris 获取任务的文件下载源
func (r *RequestBody) GetUris(gid string) *RequestBody {
	if r.errorInfo != nil {
		return r
	}

	r.Method = "aria2.getUris"
	r.Params = append(r.Params, gid)
	return r
}

func (r *RequestBody) GetFiles(gid string) *RequestBody {
	if r.errorInfo != nil {
		return r
	}

	r.Method = "aria2.getFiles"
	r.Params = append(r.Params, gid)
	return r
}


