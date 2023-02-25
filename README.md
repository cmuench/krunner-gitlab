# KRunner GitLab backend

This is a fork of the [krunner-gitlab](https://github.com/shochdoerfer/krunner-gitlab) plugin of my friend [Stephan Hochdörfer](https://github.com/shochdoerfer).

This package provides a [KRunner](https://blog.davidedmundson.co.uk/blog/cross-process-runners/) backend which will use a GitLab instance as a search backend. Currently only project names are searched for, this might change in the future.

![Maintenance Badge](https://img.shields.io/maintenance/yes/2023.svg)                                      
[![Go Report Card](https://goreportcard.com/badge/github.com/cmuench/krunner-gitlab)](https://goreportcard.com/report/github.com/cmuench/krunner-gitlab)
[![Go Github Action Workflow](https://github.com/cmuench/krunner-gitlab/workflows/Go/badge.svg)](https://github.com/cmuench/krunner-gitlab/actions?query=workflow%3AGo)

## Installation

Clone this repository

```
git clone https://github.com/cmuench/krunner-gitlab
```

Build the go application

```
go install
```

Register the runner in KDE by storing a file called `krunner-gitlab.desktop` in `$HOME/.local/share/kservices5` and then restart the rkunner process.

See: [krunner-gitlab.desktop](krunner-gitlab.desktop)

To configure krunner-gitlab with the url and the access token for your own GitLab instance, create a file `~/.config/krunner-gitlab/config.yaml` like this:

```
url: https://your-gitlab-server/api/v4
token: your-token
items_to_show: 10
query_min_length: 4
query_prefix: gitlab

```

It is important to note that the url needs to point to the GitLab API url!

## Run the application

Run `krunner-gitlab` in your `go/bin` directory. Invoke KRunner and start searching for GitLab projects. 

## systemd Integration

Add a user service in your home directory.

Example file: `~/.config/systemd/user/krunner-gitlab.service`

Check that the krunner-gitlab binary exists in your `~/go/bin` directory. Change path if the binary was installed
in another location.

```ini
[Unit]
Description=krunner Gitlab Service
ConditionPathExists=%h/.config/krunner-gitlab/config.yaml
After=default.target

[Service]
Type=simple
Restart=on-failure
RestartSec=5s
ExecStart=%h/go/bin/krunner-gitlab
SyslogIdentifier=krunner-gitlab

[Install]
WantedBy=default.target
```

Activate the user service with:

```bash
systemctl --user daemon-reload
systemctl --user enable krunner-gitlab.service
systemctl --user start krunner-gitlab.service
```

## License

KRunner GitLab is released under the Apache 2.0 license.
