FROM golang:1.17

RUN mkdir /build
WORKDIR /build

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 3000

CMD ["httpServer"]