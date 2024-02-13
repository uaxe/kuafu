package provider

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"time"

	"github.com/uaxe/infra/zhttp"
	"github.com/uaxe/kuafu/provider/modem"
	"golang.org/x/net/html"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/transform"
)

func New(ctx context.Context, providerType string) Provider {
	return &_provider{
		ctx:          ctx,
		providerType: providerType,
	}
}

type _provider struct {
	ctx          context.Context
	providerType string
}

func (p *_provider) Type() string {
	return p.providerType
}

var (
	RegCharset = regexp.MustCompile(`<meta http-equiv="Content-Type" content="text/html; charset=(\w+)" />`)
)

func (p *_provider) SuperAdminProvider(f *modem.AdminFlag) (modem.Provider, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	var root *html.Node
	for range zhttp.NewRetryTimer(ctx) {
		resp, err := http.DefaultClient.Get(fmt.Sprintf("http://%s", f.Host))
		if err != nil {
			return nil, err
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

		label := "utf8"
		if lables := RegCharset.FindSubmatch(raw); len(lables) > 0 {
			label = string(lables[1])
		}
		e, _ := charset.Lookup(label)
		reader := transform.NewReader(bytes.NewReader(raw), e.NewDecoder())
		root, err = html.Parse(reader)
		if err != nil {
			_ = resp.Body.Close()
			continue
		}
		_ = resp.Body.Close()
		break
	}

	pp, err := modem.Get(func(s modem.Provider) bool { return s.IsMe(root) })
	if err != nil {
		return nil, err
	}

	return pp.New(p.ctx, f)
}
