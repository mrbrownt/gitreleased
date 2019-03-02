FROM golang:1.11 as build

WORKDIR /go/src/gitlab.com/mrbrownt/gitreleased.app/backend
COPY . /go/src/gitlab.com/mrbrownt/gitreleased.app/backend

ENV GO111MODULE=on

RUN go mod vendor && \
    go install

FROM gcr.io/distroless/base
COPY --from=build /go/bin/backend /
CMD ["backend"]
