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

