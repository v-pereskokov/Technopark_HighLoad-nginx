docker rm -f container
docker build -t vladdos .
docker run -p 80:80 -c 4 --name container -t vladdos
