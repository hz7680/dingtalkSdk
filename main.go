package main

import (
	"fmt"
	"github.com/hz7680/dingtalkSdk/client"
)

func main()  {
	var _client = client.NewClient("ding55eb045fd2ec4e54f5bf40eda33b7ba0","1044003882","dingwyz0renb6q7srhsy","tF5qW3M4g_JNGgNUGxe97OYpNdf5JN-BHLkgj8Kz7nabTl_dQxlPBsmR-jZBrhVz","4g5j64qlyl3zvetqxz5jiocdr586fn2zvjpa8zls3ij","ding55eb045fd2ec4e54f5bf40eda33b7ba0","123456")
	token, err := _client.GetAccessToken()
	fmt.Println(token)
	fmt.Println(err)
}