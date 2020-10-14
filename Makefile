ROOT_DIR:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
AT2K_NGINX_PORT ?= 8000
AT2K_TLS_PORT ?= 444
AT2K_NGINX_CONF_FILE ?= "nginx.dev.conf"
AT2K_SSL_CERTS_PATH ?= "/tmp"
PROJECT_NAME = "at2k"

dev.push_all:
	git add .
	git commit -m "$m"
	git push

env.recreate_db_file:
	bash $(ROOT_DIR)/scripts/recreate_db_file.sh $(ROOT_DIR)

env.fill_env_file:
	bash $(ROOT_DIR)/scripts/fill_env_file.sh $(ROOT_DIR)

env.prepare_test_files:
	bash $(ROOT_DIR)/scripts/prepare_test_files.sh $(ROOT_DIR)

env.create_proto_files:
	bash $(ROOT_DIR)/scripts/create_proto_files.sh $(ROOT_DIR)

env.install_backend_libs:
	bash $(ROOT_DIR)/scripts/prepare_backend_libs.sh $(ROOT_DIR)

env.install_frontend_libs:
	bash $(ROOT_DIR)/scripts/prepare_frontend_libs.sh $(ROOT_DIR)

env.prepare_workspace: \
	env.recreate_db_file \
	env.fill_env_file \
	env.prepare_test_files \
	env.install_backend_libs \
	env.install_frontend_libs \
	env.create_proto_files

util.calc_go_lines:
	bash $(ROOT_DIR)/scripts/calc_go_lines.sh $(ROOT_DIR)

util.calc_ng_lines:
	bash $(ROOT_DIR)/scripts/calc_ng_lines.sh $(ROOT_DIR)

tests.test_runner:
	bash $(ROOT_DIR)/scripts/test_runner_tests.sh $(ROOT_DIR)

tests.api:
	bash $(ROOT_DIR)/scripts/api_tests.sh $(ROOT_DIR)

tests.timers:
	bash $(ROOT_DIR)/scripts/timers_tests.sh $(ROOT_DIR)

tests.backend_libs:
	bash $(ROOT_DIR)/scripts/backend_libs_tests.sh $(ROOT_DIR)

tests.backend: tests.test_runner tests.api tests.timers tests.backend_libs

build.backend:
	bash $(ROOT_DIR)/scripts/build_backend.sh $(ROOT_DIR)

build.frontend:
	bash $(ROOT_DIR)/scripts/build_frontend.sh $(ROOT_DIR)

build.containers:
	AT2K_NGINX_PORT=$(AT2K_NGINX_PORT) \
	AT2K_TLS_PORT=$(AT2K_TLS_PORT) \
	NGINX_CONF_FILE=$(AT2K_NGINX_CONF_FILE) \
	SSL_CERTS_PATH=$(AT2K_SSL_CERTS_PATH) \
	docker-compose -p $(PROJECT_NAME) build

build.project: build.backend build.frontend build.containers

start.project:
	AT2K_NGINX_PORT=$(AT2K_NGINX_PORT) \
	AT2K_TLS_PORT=$(AT2K_TLS_PORT) \
	NGINX_CONF_FILE=$(AT2K_NGINX_CONF_FILE) \
	SSL_CERTS_PATH=$(AT2K_SSL_CERTS_PATH) \
 	docker-compose -p $(PROJECT_NAME) up -d

stop.project:
	AT2K_NGINX_PORT=$(AT2K_NGINX_PORT) \
	AT2K_TLS_PORT=$(AT2K_TLS_PORT) \
	NGINX_CONF_FILE=$(AT2K_NGINX_CONF_FILE) \
	SSL_CERTS_PATH=$(AT2K_SSL_CERTS_PATH) \
	docker-compose -p $(PROJECT_NAME) down

restart.project: stop.project start.project
