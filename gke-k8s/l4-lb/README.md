# Setup and Test L4 GRPC LB on GKE cluster

## Create 3 node cluster
```
$ gcloud container  clusters create grpclb-cluster --zone us-central1-a  --num-nodes 3 --enable-ip-alias
NAME            LOCATION       MASTER_VERSION  MASTER_IP      MACHINE_TYPE   NODE_VERSION    NUM_NODES  STATUS
grpclb-cluster  us-central1-a  1.14.10-gke.17  34.68.140.122  n1-standard-1  1.14.10-gke.17  3          RUNNING
```

## To test load balancer service, allow tcp port 50051
```
$ gcloud compute firewall-rules create grpc-lb-fw --allow tcp:50051
```

## Query nodes
```
$ kubectl get nodes
NAME                                            STATUS   ROLES    AGE     VERSION
gke-grpclb-cluster-default-pool-09ead34e-8s0c   Ready    <none>   6m11s   v1.14.10-gke.17
gke-grpclb-cluster-default-pool-09ead34e-b08r   Ready    <none>   6m10s   v1.14.10-gke.17
gke-grpclb-cluster-default-pool-09ead34e-qvpx   Ready    <none>   6m11s   v1.14.10-gke.17
```

## Deploy servers (github.com/nileshsimaria/grpclb/gke-k8s/l4-lb)
```
$ kubectl apply -f server-deployment.yaml
```

## Query pods (we set replica to 2)
```
$ kubectl get pods
NAME                                 READY   STATUS    RESTARTS   AGE
grpclb-deployment-844c7c8b6b-2dvfs   1/1     Running   0          14m
grpclb-deployment-844c7c8b6b-4q2ws   1/1     Running   0          14m
```

## Create external load balancing service
```
$ kubectl apply -f server-lb-service.yaml
```

## Query our service
```
$ kubectl get svc
NAME                        TYPE           CLUSTER-IP      EXTERNAL-IP     PORT(S)          AGE
grpclb-deployment-service   LoadBalancer   10.15.241.197   34.69.243.120   8080:32660/TCP   12m
kubernetes                  ClusterIP      10.15.240.1     <none>          443/TCP          46m
```

## Test it using the client (only L4 LB is happening)
```
$ docker run nileshsimaria/timeclient:v1 --host 34.69.243.120:8080 --count 3
2020/02/21 19:54:24 Reply: time:"[time=2020-02-21 19:54:24.378684874 +0000 UTC m=+1206.758784680] [host=grpclb-deployment-844c7c8b6b-4q2ws]"
2020/02/21 19:54:24 Reply: time:"[time=2020-02-21 19:54:24.433850633 +0000 UTC m=+1206.813950486] [host=grpclb-deployment-844c7c8b6b-4q2ws]"
2020/02/21 19:54:24 Reply: time:"[time=2020-02-21 19:54:24.489856997 +0000 UTC m=+1206.869956853] [host=grpclb-deployment-844c7c8b6b-4q2ws]"
❯ docker run nileshsimaria/timeclient:v1 --host 34.69.243.120:8080 --count 3
2020/02/21 19:54:26 Reply: time:"[time=2020-02-21 19:54:26.578546293 +0000 UTC m=+1208.949625086] [host=grpclb-deployment-844c7c8b6b-2dvfs]"
2020/02/21 19:54:26 Reply: time:"[time=2020-02-21 19:54:26.633925254 +0000 UTC m=+1209.005003967] [host=grpclb-deployment-844c7c8b6b-2dvfs]"
2020/02/21 19:54:26 Reply: time:"[time=2020-02-21 19:54:26.689948661 +0000 UTC m=+1209.061027417] [host=grpclb-deployment-844c7c8b6b-2dvfs]"
❯ docker run nileshsimaria/timeclient:v1 --host 34.69.243.120:8080 --count 3
2020/02/21 19:54:28 Reply: time:"[time=2020-02-21 19:54:28.707188141 +0000 UTC m=+1211.078266942] [host=grpclb-deployment-844c7c8b6b-2dvfs]"
2020/02/21 19:54:28 Reply: time:"[time=2020-02-21 19:54:28.76415319 +0000 UTC m=+1211.135231980] [host=grpclb-deployment-844c7c8b6b-2dvfs]"
2020/02/21 19:54:28 Reply: time:"[time=2020-02-21 19:54:28.819968313 +0000 UTC m=+1211.191047098] [host=grpclb-deployment-844c7c8b6b-2dvfs]"
❯ docker run nileshsimaria/timeclient:v1 --host 34.69.243.120:8080 --count 3
2020/02/21 19:54:30 Reply: time:"[time=2020-02-21 19:54:30.820557218 +0000 UTC m=+1213.200657097] [host=grpclb-deployment-844c7c8b6b-4q2ws]"
2020/02/21 19:54:30 Reply: time:"[time=2020-02-21 19:54:30.876588625 +0000 UTC m=+1213.256688801] [host=grpclb-deployment-844c7c8b6b-4q2ws]"
2020/02/21 19:54:31 Reply: time:"[time=2020-02-21 19:54:30.931812946 +0000 UTC m=+1213.311912829] [host=grpclb-deployment-844c7c8b6b-4q2ws]”
```

## Clean up - delete service, deployment and finally the cluster to avoid charging from google cloud
```
$ kubectl delete service grpclb-deployment-service
$ kubectl delete deployment grpclb-deployment
$ gcloud container clusters delete grpclb-cluster
```
