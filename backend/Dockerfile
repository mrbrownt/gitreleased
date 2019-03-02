FROM golang:1.12 as build

WORKDIR /app
COPY . /app

ENV GO111MODULE=on

RUN go mod vendor && \
    go install -mod=vendor

FROM gcr.io/distroless/base
COPY --from=build /go/bin/backend /
CMD ["backend"]
