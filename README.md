# go-wecom
this is a go wecom robot report demo
# Installing
>> go get github.com/Jacksmall/go-wecom
# Example
```
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Jacksmall/go-wecom/wecombot"
)

const KEY = "XXXXXXXXXXX"

func main() {
	client := &wecombot.Client{Key: KEY}

	title := "消息提醒【dev】" + time.Now().Format("Y-m-d H:i:s")
	subText := "hello world"
	robot := "default"
	data := wecombot.ReportData{
		Msgtype: "text",
		Text: map[string]string{
			"content": title + "\r\n" + subText,
		},
	}
	resp, err := client.SendSingleBotWarningReport(data, robot)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("send success string: %v\n", resp)
}
```

# Output
>> go run main.go
```
send success string: {"errcode":0,"errmsg":"ok"}
```

