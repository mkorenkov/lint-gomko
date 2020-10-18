# lint-gomko

opionated golang linters

## Linters

- `pkg/linters/appendr` - `append` statements usage linter
- `pkg/linters/elser` - `else` statements usage linter
- `pkg/linters/readall` - `ioutil.ReadAll` usage linter

# Latest release instructions

```
go get -u github.com/mkorenkov/lint-gomko/cmd/lint-gomko-all
go vet -vettool ~/bin/lint-gomko-all ./testdata/*
```

## Running from source instructions

```
go build -o bin/lint-gomko-all cmd/lint-gomko-all/main.go
go vet -vettool bin/lint-gomko-all ./testdata/*
```
