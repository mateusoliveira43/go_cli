# Go CLI

TODO finish refactor CLI
TODO finish tutorials
TODO https://go.dev/doc/articles/wiki/ follow tutorial

A Command Line Interface (CLI) written in Golang to learn the language 😄

- [ ] Containerized development environment
- [ ] CI/CD pipeline

## Quality

TODO

TODO use `go get` in development tools so they are marked as indirect? (will this increase executable size?)

### Tests

TODO
### Type checking

TODO

### Linters

To run Go code linter, run
```
TODO
```

To run Shell script linter, run
```
TODO shellcheck
```

To run Containerfile linter, run
```
TODO hadolint
```

### Code formatters

To check Go code format, run
```
TODO
```

To format Go code, run
```
TODO
```

To check all repository's files format, run
```
TODO https://github.com/editorconfig/editorconfig-core-go
```

File format configuration in [`.editorconfig`](.editorconfig) file.

### Security vulnerability scanners

To check common security issues in Go code, run
```
TODO
```

To check known security vulnerabilities in Go dependencies, run
```
TODO
```

To scan Container Image, run
```
TODO
```

### Documentation

To install Go documentation generator tool, run
```sh
go install -v golang.org/x/pkgsite/cmd/pkgsite@latest
```

- [ ] any way to check if there are any errors on the generation of the documentation?
- [ ] any way to see generated HTML files?

To generate Go documentation, run
```sh
pkgsite
```
To see the documentation, access [`http://localhost:8080`](http://localhost:8080).

> Go code documentation is called "Doc comments" (the equivalent of doctrings for Python). For more information, check [Go Doc Comments](https://go.dev/doc/comment).

## TODO Pre-commit

To configure pre-commit automatically when cloning this repo, run
```
git config --global init.templateDir ~/.git-template
pre-commit init-templatedir --hook-type commit-msg --hook-type pre-commit ~/.git-template
```
pre-commit must be installed globally.

To configure pre-commit locally, run
```
pre-commit install --hook-type commit-msg --hook-type pre-commit
```

To test it, run
```
pre-commit run --all-files
```

pre-commit configuration in [`.pre-commit-config.yaml`](.pre-commit-config.yaml) file.

## License

This repository is licensed under the terms of [MIT License](LICENSE).

## Learning Go

## Steps

Followed https://spf13.com/presentation/building-an-awesome-cli-app-in-go-oscon/

```
go install github.com/spf13/cobra-cli@latest
go mod init github.com/mateusoliveira43/go_cli
cobra-cli init .
cobra-cli add add
cobra-cli add list
cobra-cli add done
go mod tidy
go run main.go <flag/command>
```

### [Getting started](https://go.dev/doc/tutorial/getting-started)
### [Create Module](https://go.dev/doc/tutorial/create-module)
### Stopped at https://go.dev/doc/tutorial/random-greeting

To create `go.mod` file, run
```
go mod init <path/name>
```

To run go file run
```
go run .
```

To create `go.sum`, update `go.mod` file and add dependency from [pkg.go.dev](https://pkg.go.dev/), run
```
go mod tidy
```

#### Summary

- `go.mod` file stores the **module** name and its dependencies (like Go version and Go libraries versions). Like Python's `pyproject.toml`.

- `go.sum` file stores the dependencies hashes, for security. Like Python's `poetry.lock`.

- Module name must be a path from which your module can be downloaded by Go tools. Ex.: `github.com/mateusoliveira43/go_cli`.

- Add dependency to your Module `go get <ModuleName>`. Ex.: `go get github.com/spf13/cobra`.
