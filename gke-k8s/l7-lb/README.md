```
# yaml files in github.com/nileshsimaria/grpclb/gke-k8s/l7-lb

# deploy both l7 and l4 lb files
$ kubectl apply -f server-deployment.yaml -f server-ingress-service.yaml -f server-ingress.yaml -f server-lb-service.yaml -f server-secret.yaml
deployment.apps/grpclb-deployment created
service/server-ingress-service created
ingress.extensions/server-ingress created
service/server-lb-service created
secret/server-secret unchanged

Note := server-lb-service.yaml is just l4 lb. For l7 lb, the service is nodePort (server-ingress-service.yaml) and the ingress controller which is server-ingress.yaml

# find LB and Ingress IP
$ kubectl get pods
NAME                                 READY   STATUS    RESTARTS   AGE
grpclb-deployment-84fc8f9449-5wlmb   1/1     Running   0          2m2s
grpclb-deployment-84fc8f9449-9kzmx   1/1     Running   0          2m1s
grpclb-deployment-84fc8f9449-p6lqc   1/1     Running   0          2m1s

$ kubectl get services
NAME                     TYPE           CLUSTER-IP    EXTERNAL-IP     PORT(S)           AGE
kubernetes               ClusterIP      10.0.0.1      <none>          443/TCP           8h
server-ingress-service   NodePort       10.0.0.210    <none>          50051:31299/TCP   2m7s
server-lb-service        LoadBalancer   10.0.10.123   35.202.224.46   50051:30140/TCP   2m7s

$ kubectl get ingress
NAME             HOSTS   ADDRESS         PORTS     AGE
server-ingress   *       34.107.247.95   80, 443   2m12s

As you can see use IP 35.202.224.46 from server-lb-service for l4 LB and use IP 34.107.247.9 from server-ingress for l7 LB.

# test L4 LB
$ docker run nileshsimaria/timeclient:v4 --host 35.202.224.46:50051  --count 20

# test L7 LB
$ docker run nileshsimaria/timeclient:v4 --host 34.107.247.95:443  --count 20
2020/02/22 07:08:12 Reply: time:"[time=2020-02-22 07:08:12.887172982 +0000 UTC m=+312.813041597] [host=grpclb-deployment-84fc8f9449-5wlmb]"
2020/02/22 07:08:13 Reply: time:"[time=2020-02-22 07:08:12.951555165 +0000 UTC m=+312.323720352] [host=grpclb-deployment-84fc8f9449-9kzmx]"
2020/02/22 07:08:13 Reply: time:"[time=2020-02-22 07:08:13.014906063 +0000 UTC m=+312.940774730] [host=grpclb-deployment-84fc8f9449-5wlmb]"
2020/02/22 07:08:13 Reply: time:"[time=2020-02-22 07:08:13.079048667 +0000 UTC m=+313.004917264] [host=grpclb-deployment-84fc8f9449-5wlmb]"
2020/02/22 07:08:13 Reply: time:"[time=2020-02-22 07:08:13.145606731 +0000 UTC m=+313.071475326] [host=grpclb-deployment-84fc8f9449-5wlmb]"
2020/02/22 07:08:13 Reply: time:"[time=2020-02-22 07:08:13.207100121 +0000 UTC m=+313.132968758] [host=grpclb-deployment-84fc8f9449-5wlmb]"
2020/02/22 07:08:13 Reply: time:"[time=2020-02-22 07:08:13.266674541 +0000 UTC m=+312.638839730] [host=grpclb-deployment-84fc8f9449-9kzmx]"
2020/02/22 07:08:13 Reply: time:"[time=2020-02-22 07:08:13.330297297 +0000 UTC m=+313.256165922] [host=grpclb-deployment-84fc8f9449-5wlmb]"
2020/02/22 07:08:13 Reply: time:"[time=2020-02-22 07:08:13.389727639 +0000 UTC m=+312.761892884] [host=grpclb-deployment-84fc8f9449-9kzmx]"
2020/02/22 07:08:13 Reply: time:"[time=2020-02-22 07:08:13.463603308 +0000 UTC m=+312.835768520] [host=grpclb-deployment-84fc8f9449-9kzmx]"
2020/02/22 07:08:13 Reply: time:"[time=2020-02-22 07:08:13.530236791 +0000 UTC m=+314.924982655] [host=grpclb-deployment-84fc8f9449-p6lqc]"
2020/02/22 07:08:13 Reply: time:"[time=2020-02-22 07:08:13.587296473 +0000 UTC m=+314.982042320] [host=grpclb-deployment-84fc8f9449-p6lqc]"
2020/02/22 07:08:13 Reply: time:"[time=2020-02-22 07:08:13.650779652 +0000 UTC m=+313.022944857] [host=grpclb-deployment-84fc8f9449-9kzmx]"
2020/02/22 07:08:13 Reply: time:"[time=2020-02-22 07:08:13.719150517 +0000 UTC m=+313.091315704] [host=grpclb-deployment-84fc8f9449-9kzmx]"
2020/02/22 07:08:13 Reply: time:"[time=2020-02-22 07:08:13.774724343 +0000 UTC m=+313.146889527] [host=grpclb-deployment-84fc8f9449-9kzmx]"
2020/02/22 07:08:13 Reply: time:"[time=2020-02-22 07:08:13.833666822 +0000 UTC m=+315.228412714] [host=grpclb-deployment-84fc8f9449-p6lqc]"
2020/02/22 07:08:13 Reply: time:"[time=2020-02-22 07:08:13.894665852 +0000 UTC m=+313.820534465] [host=grpclb-deployment-84fc8f9449-5wlmb]"
2020/02/22 07:08:14 Reply: time:"[time=2020-02-22 07:08:13.983773853 +0000 UTC m=+315.378519695] [host=grpclb-deployment-84fc8f9449-p6lqc]"
2020/02/22 07:08:14 Reply: time:"[time=2020-02-22 07:08:14.044909226 +0000 UTC m=+313.417074419] [host=grpclb-deployment-84fc8f9449-9kzmx]"
2020/02/22 07:08:14 Reply: time:"[time=2020-02-22 07:08:14.111752157 +0000 UTC m=+315.506497996] [host=grpclb-deployment-84fc8f9449-p6lqc]”

Note :- in this test, I have used TLS.
```
