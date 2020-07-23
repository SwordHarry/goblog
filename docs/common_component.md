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
    