package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"strconv"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)
type hostInfo struct {
  Hostname string 
  Date string 
}

func main() {
	err := godotenv.Load() 
	if err != nil {
		log.Println(err)
	}
  	botToken := os.Getenv("TOKEN_ID")
  	tgChatId := os.Getenv("CHAT_ID")
  	
	Info := setInfo()
  
  	chatID, err := strconv.ParseInt(tgChatId,10,64)
  	if err != nil {
    	log.Fatal(err)
  	}

  if checkId() {
    send := autoUpdate()
    sendTelegram(send, Info, botToken, chatID)
  } else {
    return 

  } 
}

func setInfo() hostInfo{
  hostnameInfo, err := os.Hostname()
  if err != nil {
    log.Fatal(err)
  }
  
  t := time.Now()
  
  host := hostInfo {
    Hostname: hostnameInfo,
    Date: t.Format("2006-01-02 15:04:05"),
  }
  return host
}


func checkId() bool {
  currentUser, err := user.Current()
  if err != nil { 
    log.Fatal(err) 
  }

  uid := currentUser.Uid
  if uid == "0" {
    return true 
  } else {
    return false
  }
}

func autoUpdate() string {
  cmd := exec.Command("apt-get", "update", "-y")
  stdout, err := cmd.CombinedOutput()
  if err != nil {
    log.Fatal(err)
  }
  aptResult := string(stdout)
  return aptResult
}

func sendTelegram(data string, h hostInfo, botToken string, tgChatId int64) {
  bot, err := tgbotapi.NewBotAPI(botToken)
  if err != nil {
    log.Fatal(err)
    return 
  }
  msg := fmt.Sprintf("Actualización de: %s a las %s\n%s",h.Hostname , h.Date , data) 
  sendMsg := tgbotapi.NewMessage(tgChatId, msg)
  bot.Send(sendMsg)
}
