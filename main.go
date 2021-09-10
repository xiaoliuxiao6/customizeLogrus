package customizeLogrus

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

// 根据环境变量来设置 Logger
func InitLogger(log *logrus.Logger) {
	//////////////////////////////////////////////////////////////////////////////////////
	// 设置日志格式化程序
	//////////////////////////////////////////////////////////////////////////////////////
	Formatter, ok := os.LookupEnv("LOG_Formatter")
	if !ok || strings.ToLower(Formatter) == "text" || strings.ToLower(Formatter) == "" {
		// 如果没有设置则采用默认的文本格式
		log.Formatter = new(logrus.TextFormatter) // 格式化成文本模式（默认）
		// log.Formatter.(*logrus.TextFormatter).DisableLevelTruncation = true // 禁用将 Level 文本截断为4个字符
		// log.Formatter.(*logrus.TextFormatter).DisableColors = true // 禁用控制台颜色输出
		log.Formatter.(*logrus.TextFormatter).DisableTimestamp = false // 如果格式化为文本格式，则删除时间戳
		log.Formatter.(*logrus.TextFormatter).FullTimestamp = true     // 显示当前时间戳，而不是程序运行时间时长
		log.Formatter.(*logrus.TextFormatter).PadLevelText = true      // 填充 Level 文本，以便以相同长度输出所有级别
	} else {
		log.Formatter = new(logrus.JSONFormatter) // 格式化成 JSON 格式
	}

	//////////////////////////////////////////////////////////////////////////////////////
	// 定义日志输出位置
	//////////////////////////////////////////////////////////////////////////////////////
	LOG_OUT, ok := os.LookupEnv("LOG_OUT")
	if !ok || strings.ToLower(Formatter) == "stdout" || strings.ToLower(Formatter) == "" {
		// 如果指定输入位置，默认输出到 Stdout
		log.Out = os.Stdout
	} else {
		// 否则输出到文件
		file, err := os.OpenFile(LOG_OUT, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
		if err == nil {
			log.Out = file
		} else {
			log.Panic("Failed to log to file, using default stderr")
		}

		// 同时输出到文件和 Stdout
		LOG_OUT_ALL, ok := os.LookupEnv("LOG_OUT_ALL")
		if ok && strings.ToLower(LOG_OUT_ALL) == "true" {
			// （两种写法均可）
			// log.SetOutput(io.MultiWriter(os.Stdout, file))
			log.Out = (io.MultiWriter(os.Stdout, file))
		}
	}

	//////////////////////////////////////////////////////////////////////////////////////
	// 设置日志级别
	//////////////////////////////////////////////////////////////////////////////////////
	lvl, ok := os.LookupEnv("LOG_LEVEL")
	// 如果没有定义，则使用默认的 Info 级别
	if !ok || strings.ToLower(lvl) == "" {
		log.Level = logrus.InfoLevel
	} else {
		// 如果设置了，则设置成指定级别
		level, err := logrus.ParseLevel(lvl)
		if err != nil {
			fmt.Println("--------------", level)
			log.Warning("指定日志 Level 错误，将使用 Debug Level")
			level = logrus.DebugLevel
		}
		// 设置全局日志级别
		log.SetLevel(level)
	}

	// 如果日志级别 >= Debug 模式将调用方法作为字段添加到日志行，方便调试
	if log.IsLevelEnabled(logrus.DebugLevel) {
		log.SetReportCaller(true)
	}
}
