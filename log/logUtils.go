package log

import (
	"log"
	"os"
	"sync"
)

var wg sync.WaitGroup

var fileName = "GoCN-news.log"

func Info(message interface{}) {
	wg.Wait()
	wg.Add(1)
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	//完成后，延迟关闭
	defer f.Close()
	// 设置日志输出到文件
	log.SetOutput(f)
	// 写入日志内容
	Logging := log.New(f, "TRACE: ", log.Ldate|log.Ltime)
	Logging.Println(message)
	wg.Done()
}

func Infof(message string, format ...interface{}) {
	wg.Wait()
	wg.Add(1)
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	//完成后，延迟关闭
	defer f.Close()
	// 设置日志输出到文件
	log.SetOutput(f)
	// 写入日志内容
	Logging := log.New(f, "TRACE: ", log.Ldate|log.Ltime)
	Logging.Printf(message, format...)
	wg.Done()
}

func Fatal(message interface{}) {
	wg.Wait()
	wg.Add(1)
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	//完成后，延迟关闭
	defer f.Close()
	// 设置日志输出到文件
	log.SetOutput(f)
	// 写入日志内容
	Logging := log.New(f, "TRACE: ", log.Ldate|log.Ltime)
	Logging.Fatal(message)
	wg.Done()
}

func Panicf(message string, format ...interface{}) {
	wg.Wait()
	wg.Add(1)
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	//完成后，延迟关闭
	defer f.Close()
	// 设置日志输出到文件
	log.SetOutput(f)
	// 写入日志内容
	Logging := log.New(f, "TRACE: ", log.Ldate|log.Ltime)
	Logging.Panicf(message, format...)
	wg.Done()
}
