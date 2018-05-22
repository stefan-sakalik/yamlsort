# yamlsort

This utility sorts YAML map keys by serializing and deserializing it in go. You can use tool to diffing unordered YAML files.

## Usage

```
export GOPATH="${GOPATH:-$HOME/go}"
export GOBIN="${GOBIN:-$GOPATH/bin}"
export PATH="$GOBIN:$PATH"
go get github.com/stefan-sakalik/yamlsort

# Just simple sorting
yamlsort < path/to/file1.yaml

# Diff two YAMLs
vimdiff <(yamlsort < path/to/file1.yaml) <(yamlsort < path/to/file2.yaml)
```
