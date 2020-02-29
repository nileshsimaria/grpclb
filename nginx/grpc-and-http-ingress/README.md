# One ingress controller - two ingress resource (grpc and http)

Let's see how to setup multiple ingress resources with just one ingress controller. This is a practical use case of HB where we will need http ingress for http api calls and grpc ingress for telemetry. 

Similarly, we will need UDP ingest for SNMP, flow, Native GPB telemetry and TCP ingress for outbound netconf. In this example, I am focusing only on http and grpc.

## GRPC

Follow all the instruction to setup grpc ingress from https://github.com/nileshsimaria/grpclb/tree/master/nginx/time-server

## HTTP

See K8S yaml files in github.com/nileshsimaria/grpclb/nginx/grpc-and-http-ingress

In this directory you will see web-deployment.yaml to start a web app. web-service.yaml to expose the app and web-ingress.yaml for ingress resource rule. Note that we will have only one nginx controlle running with two ingress resources. 

Before deploying web-*.yaml files, just see what we have go so far.

```
❯ kubectl get pods,svc,ing -n ingress-nginx
NAME                                            READY   STATUS    RESTARTS   AGE
pod/nginx-ingress-controller-7fbc8f8d75-lmqvq   1/1     Running   0          98m
pod/time-app-85879d5f75-wk5ls                   1/1     Running   0          94m

NAME                    TYPE           CLUSTER-IP      EXTERNAL-IP     PORT(S)                      AGE
service/ingress-nginx   LoadBalancer   10.15.253.117   35.184.58.179   80:31699/TCP,443:32038/TCP   98m
service/time-service    ClusterIP      10.15.241.76    <none>          50051/TCP                    94m

NAME                             HOSTS                ADDRESS         PORTS     AGE
ingress.extensions/time-server   timeserver.gke.net   35.184.58.179   80, 443   93m
```

Make sure GRPC is working

```
❯ docker run --add-host timeserver.gke.net:35.184.58.179 nileshsimaria/timeclient:v5 --host timeserver.gke.net:443 -api 3 -sin 500 -count 1

Calling GetTimeSIn
2020/02/29 00:33:12 [#0]StreamIn Reply: time:"[InStreamMsg count #500][time=2020-02-29 00:33:12.744153672 +0000 UTC m=+5667.886706450] [host=time-app-85879d5f75-wk5ls]"
```


Now try to make HTTP request to LB IP. 

```
❯ curl 35.184.58.179
<html>
<head><title>404 Not Found</title></head>
<body>
<center><h1>404 Not Found</h1></center>
<hr><center>nginx/1.17.8</center>
</body>
</html>
```

Above is expected as we have not started web-*.yaml so it goes to the default nginx handler which is suppose to give back 404.

## Deploy web app, expose its service and create ingress resource for the same

```
kubectl apply -f web-deployment.yaml -f web-service.yaml -f web-ingress.yaml
```

## See what you have got after starting web stuff

```
❯ kubectl get pods,svc,ing -n ingress-nginx
NAME                                            READY   STATUS    RESTARTS   AGE
pod/nginx-ingress-controller-7fbc8f8d75-lmqvq   1/1     Running   0          100m
pod/time-app-85879d5f75-wk5ls                   1/1     Running   0          96m
pod/web-app-bddb8c597-786r4                     1/1     Running   0          35s

NAME                    TYPE           CLUSTER-IP      EXTERNAL-IP     PORT(S)                      AGE
service/ingress-nginx   LoadBalancer   10.15.253.117   35.184.58.179   80:31699/TCP,443:32038/TCP   100m
service/time-service    ClusterIP      10.15.241.76    <none>          50051/TCP                    96m
service/web-service     ClusterIP      10.15.251.169   <none>          8080/TCP                     35s

NAME                             HOSTS                ADDRESS         PORTS     AGE
ingress.extensions/time-server   timeserver.gke.net   35.184.58.179   80, 443   95m
ingress.extensions/web-ingress   *                    35.184.58.179   80        34s
```


## now test web using LB IP

```
❯ curl 35.184.58.179
Hello, world!
Version: 1.0.0
Hostname: web-app-bddb8c597-786r4
```

That means using one LB IP, we are able to do both HTTP and GRPC. 

## Delete web-ingress resource

```
❯ kubectl delete ingress web-ingress -n ingress-nginx
ingress.extensions "web-ingress" deleted
```

## Now test http again (it should give you 404)

```
❯ curl 35.184.58.179
<html>
<head><title>404 Not Found</title></head>
<body>
<center><h1>404 Not Found</h1></center>
<hr><center>nginx/1.17.8</center>
</body>
</html>
```
