package main

import (
	"fmt"
	"gql-go-todo/dataloader"
	"gql-go-todo/graph"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/api"
	"github.com/99designs/gqlgen/codegen/config"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/99designs/gqlgen/plugin/modelgen"
	"github.com/samonzeweb/godb"
	"github.com/samonzeweb/godb/adapters/sqlite"
)

const defaultPort = "8080"

func addTag(b *modelgen.ModelBuild) *modelgen.ModelBuild {
	for _, model := range b.Models {
		for _, field := range model.Fields {
			field.Tag += fmt.Sprintf(` db:"%s"`, field.Name)
		}
	}

	return b
}

func main() {
	// If generate
	if len(os.Args) > 1 {
		if os.Args[1] == "gen" || os.Args[1] == "generate" {
			p := modelgen.Plugin{
				MutateHook: addTag,
			}

			cfg, _ := config.LoadConfigFromDefaultLocations()
			api.Generate(cfg, api.NoPlugins(), api.AddPlugin(&p))
			return
		}
	}
	// Otherwise server

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db, _ := godb.Open(sqlite.Adapter, "db.db")
	srv := handler.NewDefaultServer(graph.NewSchema(db))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	handler := dataloader.Middleware(db, srv)
	http.Handle("/query", handler)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
