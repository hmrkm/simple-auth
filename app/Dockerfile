FROM golang

WORKDIR /workspace

COPY . .

RUN apt update && apt upgrade -y

RUN go install

RUN go get github.com/golang/mock/mockgen && \
    go get github.com/deepmap/oapi-codegen/cmd/oapi-codegen

RUN go get -u github.com/cosmtrek/air && \
    go build -o /go/bin/air github.com/cosmtrek/air

ENTRYPOINT ["air", "-c", ".air.toml"]