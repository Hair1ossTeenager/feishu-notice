package tool

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// 发送飞书机器人消息
func SendFeishuMessage(webHook string, buff bytes.Buffer) bool {
	preUrl := "https://open.feishu.cn/open-apis/bot/v2/hook/"
	url := preUrl + webHook
	resp, err := http.Post(url, "application/json", &buff)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
	return true
}
