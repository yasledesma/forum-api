# Forum API

This is a solution to the [first lecture on Cornell University's Backend Course](https://backend-course.cornellappdev.com/chapters/routes/api). A sort of simplified Reddit clone.

## Getting Started
Setting up the project locally:

### Pre-requisites
This project has Docker as its only depency. You can install it by following the official [installation guide.](https://docs.docker.com/engine/install/)

### Installation
1. Clone the project locally.
2. Change into the repository and run `make init` to start the project.

This will start the project on development mode, where you can run all other commands defined in the Makefile.

For building and/or running a binary outside the container, you can use `make build` and `make run`.

Commands that are not part of the Makefile can be run by spinning up a terminal inside the container with `make term`.

## Usage
### Testing
All test suites can be run at the same time with:

```
make test
```

You can generate coverage reports by running:

```
make coverage
```

Also, if you want to run an specifict test suite or test, you can spin-up a terminal inside the container and run:

```
go test -v -run <test_suite>/<test>
```

### Deployment
To trigger the deployment pipeline, you need to create and push a new tag with the format:

```
v*.*.*
```

This will call the Render hook and create a new release with auto-generated notes.

