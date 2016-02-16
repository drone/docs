// +build ignore

// This program runs the template generator for all plugins.

package main

import (
	"flag"
	"log"
	"os"
	"os/exec"

	"golang.org/x/oauth2"
	"github.com/google/go-github/github"
)

var (
	token = flag.String("token", os.Getenv("GITHUB_API_TOKEN"), "github api token")
	owner = flag.String("owner", "drone-plugins", "github account")
)

func main() {
	flag.Parse()

	tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: *token})
	httpClient := oauth2.NewClient(oauth2.NoContext, tokenSource)
	client := github.NewClient(httpClient)

	opts := &github.RepositoryListByOrgOptions{}
	opts.PerPage = 100
	repos, _, err := client.Repositories.ListByOrg(*owner, opts)
	if err != nil {
		log.Fatal(err)
	}

	for _, repo := range repos {
		log.Printf("generating docs for %s", *repo.Name)

		cmd := exec.Command("go", "run", "contrib/plugin-gen.go", "--repo", *repo.Name)
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
