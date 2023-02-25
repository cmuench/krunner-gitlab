package runner

import "github.com/xanzy/go-gitlab"

type matchOut struct {
	ID, Text, IconName string
	Type               int32
	Relevance          float64
	Properties         map[string]interface{}
}

// http://blog.davidedmundson.co.uk/blog/cross-process-runners/
type Runner struct {
	client         *gitlab.Client
	queryPrefix    string
	queryMinLength int
}

func NewRunner(client *gitlab.Client, queryPrefix string, queryMinLength int) *Runner {
	return &Runner{
		client,
		queryPrefix,
		queryMinLength,
	}
}
