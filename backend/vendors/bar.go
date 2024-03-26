package vendors

import (
	"backend/getter"
	"backend/fetcher"
)

func BarAPI(cfg map[string]string) fetcher.API {
	return fetcher.API {
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
