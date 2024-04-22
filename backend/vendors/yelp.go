package vendors

import (
	"fmt"
	"backend/getter"
)


type Yelp struct {
	Get getter.Request
	Key string
}


func YelpAPI(cfg map[string]string) Yelp {
	return Yelp {
		Get: getter.Request{
			Hos: "https://api.yelp.com",
			Pat: "v3/businesses/search",
			Map: getter.Mapping{
				{Hed: "Authorization", Val: "Bearer"},
				{Arg: "key", Hed: "Authorization", App: " "},
				{Hed: "accept", Val: "application/json"},
				{Qry: "sort_by", Val: "best_match"},
				{Qry: "limit", Val: "4"},
				{Qry: "categories", Val: "Active Life,Food"},
				{Qry: "open_now", Val: "true"},
				{Arg: "latitude", Qry: "latitude"},
				{Arg: "longitude", Qry: "longitude"},
				{Arg: "distance", Equ: "0", Qry: "radius", Val: "4000"},
				{Arg: "distance", Equ: "1", Qry: "radius", Val: "8000"},
				{Arg: "distance", Equ: "2", Qry: "radius", Val: "16000"},
				{Arg: "distance", Equ: "3", Qry: "radius", Val: "32000"},
				{Arg: "expense", Equ: "0", Qry: "price", Val: "1"},
				{Arg: "expense", Equ: "1", Qry: "price", Val: "1,2"},
				{Arg: "expense", Equ: "2", Qry: "price", Val: "2,3"},
				{Arg: "expense", Equ: "3", Qry: "price", Val: "3,4"},
			},
		},
		Key: cfg["YELP_API_KEY"],
	}
}


func (y Yelp) Fetch(arg map[string]string) []map[string]string {
	arg["key"] = y.Key
	res := y.Get.Fetch(arg)
	fmt.Println(string(res))
	return []map[string]string{
		map[string]string{
			"Name": "SugarHouse Pub",
		},
		map[string]string{
			"Name": "Piper Down Pub",
		},
		map[string]string{
			"Name": "Alibi Bar & Place",
		},
	}
}
