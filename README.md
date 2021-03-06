[![Go Report Card](https://goreportcard.com/badge/github.com/cameronbrill/create-go-app)](https://goreportcard.com/report/github.com/cameronbrill/create-go-app)
[![GoDoc](https://godoc.org/github.com/cameronbrill/create-go-app?status.svg)](https://godoc.org/github.com/cameronbrill/create-go-app)

# create-go-app

This is a [create-react-app](https://github.com/facebook/create-react-app) for Go projects, based on [go-project-template](https://github.com/cameronbrill/go-project-template).

## installation

Make sure you have your `go/bin` added to your `PATH` variable, such as in your `~/.zshrc` or `~/.bash_profile`

```
export PATH="$HOME/go/bin:$PATH"
```

In your terminal, run this command

```
go install github.com/cameronbrill/create-go-app@latest
```

## usage

Running `create-go-app` will prompt you to choose a project name and template.

Optionally, you can be explicit with the name and template. The following produces a bare-bones cli app called `project-name`:

```
create-go-app project-name --cli
```

## reach

- [ ] inject projects into current go module at `./cmd/<project>`


### disclaimer

This repository is not associated with [create-react-app](https://github.com/facebook/create-react-app) or [create-go-app/cli](https://github.com/create-go-app/cli) in any way. This is my personal side project.
