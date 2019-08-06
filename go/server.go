package main

import (
	"fmt"
	"log"
	"net/http"
	"sensedata"
	graphql "github.com/graph-gophers/graphql-go"
        "github.com/graph-gophers/graphql-go/relay"
)

func main() {
	schema := graphql.MustParseSchema(sensedata.Schema, &sensedata.Resolver{})

	http.Handle("/graphql", &relay.Handler{Schema: schema})

	fmt.Println("[graphql] up and running! at http://localhost:4002")

	log.Fatal(http.ListenAndServe(":4002", nil))
}
