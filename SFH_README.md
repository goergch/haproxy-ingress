# HOW TO BUILD THE IMAGE
```bash
make docker-build
docker tag docker.io/goergch/haproxy-ingress:latest tw552sfhcontainerregistry.azurecr.io/edge/haproxy-ingress:0.14.2.sfh.0
docker push tw552sfhcontainerregistry.azurecr.io/edge/haproxy-ingress:0.14.2.sfh.0
```