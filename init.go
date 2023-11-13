package meituan

// 美团api接口地址前缀
const ApiUrl = "https://waimaiopen.meituan.com/api/v1/"

type MT struct {
	AppID       string
	AppSecret   string
	CallbackUrl string
}

// 初始化结构体
func Init(appID, appSecret, callbackUrl string) MT {
	return MT{
		AppID:       appID,
		AppSecret:   appSecret,
		CallbackUrl: callbackUrl,
	}
}
