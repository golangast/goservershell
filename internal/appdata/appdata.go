package appdata

import (
	"log/slog"
	"runtime"
	"strconv"

	"github.com/golangast/goservershell/internal/loggers"
	"github.com/mitchellh/go-ps"
	"github.com/spf13/viper"
)

func Getpidstring(name string) (string, string, string, error) {

	list, err := ps.Processes()
	if err != nil {
		panic(err)
	}
	//check for windows for the .exe
	if runtime.GOOS == "windows" {
		nameos := name + ".exe"
		for _, p := range list {
			//log.Printf("Process %s with PID %d and PPID %d", p.Executable(), p.Pid(), p.PPid())

			if p.Executable() == nameos {
				name := p.Executable()
				pid := strconv.Itoa(p.Pid())
				ppid := strconv.Itoa(p.PPid())
				return name, pid, ppid, err
			}
		}
		return "", "", "", err
		//otherwise its linux
	} else {
		for _, p := range list {
			// log.Printf("Process %s with PID %d and PPID %d", p.Executable(), p.Pid(), p.PPid())

			if p.Executable() == "goservershell" {

				name := p.Executable()
				pid := strconv.Itoa(p.Pid())
				ppid := strconv.Itoa(p.PPid())

				return name, pid, ppid, err
			}
		}
		return "", "", "", err

	}

}

func GetAppData() (Stats, error) {
	logger := loggers.CreateLogger()

	viper.SetConfigName("assetdirectory") // name of config file (without extension)
	viper.SetConfigType("yaml")           // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./optimize/")    // path to look for the config file in
	err := viper.ReadInConfig()           // Find and read the config file
	if err != nil {
		logger.Error(
			"reading config file",
			slog.String("error: ", err.Error()),
		)
	}
	//get paths of asset folders from config file
	appname := viper.GetString("opt.appname")
	//logger := loggers.CreateLogger()

	exe, pid, ppid, err := Getpidstring(appname)

	// logger.Error(
	// 	"trying to get pids for application"+exe+pid+ppid,
	// 	slog.String("error: ", err.Error()),
	// )

	Stat := Stats{Exe: exe, Pid: pid, Ppid: ppid}

	return Stat, err
}

type Stats struct {
	Exe  string
	Pid  string
	Ppid string
}
