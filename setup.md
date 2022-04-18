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

### First, have a container runtime

Since March 2022, docker desktop is a paid subscription, so you can replace it with another container runtime through VM:

* [rancher desktop](https://rancherdesktop.io/) (STILL SOME ISSUES WITH certs)
* [lima](https://github.com/lima-vm/lima)
* [colima](https://github.com/abiosoft/colima)

### Running your cluster on docker

Via [k3d](https://k3d.io/v5.4.1/), you can setup your cluster with a simple config file, whatever the container runtime you have (almost).

* Install k3d, kubectl, docker, helm via brew

```bash
brew install kubectl helm k3d docker
```

### If you want to test your services with local https

* mkcert: generate certs valid for your local computer (to serve and access your service via https for ex)

## Conclusion

You're all set!

Now, go look at the different service examples in github.
