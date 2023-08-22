---- CLI app that tracks time

-- uses UI widgets rendered inside the terminal

-- use build tags to switch between repositories either inmemory or sqlite

go build -tags=inmemory ./ --> to store the data in memory without compiling the sqlite files

TODO
golang parser/linter/formatter from vscode doesn't work properly when using build tags
