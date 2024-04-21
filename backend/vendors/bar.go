package vendors

import (
	"backend/getter"
)


type Bar struct {
	Get getter.Request
	Tok string
}


func BarAPI(cfg map[string]string) Bar {
	return Bar {
		Get: getter.Request{
			Hos: "bar.com",
			Pat: "foo",
			Map: getter.Mapping{
				{Arg: "token", Hed: "AuthToken"},
				{Hed: "Content-Type", Val: "Application/json"},
				{Arg: "latitude", Qry: "lat"},
				{Arg: "longitude", Qry: "long"},
				{Qry: "range", Val: "100"},
				{Arg: "distance", Equ: "0", Qry: "range", Val: "5"},
				{Arg: "distance", Equ: "1", Qry: "range", Val: "10"},
				{Arg: "distance", Equ: "2", Qry: "range", Val: "20"},
				{Qry: "money", Val: "$$$$"},
				{Arg: "expense", Equ: "0", Qry: "money", Val: "$"},
				{Arg: "expense", Equ: "1", Qry: "money", Val: "$$"},
				{Arg: "expense", Equ: "2", Qry: "money", Val: "$$$"},
			},
		},
		Tok: cfg["BAR_TOKEN"],
	}
}


func (b Bar) Fetch(arg map[string]string) []map[string]string {
	arg["token"] = b.Tok
	b.Get.Fetch(arg)
	return []map[string]string{
		map[string]string{
			"Name": "SugarHouse Pub",
			"Address": "1994 1100 E, Salt Lake City, UT 84106",
			"Hours": "Monday, Closed\n" +
				"Tuesday, 5 PM–2 AM\n" +
				"Wednesday, 5 PM–2 AM\n" +
				"Thursday, 5 PM–2 AM\n" +
				"Friday, 5 PM–2 AM\n" +
				"Saturday, 5 PM–2 AM\n" +
				"Saturday, Closed",
			"Phone": "+18014132857",
			"Website": "https://www.facebook.com/sugarhousepub",
			"Rating": "****",
			"Price": "$",
			"Description": "A little gem.",
		},
		map[string]string{
			"Name": "Piper Down Pub",
			"Address": "1492 S State St, Salt Lake City, UT 84115",
			"Hours": "Monday, 11 AM–1 AM\n" +
				"Tuesday, 11 AM–1 AM\n" +
				"Wednesday, 11 AM–1 AM\n" +
				"Thursday, 11 AM–1 AM\n" +
				"Friday, 11 AM–1 AM\n" +
				"Saturday, 10 AM–1 AM\n" +
				"Saturday, 10 AM–1 AM",
			"Phone": "+18014681492",
			"Website": "http://www.piperdownpub.com/",
			"Rating": "",
			"Price": "",
			"Description": "Located in the Ballpark Neighborhood, " +
				"Piper Down Olde World Pub is a taste of Ireland " +
				"without the price of a plane ticket. Join us for " +
				"weekly events, incredible local live music and a " +
				"couple of pints.",
		},
		map[string]string{
			"Name": "Alibi Bar & Place",
			"Address": "369 S Main St, Salt Lake City, UT 84111",
			"Hours": "Monday, Closed\n" +
				"Tuesday, 4 PM–2 AM\n" +
				"Wednesday, 4 PM–2 AM\n" +
				"Thursday, 4 PM–2 AM\n" +
				"Friday, 4 PM–2 AM\n" +
				"Saturday, 4 PM–2 AM\n" +
				"Saturday, 4 PM–2 AM",
			"Phone": "+13852590616",
			"Website": "https://instagram.com/alibislc?utm_medium=copy_link",
			"Rating": "*****",
			"Price": "$$",
			"Description": "Funky venue decorated with plants & art, " +
					"serving a selection of cocktails & light " +
					"bar meals.",
		},
	}
}
