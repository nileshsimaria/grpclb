# grpclb
GRPC Load Balancing 

To see how to run example code locally on your laptop and using docker follow
https://github.com/nileshsimaria/grpclb/blob/master/example-code/README.md

The server implements unary RPC as of now. The client makes number of RPC calls on the same gRPC channel to test L4 and L7 load balancing. The server reply with the time and hostname. 
