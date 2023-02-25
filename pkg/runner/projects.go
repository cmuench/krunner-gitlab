package runner

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/xanzy/go-gitlab"
	"time"
)

func (r Runner) searchProjects(query string, matches []matchOut) ([]matchOut, error) {
	opt := &gitlab.ListProjectsOptions{
		Search:           gitlab.String(query),
		SearchNamespaces: gitlab.Bool(true),
		Membership:       gitlab.Bool(true),
		Simple:           gitlab.Bool(true),
		Archived:         gitlab.Bool(false),
		Statistics:       gitlab.Bool(false),
		OrderBy:          gitlab.String("last_activity_at"),
		Sort:             gitlab.String("desc"),
		ListOptions: gitlab.ListOptions{
			PerPage: viper.GetInt("items_to_show"),
			Page:    1,
		},
	}

	projects, _, err := r.client.Projects.ListProjects(opt)
	if err != nil {
		return matches, err
	}

	matches = make([]matchOut, len(projects))
	for i := 0; i < len(projects); i++ {
		// the longer the last acitivity has been, the least important is the project
		relevance := (float64(time.Now().Unix()) / float64(projects[i].LastActivityAt.Unix()))
		if relevance > 1 {
			relevance = 1 - ((relevance - float64(1)) * float64(10))
		}

		description := projects[i].Description
		if len(description) > 0 {
			description += "<br>"
		}

		matches[i].ID = projects[i].WebURL
		matches[i].Text = fmt.Sprintf(
			"<strong>%s</strong><br>%s<i>%s</i><br>[ID: %d]",
			projects[i].Name,
			description,
			projects[i].PathWithNamespace,
			projects[i].ID,
		)
		matches[i].IconName = "hicolor/128x128/apps/krunner-gitlab.png"
		matches[i].Type = 100
		matches[i].Relevance = 1
		matches[i].Properties = map[string]interface{}{
			"multiline": true,
		}
	}

	return matches, nil
}
