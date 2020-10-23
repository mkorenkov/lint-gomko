# lint-gomko

opinionated golang linters

## Linters

- `pkg/linters/appendr` - `append` statements usage linter
- `pkg/linters/elser` - `else` statements usage linter
- `pkg/linters/readall` - `ioutil.ReadAll` usage linter

# ioutil.ReadAll linter only

```
go get -u github.com/mkorenkov/lint-gomko/cmd/lint-gomko
lint-gomko ./testdata/*
```

# All linters

```
go get -u github.com/mkorenkov/lint-gomko/cmd/lint-gomko-all
go vet -vettool ~/bin/lint-gomko-all ./testdata/*
```

# Skip file patterns

Set `IGNORE` environment variable to comma-separated [glob](https://golang.org/pkg/path/filepath/#Glob) file patterns.

```
IGNORE="*.twirp.go,*_test.go" go vet -vettool ~/bin/lint-gomko-all ./testdata/*
```

## Running from source instructions

```
go build -o ~/bin/lint-gomko-all cmd/lint-gomko-all/main.go
go vet -vettool ~/bin/lint-gomko-all ./testdata/*
```
