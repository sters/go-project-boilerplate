# go-project-boilerplate
My go project boilerplate.

## Includes

- Makefile
  - run
  - test
  - cover
  - Tools install from `./tools/tools.go`
- Github Actions
  - Go
    - Lint by golangcilint
    - Run test and upload test coverage to codecov
  - Release
    - Make release when vX.X.X tag is added by goreleaser.

## TODO when use this

- [ ] Change `cmd/xxxxx` directory name
- [ ] Change run task in `Makefile`
- [ ] Change package name in `go.mod`
- [ ] Change main in `.goreleaser.yml`
- [ ] Update `README.md`
