# aria2-go

使用 jsonrpc 和 aria2 通信的 sdk

完整实现 文档中 jsonrpc 的 method 请求构造

预置常用的操作方法

下载文件
```go
package main

import (
    "fmt"
    aria2go "github.com/gldsly/aria2-go"
)

func main() {
    client := aria2go.NewAria2Client("thanks")
    gid, err := client.Download("https://dl.google.com/go/go1.18.4.linux-amd64.tar.gz")
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(gid)
}
```