
NAME = local-cluster

kind: create build load deploy

build:
	@echo "Build Images using Docker Compose"
	docker compose build 

up:
	@echo "Start Docker Compose"
	docker compose up -d 

down:
	@echo "Stop Docker Compose"
	docker compose down

create:
	@echo "Creating Kind Cluster"
	kind create cluster --name ${NAME} --config kind-cluster-config.yml

destroy:
	@echo "Delete Kind Cluster"
	kind delete cluster --name ${NAME}

load:
	@echo "Load Docker imges into Kind Cluster"
	kind load docker-image --name ${NAME} frontend
	kind load docker-image --name ${NAME} go-server
	kind load docker-image --name ${NAME} node-server
	kind load docker-image --name ${NAME} python-server

deploy:
	@echo "Deploying to Kind Cluster"
	kubectl apply -f deployment

