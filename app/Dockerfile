FROM golang:1.16

WORKDIR /workspace

COPY . .

RUN apt update && apt upgrade -y

RUN go install

RUN go get -u github.com/cosmtrek/air && \
    go build -o /go/bin/air github.com/cosmtrek/air

CMD ["air", "-c", ".air.toml"]