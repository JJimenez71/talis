package vendors

import (
	"fmt"
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
	fmt.Println(b.Get.Fetch(arg))
	return fetcher.Activities {
		fetcher.Activity {
			Name: "Bar A",
		},
		fetcher.Activity {
			Name: "Bar B",
		},
		fetcher.Activity {
			Name: "Bar C",
		},
	}
}
