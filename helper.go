package goplay

import (
	"unsafe"
)

const (
	FmtUrl            = "/fmt"
	ShareUrl          = "/share"
	VersionUrl        = "/version"
	CompileUrl        = "/compile"
	HealthUrl         = "/_ah/health"
	ViewUrl           = "/p/%s.go"
	DefaultPlayground = "https://play.golang.org"
)

func b2str(b []byte) string {
	return unsafe.String(unsafe.SliceData(b), len(b))
}
