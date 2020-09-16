package authlite

import(
	"github.com/gin-gonic/gin"
	"micro-common/plugins/authlite/apitoken"
	"micro-common/plugins/authlite/signparam"
	"micro-common/utils"
	"time"
)

const (
	accessTokenExpire = 3600 * 1
	//refreshTokenExpire = 3600 * 24
)

//验证登陆权限
func CheckAuth(c *gin.Context){
	accessToken := c.GetHeader("accessToken")
	if accessToken == ""{
		utils.ErrorResp(c,"error_0000","access_token不能为空")
	}
	m := apitoken.NewApiToken()
	info,_ := m.CGetAccessToken(accessToken)
	if info.Id > 0{
		if time.Now().Unix() - info.Mtime > accessTokenExpire{
			utils.ErrorResp(c,"error_0002","access_token已过期")
		}
	}else{
		utils.ErrorResp(c,"error_0001","access_token错误")
	}
	c.Set("userId",info.UserId)
}

//验证参数加密
func CheckSignParam(c *gin.Context){
	if !signparam.CheckHeader(c){
		utils.ErrorResp(c,"error_0003","非法访问")
	}
}




