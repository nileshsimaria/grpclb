# GRPC LB using NGINX ingress controller

## Let's begin with creating secrets.
```
$ openssl req -x509 -sha256 -nodes -days 365 -newkey rsa:2048 -keyout tls.key -out tls.crt -subj "/CN=fortune-teller.stack.build/O=fortune-teller.stack.build”
Generating a RSA private key
...............................................+++++
.......................................+++++
writing new private key to 'tls.key'
-----
```
## Apply secrets
```
$ kubectl create secret tls tls-secret --key tls.key --cert tls.crt -n ingress-nginx
```
## start nginx and expose it 
```
kubectl apply -f nginx.yaml
kubectl apply -f nginx-service.yaml
```

## start grpc server and its service
```
kubectl apply -f app.yaml -f app-service.yaml
```

## finally create nginx ingress
```
kubectl apply -f ingress.yaml
```

## Verify 
```
kubectl get pods,services,ingress -n ingress-nginx
NAME                                            READY   STATUS    RESTARTS   AGE
pod/fortune-teller-app-669f4bcbf8-zpvpn         1/1     Running   0          72s
pod/nginx-ingress-controller-7fbc8f8d75-z5p7d   1/1     Running   0          2m9s


NAME                             TYPE           CLUSTER-IP      EXTERNAL-IP     PORT(S)                      AGE
service/fortune-teller-service   ClusterIP      10.27.249.234   <none>          50051/TCP                    71s
service/ingress-nginx            LoadBalancer   10.27.253.38    35.238.50.131   80:30763/TCP,443:32388/TCP   2m4s


NAME                                 HOSTS                        ADDRESS         PORTS     AGE
ingress.extensions/fortune-ingress   fortune-teller.stack.build   35.238.50.131   80, 443   45s
```

## update /etc/hosts
```
35.238.50.131   fortune-teller.stack.build
```

## test
```
grpcurl --insecure fortune-teller.stack.build:443 build.stack.fortune.FortuneTeller/Predict
{
  "message": "Come live with me, and be my love,\nAnd we will some new pleasures prove\nOf golden sands, and crystal brooks,\nWith silken lines, and silver hooks.\n\t\t-- John Donne"
}
```
