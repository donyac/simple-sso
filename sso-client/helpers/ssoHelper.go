package helpers

import (
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"github.com/mikemintang/go-curl"
)

type UserInfo struct {
	Name string
}

var appKey = "app1"

func IsLogin(ticket string) bool {
	userInfoPtr := GetUserInfo(ticket)
	if userInfoPtr == nil {
		return false
	} else if userInfoPtr.Name == "" {
		return false
	}
	return true
}

func GetUserInfo(ticket string) *UserInfo {

	url := "http://localhost:1080/user"
	queries := map[string]string{
		"appKey": appKey,
		"ticket": ticket,
	}
	req := curl.NewRequest()
	resp, err := req.SetUrl(url).SetQueries(queries).Get()

	if err != nil {
		logs.Debug(err)
		return nil
	} else if resp.IsOk() {
		logs.Debug(resp.Body)
		var jsonBlob = []byte(resp.Body)
		userInfoPtr := new(UserInfo)
		err = json.Unmarshal(jsonBlob, userInfoPtr)
		if err == nil {
			return userInfoPtr
		}
	} else {
		logs.Debug(resp.Raw)
		return nil
	}
	return nil
}

func Logout(ticket string) {
	url := "http://localhost:1080/logout"
	queries := map[string]string{
		"appKey": appKey,
		"ticket": ticket,
	}
	req := curl.NewRequest()
	req.SetUrl(url).SetQueries(queries).Post()
}
