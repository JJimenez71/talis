package vendors

import (
	"backend/fetcher"
	"backend/getter"
)


type Foo struct {
	Get getter.Request
}


func FooAPI(cfg map[string]string) Foo {
	return Foo {
		Get: getter.Request {
			Hos: "foo.com",
			Pat: "bar",
			Rul: getter.Rules{
	                        {"", "", "apiKey", cfg["FOO_API_KEY"], ""},
	                        {"latitude", "", "coords", "", ","},
	                        {"longitude", "", "coords", "", ","},
	                        {"", "", "range", "100", ""},
	                        {"distance", "0", "range", "5", ""},
	                        {"distance", "1", "range", "10", ""},
	                        {"distance", "2", "range", "20", ""},
	                        {"", "", "minCost", "$", ""},
	                        {"", "", "maxCost", "$$$", ""},
	                        {"expense", "0", "minCost", "$", ""},
	                        {"expense", "0", "maxCost", "$$", ""},
	                        {"expense", "1", "minCost", "$$", ""},
	                        {"expense", "1", "maxCost", "$$$", ""},
	                        {"expense", "2", "minCost", "$$$", ""},
	                        {"expense", "2", "maxCost", "$$$$", ""},
	                },
		},
        }
}


func (f Foo) Fetch(arg map[string]string) fetcher.Activities {
	f.Get.Fetch(arg)
	return fetcher.Activities {
	}
}

