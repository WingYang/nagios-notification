package main

import(
	"net/http"
	"io/ioutil"
	"strings"
	"encoding/json"
	"bytes"
	"flag"
)

const (
	AccessTokenUrl = "https://qyapi.weixin.qq.com/cgi-bin/gettoken"
	SendMessageUrl = "https://qyapi.weixin.qq.com/cgi-bin/message/send"
)

type MessageContent struct {
	Content string `json:"content"`
}

type NoticeMessage struct {
	ToUser		string			`json:"touser"`
	AgentId		int				`json:"agentid"`
	MessageType	string			`json:"msgtype"`
	Message 	MessageContent	`json:"text"`
}

func GetAccessToken(id string,secret string) string {
	RequestUrl := strings.Join([]string{AccessTokenUrl,"?corpid=",id,"&corpsecret=",secret},"")
	Request,err := http.Get(RequestUrl)
	if err != nil {
		//
	}
	defer Request.Body.Close()
	ResponseBody,err := ioutil.ReadAll(Request.Body)
	if err != nil {
		//
	}
	Response := strings.Split(string(ResponseBody),"\"")
	AccessToken := Response[3]
	return AccessToken
}

func SendMsg(token string,user string,id int,message string){
	NoticeMsg := &NoticeMessage{
		ToUser:			user,
		AgentId:		id,
		MessageType:	"text",
		Message:		MessageContent{Content:	message},
	}
	datas,err := json.MarshalIndent(NoticeMsg," "," ")
	if err != nil {
		//return err
	}
	PostRequeste,err := http.NewRequest("POST",strings.Join([]string{SendMessageUrl,"?access_token=",token},""),bytes.NewReader(datas))
	if err != nil {
		//return err
	}
	PostRequeste.Header.Set("Content-Type","application/json;encoding=utf-8")
	Client := &http.Client{}
	PostResponse,err := Client.Do(PostRequeste)
	if err != nil {
		//return err
	}
	PostResponse.Body.Close()
}

func main(){
	CorpId := flag.String("cd","","Enterprise CorpId.")
	CorpSecret := flag.String("cs","","Certificate key for management group.")
	User := flag.String("u","@all","Sent to the specified user,if the value is null,it will send to all users.")
	AgentId := flag.Int("id",0,"Enterprise application ID.")
	Message := flag.String("msg","Everything looks good!","Message Content.")
	flag.Parse()
	AccessToken := GetAccessToken(*CorpId,*CorpSecret)
	SendMsg(AccessToken,*User,*AgentId,*Message)
}