# 设计

代码不多。直接在bulkdo.go里写完方法。

我直接抓狂！！！tmd这个项目结构还是没学明白！

#### 1、readItems

读取csv文件为一个items的map
1. 创建bulkdo文件
2. 创建readItems方法
3. 编写测试方法
4. 完成readItems方法
5. 异常测试用例，完善readItems方法

#### 2、parseCommands

根据模板和items map得到执行的脚本文件
1. 创建 parseCommands 方法声明，接受两个参数，1个是模板字符串reader, 一个是items, 返回commands string slice, 还有error。
2. 编写测试方法
3. 完成 parseCommands

#### 3、execCommands

执行脚本文件
1. 创建 execCommands 方法声明，接受1个参数，commands string slice, 返回 output string slice, 和error。
2. 编写测试方法
3. 完成 execCommands

#### 4、组合这些方法

BulkDo 和命令行
BulkDo

1. 创建 BulkDo 方法声明，接受2个参数，然后是一个 模板 reader, 一个参数的reader, 返回2个值, output string slice 和 error
2. 编写测试方法
3. 完成 BulkDo

然后是命令行，命令行的运行方式是:

```
bulkdo myecho1
```

读取 ~/.bulkdao/下的命令行模板，和当前路径下的 .bulkdoitems 参数，然后调用 BulkDo 方法