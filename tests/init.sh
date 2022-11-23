#!/bin/bash

mkdir -p out/tests/hello/

cat << EOF | riscv64-linux-gnu-gcc -o out/tests/hello/a.o -c -xc -
#include <stdio.h>

int main(void) {
    printf("Hello world\n");
    return 0;
}
EOF

./gorvld out/tests/hello/a.o  
# 导入文件路径