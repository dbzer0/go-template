#
# Контейнер сборки
#
FROM golang:1.14 as builder

ARG DRONE
ARG DRONE_TAG
ARG DRONE_COMMIT
ARG DRONE_BRANCH

ENV CGO_ENABLED=0

COPY . /go/src/github.com/dbzer0/go-template
WORKDIR /go/src/github.com/dbzer0/go-template
RUN \
    if [ -z "$DRONE" ] ; then echo "no drone" && version=`git describe --abbrev=6 --always --tag`; \
    else version=${DRONE_TAG}${DRONE_BRANCH}-`echo ${DRONE_COMMIT} | cut -c 1-7` ; fi && \
    echo "version=$version" && \
    cd app && \
    go build -a -tags PROJECTNAME -installsuffix PROJECTNAME -ldflags "-X main.version=${version} -s -w" -o /go/bin/PROJECTNAME

#
# Контейнер для получения актуальных SSL/TLS сертификатов
#
FROM alpine as alpine
COPY --from=builder /etc/ssl/certs /etc/ssl/certs
RUN addgroup -S PROJECTNAME && adduser -S PROJECTNAME -G PROJECTNAME

ENTRYPOINT [ "/bin/PROJECTNAME" ]

#
# Контейнер рантайма
#
FROM scratch
COPY --from=builder /go/bin/PROJECTNAME /bin/PROJECTNAME

# копируем сертификаты из alpine
COPY --from=alpine /etc/ssl/certs /etc/ssl/certs

# копируем документацию
COPY --from=alpine /usr/share/PROJECTNAME /usr/share/PROJECTNAME

# копируем пользователя и группу из alpine
COPY --from=alpine /etc/passwd /etc/passwd
COPY --from=alpine /etc/group /etc/group

USER PROJECTNAME

ENTRYPOINT ["/bin/PROJECTNAME"]



