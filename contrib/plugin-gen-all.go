// +build ignore

// This program runs the template generator for all plugins.

package main

import (
	"flag"
	"log"
	"os"
	"os/exec"

	"code.google.com/p/goauth2/oauth"
	"github.com/google/go-github/github"
)

var (
	token = flag.String("token", os.Getenv("GITHUB_API_TOKEN"), "github api token")
	owner = flag.String("owner", "drone-plugins", "github account")
)

func main() {
	flag.Parse()

	t := &oauth.Transport{
		Token: &oauth.Token{AccessToken: *token},
	}
	client := github.NewClient(t.Client())

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
