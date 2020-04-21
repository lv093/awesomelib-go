package main


import (
	"github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"github.com/tietang/go-utils"
	"github.com/x-cray/logrus-prefixed-formatter"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)


var log_level = "debug" // trace debug info warn error fatal panic
var log_path = "logs"
var log_test_path = ""

func main() {

	RegisterLogger()
	//测试
	logrus.Info("dasdsadsa")
}


func RegisterLogger(){
	setLogLevel(log_level)
	setConsoleFormater()
	setFileFormater(log_path,log_test_path)
}

func setLogLevel(level string) {
	//日志级别
	parseLevel, e := logrus.ParseLevel(level)
	if e != nil {
		panic(e)
	}
	logrus.SetLevel(parseLevel)
	logrus.SetReportCaller(true)
}


func setConsoleFormater() {
	//formatter :=&logrus.TextFormatter{}// logrus 自带的文本格式
	formatter := &prefixed.TextFormatter{}
	//引用第三方的文本格式
	formatter.FullTimestamp = true
	formatter.TimestampFormat = "2006-01-02.15:04:05.000000"
	formatter.ForceFormatting = true
	formatter.SetColorScheme(&prefixed.ColorScheme{
		InfoLevelStyle:  "green",
		WarnLevelStyle:  "yellow",
		ErrorLevelStyle: "red",
		FatalLevelStyle: "41",
		PanicLevelStyle: "41",
		DebugLevelStyle: "blue",
		PrefixStyle:     "cyan",
		TimestampStyle:  "37",
	})
	logrus.SetFormatter(formatter)
}



//初始化log配置，配置logrus日志文件滚动生成和
func setFileFormater(logDir ,logTestDir string) {

	//设置日志文件输出的是否带文件和函数
	showfileLength()

	//配置日志输出目录
	if logTestDir != ""{
		logDir = logTestDir
	}

	logFilePath, _ := filepath.Abs(logDir)
	logrus.Infof("logrus abs path: %s", logFilePath)

	logFileName := "log"
	maxAge :=  time.Hour*24
	rotationTime := time.Hour*1
	_=os.MkdirAll(logDir, os.ModePerm)

	baseLogPath := path.Join(logDir, logFileName)

	//设置滚动日志输出writer
	writer, err := rotatelogs.New(
		strings.TrimSuffix(baseLogPath, ".log")+".%Y%m%d%H",
		rotatelogs.WithLinkName(baseLogPath),      // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(maxAge),             // 文件最大保存时间
		rotatelogs.WithRotationTime(rotationTime), // 日志切割时间间隔
	)
	if err != nil {
		logrus.Errorf("config local file system logger error. %+v", err)
	}

	//设置日志文件输出的日志格式
	//formatter := &logrus.JSONFormatter{}

	formatter :=&prefixed.TextFormatter{} //引用第三方的文本格式
	//控制台高亮显示
	formatter.DisableColors=true
	formatter.FullTimestamp=true
	formatter.TimestampFormat="2006-01-02.15:04:05.000000"
	formatter.ForceFormatting=true

	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer, // 为不同级别设置不同的输出目的
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, formatter)

	logrus.AddHook(lfHook)

}

func showfileLength() {
	lfh := utils.NewLineNumLogrusHook()
	lfh.EnableFileNameLog = true
	level, e := logrus.ParseLevel(log_level)
	if e != nil {
		panic(e)
	}
	if level == logrus.DebugLevel {
		lfh.EnableFuncNameLog = true
	}else{
		lfh.EnableFuncNameLog = false
	}
	logrus.AddHook(lfh)
}