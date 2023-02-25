#!/bin/bash

go install

if [ ! -d ~/.config/krunner-gitlab ]; then
  mkdir -p ~/.config/krunner-gitlab
fi

if [ ! -f ~/.config/krunner-gitlab/config.yaml ]; then
  cp config.yaml.dist ~/.config/krunner-gitlab/config.yaml
fi

if [ ! -f ~/.local/share/kservices5/krunner-gitlab.desktop ]; then
  cp krunner-gitlab.desktop ~/.local/share/kservices5/krunner-gitlab.desktop
fi
