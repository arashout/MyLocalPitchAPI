package mlpapi

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	p := Pitch{
		VenueID:   "34933",
		VenuePath: "three-corners-adventure-playground/football-5-a-side-34933",
		City:      "london",
	}
	t1 := time.Now()
	t2 := t1.AddDate(0, 0, 14) //Two weeks from now
	slots := GetPitchSlots(p, &http.Client{}, t1, t2)
	for _, slot := range slots {
		fmt.Println(slot)
	}
	fmt.Println(GetSlotCheckoutLink(slots[0], p))

}
