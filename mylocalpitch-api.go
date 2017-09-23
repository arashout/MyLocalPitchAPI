package main

import (
	"log"
	"net/http"
	"net/url"
	"time"
)

type Pitch struct {
	VenueID   string
	VenuePath string
	City      string
}

type MLPResponse struct {
	Meta struct {
		TotalItems int `json:"total_items"`
		Filter     struct {
			Starts string `json:"starts"`
			Ends   string `json:"ends"`
		} `json:"filter"`
	} `json:"meta"`
	Data MLPData `json:"data"`
}

type MLPData []struct {
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
}

const (
	APIEndpoint = "https://api-v2.mylocalpitch.com"
	BaseURL     = "https://www.mylocalpitch.com"
)

func GetPitchSlots(pitch Pitch, client *http.Client, starts time.Time, ends time.Time) MLPData {
	u, err := url.Parse(APIEndpoint + "/pitches/" + pitch.VenueID + "/slots")
	if err != nil {
		log.Fatal(err)
	}
	u.Scheme = "https"

	const layout = "2006-01-02"
	q := u.Query()
	q.Set("filter[starts]", starts.Format(layout))
	q.Set("filter[ends]", ends.Format(layout))
	u.RawQuery = q.Encode()

	// Add request headers
	req, err := http.NewRequest("GET", u.String(), nil)
	req.Host = APIEndpoint
	req.Header.Set("Origin", BaseURL)
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Accept-Language", "en-US,en;q=0.8")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.91 Safari/537.36")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Referer", BaseURL+"/"+pitch.City+"/venue/"+pitch.VenuePath)
	req.Header.Set("Connection", "keep-alive")

	response, err := client.Get(u.String())
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	mlpResponse := MLPResponse{}
	GetJSON(response, &mlpResponse)

	return mlpResponse.Data
}
