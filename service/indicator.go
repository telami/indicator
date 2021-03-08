package service

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const CarIndicatorUrl = "http://apply.xkctk.jtys.tj.gov.cn/apply/norm/personQuery.html"

// 恭喜，中签！
func Congratulate(code string, name string) bool {
	res := GetIndicatorInfo(code)
	if strings.Contains(res, name) {
		return true
	}
	return false
}

// 请求网站数据
func GetIndicatorInfo(appleCode string) string {

	params := url.Values{}
	params.Set("pageNo", "1")
	params.Set("issueNumber", "000000")
	params.Set("applyCode", appleCode)
	res, err := http.PostForm(CarIndicatorUrl, params)
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
