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
			Links []ProfileLink `yaml:"links"`
		} `yaml:"profile"`
		Search struct {
			Providers []SearchProvider `yaml:"providers"`
		} `yaml:"search"`
	} `yaml:"navbar"`
	Pages struct {
		Images *ImagePage `yaml:"images"`
	} `yaml:"pages"`
}

type ImagePage struct {
	Counters struct {
		Pulls bool `yaml:"pulls"`
		Stars bool `yaml:"stars"`
	} `yaml:"counters"`
}

func (p *ImagePage) HasCounters() bool {
	return p.Counters.Pulls
}

type ProfileLink struct {
	Title string `yaml:"title"`
	Icon  string `yaml:"icon"`
	URL   string `yaml:"url"`
	Form  struct {
		Method string `yaml:"method"`
		Inputs []struct {
			Name  string `yaml:"name"`
			Value string `yaml:"value"`
		} `yaml:"inputs"`
	} `yaml:"form"`
}

func (l *ProfileLink) HasForm() bool {
	return l.Form.Inputs != nil && len(l.Form.Inputs) > 0
}
