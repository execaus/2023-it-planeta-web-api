sqlc:
	generate sqlc

docker_postgres_build:
	docker build \
    		-t it_planeta_postgres \
    		-f ./docker/Dockerfile_postgres .

docker_api_build:
	docker build \
		-t it_planeta_api \
		-f ./docker/Dockerfile_api .