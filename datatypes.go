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

// Configuration struct
type Configuration struct {
	Address      string `yaml:"address"`
	TimeInterval int    `yaml:"timeinterval-to-pull"`
	Logging      struct {
		Debug bool `yaml:"debug"`
	} `yaml:"logging"`
}
