package oauth

import (
	"fmt"
	"github.com/FlashFeiFei/my-gin/common-lib/crypto"
	"strconv"
	"time"
)

//生成ClientId
func GenerateClientId(now_time time.Time) (string, error) {
	//设置时区
	var cstSh, err = time.LoadLocation("Asia/Shanghai") //上海
	now_time.In(cstSh)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Oauth%s%d", now_time.Format("20060102"), now_time.Unix()), nil
}

func GenerateClientSecret(now_time time.Time) (string, error) {
	//设置时区
	var cstSh, err = time.LoadLocation("Asia/Shanghai") //上海
	now_time.In(cstSh)
	if err != nil {
		return "", err
	}
	unix_str := strconv.FormatInt(now_time.Unix(), 10)

	return crypto.SHA1(unix_str), nil
}