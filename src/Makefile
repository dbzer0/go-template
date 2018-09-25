# This how we want to name the binary output
BINARY=PROJECTNAME

# These are the values we want to pass for VERSION and BUILD
VERSION=0.0.1
REVISION=`git rev-parse HEAD`

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS=-ldflags "-X main.version=${VERSION} -X main.revision=${REVISION}"

all	: build

# Builds the project
build:
	go build ${LDFLAGS} -o ../bin/${BINARY}

# Cleans our project: deletes binaries
clean:
	if [ -f ../bin/${BINARY} ] ; then rm ../bin/${BINARY} ; fi

.PHONY: all build clean
