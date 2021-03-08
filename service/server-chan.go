package service

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const serverUrl = "https://sctapi.ftqq.com/%s.send?title=%s"

func SendWechatNotification(secret string, content string) string {

	res, err := http.Get(fmt.Sprintf(serverUrl, secret, content))
	if err != nil {
		return ""
	}
	robots, err := ioutil.ReadAll(res.Body)
	_ = res.Body.Close()
	if err != nil {
		return ""
	}
	return string(robots)
}
