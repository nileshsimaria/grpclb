
ldocker-build:
	docker build -t timeclient .

docker-build:
	docker build -t nileshsimaria/timeclient:v1 .

ldocker-run:
	docker run --net=host timeclient:latest --host localhost:50051 --count 5
rdocker-run:
	docker run --net=host nileshsimaria/timeclient:v1 --host localhost:50051 --count 2	

genpb:
	protoc --go_out=plugins=grpc:. timep/timep.proto

clean:
	rm timep/timep.pb.go

.PHONY: genpb ldocker-build ldocker-run docker-build rdocker-run
