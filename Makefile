IMAGE_NAME = dego-ino-backend
TEST_CONTAINER_NAME = dego-ino-backend-test

build:
	docker build -t $(IMAGE_NAME) .

test:
	docker run -d --name $(TEST_CONTAINER_NAME) -p 8098:80 $(IMAGE_NAME)

clean:
	docker ps -a | grep $(IMAGE_NAME) | awk  '{print $1}' | xargs docker stop 
	docker ps -a | grep $(IMAGE_NAME) | awk  '{print $1}' | xargs docker rm