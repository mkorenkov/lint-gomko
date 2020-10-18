# lint-gomko

opinionated golang linters

## Linters

- `pkg/linters/appendr` - `append` statements usage linter
- `pkg/linters/elser` - `else` statements usage linter
- `pkg/linters/readall` - `ioutil.ReadAll` usage linter

# Opinionated instructions

```
go get -u github.com/mkorenkov/lint-gomko/cmd/lint-gomko-all
go vet -vettool ~/bin/lint-gomko-all ./testdata/*
```

# Less opinionated linters

```
go get -u github.com/mkorenkov/lint-gomko/cmd/ioutil-readall
ioutil-readall ./testdata/*
```

## Running from source instructions

```
go build -o bin/lint-gomko-all cmd/lint-gomko-all/main.go
go vet -vettool bin/lint-gomko-all ./testdata/*
```
