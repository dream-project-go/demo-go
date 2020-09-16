package signparam

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"log"
	//"time"
)

//生成加密字符串
func GenAuthToken(paramStr string, headerKey string) string {
	s := []byte(paramStr)
	key := []byte(headerKey)
	m := hmac.New(sha256.New, key)
	m.Write(s)
	signature := base64.StdEncoding.EncodeToString(m.Sum(nil))
	return signature
}

//header参数验证
func CheckHeader(c *gin.Context) bool {
	var err error
	var d []byte
	d, err = c.GetRawData()
	if err != nil {
		log.Fatal(err)
	}
	authStr := c.GetHeader("MAuthorization")
	postStr := string(d)
	queryStr := c.Request.RequestURI
	method := c.Request.Method
	var paramsStr string
	switch method {
	case "GET":
		paramsStr = queryStr
	case "POST":
		paramsStr = postStr
	default:
		paramsStr = postStr
	}
	key := "SwYNTwt5qsOHN0ms46Ms9Atyi0" //参数加密密钥读取配置文件
	genAuthStr := GenAuthToken(paramsStr, key)
	if authStr != genAuthStr {
		log.Println("authStr", authStr, "genAuthStr", genAuthStr, "queryStr", queryStr, "postStr", postStr, "paramsStr", paramsStr, c.Request.Method)
		return false
	}
	return true
}
