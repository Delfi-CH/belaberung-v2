# Configuration

# Generic Configuration

MAKEFLAGS := --jobs=$(shell nproc)

# Backend Configuration

BACKEND_BINARY_NAME=belaberung-server
GOFLAGS := -C belaberung-backend -race -v -work -x -compiler gc -gcflags=$(GO_GOCFLAGS) -ldflags=$(GO_LDFLAGS)
GOFLAGS_GCCGO := -C belaberung-backend -race -v -work -x -compiler gccgo -asmflags=$(GO_ASMFLAGS) -gccgoflags=$(GO_GCCGOFLAGS) -ldflags=$(GO_LDFLAGS)
GO_ASMFLAGS:=
GO_GOCFLAGS:=
GO_GCCGOFLAGS:=
GO_LDFLAGS:=
GO_CLEANFLAGS := -x -cache -testcache -modcache

# Directory Configuration

ifeq ($(PREFIX),)
    PREFIX := /usr
endif

BINDIR=$(PREFIX)/bin
LIBDIR=$(PREFIX)/lib
SHAREDIR=$(PREFIX)/share
WEBDIR=$(SHAREDIR)/nginx/html
CONFIGDIR=/etc
TMPDIR=/tmp
SYSTEMD_SERVICE_DIR=$(LIBDIR)/systemd/system/

OUTDIR=./dist
OUTDIR_BACKEND=$(OUTDIR)/backend
OUTDIR_WEBUI=$(OUTDIR)/webui
OUTDIR_CLIENT_LIBS=$(OUTDIR)/client-libs
OUTDIR_DESKTOP=$(OUTDIR)/desktop

# WebUI Configuration

BELABERUNG_DOMAIN=chat.example.com

# Build every Target

all: backend client-libs webui desktop 

docker: backend-docker webui-docker

# Build the backend

backend:
	go build $(GOFLAGS) -o bin/${BACKEND_BINARY_NAME}
	mkdir --parents --verbose $(OUTDIR_BACKEND)
	cp --verbose belaberung-backend/bin/${BACKEND_BINARY_NAME} $(OUTDIR_BACKEND)/
	cp --verbose belaberung-backend/belaberung-backend.service $(OUTDIR_BACKEND)/
	cp --verbose belaberung-backend/.env.example $(OUTDIR_BACKEND)/belaberung-server-enviroment

backend-docker: 
	docker buildx build -f Dockerfiles/Dockerfile.backend belaberung-backend -t belaberung-backend:latest

# Build the client libraries

