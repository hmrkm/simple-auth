FROM golang:1.16

WORKDIR /workspace

COPY go.mod .
COPY go.sum .
COPY .air.toml .

RUN go install

CMD ["air", "-c", ".air.toml"]