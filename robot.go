package main

import (
    "context"
    "io/ioutil"
    "log"
    "os"
    "strings"
    "time"

    "github.com/tencent-connect/botgo"
    "github.com/tencent-connect/botgo/dto"
    "github.com/tencent-connect/botgo/openapi"
    "github.com/tencent-connect/botgo/token"
    "github.com/tencent-connect/botgo/websocket"
    "github.com/tencent-connect/botgo/event"
    yaml "gopkg.in/yaml.v2"
)


type Config struct {
    AppID uint64 `yaml:"appid"` 
    Token string `yaml:"token"` 
}

var config Config
var api openapi.OpenAPI
var ctx context.Context


func init() {
    content, err := ioutil.ReadFile("config.yaml")
    if err != nil {
        log.Println("读取配置文件出错， err = ", err)
        os.Exit(1)
    }

    err = yaml.Unmarshal(content, &config)
    if err != nil {
        log.Println("解析配置文件出错， err = ", err)
        os.Exit(1)
    }
    log.Println(config)
}


func atMessageEventHandler(event *dto.WSPayload, data *dto.WSATMessageData) error {
    if strings.HasSuffix(data.Content, "你好") { 
        api.PostMessage(ctx, data.ChannelID, &dto.MessageToCreate{MsgID: data.ID, Content: "你好"})
    } else if strings.HasSuffix(data.Content, "顺") { 
        api.PostMessage(ctx, data.ChannelID, &dto.MessageToCreate{MsgID: data.ID, Content: "顺风顺水"})
    } else if strings.HasSuffix(data.Content, "水") {
        api.PostMessage(ctx, data.ChannelID, &dto.MessageToCreate{MsgID: data.ID, Content: "水到渠成"})
    } else if strings.HasSuffix(data.Content, "成") {
        api.PostMessage(ctx, data.ChannelID, &dto.MessageToCreate{MsgID: data.ID, Content: "成人之美"})
    } else if strings.HasSuffix(data.Content, "足") {
        api.PostMessage(ctx, data.ChannelID, &dto.MessageToCreate{MsgID: data.ID, Content: "足不出户"})
    }else {
        api.PostMessage(ctx, data.ChannelID, &dto.MessageToCreate{MsgID: data.ID, Content: "对不起，我还在学习"})
    }
    return nil
}

func main() {
    
    token := token.BotToken(config.AppID, config.Token) 
    
    api = botgo.NewOpenAPI(token).WithTimeout(3 * time.Second)
    
    ctx = context.Background()
    
    ws, err := api.WS(ctx, nil, "") 
    if err != nil {
        log.Fatalln("websocket错误， err = ", err)
        os.Exit(1)
    }

    var atMessage event.ATMessageEventHandler = atMessageEventHandler

    intent := websocket.RegisterHandlers(atMessage)     
    botgo.NewSessionManager().Start(ws, token, &intent) 
}