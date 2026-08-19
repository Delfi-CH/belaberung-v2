# Build every Target

all: backend webui

docker: backend-docker webui-docker

# Build the backend

backend:
	cd belaberung-backend && make build
	mkdir --parents --verbose dist
	cp --verbose belaberung-backend/bin/belaberung-server dist/

backend-docker: 
	docker buildx build -f Dockerfiles/Dockerfile.backend belaberung-backend -t belaberung-backend:latest

# Build the client libraries

client-libs:
	cd belaberung-client-libs && pnpm build
	cp --verbose --recursive belaberung-client-libs dist/

# Build the static WebUI

webui: client-libs
	cd belaberung-webui && pnpm build
	mkdir --parents --verbose dist
	cp --recursive --verbose belaberung-webui/build dist/www

webui-docker:
	docker buildx build -f Dockerfiles/Dockerfile.webui . -t belaberung-webui:latest

webui-docker-rootless: 
	docker buildx build -f Dockerfiles/Dockerfile.webui.rootless . -t belaberung-webui-rootless:latest

# Clean the repository: remove downloaded libraries and build artifacts

clean: backend-clean client-libs-clean webui-clean
	rm --recursive --force dist

backend-clean: 
	rm --recursive --force dist/belaberung-server
	cd belaberung-backend && make clean

client-libs-clean:
	rm --recursive --force belaberung-client-libs/dist belaberung-client-libs/node_modules dist/belaberung-client-libs
 
webui-clean:
	rm --recursive --force belaberung-webui/build belaberung-webui/node_modules dist/www

# Install all dependencies

prepare: backend-prepare client-libs-prepare webui-prepare

backend-prepare:
	cd belaberung-backend && go install

client-libs-prepare:
	cd belaberung-client-libs && pnpm install

webui-prepare:
	cd belaberung-webui && pnpm install