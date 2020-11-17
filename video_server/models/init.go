package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"strings"
	"time"
)

var dsn = "yx:wddlmm123@tcp(47.115.134.61:3306)/stream_media?charset=utf8mb4&parseTime=True&loc=Local"
var DB *gorm.DB
var logPath = "E:\\log\\gromLog\\"
var currentDate string
var currentLogFile *os.File

func init()  {
	NewLogger := beginLog()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: NewLogger,
	})
	if err != nil {
		log.Fatal(err)
	}
	DB = db
}

func beginLog()  logger.Interface{
	//获取日期
	timeStr := time.Now().String()
	dateStr := strings.Split(timeStr, " ")
	//设置当前date
	currentDate = dateStr[0]
	//logFile全路径
	logFileName := logPath + dateStr[0]
	logFile, err := os.OpenFile(logFileName, os.O_APPEND | os.O_CREATE, 0777 )
	if err != nil {
		log.Fatal(err)
	}
	//设置当前logFile
	currentLogFile = logFile
	logW := log.New(logFile, "\r\n", log.LstdFlags) // io writer
	//配置logger
	newLogger := logger.New(
		logW, // io writer
		logger.Config{
			SlowThreshold: time.Second,   // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      false,         // Disable color
		},
	)
	//开一个协程来每天更换logFile
	go checkDateAndChangeLogFile(logW)
	return newLogger
}

//每分钟检查一下时间，如果是新的一天，就新建一个file，并把log的output指向这个新文件
func checkDateAndChangeLogFile(logger *log.Logger)  {
	for  {
		time.Sleep(time.Minute * 1)
		//获取日期
		timeStr := time.Now().String()
		dateStr := strings.Split(timeStr, " ")
		if dateStr[0] != currentDate {	//说明新的一天开始了，那么就更换log文件
			logFileName := logPath + dateStr[0]	//logFile全路径
			logFile, err := os.OpenFile(logFileName, os.O_APPEND | os.O_CREATE, 0777 )
			if err != nil {
				log.Fatal(err)
			}
			//设置新的log文件
			logger.SetOutput(logFile)
			//关闭旧的log文件
			err = currentLogFile.Close()
			if err != nil {
				log.Fatal(err)
			}
			//设置新的当前logFile
			currentLogFile = logFile
		}
	}

}


