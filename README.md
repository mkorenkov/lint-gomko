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
go vet -vettool ~/bin/lint-gomko-all
```

# Skip files/directories

Set `IGNORE` environment variable to comma-separated [glob](https://golang.org/pkg/path/filepath/#Glob) file patterns.
Glob is path-dependent and since `go vet` can vet the package in the current directory, dirs can be ignored using this invocation:

```
IGNORE="testdata/*" go vet -vettool ~/bin/lint-gomko-all
```

Otherwise, package path can be passed to `go vet` and glob patterns would be dependant on that path

```
IGNORE="*.twirp.go,*_test.go" go vet -vettool ~/bin/lint-gomko-all ./testdata/*
```

## Building and running from source

```
go build -o ~/bin/lint-gomko-all cmd/lint-gomko-all/main.go
go vet -vettool ~/bin/lint-gomko-all
```
