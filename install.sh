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

if [ ! -f ~/.local/share/icons ]; then
  mkdir -p ~/.local/share/icons;
fi

xdg-icon-resource install --size 128 ./icons/krunner-gitlab.png
xdg-icon-resource install --size 256 ./icons/krunner-gitlab.png
xdg-icon-resource install --size 512 ./icons/krunner-gitlab.png
