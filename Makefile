all:
	go build

install:
	./install.sh

test:
	go test ./...

recreate-systemd: install
	systemctl --user restart krunner-gitlab.service
