# Run example code locally on your laptop

## Start server
```
$ cd github.com/nileshsimaria/grpclb/example-code/timeserver
$ go run main.go --help
Usage of /var/folders/cq/lzr00k1145s82jhsdgz9v1k40000gn/T/go-build144705154/b001/exe/main:
  -insecure
    	startup without TLS
  -port string
    	grpc server port number (default ":50051")
exit status 2
$ go run main.go
2020/02/21 23:18:53 Starting gRPC server(port=:50051)

```

By default it runs with insecure set to off.

## Start client
```
$ cd github.com/nileshsimaria/grpclb/example-code/timeclient
$ go run main.go --help
Usage of /var/folders/cq/lzr00k1145s82jhsdgz9v1k40000gn/T/go-build582306951/b001/exe/main:
  -cacert string
    	CACert for server (default "CA.crt")
  -count int
    	number of rpc calls (default 1)
  -host string
    	grpc server host:port (default "localhost:50051")
  -insecure
    	connect without TLS
  -servername string
    	CACert for server (default "timeserver.gke.net")
exit status 2

$ go run main.go
2020/02/21 23:20:32 Reply: time:"[time=2020-02-21 23:20:32.208941 -0800 PST m=+98.407762621] [host=nsimaria-mbp]"
$ go run main.go --count 3
2020/02/21 23:20:37 Reply: time:"[time=2020-02-21 23:20:37.168126 -0800 PST m=+103.366909075] [host=nsimaria-mbp]"
2020/02/21 23:20:37 Reply: time:"[time=2020-02-21 23:20:37.170695 -0800 PST m=+103.369478058] [host=nsimaria-mbp]"
2020/02/21 23:20:37 Reply: time:"[time=2020-02-21 23:20:37.173266 -0800 PST m=+103.372048705] [host=nsimaria-mbp]"
```

Note : By default it runs with ssl/tls on. The defaults command line options would give you all necessary stuff to talk to server over TLS.

Docker run (server)
```
$ docker run -p 50051:50051 nileshsimaria/timeserver:v4
```
Docker run (client)
```
$ docker run --net=host nileshsimaria/timeclient:v4 --count 1
2020/02/22 07:24:04 Reply: time:"[time=2020-02-22 07:24:04.1015099 +0000 UTC m=+31.318027001] [host=281998716c79]"
$ docker run --net=host nileshsimaria/timeclient:v4 --count 3
2020/02/22 07:24:06 Reply: time:"[time=2020-02-22 07:24:06.6153686 +0000 UTC m=+33.831885201] [host=281998716c79]"
2020/02/22 07:24:06 Reply: time:"[time=2020-02-22 07:24:06.617085 +0000 UTC m=+33.833604101] [host=281998716c79]"
2020/02/22 07:24:06 Reply: time:"[time=2020-02-22 07:24:06.6188973 +0000 UTC m=+33.835414401] [host=281998716c79]"
```

Also explore Makefile for both client and server to get idea of how to build docker containers. 

# Streaming RPCs

The example code (grpc server and client) also supports streaming RPCs. In total it has 4 RPCs. One Unary and 3 streaming RPCs as shown below. For streaming, it has one each for input stream, output stream and bi-directional stream. Read timep.proto to see what's latest and greatest in there.

```
service TimeServer {
  rpc GetTime (TimeRequest) returns (TimeReply) {}
  rpc GetTimeSOut (TimeRequest) returns (stream TimeReply) {}
  rpc GetTimeSIn (stream TimeRequest) returns (TimeReply) {}
  rpc GetTimeSInSOut (stream TimeRequest) returns (stream TimeReply) {}
}
```

- GetTime is Unary
- GetTimeSOut is output streaming
- GetTimeSIn is input stream
- GetTimeSInSOut is bidirectional streaming

To test various APIs, you can you the command line options of the client.

```
❯ ./timeclient --help
Usage of ./timeclient:
  -api int
    	GetTime(1), GetTimeSOut(2), GetTimeSIn(3), GetTimeSInSOut(4) (default 1)
  -cacert string
    	CACert for server (default "CA.crt")
  -count int
    	number of rpc calls (default 1)
  -host string
    	grpc server host:port (default "localhost:50051")
  -insecure
    	connect without TLS
  -servername string
    	CACert for server (default "timeserver.gke.net")
  -sin int
    	number of messages in GetTimeSIn (input stream) rpc calls (default 1)
```

Use 'api' option to specify which rpc you want to test. Also see 'sin' option, you can specify number of streaming messages for streaming RPCs. 
