package main

import (
	"context"
	"fmt"

	"github.com/google/go-github/v32/github"
)

func main() {
	client := github.NewClient(nil)
	user := "drone-plugins"
	opt := &github.RepositoryListOptions{Type: "owner", ListOptions: github.ListOptions{PerPage: 100}}

	repos, _, err := client.Repositories.List(context.Background(), user, opt)
	if err != nil {
		fmt.Println(err)
	}

	for _, repo := range repos {
		var license string
		if repo.License != nil && repo.License.Key != nil {
			license = *repo.License.Key
		}
		fmt.Printf("%s,%s,%s\n", *repo.FullName, *repo.HTMLURL, license)
	}
}
