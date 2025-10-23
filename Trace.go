package common

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"runtime/debug"
	"strings"

	"github.com/dahaipublic/common/logs"
)

/*
*@note 得到当前exe程序执行的不目录
 */
func GetCurrPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))
	ret := path[:index]
	return ret
}

func init() {
	if runtime.GOOS != "linux" {
		Logger.SetLogger("console", "")
	}
}

type CLogger struct {
	*logs.BeeLogger
}

// 实现 Writer interface
// func (this *CLogger) Write(p []byte) (n int, err error) {

// }

func (this *CLogger) Start(logPath, filename string) {
	if len(filename) == 0 {
		filename = filepath.Base(os.Args[0])
		if filepath.Ext(filename) != "" {
			po := strings.LastIndex(filename, ".")
			filename = filename[:po]
		}
	}

	//if runtime.GOOS == "windows" {
	this.setFile(logPath, filename)
	//} else {
	//	this.setFile(GetCurrPath()+"/"+logPath, filename)
	//}
	this.Info("start " + filename + " ...")
}

/*
 *@param deep可选参数,兼容无参调用.默认为4,定位调用者代码位置,业务层有再次包装可继续加大深度
 *@note 让日志记录文件名和行号
 */
func (this *CLogger) EnableFileLine(deep ...int) {
	this.EnableFuncCallDepth(true)
	deepLevel := 4
	if len(deep) > 0 {
		deepLevel = deep[0]
	}
	this.SetLogFuncCallDepth(deepLevel)
}

func (this *CLogger) setFile(path, filename string) {
	if _, err := os.Stat(path); err != nil {
		os.Mkdir(path, 0600)
	}

	//format := `{"perm":"0666","maxsize":30000000,"maxDays":180,"daily":false,"filename":"%s/%s.log"}`
	format := map[string]any{
		"perm":     "0666",
		"maxsize":  30 * 1024 * 1024,
		"maxDays":  180,
		"daily":    false,
		"filename": fmt.Sprintf("%s%s%s.log", path, string(os.PathSeparator), filename),
	}
	config, _ := json.Marshal(format)
	Logger.SetLogger("file", string(config))
}

var Logger = &CLogger{BeeLogger: logs.NewLogger(10000)}

// ------------  LOG Func  --------------
func Debug(format string, v ...interface{}) {
	Logger.Debug(format, v...)
}

func Info(format string, v ...interface{}) {
	Logger.Info(format, v...)
}

func Notice(format string, v ...interface{}) {
	Logger.Notice(format, v...)
}

func Warning(format string, v ...interface{}) {
	Logger.Warn(format, v...)
}

func Error(format string, v ...interface{}) {
	// 把stack打印出来，略掉前2行
	stackInfo := string(debug.Stack())
	slice1 := stackInfo[strings.Index(stackInfo, "\n")+1:]
	stackInfoOut := "\n" +
		strings.Repeat("*", 70) +
		"\n" +
		slice1[strings.Index(slice1, "\n")+1:] +
		strings.Repeat("*", 70)

	Logger.Error(format+stackInfoOut, v...)
}
