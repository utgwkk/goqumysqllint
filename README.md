# goqumysqllint

A linter for goqu query builder with MySQL dialect

## Example

See [testdata/src/](./testdata/src/) directory.

## Usage

### As a vettool

```
$ go install github.com/utgwkk/goqumysqllint/cmd/goqumysqllint@latest
$ go vet -vettool=`which goqumysqllint` ./...
```

### As a golangci-lint plugin

You can use goqumysqllint via [golangci-lint's module plugin system](https://golangci-lint.run/plugins/module-plugins/).

```yaml
version: v1.62.2
plugins:
  - module: 'github.com/utgwkk/goqumysqllint'
    version: latest
```
