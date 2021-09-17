API_IMAGE_NAME = goal-api
API_CONTAINER_NAME = goal-api-service

ADMIN_IMAGE_NAME = goal-admin
ADMIN_CONTAINER_NAME = goal-admin-service

build-api:
	docker build -t $(API_IMAGE_NAME) .

build-admin:
	docker build -t $(ADMIN_IMAGE_NAME) .

run-api:
	docker run -d --env ENV_NAME=.env --env APP_SERVICE=api --name $(API_CONTAINER_NAME) -p 8080:8080 $(API_IMAGE_NAME)

run-admin:
	docker run -d --env ENV_NAME=.env --env APP_SERVICE=admin --name $(ADMIN_CONTAINER_NAME) -p 8081:8080 $(ADMIN_IMAGE_NAME)

run-api-local:
	docker run -d --network backend_network --env ENV_NAME=.env --env APP_SERVICE=api --name $(API_CONTAINER_NAME) -p 8080:8080 $(API_IMAGE_NAME)

run-admin-local:
	docker run -d --network backend_network --env ENV_NAME=.env --env APP_SERVICE=admin --name $(ADMIN_CONTAINER_NAME) -p 8081:8080 $(ADMIN_IMAGE_NAME)

clean:
	docker ps -a | grep $(API_IMAGE_NAME) | awk  '{print $$1}' | xargs docker stop
	docker ps -a | grep $(API_IMAGE_NAME) | awk  '{print $$1}' | xargs docker rm
	docker ps -a | grep $(ADMIN_IMAGE_NAME) | awk  '{print $$1}' | xargs docker stop
	docker ps -a | grep $(ADMIN_IMAGE_NAME) | awk  '{print $$1}' | xargs docker rm