# 基于zookeeper的分布式配置中心 Go 语言实现单文件版本

思路来源 [superconf](https://github.com/huitouche/superconf)

使用方法见 config_test

### 特点

- 一处修改，所有机器实时同步更新
- 安装部署方便，使用简单

### 思路

通过 监听 节点的数据变更, 获取变更回调后修改本地对应的数据


### 注意

- 未进行错误处理, 会导致 zk.conn 重复释放
- 未对 对应数据加锁, 可能导致配置被修改