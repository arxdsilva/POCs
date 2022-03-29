package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/google/go-github/v43/github"
	"golang.org/x/oauth2"
)

func main() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	workflows, rep, err := client.Actions.ListWorkflows(ctx,
		"arxdsilva", "golang-ifood-sdk", &github.ListOptions{})
	fmt.Println("err: ", err)
	fmt.Printf("rep: %+v\n", rep)
	for i, workflow := range workflows.Workflows {
		if i > 4 {
			break
		}
		fmt.Println("")

		fmt.Println("ID: ", *workflow.ID)
		fmt.Println("Name: ", workflow.GetName())
		fmt.Println("GetPath: ", workflow.GetPath())
		fmt.Println("Path: ", *workflow.Path)
		fmt.Println("EventsURL: ", workflow.GetURL())
		fmt.Println("GetState: ", workflow.GetState())
		fmt.Println("HTMLURL: ", *workflow.HTMLURL)

		fmt.Println("")
	}

	repo, resp, err := client.Repositories.Dispatch(
		context.Background(), "arxdsilva", "golang-ifood-sdk",
		github.DispatchRequestOptions{
			EventType: "PushEvent",
		})
	fmt.Println("err: ", err)
	fmt.Printf("repo: %+v\n", repo)
	fmt.Printf("resp: %+v\n", resp)
	result, err := ioutil.ReadAll(resp.Body)
	fmt.Println("err: ", err)
	fmt.Printf("status: %+v\n", string(result))
}
