FROM golang:1.9.5

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./... && \
    go install -v ./... && \
    go build ./main.go

CMD ["./main"]