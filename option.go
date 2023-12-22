package goplay

import "time"

type Options struct {
	// playground server address
	Address string
	// proxy url
	Proxy string

	// http request timeout duration
	Timeout time.Duration
}

type Option func(opt *Options)

func WithAddress(addr string) Option {
	return func(opt *Options) {
		opt.Address = addr
	}
}

func WithProxy(proxy string) Option {
	return func(opt *Options) {
		opt.Proxy = proxy
	}
}
