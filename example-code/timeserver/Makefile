ldocker-build:
	docker build -t timeserver .

docker-build:
	docker build -t nileshsimaria/timeserver:v4 .

ldocker-run:
	docker run -p 50051:50051 -it timeserver:latest
rdocker-run:
	docker run -p 50051:50051  nileshsimaria/timeserver:v4
docker-push:
	docker push nileshsimaria/timeserver:v4

genpb:
	protoc --go_out=plugins=grpc:. timep/timep.proto

clean:
	rm timep/timep.pb.go

.PHONY: genpb ldocker-run ldocker-build docker-build rdocker-run docker-push
