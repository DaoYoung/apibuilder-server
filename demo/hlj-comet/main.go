package main

import (
	_ "net/http/pprof"
	"hlj-comet/app"
	"hlj-comet/core"
	"hlj-comet/handler"
	log "github.com/sirupsen/logrus"
	"os"
	"flag"
	"fmt"
	"strconv"
	"os/signal"
	"syscall"
)

var versionFlag = flag.Bool("version", false, "Version info")

func main() {
	// 运行命令
	flag.Parse()
	if *versionFlag {
		fmt.Printf("%v\n", app.Version)
		return
	}

	// 初始化
	if err := app.InitConfig("config.yml"); err != nil {
		log.Fatal(err)
	}
	if err := app.InitDb(); err != nil {
		log.Fatal(err)
	}
	if err := app.InitEc(); err != nil {
		log.Fatal(err)
	}

	defer app.Db.Close()
	defer app.Ec.Close()
	defer setupLog()()
	defer createPid()()

	core.InitComet(app.Ec, core.CometOption{
		MsgBufSize:   app.Config.Comet.MsgBufSize,
		MsgTrunkSize: app.Config.Comet.MsgTrunkSize,
		GcInterval:   app.Config.Comet.GcInterval,
	})

	go core.GetComet().Run()
	go handler.Serve()

	// 监听系统退出信号，确保defer执行
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-signalChan
}

func createPid() func() {
	f, err := os.OpenFile("pid", os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	f.WriteString(strconv.Itoa(os.Getpid()))
	f.Close()
	return func() {
		os.Remove("pid")
	}
}

func setupLog() func() {
	openedFiles := make([]*os.File, 0)

	logFile := openLogFile(app.Config.LogPath)
	openedFiles = append(openedFiles, logFile)
	log.SetLevel(log.Level(app.Config.LogLevel))
	log.SetOutput(logFile)

	if app.Config.LogSql {
		sqlLogFile := openLogFile(app.Config.SqlLogPath)
		openedFiles = append(openedFiles, sqlLogFile)

		sqlLogger := log.New()
		sqlLogger.Out = sqlLogFile

		app.Db.LogMode(true)
		app.WedDb.LogMode(true)
		app.Db.SetLogger(sqlLogger)
		app.WedDb.SetLogger(sqlLogger)
	}

	return func() {
		for _, file := range openedFiles {
			file.Close()
		}
	}
}

func openLogFile(path string) *os.File {
	logFile, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND|os.O_CREATE, 777)
	if err != nil || logFile == nil {
		log.Fatal(err)
	}

	return logFile
}
