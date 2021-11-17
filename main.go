package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)
import "github.com/gin-contrib/cors"

type DingTalkMsg struct {
	At struct {
		IsAtAll   bool     `json:"isAtAll"`
		AtMobiles []string `json:"atMobiles"`
	} `json:"at"`
	Text struct {
		Content string `json:"content"`
	} `json:"text"`
	Msgtype string `json:"msgtype"`
}

func alert(accessToken, mobile string) {
	url := "https://oapi.dingtalk.com/robot/send?access_token=" + accessToken
	method := "POST"

	msg := DingTalkMsg{
		At: struct {
			IsAtAll   bool     `json:"isAtAll"`
			AtMobiles []string `json:"atMobiles"`
		}{false, []string{mobile}},
		Text: struct {
			Content string `json:"content"`
		}{"!视频未播放"},
		Msgtype: "text",
	}

	marshal, err := json.Marshal(msg)
	if err != nil {
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewReader(marshal))

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	_, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {

	router := gin.Default()
	// CORS for https://foo.com and https://github.com origins, allowing:
	// - PUT and PATCH methods
	// - Origin header
	// - Credentials share
	// - Preflight requests cached for 12 hours
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://www.bjjnts.cn"},
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))
	router.GET("/alert", func(c *gin.Context) {
		mobile, _ := c.GetQuery("mobile")
		accessToken, _ := c.GetQuery("access_token")
		alert(accessToken, mobile)
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run()
}
