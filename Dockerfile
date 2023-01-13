FROM golang:1.19

WORKDIR /usr/src/app

COPY . .

RUN go get .
RUN go build -o server.bin

EXPOSE 8080

CMD ["./server.bin"]