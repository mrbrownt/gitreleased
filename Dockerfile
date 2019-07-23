# GO build
FROM golang:1.12 as go-build

WORKDIR /app
COPY backend /app

ENV GO111MODULE=on
ENV GOPROXY=https://proxy.golang.org

RUN go install

# Yarn build
FROM node:11 as js-build
WORKDIR /app
COPY frontend /app

RUN yarn install --pure-lockfile && \
    yarn build

# Server container
FROM gcr.io/distroless/base
COPY --from=go-build /go/bin/backend /
COPY backend/migrations /migrations
COPY --from=js-build /app/dist /dist

CMD ["/backend"]
