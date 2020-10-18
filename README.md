# lint-gomko

opionated golang linters

## Linders

- `pkg/linters/appendr` - `append` statements usage linter
- `pkg/linters/elser` - `else` statements usage linter
- `pkg/linters/readall` - `ioutil.ReadAll` usage linter

## Building

```
go build -o bin/lint-gomko-all cmd/lint-gomko-all/main.go
go vet -vettool bin/lint-gomko-all ./testdata/*
```
