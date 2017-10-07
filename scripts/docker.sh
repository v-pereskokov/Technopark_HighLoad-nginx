docker rm -f $1
docker build -t vlados .
docker run -p 80:80 -c 4 --name $1 -t vlados
