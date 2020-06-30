eagle

##### eagle工具是为了迁移当前java项目接入到vitess的辅助项目，vitess经过分析整理，有很多不能执行的sql,之前有pdf文件 [vitess不支持语法](http://gitlab.yzf.net/PaaS-public/PaaS-Board/PaaS/Vitess/blob/master/vitess_sql%E4%B8%8D%E6%94%AF%E6%8C%81%E6%80%BB%E7%BB%93.pdf) 使用eagle就能快速扫描mybatis项目的sql，把问题sql，或者关键字做命令行提醒

### eagle特性

* 扫描mybatis xml文件
* 查询xml文件中的问题关键字
* 解析xml文件中的子查询sql,并且打印sql
### 如何使用
