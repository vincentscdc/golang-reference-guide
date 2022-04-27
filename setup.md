# Install golang environment

## golang

make sure you have [brew](https://brew.sh/) installed or another package manager.

```bash
brew install golang
```

## MUST installs

### golang linter

```bash
brew install golangci-lint
```

In order to have a somewhate reproducible linting experience, you MUST use a config file, .golangci.toml.
You can access the reference one in the reference folder.

### Git changelog

```bash
brew install git-cliff
```

Git cliff allows you to automatically parse your git commit messages to generate a great changelog.

It also forces you to follow [git conventional commits](https://www.conventionalcommits.org).

### vscode and plugins

* Install [vscode](https://code.visualstudio.com/download)
* Install the golang extension (look in the plugins and search "@popular golang", it should be the first one).
* In the vscode preferences, search for go lint tool and select golangci-lint.
* In the vscode preferences, search for go format tool and select gofumports.

## nice to have companions

* k6: bench your API, locally or anywhere. Can be useful to profile your API.
* trivy: security scanner (docker image and source) that can be useful

```bash
brew install k6 trivy
```

## (OPTIONAL) a local kubernetes cluster and docker build locally

Yes, you want one.

You can follow the setup in the [infra-local-k8s repo](https://github.com/monacohq/infra-local-k8s)

## Conclusion

You're all set!

Now, go look at the different service examples in github.
