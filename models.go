package aria2go

// DownloadRequestOption 下载请求参数配置
type DownloadRequestOption struct {
	Dir string `json:"dir"`
	Out string `json:"out"`
}

// Response aria2 通常响应
type Response struct {
	ID      string        `json:"id"`
	JSONRPC string        `json:"jsonrpc"`
	Result  string        `json:"result"`
	Error   *ResponseError `json:"error"`
}

// ResponseError Response 中 Error 字段
type ResponseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
