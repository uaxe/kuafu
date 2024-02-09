package cmcc

import (
	"context"
	"github.com/uaxe/kuafu/provider/modem"
	"net/http"
)

func init() {
	if err := modem.Register(&CMCCProvider{
		providerType: modem.CMCCProvider}); err != nil {
		panic(err)
	}
}

var _ modem.Provider = (*CMCCProvider)(nil)

type CMCCProvider struct {
	providerType string
	options
}

type OptionFunc func(opt *options) error

type options struct {
	ctx       context.Context
	hclient   *http.Client
	host      string
	port      int
	macAddr   string
	telnetURL string
}

func WithContext(ctx context.Context) OptionFunc {
	return func(opt *options) error {
		opt.ctx = ctx
		return nil
	}
}

func WithHclient(hclient *http.Client) OptionFunc {
	return func(opt *options) error {
		opt.hclient = hclient
		return nil
	}
}

func WithHost(host string) OptionFunc {
	return func(opt *options) error {
		opt.host = host
		return nil
	}
}

func WithMacAddr(addr string) OptionFunc {
	return func(opt *options) error {
		opt.macAddr = addr
		return nil
	}
}

func WithTelnetURL(uri string) OptionFunc {
	return func(opt *options) error {
		opt.telnetURL = uri
		return nil
	}
}

func defaultOptions() options {
	return options{
		ctx:       context.Background(),
		hclient:   http.DefaultClient,
		host:      "192.168.1.1",
		port:      23,
		telnetURL: "http://%s/cgi-bin/telnetenable.cgi?telnetenable=1&key=%s",
	}
}

func (s *CMCCProvider) New(ctx context.Context, f *modem.AdminFlag) (modem.Provider, error) {
	ss := &CMCCProvider{options: defaultOptions()}
	if ctx != nil {
		ss.ctx = ctx
	}
	if f.Host != "" {
		ss.host = f.Host
	}
	if f.MacAddr != "" {
		ss.macAddr = f.MacAddr
	}
	if f.Port != 0 {
		ss.port = f.Port
	}
	return ss, nil
}

func (s *CMCCProvider) Type() string {
	return s.providerType
}
