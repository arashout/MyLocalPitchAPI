package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"time"
)

type Pitch struct {
	VenueID   string
	VenuePath string
}

type ResponseData struct {
	Meta struct {
		TotalItems int `json:"total_items"`
		Filter     struct {
			Starts string `json:"starts"`
			Ends   string `json:"ends"`
		} `json:"filter"`
	} `json:"meta"`
	Slots []struct {
		Type       string `json:"type"`
		ID         string `json:"id"`
		Attributes struct {
			Starts         time.Time `json:"starts"`
			Ends           time.Time `json:"ends"`
			Price          string    `json:"price"`
			AdminFee       string    `json:"admin_fee"`
			Currency       string    `json:"currency"`
			Availabilities int       `json:"availabilities"`
		} `json:"attributes"`
	} `json:"data"`
}

const (
	url_api = "https://api-v2.mylocalpitch.com"
	city    = "london"
)

func getJson(r *http.Response, target interface{}) error {
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}

func GetPitchSlots(pitch Pitch, client *http.Client) {
	u, err := url.Parse("https://api-v2.mylocalpitch.com/pitches/" + pitch.VenueID + "/slots")
	if err != nil {
		log.Fatal(err)
	}
	u.Scheme = "https"

	q := u.Query()
	q.Set("filter[starts]", "2017-09-20")
	q.Set("filter[ends]", "2017-09-26")
	u.RawQuery = q.Encode()
	log.Println(u)
	log.Println(u.String())

	req, err := http.NewRequest("GET", u.String(), nil)
	req.Host = "api-v2.mylocalpitch.com"
	req.Header.Set("Origin", "https://www.mylocalpitch.com")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Accept-Language", "en-US,en;q=0.8")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.91 Safari/537.36")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Referer", "https://www.mylocalpitch.com/"+city+"/venue/"+pitch.VenuePath)
	req.Header.Set("Connection", "keep-alive")

	response, err := client.Get(u.String())
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	respData := ResponseData{}
	getJson(response, &respData)
	log.Print("succes")
	log.Println(respData.Slots[0].Type)
}

func main() {
	p := Pitch{
		VenueID:   "34933",
		VenuePath: "three-corners-adventure-playground/football-5-a-side-34933",
	}
	GetPitchSlots(p, &http.Client{})

}
