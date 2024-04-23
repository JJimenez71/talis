package vendors

import (
	"fmt"
	"strings"
	"encoding/json"
	"backend/getter"
)


type Yelp struct {
	Get getter.Request
	Key string
}


func YelpAPI(cfg map[string]string) Yelp {
	return Yelp {
		Key: cfg["YELP_API_KEY"],
		Get: getter.Request{
			Hos: "https://api.yelp.com",
			Pat: "v3/businesses/search",
			Map: getter.Mapping{
				// form headers using arg values
				{Hed: "Authorization", Val: "Bearer"},
				{Hed: "Authorization", App: " ", Arg: "key"},
				// hardcoded headers
				{Hed: "accept", Val: "application/json"},
				// hardcoded query parameters
				{Qry: "sort_by", Val: "best_match"},
				{Qry: "limit", Val: "20"},
				{Qry: "categories", Val: "Active Life,Food"},
				{Qry: "open_now", Val: "true"},
				// form query parameters using arg values
				{Qry: "latitude", Arg: "latitude"},
				{Qry: "longitude", Arg: "longitude"},
				// form query parameters by mapping arg values
				{Qry: "radius", Val: "4000", Arg: "distance", Equ: "0"},
				{Qry: "radius", Val: "8000", Arg: "distance", Equ: "1"},
				{Qry: "radius", Val: "16000", Arg: "distance", Equ: "2"},
				{Qry: "radius", Val: "32000", Arg: "distance", Equ: "3"},
				{Qry: "price", Val: "1", Arg: "expense", Equ: "0"},
				{Qry: "price", Val: "2", Arg: "expense", Equ: "1"},
				{Qry: "price", Val: "3", Arg: "expense", Equ: "2"},
				{Qry: "price", Val: "4", Arg: "expense", Equ: "3"},
			},
		},
	}
}


func (y Yelp) Fetch(arg map[string]string) []map[string]string {
	arg["key"] = y.Key
	res := y.Get.Fetch(arg)
	if res == nil {
		return []map[string]string{}
	}
	body := struct{
		Businesses []struct{
			Alias string `json:"alias"`
			Categories []struct{
				Alias string `json:"alias"`
				Title string `json:"title"`
			} `json:"categories"`
			Coordinates struct{
				Latitude float64 `json:"latitude"`
				Longitude float64 `json:"longitude"`
			} `json:"coordinates"`
			DisplayPhone string `json:"display_phone"`
			Distance float64 `json:"display"`
			Id string `json:"id"`
			ImageUrl string `json:"image_url"`
			IsClosed bool `json:"is_closed"`
			Location struct{
				Address1 string `json:"address1"`
				Address2 string `json:"address2"`
				Address3 string `json:"address3"`
				City string `json:"city"`
				Country string `json:"country"`
				DisplayAddress []string `json:"display_address"`
				State string `json:"state"`
				ZipCode string `json:"zip_code"`
			} `json:"location"`
			Name string `json:"name"`
			Phone string `json:"phone"`
			Price string `json:"price"`
			Rating float64 `json:"rating"`
			ReviewCount int `json:"review_count"`
			Transactions []string  `json:"transactions"`
			Url string `json:"url"`
		} `json:"businesses"`
		Region struct{
			Center struct{
				Latitude float64 `json:"latitude"`
				Longitude float64 `json:"longitude"`
			} `json:"center"`
		} `json:"region"`
		Total int `json:"total"`
	}{}
	if err := json.Unmarshal(res, &body); err != nil {
		fmt.Println("Error:", err)
		return []map[string]string{}
	}
	ret := []map[string]string{}
	for _, b := range body.Businesses {
		r := map[string]string{
			"Name": b.Name,
			"Address": strings.Join(b.Location.DisplayAddress, "\n"),
			"phone": b.DisplayPhone,
			"website": b.Url,
			"rating": fmt.Sprintf("%f", b.Rating),
			"image": b.ImageUrl,
		}
		ret = append(ret, r)
	}
	return ret
}

