package apitoken

import (
	"encoding/json"
	"fmt"
	"github.com/micro/go-micro/util/log"
	"math/rand"
	"micro-common/library/commom/redis"
	"micro-common/plugins/db"
	"time"
	"unicode/utf8"
)

type ApiToken struct {
	Id           int64  `gorm:"column:id" json:"id"`
	UserId       int64  `gorm:"column:user_id" json:"user_id"`             // 用户id
	Device       string `gorm:"column:device" json:"device"`               // 设备类型：1android  2ios
	AccessToken  string `gorm:"column:access_token" json:"access_token"`   // 访问令牌
	RefreshToken string `gorm:"column:refresh_token" json:"refresh_token"` // 刷新令牌
	Mtime        int64  `gorm:"column:mtime" json:"mtime"`                 // 修改时间
}

const (
	ApiTokenTable = "gugu_api_token"
	DBName        = "chumanlite"
)

func NewApiToken() *ApiToken {
	return &ApiToken{}
}

func (a *ApiToken) Model() *db.Model {
	return &db.Model{
		DB:    DBName,
		Table: ApiTokenTable,
	}
}

//通过accessToken获取令牌详情
func (a *ApiToken) CGetAccessToken(accessToken string) (ApiToken, error) {
	key := fmt.Sprintf("ApiTokenInfo:%s", accessToken)
	client := redis.GetRedis("chumanlite")
	var err error
	var d ApiToken
	if !client.Exists(key) {
		err = a.Model().Slave().Table(ApiTokenTable).Where("access_token = ?", accessToken).First(&d).Error
		if db.CheckIsError(err) {
			log.Info(err)
		}
		if d.Id > 0 {
			bytes, _ := json.Marshal(d)
			client.Set(key, string(bytes), time.Second*3600*24*1)
		} else { //防止穿透
			d.Id = 0
			bytes, _ := json.Marshal(d)
			client.Set(key, string(bytes), time.Second*360)
		}
	}
	if cacheData := client.Get(key); cacheData != "" {
		err = json.Unmarshal([]byte(cacheData), &d)
		if err != nil {
			log.Infof("error: %v", err)
		}
	}
	return d, err
}

//更新令牌
func (a *ApiToken) GenAccessRefreshByUidDevice(userId int64, device int64) (string, string, error) {
	accessToken := GenRand(25, false, false)
	refreshToken := GenRand(25, false, false)
	now := time.Now().Unix()
	sql := fmt.Sprintf("INSERT INTO `chumanlite`.`gugu_api_token` (`user_id`,`device`,`access_token`,`refresh_token`,`mtime`) VALUES('%d','%d','%s','%s','%d') ON DUPLICATE KEY UPDATE `access_token` = '%s',`refresh_token` = '%s',`mtime`= %d;",
		   userId,device,accessToken,refreshToken,now,accessToken,refreshToken,now)
	err := a.Model().Master().Exec(sql).Error
	if db.CheckIsError(err) {
		log.Info(err)
	}
	return accessToken, refreshToken, err
}

//生成随机字符串
func GenRand(n int, isNumber bool, isMixedLetter bool) string {
	var letter string
	if !isNumber && isMixedLetter {
		letter = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	} else if isNumber{
		letter = "0123456789"
	}else{
		letter = "0123456789abcdefghijklmnopqrstuvwxyz"
	}
	b := make([]byte, n)
	rand.Seed(time.Now().UnixNano())
	for i := range b {
		b[i] = letter[rand.Intn(utf8.RuneCountInString(letter))]
	}
	return string(b)
}
