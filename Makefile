ROOT_DIR:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

dev.fill_env_file:
	bash $(ROOT_DIR)/scripts/fill_env_file.sh $(ROOT_DIR)

dev.prepare_test_files:
	bash $(ROOT_DIR)/scripts/prepare_test_files.sh $(ROOT_DIR)

dev.create_proto_files:
	bash $(ROOT_DIR)/scripts/create_proto_files.sh $(ROOT_DIR)

dev.install_backend_libs:
	bash $(ROOT_DIR)/scripts/prepare_backend_libs.sh $(ROOT_DIR)

dev.prepare_workspace: dev.fill_env_file dev.prepare_test_files dev.install_backend_libs

dev.calc_go_lines:
	bash $(ROOT_DIR)/scripts/calc_go_lines.sh $(ROOT_DIR)

dev.test_runner_tests:
	bash $(ROOT_DIR)/scripts/test_runner_tests.sh $(ROOT_DIR)

dev.api_tests:
	bash $(ROOT_DIR)/scripts/api_tests.sh $(ROOT_DIR)

dev.backend_libs_tests:
	bash $(ROOT_DIR)/scripts/backend_libs_tests.sh $(ROOT_DIR)

dev.backend_tests: dev.test_runner_tests dev.api_tests dev.backend_libs_tests
