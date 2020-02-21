ldocker-build:
	docker build -t timeserver .

ldocker-run:
	docker run -p 50051:50051 -it timeserver:latest
genpb:
	protoc --go_out=plugins=grpc:. timep/timep.proto

clean:
	rm timep/timep.pb.go

.PHONY: genpb ldocker-run ldocker-build
