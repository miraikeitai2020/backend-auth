GOCMD=go
DOCKERCMD=docker
GO_RUN=$(GOCMD) run
GO_BUILD=$(GOCMD) build
DOCKER_BUILD=$(DOCKERCMD) build
DOCKER_RUN=$(DOCKERCMD) run

all:
compose-build:
	docker-compose -f build/docker-compose.dev.yml up -d --build
compose-clean:
	docker-compose -f build/docker-compose.dev.yml down --rmi all --volumes