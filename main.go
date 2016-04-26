package main

import "fmt"
import "net/http"
import "io/ioutil"
import "encoding/json"

const Meetupurl = "http://api.meetup.com/%s/events"
var communities = []string{
  "PHP-MG",
  "GDG-BH"}

type Event struct {
  Id string `json:"id"`
  Name string `json:"name"`
  Link string `json:"link"`
  Description string `json:"description"`
  WaitlistCount int64 `json:"waitlist_count"`
  YesCount int64 `json:"yes_rsvp_count"`
  Limit int64 `json:"rsvp_limit"`
}

func getEvents(community string) *[]Event {
  var eventsUrl = fmt.Sprintf(Meetupurl, community)

  resp, err := http.Get(eventsUrl)
  if err != nil {
    panic(err.Error())
  }
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)

  var s = new([]Event)
  err = json.Unmarshal(body, &s)
  if(err != nil){
      fmt.Println("whoops:", err)
  }
  return s
}

func main() {
  var length, i int

  length = len(communities)
  for i = 0; i < length; i++ {
    fmt.Println(getEvents(communities[i]))
  }
}
