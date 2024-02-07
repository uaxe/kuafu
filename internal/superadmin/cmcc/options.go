package cmcc

import (
	"context"
	"net/http"
)

type OptionFunc func(opt *option) error

type option struct {
	ctx       context.Context
	hclient   *http.Client
	name      string
	addr      string
	telnetURL string
}

func WithHclient(hclient *http.Client) OptionFunc {
	return func(opt *option) error {
		opt.hclient = hclient
		return nil
	}
}

func WithName(name string) OptionFunc {
	return func(opt *option) error {
		opt.name = name
		return nil
	}
}

func WithAddr(addr string) OptionFunc {
	return func(opt *option) error {
		opt.addr = addr
		return nil
	}
}

func WithTelnetURL(uri string) OptionFunc {
	return func(opt *option) error {
		opt.telnetURL = uri
		return nil
	}
}

func defaultOption(ctx context.Context) *option {
	return &option{
		ctx:       ctx,
		hclient:   http.DefaultClient,
		addr:      "192.168.1.1",
		telnetURL: "http://%s/cgi-bin/telnetenable.cgi?telnetenable=1&key=%s",
	}
}
