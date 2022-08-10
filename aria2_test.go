package aria2go

import (
	"fmt"
	"testing"
)

var client *Aria2Client

func TestMain(m *testing.M) {
	client = NewAria2Client("thanks")
	m.Run()
}

func TestAria2Download(t *testing.T) {
	downloadFileUri := "https://dl.google.com/go/go1.18.4.linux-amd64.tar.gz"
	downloadRequest, id, err := NewRequest().SetToken(client.Token).AddUri([]string{downloadFileUri}, &Option{
		Out: "bbb.tar.gz",
	}).Create()
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Println("request id: ", id)

	result, err := client.SendRequest(downloadRequest)
	if err != nil {
		t.Error(err.Error())
		return
	}

	fmt.Println("download result: ", result)
}

func TestAria2Shutdown(t *testing.T) {
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

	fmt.Printf("request id: %s, response: %s\n", id, result.Result)
}
