package web

type Options struct {
	Address      string `json:"address" toml:"address"`
	Port         int    `json:"port" toml:"port"`
	ReadTimeout  int    `json:"read_timeout" toml:"read_timeout"`
	WriteTimeout int    `json:"write_timeout" toml:"write_timeout"`
	BindRoutes   GinRouteFunc
}
