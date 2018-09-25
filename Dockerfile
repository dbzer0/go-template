FROM golang as builder

COPY . /go/src/device.ufo/PROJECTNAME
WORKDIR /go/src/device.ufo/PROJECTNAME/src

RUN CGO_ENABLED=0 go build -v -a -tags PROJECTNAME -installsuffix PROJECTNAME -ldflags "-X main.version=${VERSION} -X main.revision=${BUILD}" -o /go/bin/PROJECTNAME

# контейнер окружения
FROM scratch
COPY --from=builder /go/bin/PROJECTNAME /
#VOLUME ["/data"]
ENTRYPOINT ["/PROJECTNAME", "-config", "/conf/PROJECTNAME.conf"]
