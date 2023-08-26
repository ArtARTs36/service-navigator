package config

type Frontend struct {
	AppName string `yaml:"app_name"`
	Navbar  struct {
		Links []struct {
			Title string `yaml:"title"`
			Icon  string `yaml:"icon"`
			URL   string `yaml:"url"`
		} `yaml:"links"`
		Profile struct {
			Links []struct {
				Title string `yaml:"title"`
				Icon  string `yaml:"icon"`
				URL   string `yaml:"url"`
			} `yaml:"links"`
		} `yaml:"profile"`
		Search struct {
			Providers []SearchProvider `yaml:"providers"`
		} `yaml:"search"`
	} `yaml:"navbar"`
}
