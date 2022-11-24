
# gorvld


```
从零开始使用 Go 语言实现一个链接器，可以正确地链接相对简单的 C 程序

PS：限定版 minimal.o，因为发现可以省掉太多繁琐细节。

```


## ELF 文件构成
```

ELF header
program header table （对 section，一对多映射，我的做法是优化掉，冗余了）
.init
.text
...       各类 Section 组成
.strtab
section header table （对 section，一对一映射）

```



```

第三课：解析链接器参数
在本节课中，我们完成了链接器参数的解析。


```

```
第四课：解析静态链接库文件
在本节课中，我们完成了静态链接库文件的解析。


第五课：解析未定义符号，移除未使用文件
在本节课中，我们根据前面课程中已经读取到的信息解析未定义 symbol 的位置以及删除掉无用的 object 文件。


第六课：处理 Mergeable Sections
在本节课中，我们开始处理 mergeable sections。


第七课：开始写文件
从本节课开始，我们开始可执行文件的写入。在本节课中，我们完成了 ELF header 的写入。

```


gcc minimal.S -c && ld minimal.o
