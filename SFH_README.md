# HOW TO BUILD THE IMAGE
```bash
make docker-build
docker tag docker.io/goergch/haproxy-ingress:latest sfhcontainerregistry.azurecr.io/edge/haproxy-ingress:0.14.8.sfh.0
docker push sfhcontainerregistry.azurecr.io/edge/haproxy-ingress:0.14.8.sfh.0
```