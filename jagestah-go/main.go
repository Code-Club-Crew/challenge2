// User provides film number via CLI or Env Variable
// Pull that film number from SWAPI
// Pull all Ships in that Film
// Pull all Pilots of those Ships
// Print results in JSON
// Need to Define a struct that will be converted to JSON at the end
//functions for getting relevant film details
//functions for getting relevant ship details
//functions for getting relevant pilot details
package main

import (
	"fmt"
  "encoding/json"
  "os"
  // "strconv"
  "net/http"
  "io/ioutil"

)

var filmDeets map[string][]string

func main() {
  filmOutput := getFilmDeets(os.Args[1])
  for _, value := range filmOutput["starships"] {
    fmt.Println(value)
    fmt.Println(swapiRequest(value))
  }
}

func getFilmDeets(filmID string) (filmDeets map[string][]string){
  filmUrl := fmt.Sprintf("https://swapi.co/api/films/%s/", filmID)
  fmt.Println("Gathering details from:", filmUrl)
  // fmt.Println(filmUrl)
  filmDeets = swapiRequest(filmUrl)
  // httpReq, _ := http.NewRequest("GET", filmUrl, nil)
  // res, _ := http.DefaultClient.Do(httpReq)
  // body, _ := ioutil.ReadAll(res.Body)
  // json.Unmarshal(body, &filmDeets)
  // fmt.Println(filmDeets)
  return filmDeets
}
func swapiRequest(requestUrl string) (requestResults map[string][]string) {
  httpReq, err := http.NewRequest("GET", requestUrl, nil)
  errorHandle(err)
  res, err := http.DefaultClient.Do(httpReq)
  errorHandle(err)
  body, err := ioutil.ReadAll(res.Body)
  errorHandle(err)
  json.Unmarshal(body, &requestResults)
  return requestResults
}

func getShipDeets(shipUrl string) (shipDeets map[string]string) {
  return
}


func errorHandle (err error) {
  if err != nil {
    fmt.Println(err)
  }
}
