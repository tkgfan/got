# CHANGELOG

## v1.1.6 2024/02/19

- 增加
  - dsl
- 修改
  - structs.CopyFields 抛出 error 改 panic
  - 升级 go 版本

## v1.1.5 2024/01/29

- 增加
  - 固定大小并发池增加默认策略

## v1.1.4

- 修复
  - 栈日志无法序列化 Cause 问题

## v1.1.3

- 新增
  - 变更日志
  - 链路日志增加 ip、hasErr 信息
- 修改
  - errors 移除自定义 JSON 序列化方法