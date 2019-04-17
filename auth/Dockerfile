FROM golang:1.12 as build

WORKDIR /app
COPY . .

RUN go install -mod=vendor

FROM gcr.io/distroless/base
COPY --from=build /go/bin/auth /
CMD ["/auth"]
