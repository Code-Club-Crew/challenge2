// User provides film number via CLI arg
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
  "net/http"
  "io/ioutil"
)

type filmStruct struct {
	Title string
	Episode_id int
	Opening_crawl string
	Director string
	Producer string
	Release_date string
	Characters []string
	Planets []string
	Starships []string
	Vehicles []string
	Species []string
	Created string
	Edited string
	Url string
}

type shipStruct struct {
	Name string
	Model string
	Manufacturer string
	Cost_in_credits string
	Length string
	Max_atmosphering_speed string
	Crew string
	Passengers string
	Cargo_capacity string
	Consumables string
	Hyperdrive_rating string
	Mglt string
	Starship_class string
	Pilots []string
	Films []string
	Created string
	Edited string
	Url string
}

func main() {
	var filmOutput filmStruct
  filmOutput = getFilmDeets(os.Args[1])
	fmt.Println("Film Title:",filmOutput.Title)
  for key, value := range filmOutput.Starships {
    shipOutput := getShipDeets(value)
		shipKey := (key + 1)
		// fmt.Println(shipKey, value)
		fmt.Printf("Ship number %d is %s\n", shipKey, shipOutput.Name)
		// fmt.Println(s)
  }
}

func getFilmDeets(filmID string) (filmDeets filmStruct){
  filmUrl := fmt.Sprintf("https://swapi.co/api/films/%s/", filmID)
  filmRequest := swapiRequest(filmUrl)
	// fmt.Println(filmRequest)
	json.Unmarshal(filmRequest, &filmDeets)
	// fmt.Println(filmDeets)
  return filmDeets
}

func getShipDeets(shipUrl string) (shipDeets shipStruct) {
	shipRequest := swapiRequest(shipUrl)
	json.Unmarshal(shipRequest, &shipDeets)
  return shipDeets
}

func swapiRequest(requestUrl string) (requestResults []uint8) {
	fmt.Println("Gathering details from:", requestUrl)
  httpReq, err := http.NewRequest("GET", requestUrl, nil)
  errorHandle(err)
  res, err := http.DefaultClient.Do(httpReq)
  errorHandle(err)
  requestResults, err = ioutil.ReadAll(res.Body)
  errorHandle(err)
  return requestResults
}

func errorHandle (err error) {
  if err != nil {
    fmt.Println(err)
  }
}
