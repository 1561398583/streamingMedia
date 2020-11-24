package context

import (
	"video_server/config"
	"video_server/loggo"
)


var Logger = loggo.New(config.Log_path, "", loggo.LstdFlags, loggo.Debug)
