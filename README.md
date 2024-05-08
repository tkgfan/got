# got

> go 常用工具库。

## 安装

```bash
go get github.com/tkgfan/got
```

## 简要说明

### core

此文件加下的函数不依赖于第三方库，简要说明：
- structs
  - CopyFields(dst, src any) (err error): 将 src 上的字段复制到 dst 上。
  - IsNil(val any) bool: 判断 val 是否为 nil。
  - IsSerializable(val any): 判断 val 是否可序列化。
  - IsBasicType(val any): 判断 val 是否为基本类型。
- slices
  - ToInterfaceSlice(val any) (res []any): 将 val 转换为 interface 切片。
- errs
  - New(msg string) error: 创建一个包含堆栈信息的 error。
  - Wrap(err error) error: 返回包含栈信息的 error。
  - Wrapf(err error, format string, args ...any) error: 返回包含栈信息的 error。
  - Cause(err error) error: 返回由 errors.New、errors.Wrap、errors.Wrapf 中包裹的 cause error。普通 error 则返回其本身。
- strs
  - Rand(len int) string: 生成随机 token 字符串，len 为 token 长度。
  - WrapByFn：基于 AC 自动机进行模式串匹配并自定义处理模式串。
- logx：日志工具。
- concurrent：并发工具。
- dsl：可用于快速构建 ElasticSearch、ZincSearch 查询条件语句。
- env：环境变量工具。
- fs: 文件工具
- maths：数学工具。

### data

此文件夹下为与数据库相关的工具。简要介绍：
- mongos：mongo 工具。
- redis：redis 工具。
