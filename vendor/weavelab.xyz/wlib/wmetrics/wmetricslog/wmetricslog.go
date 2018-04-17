package wmetricslog

import "weavelab.xyz/wlib/wlog"

//Logger is used by all wmetrics packages for logging. It can be replaced with a custom logger if desired
var Logger = wlog.NewWLogger(wlog.WlogdLogger)
