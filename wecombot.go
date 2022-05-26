package wecombot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	Key string
}

type ReportData struct {
	Msgtype string            `json:"msgtype"`
	Text    map[string]string `json:"text"`
}

const Default = "default"

func (c *Client) SendSingleBotWarningReport(data ReportData, robot string) (string, error) {
	switch robot {
	case Default:
		robot = Default
	}
	query := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=%s", c.Key)
	body, _ := json.Marshal(data)
	resp, err := http.Post(query, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
