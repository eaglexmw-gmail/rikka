package main

import (
	"github.com/7sDream/rikka/common/logger"
	"github.com/7sDream/rikka/common/util"
	"github.com/7sDream/rikka/plugins"
	"github.com/7sDream/rikka/server"
)

// Logger of this package
var (
	l = logger.NewLogger("[Entry]")
)

// Main entry point
func main() {
	// print launch args
	l.Info("Start rikka with arg:")
	l.Info("\t bind to socket", socket)
	l.Info("\t password", util.MaskString(*argPassword, 3))
	l.Info("\t max file size", *argMaxSizeByMB, "MB")
	l.Info("\t plugin", *argPluginStr)
	l.Info("\t log level", *argLogLevel)

	l.Info("Load plugin...")
	plugins.Load(thePlugin)

	// Set SubFolder
	util.SetSubFolder(*argSubFolder)

	// start Rikka servers (this call is Sync)
	server.StartRikka(socket, *argPassword, *argMaxSizeByMB, *argHTTPS, *argCertDir, *argAllowOrigin)
}
