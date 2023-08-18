clean: 
	rm -rf gokube

network:
	docker network create gokube-network

build:
	docker build --no-cache -t ghcr.io/techierishi/gokube:latest .

run:
	docker run --privileged --name gokube-container -d \
        --network gokube-network --network-alias docker \
        -e DOCKER_TLS_CERTDIR=/certs \
        -p 5555:5555 \
        -p 5556:5556 \
        -v gokube-network-certs-ca:/certs/ca \
        -v gokube-network-certs-client:/certs/client \
        ghcr.io/techierishi/gokube:latest

exec:
	docker exec -it gokube-container /bin/sh

panes:
	tmux split-window -h
	tmux split-window -v
