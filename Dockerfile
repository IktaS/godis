FROM golang:alpine

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64


WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .


RUN go build -o main ./cmd/godis/

WORKDIR /dist

RUN cp /build/main .
RUN cp /build/.env .
RUN cp -r /build/web .

EXPOSE 8080

CMD ["/dist/main"]