package main

import (
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
