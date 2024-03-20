package vendors

import (
	"backend/config"
	"backend/mapper"
	"backend/fetcher"
)

func BarAPI(cfg config.File) fetcher.API {
	return fetcher.API {
		Hos: "bar.com",
		Pat: "foo",
		QryRul: mapper.Rules{
			{"", "", "token", cfg.Value("BAR_TOKEN"), ""},
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
	}
}
