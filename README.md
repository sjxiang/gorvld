
# gorvld

```
从零开始使用 Go 语言实现一个 RV64GC（RISC-V 64 位）架构的链接器，可以正确地链接相对简单的 C 程序

```


```

第一课：搭建开发环境、初始化项目、开始读取 ELF 文件

$ sudo apt install -y gcc-10-riscv64-linux-gnu qemu-user
$ which riscv64-linux-gnu-gcc-10 
/usr/bin/riscv64-linux-gnu-gcc-10
$ sudo ln -sf /usr/bin/riscv64-linux-gnu-gcc-10 /usr/bin/riscv64-linux-gnu-gcc  # alias，别名


chmod 755 ./tests/init.sh


$ hexdump -C -n 8 out/tests/hello/a.o   # 打印前 8 bit
00000000  7f 45 4c 46 02 01 01 00                           |.ELF....|
00000008

$ readelf -S out/tests/hello/a.o  # 查看 sections 数量

```

# ELF 文件结构

```
1. ELF header
2. Program header table （1 对 多） 
3. Sections
    .text
    .rodata
    ...
    .data

4. Section header table （1 对 1）


```




```
第二课：继续读取 ELF 文件
在本节课中，我们继续读取并解析了 object file 中几个重要的 section 类型。

第三课：解析链接器参数
在本节课中，我们完成了链接器参数的解析。

第四课：解析静态链接库文件
在本节课中，我们完成了静态链接库文件的解析。

第五课：解析未定义符号，移除未使用文件
在本节课中，我们根据前面课程中已经读取到的信息解析未定义 symbol 的位置以及删除掉无用的 object 文件。

第六课：处理 Mergeable Sections
在本节课中，我们开始处理 mergeable sections。

第七课：开始写文件
从本节课开始，我们开始可执行文件的写入。在本节课中，我们完成了 ELF header 的写入。

```