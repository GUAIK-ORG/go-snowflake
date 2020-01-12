# ❄️ GO-Snowflake

## Snowflake简介

在单机系统中我们会使用自增id作为数据的唯一id，自增id在数据库中有利于排序和索引，但是在分布式系统中如果还是利用数据库的自增id会引起冲突，自增id非常容易被爬虫爬取数据。在分布式系统中有使用uuid作为数据唯一id的，但是uuid是一串随机字符串，所以它无法被排序。

Twitter设计了Snowflake算法为分布式系统生成ID,Snowflake的id是int64类型，它通过datacenterId和workerId来标识分布式系统，下面看下它的组成：

| 1bit | 41bit | 5bit | 5bit | 12bit |
|---|---|---|---|---|
| 符号位（保留字段） | 时间戳(当前时间-纪元时间) | 数据中心id | 机器id | 自增序列

### 算法简介

在使用Snowflake生成id时，首先会计算时间戳timestamp（当前时间 - 纪元时间），如果timestamp数据超过41bit则异常。同样需要判断datacenterId和workerId不能超过5bit(0-31)，在处理自增序列时，如果发现自增序列超过12bit时需要等待，因为当前毫秒下12bit的自增序列被用尽，需要进入下一毫秒后自增序列继续从0开始递增。

---

## 快速开始

### 安装

`git clone https://github.com/GUAIK-ORG/go-snowflake.git`

### 运行

`go run main.go`
