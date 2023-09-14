FROM golang:1.21.1

WORKDIR /src

COPY go.mod go.sum .
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o fgprof

FROM alpine:3.18.3
COPY --from=0 /src/fgprof /usr/local/bin
ENTRYPOINT [ "/usr/local/bin/fgprof" ]
