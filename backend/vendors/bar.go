package vendors

import (
	"backend/fetcher"
	"backend/getter"
)


type Bar struct {
	Get getter.Request
}


func BarAPI(cfg map[string]string) Bar {
	return Bar {
		Get: getter.Request{
			Hos: "bar.com",
			Pat: "foo",
			Rul: getter.Rules{
				{"", "", "token", cfg["BAR_TOKEN"], ""},
				{"latitude", "", "lat", "", ""},
				{"longitude", "", "long", "", ""},
				{"", "", "range", "100", ""},
				{"distance", "0", "range", "5", ""},
				{"distance", "1", "range", "10", ""},
				{"distance", "2", "range", "20", ""},
				{"", "", "money", "$", ""},
				{"expense", "2", "money", "$$", ","},
				{"expense", "3", "money", "$$$", ","},
			},
		},
	}
}


func (b Bar) Fetch(arg map[string]string) fetcher.Activities {
	b.Get.Fetch(arg)
	return fetcher.Activities {
		fetcher.Activity {
			Name: "SugarHouse Pub",
			Address: "1994 1100 E, Salt Lake City, UT 84106",
			Hours: []string{"Monday, Closed",
				"Tuesday, 5 PM–2 AM",
				"Wednesday, 5 PM–2 AM",
				"Thursday, 5 PM–2 AM",
				"Friday, 5 PM–2 AM",
				"Saturday, 5 PM–2 AM",
				"Saturday, Closed",
			},
			Phone: "+18014132857",
			Website: "https://www.facebook.com/sugarhousepub",
			Rating: "****",
			Price: "$",
			Description: "A little gem.",
		},
		fetcher.Activity {
			Name: "Piper Down Pub",
			Address: "1492 S State St, Salt Lake City, UT 84115",
			Hours: []string{"Monday, 11 AM–1 AM",
				"Tuesday, 11 AM–1 AM",
				"Wednesday, 11 AM–1 AM",
				"Thursday, 11 AM–1 AM",
				"Friday, 11 AM–1 AM",
				"Saturday, 10 AM–1 AM",
				"Saturday, 10 AM–1 AM",
			},
			Phone: "+18014681492",
			Website: "http://www.piperdownpub.com/",
			Rating: "",
			Price: "",
			Description: "Located in the Ballpark Neighborhood, " +
				"Piper Down Olde World Pub is a taste of Ireland " +
				"without the price of a plane ticket. Join us for " +
				"weekly events, incredible local live music and a " +
				"couple of pints. 21+",
		},
		fetcher.Activity {
			Name: "Alibi Bar & Place",
			Address: "369 S Main St, Salt Lake City, UT 84111",
			Hours: []string{"Monday, Closed",
				"Tuesday, 4 PM–2 AM",
				"Wednesday, 4 PM–2 AM",
				"Thursday, 4 PM–2 AM",
				"Friday, 4 PM–2 AM",
				"Saturday, 4 PM–2 AM",
				"Saturday, 4 PM–2 AM",
			},
			Phone: "+13852590616",
			Website: "https://instagram.com/alibislc?utm_medium=copy_link",
			Rating: "*****",
			Price: "$$",
			Description: "Funky venue decorated with plants & art, " +
					"serving a selection of cocktails & light " +
					"bar meals.",
		},
	}
}
