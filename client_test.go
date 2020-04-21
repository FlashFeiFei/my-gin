package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"testing"
)

//压力测试Oauth的addclient接口
//运行go test -v -bench BenchmarkOauthAddClient
func BenchmarkOauthAddClient(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data := url.Values{
			"client_name":  {"name_" + strconv.FormatInt(int64(i), 10)},
			"redirect_url": {"http://www.baidu.com"},
		}
		response, _ := http.PostForm("http://localhost:8083/oauth/client/addclient", data)
		body, _ := ioutil.ReadAll(response.Body)
		response.Body.Close()
		b.Log(string(body))
	}
}

//压力测试一下自己封装了gin之后的性能，看看会发生什么事情
func BenchmarkMyGinHttp(b *testing.B) {
	b.N = 100000
	for i := 0; i < b.N; i++ {
		//data := url.Values{
		//	"client_name":  {"name_" + strconv.FormatInt(int64(i), 10)},
		//	"redirect_url": {"http://www.baidu.com"},
		//}
		response, _ := http.Get(fmt.Sprintf("http://127.0.0.1:8083/background/user/hello_world2/%d", i))
		body, _ := ioutil.ReadAll(response.Body)
		response.Body.Close()
		b.Log(string(body))
	}
}
