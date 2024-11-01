#!/usr/bin/bash

set -x

# Install starship prompt.
cd /tmp && curl -sS https://starship.rs/install.sh | sh -s -- -y || exit -1
mkdir -p ~/.config/fish
# echo "starship init fish | source" >> ~/.config/fish/config.fish

# Install parceljs.
npm install --save-dev parcel || exit -1

# Install latest templ command.
go install github.com/a-h/templ/cmd/templ@latest
# Install latest air command.
go install github.com/air-verse/air@latest
# Run go mod tidy.
cd ../ && go mod tidy

exit 0
