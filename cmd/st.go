/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"os"
	"os/exec"
	"runtime"

	"log/slog"

	"github.com/golangast/goservershell/internal/loggers"
	"github.com/golangast/goservershell/src/server"
	"github.com/spf13/cobra"
)

// stCmd represents the st command
var stCmd = &cobra.Command{
	Use:   "st",
	Short: "to build and start program",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		opts := loggers.PrettyHandlerOptions{
			SlogOpts: slog.HandlerOptions{
				AddSource: true,
				Level:     slog.LevelDebug,
			},
		}
		handler := loggers.NewPrettyHandler(os.Stdout, opts)
		logger := slog.New(handler)
		err, out, errout := Startprograms(`go build`)
		if err != nil {
			logger.Error(
				"while trying to get files from assets",
				slog.String("error: ", err.Error()),
			)
		}
		if errout != "" && out != "" {
			logger.Info("starting application", slog.String("out", out), slog.String("errors: ", errout))
		}

		server.Server()

	},
}

func init() {
	rootCmd.AddCommand(stCmd)

}
func Startprograms(command string) (error, string, string) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", command)
	} else {
		cmd = exec.Command("bash", "-c", command)
	}

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return err, stdout.String(), stderr.String()
}
