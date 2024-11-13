package namedpipe

import (
	"runtime"

	"github.com/stenet/go-mssqldb/msdsn"
)

type namedPipeDialer struct{}

var dialer namedPipeDialer = namedPipeDialer{}

func init() {
	if runtime.GOOS == "windows" {
		msdsn.ProtocolParsers = append(msdsn.ProtocolParsers, dialer)
		msdsn.ProtocolDialers["np"] = dialer
	}
}
