// +build ignore

// This program downloads the plugin documentation from the
// repository and generates a hugo template.

package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"golang.org/x/oauth2"
	"github.com/BurntSushi/toml"
	"github.com/google/go-github/github"
	"gopkg.in/yaml.v2"
)

var (
	token = flag.String("token", os.Getenv("GITHUB_API_TOKEN"), "github api token")
	owner = flag.String("owner", "drone-plugins", "github account")
	repo  = flag.String("repo", "drone-slack", "github repository")
)

type YAML struct {
	Plugin Plugin `yaml:"plugin"`
}

type Plugin struct {
	Title string   `toml:"title"       yaml:"name"`
	Desc  string   `toml:"description" yaml:"desc"`
	User  string   `toml:"user"        yaml:"-"`
	Repo  string   `toml:"repo"        yaml:"-"`
	Image string   `toml:"image"       yaml:"image"`
	Tags  []string `toml:"tags"        yaml:"labels"`
	Kind  string   `toml:"categories"  yaml:"type"`

	Draft  bool      `toml:"draft"  yaml:"-"`
	Date   time.Time `toml:"date"   yaml:"-"`
	Menu   string    `toml:"menu"   yaml:"-"`
	Weight int       `toml:"weight" yaml:"-"`
}

func main() {
	flag.Parse()

	tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: *token})
	httpClient := oauth2.NewClient(oauth2.NoContext, tokenSource)
	client := github.NewClient(httpClient)

	dst := fmt.Sprintf("static/logos/%s.svg", strings.Replace(*repo, "drone-", "", -1))
	err := download(client, "logo.svg", dst)
	if err != nil {
		log.Fatal(err)
	}

	raw, err := downloadStr(client, ".drone.yml")
	if err != nil {
		log.Fatal(err)
	}
	conf := YAML{}
	err = yaml.Unmarshal(raw, &conf)
	if err != nil {
		log.Fatal(err)
	}
	conf.Plugin.Date = time.Now()
	conf.Plugin.Weight = 1
	conf.Plugin.User = *owner
	conf.Plugin.Repo = *repo

	docs, err := downloadStr(client, "DOCS.md")
	if err != nil {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)
	if err := toml.NewEncoder(buf).Encode(conf.Plugin); err != nil {
		log.Fatal(err)
	}

	post := strings.Replace(*repo, "drone-", "", -1)
	post = strings.Replace(post, "-", "_", -1)
	post = fmt.Sprintf("content/plugins/%s.md", post)
	f, err := os.Create(post)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	fmt.Fprintf(f, tmpl, buf.String(), string(docs))
}

func download(client *github.Client, source, dest string) error {
	rc, err := client.Repositories.DownloadContents(*owner, *repo, source, nil)
	if err != nil {
		return err
	}
	defer rc.Close()

	f, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer f.Close()
	io.Copy(f, rc)
	return nil
}

func downloadStr(client *github.Client, source string) ([]byte, error) {
	rc, err := client.Repositories.DownloadContents(*owner, *repo, source, nil)
	if err != nil {
		return nil, err
	}
	defer rc.Close()

	var buf bytes.Buffer
	io.Copy(&buf, rc)

	return buf.Bytes(), nil
}

var tmpl = `+++

# THIS DOCUMENT IS AUTOMATICALLY GENERATED. PLEASE DO NOT EDIT

%s
+++

%s
`
