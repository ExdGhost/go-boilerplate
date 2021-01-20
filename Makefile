#Declarations
# FLAG is used to notify the shell scripts that they are being invoked from a Makefile
FLAG="TRUE"
# BUILDNAME is the standard name tag used throughout, change it as required for the specific project
BUILDNAME="go-boilerplate-api"
# ENV is the standard env name passed for running in docker
ENV="docker"

#Commands
build: . # Creates binary
	go build cmd/api -o ${BUILDNAME}

clean: . # Cleans unecessary junk
	rm -Rf ${BUILDNAME} _reports

re-init: . # Completely cleans current repo (vendor, gomod, gosum and junk files) and re initializes it
	rm -Rf ${BUILDNAME} _reports vendor go.mod go.sum
	go mod init ${BUILDNAME}
	go mod vendor

vendor: . # Initializes dependencies only
	go mod vendor

lint: # Runs go linter
	./scripts/lint.sh ${FLAG}

run: # Runs app
	TIER=development go run cmd/api/main.go

test: # Runs test scripts
	./scripts/test.sh ${FLAG}

# Docker based command , needs docker to be installed on the system to work
docker-build: # Builds docker image
	docker image rm -f ${BUILDNAME}
	./build/build ${BUILDNAME} ${FLAG}

docker-run: # Runs already built docker image (will fail if image has not been built)
	./build/run ${BUILDNAME} ${ENV} ${FLAG}

docker-clean: # Cleans existing the docker image using the buildname
	docker image rm -f ${BUILDNAME}

swag-install: # Installs swagger ui package
	go get -u github.com/swaggo/swag/cmd/swag

swag-init: # Parses & initializes directory with swagger docs (vs code requires a reload sometimes, and perform go mod vendor on 1st run)
	cp cmd/api/main.go main.go
	swag init
	rm main.go
