# Bootstrapping a new project

## Application

### Reference folders

You should follow the examples in the reference example folder (TODO: create this folder and add the link).

clone it, remove the .git folder, rename the things you need to rename and you are good to go.

### Your project folder architecture (OPINION)

| folder           | usage                                                                                                  |
|------------------|--------------------------------------------------------------------------------------------------------|
| cmd/serve/       | your server is running here, usually in a main.go file, package main                                   |
| cmd/migrations/  | your migration is running here, usually in a main.go file, package main                                |
| config/          | your toml/yaml config, will be loaded by a configmap in production                                     |
| /pkg/            | all your libs and functions                                                                            |
|     /config      | your config struct and the function to load it                                                         |
|     /internal    | (optional) used to hide from external import                                                           |
|     /xxx         | entity package for xxx                                                                                 |
|         /http    | your http wrapper implementation                                                                       |
|     /yyy         | entity package for yyy                                                                                 |
|         /http    | your http wrapper implementation                                                                       |
| /infra/          | whatever is related to deployment, for ex your yaml files for your local k8s cluster                   |
| /sql/            | all things directly related to sql                                                                     |
|     /migrations/ | sql files to be run by the migration                                                                   |
|     /init        | init to be run in your local database or by dba on production                                          |  

### Tooling (MUST)

Make sure to look at the Makefile at the root of the project.
**You should be able to test, lint, bench, build.**

**Make sure you have a .golangci.toml config file in your folder.**

For other things, refer to specific topics in the [README](./README.md)

## Library (MUST)

Each library must be tested, documented, benched and released according to [go principles](https://go.dev/blog/using-go-modules).

* One repository per lib
* Proper semantic versioning
