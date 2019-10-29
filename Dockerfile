# build stage
FROM golang:1.11-alpine as backend
RUN apk add --update --no-cache bash ca-certificates curl git make tzdata

RUN mkdir -p /go/src/github.com/dgrezza/minimum-pod-scheduler
ADD Gopkg.* Makefile /go/src/github.com/dgrezza/minimum-pod-scheduler/
WORKDIR /go/src/github.com/dgrezza/minimum-pod-scheduler
RUN make vendor
ADD . /go/src/github.com/dgrezza/minimum-pod-scheduler
RUN CGO_ENABLED=0 go build -ldflags '-d -w -s' -o scheduler /go/src/github.com/dgrezza/minimum-pod-scheduler/cmd/scheduler

FROM alpine:3.7
COPY --from=backend /usr/share/zoneinfo/ /usr/share/zoneinfo/
COPY --from=backend /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=backend /go/src/github.com/dgrezza/minimum-pod-scheduler/scheduler /bin

ENTRYPOINT ["/bin/scheduler"]

