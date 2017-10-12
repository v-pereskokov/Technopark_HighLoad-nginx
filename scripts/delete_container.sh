docker rm -f container
docker rm $(docker ps -a -q) && docker rmi $(docker images -q)
