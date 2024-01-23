package main

import (
	"log"

	"github.com/gdamore/tcell/v2"
	"github.com/gliderlabs/ssh"
	"github.com/rivo/tview"
)

type Repo struct {
	Name  string
	Owner string
	Stars int
}

func main() {
	ssh.Handle(func(s ssh.Session) {
		// Initialize TUI application
		app := tview.NewApplication()

		// Fetch GitHub repositories (pseudo-function)
		repos := fetchGitHubRepos() // Implement this function

		// Create a list to display repositories
		list := tview.NewList().ShowSecondaryText(false)

		// Add repositories to the list
		for _, repo := range repos {
			list.AddItem(repo.Name, "", 0, nil)
		}

		// Add color and attributes
		list.SetSelectedBackgroundColor(tcell.ColorGreen).
			SetSelectedTextColor(tcell.ColorBlack)

		// Set up the application
		if err := app.SetRoot(list, true).Run(); err != nil {
			panic(err)
		}
	})

	log.Println("starting ssh server on port 2222...")
	log.Fatal(ssh.ListenAndServe(":2222", nil))
}

// fetchGitHubRepos returns a list of GitHub repositories.
func fetchGitHubRepos() []Repo {
	return []Repo{
		{Name: "tview", Owner: "rivo", Stars: 1725},
		{Name: "go-git", Owner: "go-git", Stars: 9484},
		{Name: "gopsutil", Owner: "shirou", Stars: 4814},
		{Name: "termui", Owner: "gizak", Stars: 7934},
		{Name: "goreleaser", Owner: "goreleaser", Stars: 6334},
	}
}
