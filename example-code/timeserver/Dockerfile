FROM golang:1.13.0-alpine3.10

EXPOSE 50051
WORKDIR /src
COPY . .
RUN GO111MODULE=on go build -mod vendor

ENTRYPOINT [ "/src/timeserver"]

