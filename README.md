# customize-logrus
以 github.com/sirupsen/logrus 为基础，完全兼容 Logrus 接口
只是在此基础上做了简单封装，实现了使用环境变量来定义某些功能和更方便使用

> 注意：如果指定的变量 Value 为空，则使用默认设置

## 1.设置日志格式化程序

```sh
## 用法（不区分大小写）：
export LOG_Formatter=JSON	## 使用 JSONFormatter 格式化程序输出为 JSON 格式
export LOG_Formatter=TEXT	## 使用 TextFormatter 格式化程序输出为 Text 格式（默认）
```

```
os.Setenv("LOG_Formatter", "JSON")
os.Setenv("LOG_Formatter", "TEXT")
```



## 2.定义日志输出位置

```sh
## 用法（不区分大小写）：
export LOG_OUT=stdout	## 输出到 Stdout（默认）
export LOG_OUT="/home/1.txt"	## 输出到具体文件
export LOG_OUT_ALL=True			## 同时输出到 Stdout 和具体文件（必须设置 LOG_OUT 值为路径）
```

```go
os.Setenv("LOG_OUT", "stdout")
os.Setenv("LOG_OUT", "/home/1.txt")
os.Setenv("LOG_OUT_ALL", "True")
```



## 3.定义日志级别

```sh
## 用法（不区分大小写）：
export LOG_LEVEL=Debug		## 如果设置为 Debug 模式，日志将同时打印调用函数的具体函数名和行数，方便调试
export LOG_LEVEL=Info		## （默认）
export LOG_LEVEL=Warn
export LOG_LEVEL=Warning	## 同 Warn
export LOG_LEVEL=Error
export LOG_LEVEL=Fatal		## 记录日志后调用 os.Exit(1)
export LOG_LEVEL=Panic		## 记录日志后触发 panic()
```

```go
os.Setenv("LOG_LEVEL", "Debug")
os.Setenv("LOG_LEVEL", "Info")
os.Setenv("LOG_LEVEL", "Warn")
os.Setenv("LOG_LEVEL", "Warning")
os.Setenv("LOG_LEVEL", "Error")
os.Setenv("LOG_LEVEL", "Fatal")
os.Setenv("LOG_LEVEL", "Panic")
```



## 4.代码内直接指定

除了使用以上正常方式指定以外，还可以直接在代码内定义环境变量，如下所示：

```go
package main

import (
	customizeLogrus "github.com/xiaoliuxiao6/customize-logrus"
)

var (
	log = customizeLogrus.InitLogger()
)

func main() {
	// 代码内指定环境变量
	// os.Setenv("LOG_Formatter", "JSON")

	log.Info("Info 内容")
	log.Debug("Debug 内容")
	log.Error("Error 内容")
}

```



# 示例用法

#### 1.使用标准库 log

```go
package main

import "log"

func main() {
	log.Print("普通日志")
}
```



#### 2.使用 Logrus

```go
package main

import (
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func main() {
	log.Print("普通日志")

	// 将支持日志级别输出
	log.Info("这是一条 info...")
	log.Debug("这个就厉害了，debug...")
	log.Print("这个是用的 print")
}
```



#### 3.使用自定义库以支持环境变量

```go
package main

import (
	"github.com/sirupsen/logrus"
	customizeLogrus "github.com/xiaoliuxiao6/customize-logrus"
)

// 初始化 Log
var log = logrus.New()

// 是否添加自定义方法以支持环境变量方式修改 Log 设置
func init() {
	// 定义日志级别
	// os.Setenv("LOG_Formatter", "JSON")
	// os.Setenv("LOG_Formatter", "TEXT")

	// 定义日志输出位置
	// os.Setenv("LOG_OUT", "/home/1.txt")
	// os.Setenv("LOG_OUT_ALL", "True")
	// os.Setenv("LOG_OUT", "stdout")

	// 定义日志级别
	// os.Setenv("LOG_LEVEL", "Debug")
	// os.Setenv("LOG_LEVEL", "Info")
	// os.Setenv("LOG_LEVEL", "Warn")
	// os.Setenv("LOG_LEVEL", "Warning")
	// os.Setenv("LOG_LEVEL", "Error")
	// os.Setenv("LOG_LEVEL", "Fatal")
	// os.Setenv("LOG_LEVEL", "Panic")

	// 根据环境变量来设置 Log
	customizeLogrus.InitLogger(log)
}

func main() {

	log.Info("这是一条 info...")
	log.Debug("这个就厉害了，debug...")
	log.Print("这个是用的 print")
	linshi()
}

func linshi() {
	log.Error("test")
}
```
