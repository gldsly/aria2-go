package aria2go

import (
	"encoding/json"
	"errors"

	"github.com/google/uuid"
)

// BasicRequestBody 基础请求结构体
type BasicRequestBody struct {
	JSONRPC string        `json:"jsonrpc"`
	ID      string        `json:"id"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
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
		Params:  []interface{}{
			"token:"+token,
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

func (b *BasicRequestBody) Create() (result []byte, id string, err error) {
	if b.errorInfo != nil {
		return nil, "", b.errorInfo
	}
	id = uuid.New().String()
	b.ID = id
	result, err = json.Marshal(b)
	return
}

// Download 下载请求
func (b *BasicRequestBody) Download(downloadSourceUri []string, option *DownloadRequestOption) *BasicRequestBody {
	if b.errorInfo != nil {
		return b
	}

	if len(downloadSourceUri) < 1 {
		b.errorInfo = errors.New("download source uri is required")
		return b
	}

	b.Method = "aria2.addUri"
	b.Params = append(b.Params, downloadSourceUri)
	if option != nil {
		b.Params = append(b.Params, option)
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

