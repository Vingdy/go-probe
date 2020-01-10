package dingding

import (
	"fmt"
	go_curl "github.com/mikemintang/go-curl"
	"probe/model"
	"strconv"
)

var (
	access_token = "29ce573267a1632a5b9dd195d8430d53722fb0055495e19749df678fefcfc436"
)

func SendPost(serverInfo []model.ServerVersionInfo) {
	url := "https://oapi.dingtalk.com/robot/send?access_token=" + access_token
	headers := map[string]string{
		"Content-Type": "application/json",
	}

	var version_text string

	for _, version := range serverInfo[0].VersionInfo {
		version_text += "> 版本号:" + version.Version + "   数量：" + strconv.Itoa(version.VersionNum) + "\n"
	}
	// 发送markdown消息
	postData := map[string]interface{}{
		"msgtype": "markdown",
		"markdown": map[string]interface{}{
			"title": "adn_tracking版本号检查",
			"text": "### adn_tracking版本号检查" + "\n" +
				"#### 服务器地区：" + serverInfo[0].ServerRegion + "   总数量:" + strconv.Itoa(serverInfo[0].Number) + "\n" +
				version_text,
		},
		"at": map[string]interface{}{
			"isAtAll": false,
		},
	}

	// 链式操作
	req := go_curl.NewRequest()
	resp, err := req.
		SetUrl(url).
		SetHeaders(headers).
		SetPostData(postData).
		Post()

	// 返回处理
	if err != nil {
		fmt.Println(err)
	} else {
		if resp.IsOk() {
			fmt.Println(resp.Body)
		} else {
			fmt.Println(resp.Raw)
		}
	}

}

/*
curl 'https://oapi.dingtalk.com/robot/send?access_token=29ce573267a1632a5b9dd195d8430d53722fb0055495e19749df678fefcfc436' \
-H 'Content-Type: application/json' \
-d '{"msgtype": "text",
"text": {
"content": "版本号我就是我, 是不一样的烟火"
}
}'*/
