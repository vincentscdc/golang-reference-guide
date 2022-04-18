# Deployment

Your service is destined to be running in a container in a kubernetes cluster.
Hence, you can test this service, in a container on a local k8s.

## Build

The build happens in a multistage docker container, either locally or remotely in the CI.
The existing Dockerfile provided in the example application folder is optimized for size and build speed. It should be work for most use cases, but feel free to modifiy it if needed.

## Local deployment

You can leverage your local k8s and the files present in the deployment folder, as well as the deploy script.

## CI/CD

We will be leveraging github actions, coveralls.
