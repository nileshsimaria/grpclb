# GRPC LB using NGINX ingress controller (timeserver)

K8S YAML files in github.com/nileshsimaria/grpclb/nginx/time-server

## start nginx pod and expose it's service
```
kubectl apply -f nginx.yaml -f nginx-service.yaml
```

## setup secrets

This timeserver has hard coded cert from github.com/nileshsimaria/grpclb/tls-certs; copy server.crt and server.tls from that directory to github.com/nileshsimaria/grpclb/nginx/time-server

The CN is timeserver.gke.net which is used as host in ingress (virtual hosts)

## Apply secrets
```
$ kubectl create secret tls tls-secret --key server.key --cert server.crt -n ingress-nginx
```

## start grpc server and its service
```
kubectl apply -f app.yaml -f app-service.yaml
```

Note that we run timeserver without TLS. 

## finally create nginx ingress
```
kubectl apply -f ingress.yaml
```

## Verify 
```
❯ kubectl get pods,services,ingress -n ingress-nginx
NAME                                            READY   STATUS    RESTARTS   AGE
pod/nginx-ingress-controller-7fbc8f8d75-lmqvq   1/1     Running   0          5m46s
pod/time-app-85879d5f75-wk5ls                   1/1     Running   0          98s

NAME                    TYPE           CLUSTER-IP      EXTERNAL-IP     PORT(S)                      AGE
service/ingress-nginx   LoadBalancer   10.15.253.117   35.184.58.179   80:31699/TCP,443:32038/TCP   5m46s
service/time-service    ClusterIP      10.15.241.76    <none>          50051/TCP                    97s

NAME                             HOSTS                ADDRESS         PORTS     AGE
ingress.extensions/time-server   timeserver.gke.net   35.184.58.179   80, 443   46s
```

## test

Unary RPC
```
docker run --add-host timeserver.gke.net:35.184.58.179 nileshsimaria/timeclient:v5 --host timeserver.gke.net:443
Calling GetTime 5 times

2020/02/28 23:04:44 Reply: time:"[time=2020-02-28 23:04:44.926750534 +0000 UTC m=+360.069303293] [host=time-app-85879d5f75-wk5ls]"
2020/02/28 23:04:45 Reply: time:"[time=2020-02-28 23:04:44.97004502 +0000 UTC m=+360.112597777] [host=time-app-85879d5f75-wk5ls]"
2020/02/28 23:04:45 Reply: time:"[time=2020-02-28 23:04:45.023142491 +0000 UTC m=+360.165695191] [host=time-app-85879d5f75-wk5ls]"
2020/02/28 23:04:45 Reply: time:"[time=2020-02-28 23:04:45.064576241 +0000 UTC m=+360.207128927] [host=time-app-85879d5f75-wk5ls]"
2020/02/28 23:04:45 Reply: time:"[time=2020-02-28 23:04:45.109688313 +0000 UTC m=+360.252240996] [host=time-app-85879d5f75-wk5ls]"
```

Output Streaming RPC
```
docker run --add-host timeserver.gke.net:35.184.58.179 nileshsimaria/timeclient:v5 --host timeserver.gke.net:443 -api 2

Calling GetTimeSOut
2020/02/28 23:04:59 StreamOut Reply: time:"[#1][time=2020-02-28 23:04:59.862776641 +0000 UTC m=+375.005329355] [host=time-app-85879d5f75-wk5ls]"
2020/02/28 23:04:59 StreamOut Reply: time:"[#2][time=2020-02-28 23:04:59.862969821 +0000 UTC m=+375.005522525] [host=time-app-85879d5f75-wk5ls]"
2020/02/28 23:04:59 StreamOut Reply: time:"[#3][time=2020-02-28 23:04:59.862996266 +0000 UTC m=+375.005548962] [host=time-app-85879d5f75-wk5ls]"
2020/02/28 23:04:59 StreamOut Reply: time:"[#4][time=2020-02-28 23:04:59.863013875 +0000 UTC m=+375.005566570] [host=time-app-85879d5f75-wk5ls]"
2020/02/28 23:04:59 StreamOut Reply: time:"[#5][time=2020-02-28 23:04:59.863052506 +0000 UTC m=+375.005605198] [host=time-app-85879d5f75-wk5ls]"
2020/02/28 23:04:59 StreamOut Reply: time:"[#6][time=2020-02-28 23:04:59.863068667 +0000 UTC m=+375.005621362] [host=time-app-85879d5f75-wk5ls]"
2020/02/28 23:04:59 StreamOut Reply: time:"[#7][time=2020-02-28 23:04:59.863085387 +0000 UTC m=+375.005638083] [host=time-app-85879d5f75-wk5ls]"
2020/02/28 23:04:59 StreamOut Reply: time:"[#8][time=2020-02-28 23:04:59.863118304 +0000 UTC m=+375.005670991] [host=time-app-85879d5f75-wk5ls]"
2020/02/28 23:04:59 StreamOut Reply: time:"[#9][time=2020-02-28 23:04:59.863136548 +0000 UTC m=+375.005689243] [host=time-app-85879d5f75-wk5ls]"
2020/02/28 23:04:59 StreamOut Reply: time:"[#10][time=2020-02-28 23:04:59.863146168 +0000 UTC m=+375.005698854] [host=time-app-85879d5f75-wk5ls]"
2020/02/28 23:04:59 Done
2020/02/28 23:04:59 EOF
```

Output Streaming RPC
```
docker run --add-host timeserver.gke.net:35.184.58.179 nileshsimaria/timeclient:v5 --host timeserver.gke.net:443 -api 3 -sin 500 -count 1

Calling GetTimeSIn
2020/02/28 23:05:19 [#0]StreamIn Reply: time:"[InStreamMsg count #500][time=2020-02-28 23:05:19.330861511 +0000 UTC m=+394.473414235] [host=time-app-85879d5f75-wk5ls]"
```

Bi-directional RPC
```
docker run --add-host timeserver.gke.net:35.184.58.179 nileshsimaria/timeclient:v5 --host timeserver.gke.net:443 -api 4 -sin 3000 -count 1

2020/02/28 23:04:21 [#0]StreamInOUT Reply: time:"[InOUTStreamMsg count #1000][time=2020-02-28 23:04:21.839572261 +0000 UTC m=+336.982124991] [host=time-app-85879d5f75-wk5ls]"
2020/02/28 23:04:21 [#0]StreamInOUT Reply: time:"[InOUTStreamMsg count #2000][time=2020-02-28 23:04:21.875763702 +0000 UTC m=+337.018316433] [host=time-app-85879d5f75-wk5ls]"
2020/02/28 23:04:21 [#0]StreamInOUT Reply: time:"[InOUTStreamMsg count #3000][time=2020-02-28 23:04:21.876556591 +0000 UTC m=+337.019109300] [host=time-app-85879d5f75-wk5ls]"
```

