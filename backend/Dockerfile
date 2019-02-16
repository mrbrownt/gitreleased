FROM golang:1.11 as build

WORKDIR /go/src/gitlab.com/mrbrownt/gitreleased.app/backend
COPY . /go/src/gitlab.com/mrbrownt/gitreleased.app/backend

RUN go get ./... && \
    go install

FROM gcr.io/distroless/base
COPY --from=build /go/bin/backend /
CMD ["backend"]
