FROM node:11 as build

COPY . /app
WORKDIR /app

RUN yarn install --pure-lockfile && \
    yarn build

FROM nginx:alpine
COPY --from=build /app/dist /usr/share/nginx/html

COPY default.conf /etc/nginx/conf.d/default.conf
