FROM golang:1.13.0-alpine3.10

WORKDIR /src
COPY . .
RUN GO111MODULE=on go build -mod vendor

ENTRYPOINT [ "/src/timeclient", "--count", "5" ]
