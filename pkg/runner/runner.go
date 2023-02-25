package runner

import (
	"github.com/godbus/dbus"
	"github.com/google/martian/log"
	"os/exec"
	"strings"
)

const CommandSearchProjects = "p"
const CommandSearchGroups = "g"

func (r Runner) Actions() ([]string, *dbus.Error) {
	return make([]string, 0), nil
}

func (r Runner) Match(query string) ([]matchOut, *dbus.Error) {
	matches := make([]matchOut, 0)

	command, query := r.separateQueryAndCommand(query)

	switch command {
	case CommandSearchProjects:
		matches, _ = r.searchProjects(query, matches)
	case CommandSearchGroups:
		matches, _ = r.searchGroups(query, matches)
	default:
		return matches, nil
	}

	return matches, nil
}

func (r Runner) separateQueryAndCommand(query string) (string, string) {
	if len(query) < len(r.queryPrefix)+r.queryMinLength {
		return "", ""
	}

	if !strings.HasPrefix(query, r.queryPrefix) {
		return "", ""
	}

	queryRaw, _ := strings.CutPrefix(query, r.queryPrefix)

	words := strings.Fields(queryRaw)
	command := words[0]

	query = strings.TrimSpace(strings.Join(words[1:], " "))

	return command, query
}

func (r Runner) Run(matchId string, actionId string) *dbus.Error {
	cmd := "xdg-open"
	args := []string{matchId}
	err := exec.Command(cmd, args...).Run()
	if err != nil {
		log.Errorf("Cannot execute command: %s", err.Error())
	}

	return nil
}
