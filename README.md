---

Note to reviewers.  I guess I was confused by what you were looking for and maybe overdid it. This "runbook" won't make sence without the code. The whole app and directions can be found here: https://github.com/jgreat/things

---

# things


A simple REST/CRUD code sample app.

Golang, PostgreSQL

Done in Go so I can learn some more Go.

## Config

| Variable | Type | Description |
| --- | --- | --- |
| DB_URL | `string` | `postgres://` url to database endpoint |

## Endpoints

List all things

Example
```
curl localhost:8080/things
```

Add a thing

Payload

```
{ "name": "string"}
```

Example

```
curl --header "Content-Type: application/json" \
  --request POST --data '{"name":"plumbus"}' \   
  http://localhost:8080/thing
```

TODO: Delete a thing

TODO: Update a thing

---

## Development

See [Setup](#setup) for requirements and creating a go development environment.

### Doing development

When you start work, source the go project `source_me.sh` file to setup your environment.

```
cd ~/gopath/things
source ./source_me.sh
```

Use `docker-compose` to build/run the app with sample DB and follow the logs:

```
cd ~/gopath/things/src/github.com/jgreat/things
docker-compose up --build
```

* App endpoint: http://127.0.0.1:8080
* DB endpoint: `postgres://things:example@db:5432/things?sslmode=disable`

When you make changes to the code the container should automatically restart and rebuild the app binary.

### Debugging

When launched from `docker-compose` the app is run with headless `dlv` listening on `:2345`. 

There is a pre-configured `.vscode/launch.json` profile ready to attach to `dlv` to remote debug.

In `vscode` select the debug option and run `Attach Remote`, set your break points and have fun.

---

## Setup

### Prerequisites

This environment is pre-configured for running/compling in docker with remote debugging and automatic rebuilds on code changes.

- `vscode` - https://code.visualstudio.com/download
- `go` 1.15 - https://golang.org/dl/
- `docker` - https://docs.docker.com/get-docker/
- `docker-compose` - https://docs.docker.com/compose/install/

### Install go

Download latest go 1.15 for your system: https://golang.org/dl/

Extract tar to `~/bin` (this will overwite the contents of the current `go`)

```
cd ~/bin
tar xvzf ~/Downloads/go1.15.3.linux-amd64.tar.gz
```

Move `go` to a versioned directory

```
mv go go-1.15.3
```

### Set up development environment

These instructions will help you create an isolated project path in you home directory.

```
mkdir -p ~/gopath/things
cd ~/gopath/things
```

Add this script to your project base and point `GOROOT` at the version of go:

`source_me.sh`

```
#!/bin/bash

export GOROOT="${HOME}/bin/go-1.15.3"
export GOPATH="$(pwd)"
export PATH="${PATH}:${GOROOT}/bin:${GOPATH}/bin"
```

Create src path and clone the repo:

```
cd ~/gopath/things
mkdir -p src/github.com/jgreat
cd src/github.com/jgreat
git clone git@github.com:jgreat/things.git
```

## TODOs and Caveats

- This is designed to work in a "native" docker experience (on Linux).  For Windows or OSX more work will be needed to wire up the networking since everything is harder (in a VM).
- There's a ton to do in error handling, monitoring, authentication, security... 
- There is a prod ready Dockerfile in the base of the project, but work is needed to wire that into a CI/CD system for building.
- Needs a Helm chart for deployment.