client-libs:
	cd belaberung-client-libs && pnpm build
	mkdir --parents --verbose $(OUTDIR_CLIENT_LIBS)
	cp --verbose --recursive belaberung-client-libs/* $(OUTDIR_CLIENT_LIBS)/

# Build the static WebUI

webui: client-libs
	cd belaberung-webui && pnpm build
	mkdir --parents --verbose $(OUTDIR_WEBUI)
	cp --recursive --verbose belaberung-webui/build $(OUTDIR_WEBUI)/www
	cp --verbose belaberung-webui/nginx.conf $(OUTDIR_WEBUI)/nginx.conf

webui-docker:
	docker buildx build -f Dockerfiles/Dockerfile.webui . -t belaberung-webui:latest

webui-docker-rootless: 
	docker buildx build -f Dockerfiles/Dockerfile.webui.rootless . -t belaberung-webui-rootless:latest

# Build the desktop app

desktop:
	cd belaberung-gui-desktop && deno task desktop
	mkdir --parents --verbose $(OUTDIR_DESKTOP)/belaberung
	cp --recursive --verbose belaberung-gui-desktop/dist/belaberung/* $(OUTDIR_DESKTOP)/belaberung/

desktop-appimage:
	cd belaberung-gui-desktop && deno task appimage
	mkdir --parents --verbose $(OUTDIR_DESKTOP)/
	cp --recursive --verbose belaberung-gui-desktop/dist/*.AppImage $(OUTDIR_DESKTOP)/

# Clean the repository: remove downloaded libraries and build artifacts, aswell as any files not in the git staging area

clean: backend-clean client-libs-clean webui-clean desktop-clean
	rm --recursive --force $(OUTDIR)
	git clean --force

backend-clean: 
	rm --recursive --force $(OUTDIR_BACKEND)
	cd belaberung-backend && go clean $(GO_CLEANFLAGS)

client-libs-clean:
	rm --recursive --force belaberung-client-libs/dist belaberung-client-libs/node_modules $(OUTDIR_CLIENT_LIBS)/belaberung-client-libs
 
webui-clean:
	rm --recursive --force belaberung-webui/build belaberung-webui/node_modules $(OUTDIR_WEBUI)/www

desktop-clean:
	rm --recursive --force belaberung-gui-desktop/build belaberung-gui-desktop/dist belaberung-gui-desktop/node_modules $(OUTDIR_DESKTOP)

# Install all dependencies

prepare: backend-prepare client-libs-prepare webui-prepare desktop-prepare

backend-prepare:
	cd belaberung-backend && go install

client-libs-prepare:
	cd belaberung-client-libs && pnpm install

webui-prepare:
	cd belaberung-webui && pnpm install

desktop-prepare: 
	cd belaberung-gui-desktop && pnpm install && deno install

# Start Applications in development mode

dev: backend-dev webui-dev

backend-dev:
	go run $(GOFLAGS) .

webui-dev: client-libs
	cd belaberung-webui && pnpm dev

# Format the source code

format: backend-format webui-format desktop-format

backend-format:
	cd belaberung-backend && make format

webui-format:
	cd belaberung-webui && pnpm format

desktop-format: 
	cd belaberung-gui-desktop && pnpm format

# Install the files

install: backend-install webui-install

backend-install:
	install -d $(DESTDIR)$(BINDIR)
	install -d $(DESTDIR)$(SYSTEMD_SERVICE_DIR)
	install -d $(DESTDIR)$(CONFIGDIR)
	install -m 755 $(OUTDIR_BACKEND)/$(BACKEND_BINARY_NAME) $(DESTDIR)$(BINDIR)/$(BACKEND_BINARY_NAME)
	install -m 644 $(OUTDIR_BACKEND)/belaberung-backend.service $(DESTDIR)$(SYSTEMD_SERVICE_DIR)/belaberung-backend.service
	install -m 600 $(OUTDIR_BACKEND)/belaberung-server-enviroment $(DESTDIR)$(CONFIGDIR)/belaberung-server-enviroment

webui-install:
	install -d $(DESTDIR)$(WEBDIR)
	install -d $(DESTDIR)$(CONFIGDIR)/nginx/conf.d
	cp $(OUTDIR_WEBUI)/nginx.conf $(TMPDIR)/nginx.conf.template
	BELABERUNG_WEBUI_PORT=80 \
	BELABERUNG_WEBUI_BACKEND_HOST='http://localhost:8081/' \
	BELABERUNG_WEBUI_SERVER_NAME='$(BELABERUNG_DOMAIN)' \
	BELABERUNG_WEBUI_ROOT='$(WEBDIR)' \
	envsubst '$${BELABERUNG_WEBUI_PORT} $${BELABERUNG_WEBUI_BACKEND_HOST} $${BELABERUNG_WEBUI_SERVER_NAME} $${BELABERUNG_WEBUI_ROOT}' < $(TMPDIR)/nginx.conf.template > $(TMPDIR)/nginx.conf
	install -m 644 $(TMPDIR)/nginx.conf $(DESTDIR)$(CONFIGDIR)/nginx/conf.d/default.conf
	cp --recursive --verbose $(OUTDIR_WEBUI)/www/* $(DESTDIR)$(WEBDIR)/