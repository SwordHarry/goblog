## 架构

internal：内部架构，请求调用链路从上往下
routers：路由
service：服务和入参校验
dao：封装调用 model
model：数据库 model
----
middleware：中间件


pkg：内容为通用组件内容，贯穿整个project，如：
- app：response 响应处理；pagination 分页处理；错误信息国际化
- errcode：通用错误码
- logger：日志
- setting：project 配置<br/>配置读入 global 中

configs：yaml 配置文件<br>
scripts：存放 sql 等脚本<br>
storage：存放日志等文件<br>
