package runner

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/xanzy/go-gitlab"
)

func (r Runner) searchGroups(query string, matches []matchOut) ([]matchOut, error) {
	opt := &gitlab.ListGroupsOptions{
		ListOptions: gitlab.ListOptions{
			PerPage: viper.GetInt("items_to_show"),
			Page:    1,
		},
		OrderBy:      gitlab.String("name"),
		Search:       gitlab.String(query),
		Sort:         gitlab.String("desc"),
		Statistics:   gitlab.Bool(false),
		TopLevelOnly: gitlab.Bool(false),
	}

	groups, _, err := r.client.Groups.ListGroups(opt)
	if err != nil {
		return matches, err
	}

	matches = make([]matchOut, len(groups))
	for i := 0; i < len(groups); i++ {
		description := groups[i].Description
		if len(description) > 0 {
			description = description + "<br>"
		}

		matches[i].ID = groups[i].WebURL
		matches[i].Text = fmt.Sprintf(
			"<strong>%s</strong><br>%s[ID: %d]",
			groups[i].FullName,
			description,
			groups[i].ID,
		)
		matches[i].IconName = "internet-web-browser"
		matches[i].Type = 100
		matches[i].Properties = map[string]interface{}{
			"multiline": true,
		}
	}

	return matches, nil
}
