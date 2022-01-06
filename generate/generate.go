package generate

import (
	"fmt"

	"github.com/99designs/gqlgen/api"
	"github.com/99designs/gqlgen/codegen/config"
	"github.com/99designs/gqlgen/plugin/modelgen"
)

func addTag(b *modelgen.ModelBuild) *modelgen.ModelBuild {
	for _, model := range b.Models {
		for _, field := range model.Fields {
			field.Tag += fmt.Sprintf(` db:%s`, field.Name)
		}
	}

	return b
}

func main() {
	p := modelgen.Plugin{
		MutateHook: addTag,
	}

	cfg, _ := config.LoadConfigFromDefaultLocations()
	api.Generate(cfg, api.NoPlugins(), api.AddPlugin(&p))
}
