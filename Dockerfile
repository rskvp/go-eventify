FROM node:20.10.0-alpine3.18 AS REACT

RUN wget https://github.com/assalielmehdi/go-eventify-ui/zipball/master/ -O eventify.zip \
  && unzip eventify.zip \
  && mv assalielmehdi-go-eventify-ui-* /eventify

WORKDIR /eventify

RUN yarn install

RUN yarn build

FROM golang:1.21.4-alpine3.18 AS GO

RUN wget https://github.com/assalielmehdi/go-eventify/zipball/master/ -O eventify.zip \
  && unzip eventify.zip \
  && mv assalielmehdi-go-eventify-* /eventify

WORKDIR /eventify

RUN mkdir templates static

COPY --from=REACT /eventify/build/index.html ./templates/index.html
COPY --from=REACT /eventify/build/static/ ./static/

RUN go build .

FROM alpine:3.18

WORKDIR /eventify

COPY --from=GO /eventify/templates ./templates
COPY --from=GO /eventify/static ./static
COPY --from=GO /eventify/.env .
COPY --from=GO /eventify/eventify .

ENV GIN_MODE=release

ENTRYPOINT ["./eventify"]