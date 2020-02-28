# grpclb (GKE native)
GRPC Load Balancing 

To see how to run example code locally on your laptop and using docker
- https://github.com/nileshsimaria/grpclb/blob/master/example-code/README.md

The server implements unary RPC as of now. The client makes number of RPC calls on the same gRPC channel to test L4 and L7 load balancing. The server reply with the time and hostname. 

How to achieve and test L4 LB on GKE
- https://github.com/nileshsimaria/grpclb/blob/master/gke-k8s/l4-lb/README.md

How to achieve and test L7 LB on GKE
- https://github.com/nileshsimaria/grpclb/blob/master/gke-k8s/l7-lb/README.md

# grpclb (NGINX on GKE)
https://github.com/nileshsimaria/grpclb/tree/master/nginx/fortune-teller
