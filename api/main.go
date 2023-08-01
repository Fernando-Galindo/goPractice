package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Geo struct {
	Ip               string  `json:"ip"`
	IspName          string  `json:"ispName"`
	Continent        string  `json:"continent"`
	Zipcode          string  `json:"zipcode"`
	ConnectionType   string  `json:"connectionType"`
	OrganizationType string  `json:"organizationip"`
	Latitude         float64 `json:"latitude"`
	Longitude        float64 `json:"longitude"`
	IsPostalEst      bool    `json:"isPostalCodeEstimate"`
	Region           string  `json:"region"`
	Cennsu           string  `json:"censusRegion"`
	Country          string  `json:"country"`
	District         string  `json:"district"`
	City             string  `json:"city"`
}
type TipResp struct {
	Geo    Geo    `json:"geo"`
	Demo   string `json:"demographics"`
	Errors string `json:"errors"`
	Ua     string `json:"ua"`
}

func main() {
	base := "https://tip-nonprod.rvohealth.com/"
	path := "api/v1/lookup"
	params := url.Values{}
	params.Add("ip", "65.184.194.43")

	urlCompiled, _ := url.ParseRequestURI(base)
	urlCompiled.Path = path
	urlCompiled.RawQuery = params.Encode()

	req, _ := http.NewRequest("GET", urlCompiled.String(), nil)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	req.Header.Add("token", "f2096417-b07d-47e3-8c31-8103872dd792")
	client := &http.Client{}

	resp, _ := client.Do(req)
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	var tip TipResp
	err := json.Unmarshal(body, &tip)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println(resp.Status)
	fmt.Printf("resp.body type: %T\nbody type: %T\nbody value: %v\n", resp.Body, body, string(body))
	fmt.Printf("tip type: %T\nbody value: %v\n", tip, tip)
	fmt.Println(tip.Geo.IspName)

}
