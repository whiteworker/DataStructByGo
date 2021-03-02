package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"time"
)

const (
	InfoLevel = iota
	WarnLevel
	ErrLevel
)

// fileLogger代表了一种级别的日志文件对象，
type fileLogger struct {
	logger      *log.Logger
	writer      io.Writer //记录打开日志文件描述符，便于后续关闭
	curOpenFile string    //当前打开文件名
	openTime    time.Time //当前日志打开时间
}

// ELog 代表一个日志对象，包含了三个日志级别的日志文件对象
type ELog struct {
	mu         sync.Mutex
	minLevel   int // 日志打印的最低等级，低于此等级的日志将被忽略
	infoLogger *fileLogger
	warnLogger *fileLogger
	errLogger  *fileLogger
}

func (l *ELog) Info(v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.minLevel > InfoLevel {
		return
	}

	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "unknow"
		line = 0
	} else {
		file = filepath.Base(file)
	}
	//这里先使用runtime.Caller获取文件名和行号，后标准化输出
	// 下面的Warn()函数直接使用log.logger.Outpu()函数指明调用者的文件名和行号，两种方式都OK
	l.infoLogger.logger.Printf("%15s:%4d - %s", file, line, v)
}

func (l *ELog) Warn(v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.minLevel > WarnLevel {
		return
	}
	// l.warnLogger.logger.Println(v)

	//这里使用Output指定callpath
	str := fmt.Sprintf("%s", v)
	l.warnLogger.logger.Output(2, str)
}

func (l *ELog) Err(v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.minLevel > ErrLevel {
		return
	}
	//这里使用Output指定callpath
	str := fmt.Sprintf("%s", v)
	l.errLogger.logger.Output(2, str)
}

//Close 目前因为3个日志写入同一个文件，故只需要关闭一次
func (l *ELog) Close() {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.infoLogger.writer.(*os.File).Close()
}

// 返回一个ELog对象，该对象实现了三种日志级别格各自的logger对象，低层写入同一个文件
func NewLogger(level int, logPath string, logFile string) *ELog {
	mkLogDir(logPath)
	file := logPath + "/" + logFile
	writer, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		panic("open log file error")
	}

	infoLog := log.New(writer, "[info]", log.Ldate|log.Ltime)
	warnLog := log.New(writer, "[warn]", log.Ldate|log.Ltime|log.Llongfile)
	errLog := log.New(writer, "[err]", log.Ldate|log.Ltime|log.Lshortfile)

	now := time.Now()

	infoLogFile := &fileLogger{
		logger:      infoLog,
		curOpenFile: file,
		writer:      writer,
		openTime:    now,
	}

	warnLogFile := &fileLogger{
		logger:      warnLog,
		curOpenFile: file,
		writer:      writer,
		openTime:    now,
	}

	errLogFile := &fileLogger{
		logger:      errLog,
		curOpenFile: file,
		writer:      writer,
		openTime:    now,
	}

	l := &ELog{
		infoLogger: infoLogFile,
		warnLogger: warnLogFile,
		errLogger:  errLogFile,
		minLevel:   level,
	}
	return l
}

// 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func mkLogDir(logPath string) {

	if !IsDir(logPath) {
		fmt.Println("dddd")
		os.Mkdir(logPath, 0777)

	}
}

// 判断所给路径文件/文件夹是否存在
func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

var logger *ELog

func initLogger() {
	logPath := "./logs"
	logFile := "1.log"
	logLevel := InfoLevel
	logger = NewLogger(logLevel, logPath, logFile)
}

func main() {

	initLogger()

	logger.Info("hello info")
	test()
}

func test() {
	logger.Info("hello info test", "kkkkkk")
	logger.Warn("hello waring", "warn1", ",warn2")
	logger.Err("hello err", "err1", "error2")
}
