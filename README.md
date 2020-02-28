# example server and client code

To see how to run example code locally on your laptop and using docker
- https://github.com/nileshsimaria/grpclb/blob/master/example-code/README.md

In v1, the server implements unary RPC and in v5 it supports unary, input stream, output stream and bi-directional streaming RPC calls. The client makes number of RPC calls on the same gRPC channel to test L4 and L7 load balancing. The server reply with the time and hostname. 


# GRPC LB without NGINX (using native google cloud constructs)

How to achieve and test L4 LB on GKE (without nginx)
- https://github.com/nileshsimaria/grpclb/blob/master/gke-k8s/l4-lb/README.md

How to achieve and test L7 LB on GKE (without nginx)
- https://github.com/nileshsimaria/grpclb/blob/master/gke-k8s/l7-lb/README.md


# GRPC LB with NGINX

## time server (nileshsimaria/grpclb/example-code/timeserver)
https://github.com/nileshsimaria/grpclb/tree/master/nginx/time-server

## fortune-teller server
https://github.com/nileshsimaria/grpclb/tree/master/nginx/fortune-teller