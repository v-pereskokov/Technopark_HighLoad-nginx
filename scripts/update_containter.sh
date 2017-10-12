docker rm -f container
docker build -t vladddos .
docker run -p 80:80 -c 4 --name container -t vladddos
