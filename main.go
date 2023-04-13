package main

import (
	"fmt"
	"os"
	"github.com/slack-go/slack"
)

func main() {

  os.Setenv("YOUR_TOKEN_HERE", "xoxe.xoxp-1-Mi0yLTUwODY3NDU2NTUzNTAtNTEwNTk3MzUyMjQ4MS01MTE0MTM2ODcwNjU3LTUwOTQ5MDkwNTU3ODItZjFmMzVmYjBlYjdkNjU1ZTRmOGJjZThjMDRmZmI1OGQ5NzVlYjI2MmExZjJiMjUwNzhhMGJiNWEyOWEyOWIxOQ")
  os.Setenv("CHANNEL_ID","C052C1V007R")

	api := slack.New(os.Getenv("YOUR_TOKEN_HERE"))
  channelid := os.Getenv("CHANNEL_ID")
  fileArr := []string{"rust.pdf"}


  attachment := slack.Attachment{
    Pretext: "Welcome",
    Text: "Hello, welcome to the channel!",
  }

  
  
  channelid, timestamp, err := api.PostMessage(
    "CHANNEL_ID", 
    slack.MsgOptionMeMessage(), 
    slack.MsgOptionAttachments(attachment), 
    slack.MsgOptionAsUser(true),
  )
 
  if err != nil{
    fmt.Printf("%s\n",err)
    return
   } 

  fmt.Printf("Message successfully sent to channel %v at %v\n", channelid, timestamp)


  for i := 0; i <len(fileArr); i++ {
    params := slack.FileUploadParameters{
      Channels: []string{channelid},
      File: fileArr[i],
    } 
   
      fileArr, err := api.UploadFile(params)
      if err != nil{
       fmt.Printf("%s\n",err)
       return
      } 
      fmt.Printf("Name: %s ,URL: %s\n", fileArr.Name, fileArr.URL)
  }
		
}