FROM golang:latest


WORKDIR $GOPATH/src/gindocker
ADD . $GOPATH/src/gindocker
RUN go build .

EXPOSE 8000

ENTRYPOINT ["./gindocker"]