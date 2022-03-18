# Docker Registry Queryer
Build: `go build -o docker-query main.go`

Usage:
```bash
./docker-query -h
# Usage of ./docker-query:
#   -addr string
#         docker registry endpoint
```

Example run:
```bash
./docker-query -addr 172.17.0.2:5000
2022/03/18 14:58:09 succeed: http://172.17.0.2:5000/v2/busybox/tags/list
 ###########################################
Available images:
172.17.0.2:5000/busybox:latest
```
