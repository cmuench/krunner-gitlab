package runner

import (
	"testing"
)

func TestSeparateQueryAndCommand(t *testing.T) {
	runner := Runner{
		queryPrefix: "gitlab",
	}

	command, query := runner.separateQueryAndCommand("gitlab p myproject")
	if command != CommandSearchProjects {
		t.Errorf("Failed to detect command, got '%s', want: '%s'.", command, "p")
	}
	if query != "myproject" {
		t.Errorf("Failed to detect query, got '%s', want: '%s'.", query, "myproject")
	}

	command, query = runner.separateQueryAndCommand("gitlab p     myproject")
	if command != CommandSearchProjects {
		t.Errorf("Failed to detect command, got '%s', want: '%s'.", command, "p")
	}
	if query != "myproject" {
		t.Errorf("Failed to detect query, got '%s', want: '%s'.", query, "myproject")
	}

	command, query = runner.separateQueryAndCommand("gitlab p name with whitespaces")
	if command != CommandSearchProjects {
		t.Errorf("Failed to detect command, got '%s', want: '%s'.", command, "p")
	}
	if query != "name with whitespaces" {
		t.Errorf("Failed to detect query, got '%s', want: '%s'.", query, "name with whitespaces")
	}
}
