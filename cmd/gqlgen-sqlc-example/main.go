package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/shilohchurch/shilohchurch/dataloaders" // update the username
	"github.com/shilohchurch/shilohchurch/gqlgen"      // update the username
	"github.com/shilohchurch/shilohchurch/pg"          // update the username
)

func main() {
	// initialize the db
	db, err := pg.Open("dbname=cducbdmcfmvbai:eb7f377b251973723caf4075b463c928f93a5d61296832a4ed7f69cdba27fc02@ec2-44-205-112-253.compute-1.amazonaws.com:5432/dutvg8v4vddc4 sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// initialize the repository
	repo := pg.NewRepository(db)

	// initialize the dataloaders
	dl := dataloaders.NewRetriever() // <- here we initialize the dataloader.Retriever

	// configure the server
	mux := http.NewServeMux()
	mux.Handle("/", gqlgen.NewPlaygroundHandler("/query"))
	dlMiddleware := dataloaders.Middleware(repo)     // <- here we initialize the middleware
	queryHandler := gqlgen.NewHandler(repo, dl)      // <- use dataloader.Retriever here
	mux.Handle("/query", dlMiddleware(queryHandler)) // <- use dataloader.Middleware here

	// run the server
	port := ":8080"
	fmt.Fprintf(os.Stdout, "🚀 Server ready at http://localhost%s\n", port)
	fmt.Fprintln(os.Stderr, http.ListenAndServe(port, mux))
}
