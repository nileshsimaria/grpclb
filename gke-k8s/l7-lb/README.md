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


# L7 only tests with streaming RPCs

## Start the cluster

```
$ gcloud container  clusters create grpclb-cluster --zone us-central1-a  --num-nodes 3 --enable-ip-alias
NAME            LOCATION       MASTER_VERSION  MASTER_IP     MACHINE_TYPE   NODE_VERSION    NUM_NODES  STATUS
grpclb-cluster  us-central1-a  1.14.10-gke.17  34.66.66.136  n1-standard-1  1.14.10-gke.17  3          RUNNING
```

Its important to use --enable-ip-alias to make VPC-native clusters in GKE.

## Connect to the newly created cluster and check the status of 3 nodes
```
$ gcloud container clusters get-credentials grpclb-cluster --zone us-central1-a --project gke-learn-1
Fetching cluster endpoint and auth data.
kubeconfig entry generated for grpclb-cluster.

$ kubectl get nodes
NAME                                            STATUS   ROLES    AGE     VERSION
gke-grpclb-cluster-default-pool-4429c7a9-3z5m   Ready    <none>   3m54s   v1.14.10-gke.17
gke-grpclb-cluster-default-pool-4429c7a9-8sx7   Ready    <none>   3m45s   v1.14.10-gke.17
gke-grpclb-cluster-default-pool-4429c7a9-ft4g   Ready    <none>   3m44s   v1.14.10-gke.17
```

For this test, we will use yaml files from github.com/nileshsimaria/grpclb/gke-k8s/l7-lb directory.

## Deploy the service ( 3 replicas, using v5 image nileshsimaria/timeserver:v5)
```
$ kubectl apply -f server-deployment.yaml
deployment.apps/grpclb-deployment created

$ kubectl get deployments
NAME                READY   UP-TO-DATE   AVAILABLE   AGE
grpclb-deployment   3/3     3            3           55s

$ kubectl get pods
NAME                                 READY   STATUS    RESTARTS   AGE
grpclb-deployment-679bd58666-glbtf   1/1     Running   0          60s
grpclb-deployment-679bd58666-j8c6p   1/1     Running   0          60s
grpclb-deployment-679bd58666-v4tl8   1/1     Running   0          60s
```

## Apply secrets
```
kubectl apply -f server-secret.yaml
secret/server-secret created
```

We need secret because we are using SSL/TLS connections from client to ingress and from ingress to our service.

### Expose our server (NodePort service)

```
$ cat server-ingress-service.yaml
---
apiVersion: v1
kind: Service
metadata:
  name: server-ingress-service
  labels:
    type: server-ingress-service
  annotations:
    service.alpha.kubernetes.io/app-protocols: '{"servport":"HTTP2"}'
    cloud.google.com/neg: '{"ingress": true}'
spec:
  type: NodePort
  ports:
  - name: servport
    port: 50051
    protocol: TCP
    targetPort: 50051
  selector:
    app: servers


$ kubectl apply -f server-ingress-service.yaml    
```

- service.alpha.kubernetes.io/app-protocols set the protocol for communication between LB and the server or app
- cloud.google.com/neg to instruct that the LB should use network endpoint group


Now check our newly created service
```
$ kubectl get services
NAME                     TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)           AGE
kubernetes               ClusterIP   10.15.240.1     <none>        443/TCP           13m
server-ingress-service   NodePort    10.15.240.173   <none>        50051:32264/TCP   5m14s
```


## Create ingress
```
$ cat server-ingress.yaml
apiVersion: extensions/v1beta1
kind: Ingress
metadata:
  name: server-ingress
  annotations:
    kubernetes.io/ingress.allow-http: "false"
spec:
  tls:
  - hosts:
    - timeserver.gke.net
    secretName: server-secret
  backend:
    serviceName: server-ingress-service
    servicePort: 50051

$ kubectl apply -f server-ingress.yaml
```

Here we are specifying the backend service name which is ofcourse our server-ingress-service. Also configure TLS, remember we have created secret server-secret earlier ?

Now check our ingress.

