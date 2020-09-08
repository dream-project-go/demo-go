package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

func CurlPost(url string, data interface{}, header map[string]string) []byte {
	//初始化client
	client := &http.Client{
		Timeout: 10 * time.Second, //超时设置
	}
	jd, err := json.Marshal(data)
	if err != nil {
		log.Fatal("json.Marshal(data)", err)
	}
	bdata := bytes.NewReader(jd)
	fmt.Print(bdata)
	req, err := http.NewRequest("POST", url, bytes.NewReader(jd))
	if len(header) > 0 { //设置请求头
		for k, v := range header {
			req.Header.Add(k, v)
		}
	}
	//发送请求
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("client.Do", err)
	}
	//从响应body里面读取返回的数据
	defer resp.Body.Close()
	ret, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("ioutil.ReadAll", err)
	}
	return ret
}

func BuildQuery(data map[string]string) {
}

func main() {
	h := make(map[string]string)
	h["Content-Type"] = "application/json"
	h["Mauthorization"] = "8cc4a3ffb5eef4b5f569ca110d390ee7:908512280b26cf78b642f8a8610069df:BZrH1lDfAayQUS7IiTpLO47guj8="
	pData := make(map[string]interface{})
	// pData["password"] = "FwNtNmO6JiGvEtzJ0JF_VrH1RAMBvPUf9ISyGxb-_qupzRo0mNiiGnIsj8y7iPcUMFV6_Y0nL8A_s0Exp_8DEmyRdrU3sb7JjmHri4ufyy1NJOu-Tlo_MAnU3llQjKk-mBpGHd4_700GRoDdjHYlEgPDxv_hsbPL-hTL7oNv6bk="
	// pData["user"] = "15013884809"
	// url := "https://api-demo.chumanapp.com/secure/?m=Api&c=User&a=login"

	var u url.URL
	url := "http://192.168.56.106:8085/?m=Api&c=User&a=login&name=cjb"
	q := u.Query()
	q.Add("password", "FwNtNmO6JiGvEtzJ0JF_VrH1RAMBvPUf9ISyGxb-_qupzRo0mNiiGnIsj8y7iPcUMFV6_Y0nL8A_s0Exp_8DEmyRdrU3sb7JjmHri4ufyy1NJOu-Tlo_MAnU3llQjKk-mBpGHd4_700GRoDdjHYlEgPDxv_hsbPL-hTL7oNv6bk=")
	q.Add("user", "15013884809")
	qStr := q.Encode()
	url += "&" + qStr

	ret := CurlPost(
		// "https://api-dev.chumanapp.com/secure/?m=Api&c=User&a=login",
		// "https://api.chumanapp.com/secure/?m=Api&c=User&a=status",
		url,
		pData,
		h,
	)
	log.Println(ret, string(ret))
}
