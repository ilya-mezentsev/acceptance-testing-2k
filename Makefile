ROOT_DIR:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

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
	env.install_frontend_libs

util.calc_go_lines:
	bash $(ROOT_DIR)/scripts/calc_go_lines.sh $(ROOT_DIR)

tests.test_runner:
	bash $(ROOT_DIR)/scripts/test_runner_tests.sh $(ROOT_DIR)

tests.api:
	bash $(ROOT_DIR)/scripts/api_tests.sh $(ROOT_DIR)

tests.backend_libs:
	bash $(ROOT_DIR)/scripts/backend_libs_tests.sh $(ROOT_DIR)

tests.backend: tests.test_runner tests.api tests.backend_libs

build.backend:
	bash $(ROOT_DIR)/scripts/build_backend.sh $(ROOT_DIR)

build.frontend:
	bash $(ROOT_DIR)/scripts/build_frontend.sh $(ROOT_DIR)

build.containers:
	docker-compose build

build.project: build.backend build.frontend build.containers
