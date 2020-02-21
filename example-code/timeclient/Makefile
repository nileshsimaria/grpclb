
ldocker-build:
	docker build -t timeclient .

ldocker-run:
	docker run --net=host timeclient:latest --host localhost:50051 --count 5

genpb:
	protoc --go_out=plugins=grpc:. timep/timep.proto

clean:
	rm timep/timep.pb.go

.PHONY: genpb ldocker-build ldocker-run
