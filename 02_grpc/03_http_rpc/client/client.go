package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	// 定义请求的 URL
	url := "http://localhost:1234/jsonrpc"
	// 定义请求体数据
	data := strings.NewReader(`{"id": 0, "params": ["shane"], "method": "HelloService.Hello"}`)
	// 创建一个新的 POST 请求
	req, err := http.NewRequest("POST", url, data)
	if err != nil {
		fmt.Println("创建请求出错:", err)
		return
	}
	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("发送请求出错:", err)
		return
	}
	// 确保在函数结束时关闭响应体
	defer resp.Body.Close()
	// 读取响应体内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应内容出错:", err)
		return
	}
	// 打印响应状态码和内容
	fmt.Println("响应内容:", string(body))
}
