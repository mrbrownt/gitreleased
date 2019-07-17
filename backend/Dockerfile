FROM golang:1.12 as build

WORKDIR /app
COPY . /app

ENV GO111MODULE=on
ENV GOPROXY=https://proxy.golang.org

RUN go install

FROM gcr.io/distroless/base
COPY --from=build /go/bin/backend /
COPY migrations /migrations
CMD ["/backend"]
