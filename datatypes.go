package main

type MAC struct {
	MAC     string
	Vendor  string
	OUI     string
	Typ     string
	Address string
}

type OUI struct {
	Vendor  string
	OUI     string
	Typ     string
	Address string
}

//
type Configuration struct {
	Address      string `yaml:"address"`
	TimeInterval int    `yaml:"timeinterval-to-pull"`
	Logging      struct {
		Dir   string `yaml:"log-dir"`
		Debug bool   `yaml:"debug"`
	}
}
