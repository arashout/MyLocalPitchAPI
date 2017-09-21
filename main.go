package main

import (
    "net/http"
    "net/url"
    "log"
    "bufio"
    "bytes"
)

type Pitch struct{
    VenueID string
    VenuePath string
}

const(
    url_api = "https://api-v2.mylocalpitch.com"
    city = "london"
)

func GetPitchSlots(pitch Pitch, client *http.Client){
    u, err := url.Parse("https://api-v2.mylocalpitch.com/pitches/" + pitch.VenueID + "/slots")
    if err != nil {
        log.Fatal(err)
    }
    u.Scheme = "https"

    q := u.Query()
    q.Set("filter[starts]", "2017-09-20")
    q.Set("filter[ends]","2017-09-26")
    u.RawQuery = q.Encode()
    log.Println(u)
    log.Println(u.String())

    req, err := http.NewRequest("GET", u.String(), nil)
    req.Host = "api-v2.mylocalpitch.com"
    req.Header.Set("Origin", "https://www.mylocalpitch.com")
    req.Header.Set("Accept-Encoding", "gzip, deflate, br")
    req.Header.Set("Accept-Language", "en-US,en;q=0.8")
    req.Header.Set("User-Agent","Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.91 Safari/537.36")
    req.Header.Set("Accept","application/json, text/plain, */*")
    req.Header.Set("Referer", "https://www.mylocalpitch.com/" + city + "/venue/" + pitch.VenuePath)
    req.Header.Set("Connection", "keep-alive")

    response, err :=client.Get(u.String())
    if err != nil {
        log.Fatal(err)
    }
    defer response.Body.Close()

    scanner := bufio.NewScanner(response.Body)
    scanner.Split(bufio.ScanRunes)
    var buf bytes.Buffer
    for scanner.Scan() {
        buf.WriteString(scanner.Text())
    }
    log.Println(buf.String())
}

func main(){
    p := Pitch{
        VenueID:"34933",
        VenuePath:"three-corners-adventure-playground/football-5-a-side-34933",
    }
    GetPitchSlots(p, &http.Client{})

}
