
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Stop representing a bus stop
type Stop struct {
	RouteID     string `json:"routeID"`
	DirectionID int    `json:"directionID"`
	StopID      int    `json:"stopID"`
	TimePointID int    `json:"timePointID"`
}

// BusStopRequest payload
type BusStopRequest struct {
	Stops []Stop `json:"stops"`
}

// Crossings represents time data for a stop
type Crossings struct {
	Cancelled   bool   `json:"cancelled"`
	SchedTime   string `json:"schedTime"`
	SchedPeriod string `json:"schedPeriod"`
	PredTime    string `json:"predTime"`
	PredPeriod  string `json:"predPeriod"`
	Countdown   string `json:"countdown"`
	Destination string `json:"destination"`
}

// StopsResponse represents a stop in a response
type StopsResponse struct {
	DirectionID     int         `json:"directionID"`
	StopID          int         `json:"stopID"`
	TimePointID     int         `json:"timePointID"`
	SameDestination bool        `json:"sameDestination"`
	Crossings       []Crossings `json:"crossings"`
}

// RouteStops represents a list of bus stops
type RouteStops struct {
	RouteID int             `json:"routeID"`
	Stops   []StopsResponse `json:"stops"`
}

// BusStopResponseData represents the data from a response
type BusStopResponseData struct {
	ErrorMessage    string       `json:"errorMessage"`
	ShowArrivals    bool         `json:"showArrivals"`
	ShowStopNumber  bool         `json:"showStopNumber"`
	ShowScheduled   bool         `json:"showScheduled"`
	ShowDestination bool         `json:"showDestination"`
	UpdateTime      string       `json:"updateTime"`
	UpdatePeriod    string       `json:"updatePeriod"`
	RouteStops      []RouteStops `json:"routeStops"`
}

// BusStopResponse represents a response object
type BusStopResponse struct {
	Data BusStopResponseData `json:"d"`
}

func main() {
	payload := BusStopRequest{
		Stops: []Stop{
			Stop{
				RouteID:     "83",
				DirectionID: 17,
				StopID:      524,
				TimePointID: 0,
			},
		},
	}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(payload)
	res, _ := http.Post("http://webwatch.lavta.org/tmwebwatch/GoogleMap.aspx/getStopTimes", "application/json", b)
	body := BusStopResponse{}
	json.NewDecoder(res.Body).Decode(&body)
	stopTimes := body.Data.RouteStops[0].Stops[0].Crossings
	fmt.Println("Next stop times: ")
	for _, time := range stopTimes {
		fmt.Printf("%v%s, ", time.PredTime, time.PredPeriod)
	}
	fmt.Print("\n")
}