package presentation

import (
	"github.com/tyler-sommer/stick"
	"github.com/tyler-sommer/stick/twig"
	"io"
)

type Renderer struct {
	twigRenderer *stick.Env
	globVars     map[string]stick.Value
}

func NewRenderer(templateDir string, globVars map[string]stick.Value) *Renderer {
	loader := stick.NewFilesystemLoader(templateDir)
	renderer := twig.New(loader)

	return &Renderer{
		twigRenderer: renderer,
		globVars:     globVars,
	}
}

func (r *Renderer) Render(templatePath string, writer io.Writer, vars map[string]stick.Value) error {
	for k, val := range r.globVars {
		vars[k] = val
	}

	return r.twigRenderer.Execute(templatePath, writer, vars)
}
