## 公共组件
### 1.错误码标准化
/pkg/errcode/
    errcode.go 错误码到状态码的转换，错误码的生成代码
    common_code.go 定义了错误码

### 2.配置管理
采用了 viper 第三方库
/configs/config.yaml 配置文件，包括 server app database 三个模块
/pkg/setting/
    setting.go 配置的绑定与生成
    section.go 配置的模块读取与模块生成
/main.go setupSetting 进行配置的创建

### 3.数据库连接
采用了 gorm 第三方库
/internal/model/model.go 生成数据库连接，其配置信息从 2.配置管理 中读取
/main.go setupDBEngine 进行数据库连接的生成

### 4.日志写入
/pkg/logger/logger.go 包括：
- 日志分级
- 日志标准化
- 设置日志等级
- 设置日志公共字段
- 设置日志上下文属性
- 设置某一层的调用栈的信息：程序计数器，文件信息和行号
- 设置当前整个调用栈信息
- 日志格式化和输出
- 日志分级输出

/main.go setupLogger 日志器的生成：采用了 lumberjack 第三方库

### 5.响应处理
与错误码一起定义好响应结果
- pkg
    - convert
        - convert.go 类型转换，将 url 中的查询字符串进行转换处理
    - app
        - pagination.go 分页处理
        - app.go 响应处理
### 6.swagger 接口文档
- docs: swag init 自动生成
    - docs.go
    - swagger.json
    - swagger.yaml

/internal/routers/api/v1/*.go
/main.go
在两者中添加相应注解，自动生成 swagger 文档

记得在 router.go 中进行路由注册
```go
// 注册 swagger 接口文档路由
r := gin.New()
r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
```
### 7.接口校验

### 8.国际化处理
/internal/middleware/translations.go
    进行 validator 验证器的注册和语言的载入
/pkg/app/form.go
    进行入参校验的二次封装，若绑定错误则使用国际化处理错误信息

### 9.jwt 鉴权
json web token
采用 jwt-go 库进行开发
secret-key 需要是 []byte 类型
payload 中不能明文存储重要信息，因为可以进行 base64 的解码

### 10. 访问日志记录
自定义 accessLogger，覆盖 c.Writer
### 11. 异常捕获处理