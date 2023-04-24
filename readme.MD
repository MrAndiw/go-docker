# HOW TO CREATE & DEPLOY IMAGE TO DOCKER | Golang

1. Build Script Dockerfile
   1.2. run : docker build --tag go-docker .
2. Create Network
   2.1 run : docker network create tulip-net
3. Run Docker Container Golang
   3.1 run : docker run -d -p 8080:8080 --name go-docker --network go-docker-network go-docker
4. Connect existing container to network (posgres)
   4.1 run : docker network connect go-docker-network [containerID]
