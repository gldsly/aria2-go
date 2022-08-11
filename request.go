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
	ID        string        `json:"id"`
	Method    string        `json:"method"`
	Params    []interface{} `json:"params"`
	errorInfo error
}

// NewRequest 创建请求
func NewRequest() *BasicRequestBody {
	return &BasicRequestBody{
		JSONRPC: DEFAULT_JSONRPC_VERSION,
		ID:      DEFAULT_ID,
		Method:  "",
		Params:  []interface{}{},
	}
}

// NewRequestWithToken 创建请求(Token)
func NewRequestWithToken(token string) *BasicRequestBody {
	return &BasicRequestBody{
		JSONRPC: DEFAULT_JSONRPC_VERSION,
		ID:      DEFAULT_ID,
		Method:  "",
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
func (b *BasicRequestBody) Create() (result []byte, id string, err error) {
	if b.errorInfo != nil {
		return nil, "", b.errorInfo
	}
	id = uuid.New().String()
	b.ID = id
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

// Shutdown 关闭 aria2
func (b *BasicRequestBody) Shutdown() *BasicRequestBody {
	if b.errorInfo != nil {
		return b
	}

	b.Method = "aria2.shutdown"
	return b
}
