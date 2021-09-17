IMAGE_NAME = goal
API_CONTAINER_NAME = goal-api-service
ADMIN_CONTAINER_NAME = goal-admin-service

build:
	docker build -t $(IMAGE_NAME) .

run-api:
	docker run -d --env ENV_NAME=.env --env APP_SERVICE=api --name $(API_CONTAINER_NAME) -p 8080:8080 $(IMAGE_NAME)

run-admin:
	docker run -d --env ENV_NAME=.env --env APP_SERVICE=admin --name $(ADMIN_CONTAINER_NAME) -p 8081:8080 $(IMAGE_NAME)

clean:
	docker ps -a | grep $(IMAGE_NAME) | awk  '{print $$1}' | xargs docker stop
	docker ps -a | grep $(IMAGE_NAME) | awk  '{print $$1}' | xargs docker rm