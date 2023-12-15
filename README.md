# Learn-BulKGo
Learning Bulk Do Jobs

在学习编写一个批量执行小工具。

## 使用

首先需要配置命令行模板文件。到~/.bulkdo/下创建命令模板文件，例如command1.sh，内容为go模板，例如：

```shell
echo {{.v.t1}} {{.v.t2}}
```

在我们希望执行bulkdo命令的目录下，有个名为.bulkdoitems的文件，里面的内容为:

```
t1, t2, t3
a1, a2, a3
b1, b2, b3
```

然后执行bulkdo command1, 就会把command1的变量替换后执行

```shell
echo a1 a2
echo b1 b2
```