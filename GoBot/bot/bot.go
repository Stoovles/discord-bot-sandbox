package bot

import (
  "../config"
  "fmt"
  "github.com/bwmarrin/discordgo"
  "strings"
  "io/ioutil"
  "net/http"
  "encoding/json"
)

var BotID string
var goBot *discordgo.Session

func Start() {
  goBot, err := discordgo.New("Bot " + config.Token)

  if err != nil {
    fmt.Println(err.Error())
    return
  }

  u, err := goBot.User("@me")

  if err != nil {
    fmt.Println(err.Error())

  }

  BotID = u.ID

  goBot.AddHandler(messageHandler)
  // goBot.AddHandler(weatherChecker)

  goBot.Open()

  if err != nil {
    fmt.Println(err.Error())
    return
  }

  fmt.Println("Bot is running!")


}

func weatherChecker() (string) {
  response, err := http.Get("http://api.openweathermap.org/data/2.5/forecast?zip=80202&appid=e703a5585d75c3f03cf5aa9b1d564cdb")

  if err != nil {
    fmt.Printf("The HTTP request failed with error %s\n", err)
  } else {
    weatherResponseData, _ := ioutil.ReadAll(response.Body)

    fmt.Println(string(weatherResponseData))

    var weatherResponseObject WeatherResponse
    json.Unmarshal(weatherResponseData, &weatherResponseObject)

    fmt.Println(weatherResponseObject)

    return weatherResponseObject.City.Name

  }
  return ""
}

// func contingentWeatherMessage() {
//
//
// }

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

  if strings.HasPrefix(m.Content, config.BotPrefix) {

    if m.Author.ID == BotID {
      return
    }

    if m.Content == "!ping" { // !ping
      privateChannel, _ := s.UserChannelCreate("189883688937324544")
      _, _ = s.ChannelMessageSend(privateChannel.ID, "pong")
    }

    if m.Content == "!weather" {
      _, _ = s.ChannelMessageSend(m.ChannelID, weatherChecker()) //goal is to send this information to all necessary ID's
    }


  }

}
