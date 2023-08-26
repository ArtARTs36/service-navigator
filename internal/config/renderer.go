package config

import (
	"github.com/artarts36/service-navigator/internal/presentation/view"
	"github.com/tyler-sommer/stick"
)

func initRenderer(env *Environment, conf *Config) *view.Renderer {
	vars := map[string]stick.Value{}
	vars["_navBar"] = conf.Frontend.Navbar
	vars["_appName"] = conf.Frontend.AppName
	vars["_username"] = env.User

	return view.NewRenderer("/app/templates", vars)
}
