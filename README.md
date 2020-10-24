# lint-gomko

opinionated golang linters

## Linters

- `pkg/linters/appendr` - `append` statements usage linter
- `pkg/linters/elser` - `else` statements usage linter
- `pkg/linters/readall` - `ioutil.ReadAll` usage linter

# ioutil.ReadAll linter only

```
go get -u github.com/mkorenkov/lint-gomko/cmd/lint-gomko
go vet -vettool ~/bin/lint-gomko
```

# All linters

```
go get -u github.com/mkorenkov/lint-gomko/cmd/lint-gomko-all
go vet -vettool ~/bin/lint-gomko-all ./...
```

# Skip files/directories

Set `IGNORE` environment variable to comma-separated glob-like file patterns.

Examples:
- `IGNORE="**/*_test.go"` ignores all the test files
- `IGNORE="**/cmd/**"` ignores everything under cmd directory

```
IGNORE="**/*_test.go" go vet -vettool ~/bin/lint-gomko-all ./...
```

NOTE: uses [github.com/gobwas/glob](https://github.com/gobwas/glob) with `'/'` separator.

## Building and running from source

```
go build -o ~/bin/lint-gomko-all cmd/lint-gomko-all/main.go
go vet -vettool ~/bin/lint-gomko-all ./...
```
