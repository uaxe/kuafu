package cucc

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/uaxe/infra/zhttp"
	"github.com/uaxe/kuafu/provider/modem"
	"golang.org/x/net/html"
)

func init() {
	if err := modem.Register(&CUCCProvider{
		providerType: modem.CUCCProvider}); err != nil {
		panic(err)
	}
}

var _ modem.Provider = (*CUCCProvider)(nil)

type CUCCProvider struct {
	providerType string
	device       string
	options
}

type OptionFunc func(opt *options) error

type options struct {
	ctx     context.Context
	host    string
	port    int
	macAddr string
}

func WithContext(ctx context.Context) OptionFunc {
	return func(opt *options) error {
		opt.ctx = ctx
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

func defaultOptions() options {
	return options{
		ctx:  context.Background(),
		host: "192.168.1.1",
		port: 23,
	}
}

func (s *CUCCProvider) New(ctx context.Context, f *modem.AdminFlag) (modem.Provider, error) {
	ss := &CUCCProvider{options: defaultOptions()}
	if ctx != nil {
		ss.ctx = ctx
	}
	if f.Host != "" {
		ss.host = f.Host
	}
	if f.Port != 0 {
		ss.port = f.Port
	}

	_ = ss.get_device()

	return ss, nil
}

func (s *CUCCProvider) Type() string {
	return s.providerType
}

func (s *CUCCProvider) IsMe(root *html.Node) bool {
	doc := goquery.NewDocumentFromNode(root)
	return strings.Contains(doc.Text(), "中国联通")
}

var (
	RegDevice = regexp.MustCompile(`<title>(\w+)</title>`)
)

func (s *CUCCProvider) get_device() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	deviceURL := fmt.Sprintf("http://%s/hidden_version_switch.gch", s.host)
	for range zhttp.NewRetryTimer(ctx) {
		resp, err := http.DefaultClient.Get(deviceURL)
		if err != nil {
			return err
		}
		if resp.StatusCode != http.StatusOK {
			_ = resp.Body.Close()
			continue
		}
		raw, err := io.ReadAll(resp.Body)
		if err != nil {
			_ = resp.Body.Close()
			continue
		}

		if lables := RegDevice.FindSubmatch(raw); len(lables) > 0 {
			s.device = string(lables[1])
		}
		_ = resp.Body.Close()
		break
	}

	return ctx.Err()
}