```
$ kubectl get ingress
NAME             HOSTS   ADDRESS         PORTS     AGE
server-ingress   *       34.107.247.95   80, 443   7h40m```
```

Use this 130.211.11.169:443 to talk to ingress over TLS. This should do L7 LB even for streaming RPCs. Please note that we are not at all creating normal L4 load balancing service for this test. It takes some time for the GKE to setup everything for ingress service. Wait for max 10 minutes before you give up :-) Meanwhile, go to console.cloud.google.com and check 'Services & Ingress' under GKE. Once ingress has been created, use 34.107.247.95:443 to make RPC calls. 

## Test output streaming API

GetTimeSOut is output streaming API. The server sends 10 messages before closing the stream. 
```
./timeclient --host 34.107.247.95:443  -api 2
Calling GetTimeSOut
2020/02/22 22:43:44 StreamOut Reply: time:"[#1][time=2020-02-23 06:43:44.829904304 +0000 UTC m=+27782.203799410] [host=grpclb-deployment-679bd58666-njxn4]"
2020/02/22 22:43:44 StreamOut Reply: time:"[#2][time=2020-02-23 06:43:44.830111149 +0000 UTC m=+27782.204006256] [host=grpclb-deployment-679bd58666-njxn4]"
2020/02/22 22:43:44 StreamOut Reply: time:"[#3][time=2020-02-23 06:43:44.83013713 +0000 UTC m=+27782.204032232] [host=grpclb-deployment-679bd58666-njxn4]"
2020/02/22 22:43:44 StreamOut Reply: time:"[#4][time=2020-02-23 06:43:44.830157711 +0000 UTC m=+27782.204052820] [host=grpclb-deployment-679bd58666-njxn4]"
2020/02/22 22:43:44 StreamOut Reply: time:"[#5][time=2020-02-23 06:43:44.830177818 +0000 UTC m=+27782.204072922] [host=grpclb-deployment-679bd58666-njxn4]"
2020/02/22 22:43:44 StreamOut Reply: time:"[#6][time=2020-02-23 06:43:44.83019955 +0000 UTC m=+27782.204094655] [host=grpclb-deployment-679bd58666-njxn4]"
2020/02/22 22:43:44 StreamOut Reply: time:"[#7][time=2020-02-23 06:43:44.830218903 +0000 UTC m=+27782.204114008] [host=grpclb-deployment-679bd58666-njxn4]"
2020/02/22 22:43:44 StreamOut Reply: time:"[#8][time=2020-02-23 06:43:44.830251277 +0000 UTC m=+27782.204146375] [host=grpclb-deployment-679bd58666-njxn4]"
2020/02/22 22:43:44 StreamOut Reply: time:"[#9][time=2020-02-23 06:43:44.830281499 +0000 UTC m=+27782.204176606] [host=grpclb-deployment-679bd58666-njxn4]"
2020/02/22 22:43:44 StreamOut Reply: time:"[#10][time=2020-02-23 06:43:44.830300743 +0000 UTC m=+27782.204195848] [host=grpclb-deployment-679bd58666-njxn4]"
2020/02/22 22:43:44 Done
2020/02/22 22:43:44 EOF
```
## Test input streaming API

GetTimeSIn is input streaming API. It sends 500 messages in this exaple run using '-sin' option. The server counts the messages fromt he client and add the number of messages in its reply. 

```
❯ ./timeclient --host 34.107.247.95:443  -api 3 -sin 500
Calling GetTimeSIn
2020/02/22 22:45:03 [#0]StreamIn Reply: time:"[InStreamMsg count #500][time=2020-02-23 06:45:03.918277913 +0000 UTC m=+27861.552974592] [host=grpclb-deployment-679bd58666-blhfx]"
```
## Test bi-directional streaming API

use '-api 4' to test it. Use '-sin' to send number of messages from the client. Server sends messages on every 1000 input messages.

```
$ ./timeclient --host 34.107.247.95:443  -api 4 -sin 3000
Calling GetTimeSInSOut
2020/02/22 22:49:25 [#0]StreamInOUT Reply: time:"[InOUTStreamMsg count #1000][time=2020-02-23 06:49:25.682510813 +0000 UTC m=+28123.317207491] [host=grpclb-deployment-679bd58666-blhfx]"
2020/02/22 22:49:25 [#0]StreamInOUT Reply: time:"[InOUTStreamMsg count #2000][time=2020-02-23 06:49:25.698705076 +0000 UTC m=+28123.333401753] [host=grpclb-deployment-679bd58666-blhfx]"
2020/02/22 22:49:25 [#0]StreamInOUT Reply: time:"[InOUTStreamMsg count #3000][time=2020-02-23 06:49:25.711756286 +0000 UTC m=+28123.346452953] [host=grpclb-deployment-679bd58666-blhfx]"
```

## Summary

Create a deployment (pods), expose them via NodePort service and create an ingress service which points to NodePort service

```
$ kubectl get pods,services,ingress
NAME                                     READY   STATUS    RESTARTS   AGE
pod/grpclb-deployment-679bd58666-5gp47   1/1     Running   0          7h41m
pod/grpclb-deployment-679bd58666-blhfx   1/1     Running   0          7h41m
pod/grpclb-deployment-679bd58666-njxn4   1/1     Running   0          7h41m

NAME                             TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)           AGE
service/kubernetes               ClusterIP   10.77.0.1    <none>        443/TCP           7h43m
service/server-ingress-service   NodePort    10.77.5.88   <none>        50051:31229/TCP   7h40m

NAME                                HOSTS   ADDRESS         PORTS     AGE
ingress.extensions/server-ingress   *       34.107.247.95   80, 443   7h40m
```





