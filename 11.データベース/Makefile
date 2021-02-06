IMAGE_NAME:=istudy-mysql:1
CONTAINER_NAME:=study-mysql
run:;	docker run --name $(CONTAINER_NAME) -d -p 3306:3306 ${IMAGE_NAME}
build:;	docker build -t $(IMAGE_NAME) . 
