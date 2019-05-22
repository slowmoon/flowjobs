FROM golang:1.12 AS build-env
ADD ./src /opt/src/first-project
WORKDIR /opt/src/first-project
RUN go build


FROM alpine
RUN apk add -U tzdata
COPY --from=build-env /opt/src/first-project/game-server /usr/local/bin
RUN chmod +x /usr/local/bin/game-server

CMD [ "game-server" ]