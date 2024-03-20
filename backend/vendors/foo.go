package vendors

import (
	"backend/config"
	"backend/mapper"
	"backend/fetcher"
)

func FooAPI(cfg config.File) fetcher.API {
	return fetcher.API {
		Hos: "foo.com",
		Pat: "bar",
		QryRul: mapper.Rules{
                        {"", "", "apiKey", cfg.Value("FOO_API_KEY"), ""},
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
        }
}

