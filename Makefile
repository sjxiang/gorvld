TESTS := $(wildcard tests/*.sh)

build:
	go build

clean:
	go clean
	rm -rf out/


test:
	@echo 'Testing' $@
	@mkdir -p out/
	@gcc -c test/minimal.S -o out/minimal.o
	@./gorvld out/minimal.o 
	@printf '\e[32mOK\e[0m\n'
	@printf '\e[32mPassed test\e[0m\n'

#  导入文件路径

.PHONY: build clean test 