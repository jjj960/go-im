user-rpc-dev:
	@make -f deploy/makefile/makefile_user_dev.mk release-test

release-test: user-rpc-dev