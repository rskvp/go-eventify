FROM node:20.10.0-alpine AS ui

RUN wget https://github.com/assalielmehdi/go-eventify-ui/zipball/master/ -O ui.zip \
  && unzip ui.zip \
  && mv assalielmehdi-go-eventify-ui-* /ui

WORKDIR /ui

RUN yarn install

RUN yarn build

FROM golang:1.21.4-alpine

RUN apk add build-base \
  && wget https://github.com/assalielmehdi/go-eventify/zipball/master/ -O server.zip \
  && unzip server.zip \
  && mv assalielmehdi-go-eventify-* /eventify

WORKDIR /eventify

RUN mkdir templates static

COPY --from=ui /ui/build/index.html ./templates/index.html
COPY --from=ui /ui/build/static/ ./static/

RUN go build .

ENV GIN_MODE=release

ENTRYPOINT ["./eventify"]