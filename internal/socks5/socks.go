package socks5

import (
	"context"
	"github.com/hardstylez72/cry/internal/lib"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	_ "github.com/armon/go-socks5"
	"github.com/hardstylez72/cry/internal/server/config"
	"github.com/pkg/errors"
	"golang.org/x/net/proxy"
)

type Config struct {
	Host     string
	Login    string
	Password string
	Disable  bool
}

type Proxy struct {
	Config *Config
	Cli    *http.Client
}

const (
	Timeout             = time.Second * 30
	KeepAlive           = time.Second * 30
	TLSHandshakeTimeout = time.Second * 10
	TestURL             = "https://eth-pokt.nodies.app"
)

const mask = "ip:port:login:password"

var ErrParseError = errors.New("parse error")

// NewSock5ProxyString
// example "168.91.123.48:5346:meqdlqu:q6mnwsgrwef2"
// mask "ip:port:login:password"
func NewSock5ProxyString(s, userAgentHeader string) (*Proxy, error) {
	c, err := ParseProxy(s)
	if err != nil {
		return nil, errors.Wrap(ErrParseError, err.Error())
	}

	if c.Disable {
		rotatingProxy := "p.webshare.io:80:" + config.CFG.WebshareProxyUser + ":" + config.CFG.WebshareProxyPassword
		c, err = ParseProxy(rotatingProxy)
		if err != nil {
			return nil, errors.Wrap(ErrParseError, err.Error())
		}
	}

	return NewSock5Proxy(c)
}

func (c Config) Validate() error {
	if c.Disable {
		return nil
	}
	ss := strings.Split(c.Host, ":")
	if len(ss) != 2 {
		return errors.New("invalid host format. " + c.Host + " mask XX.XX.XX.XX:XX")
	}

	if ss[0] == "" || ss[1] == "" {
		return errors.New("invalid host format. " + c.Host + " mask XX.XX.XX.XX:XX")
	}

	return nil
}

var proxyTypeMap = lib.NewMap[string, string]()

func NewSock5Proxy(c *Config) (*Proxy, error) {

	if err := c.Validate(); err != nil {
		return nil, err
	}

	if c.Disable {
		p := &Proxy{
			Config: c,
			Cli: &http.Client{
				Transport: http.DefaultTransport,
				Timeout:   Timeout,
			},
		}

		return p, nil
	}

	t, ok := proxyTypeMap.Get(c.Host)
	if ok {
		switch t {
		case "socks5":
			return makeSocks5Proxy(c)
		case "http":
			return makeHttpProxy(c)
		}
	}

	p, err := makeSocks5Proxy(c)
	if err == nil {
		proxyTypeMap.Set(c.Host, "socks5")
		return p, nil
	}

	p, err = makeHttpProxy(c)
	if err == nil {
		proxyTypeMap.Set(c.Host, "http")
		return p, nil
	}

	return nil, errors.Wrap(err, "failed to create proxy")
}

func makeHttpProxy(c *Config) (*Proxy, error) {

	p := &Proxy{
		Config: c,
		Cli: &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(&url.URL{
					Scheme: "http",
					User:   url.UserPassword(c.Login, c.Password),
					Host:   c.Host,
				}),
				ForceAttemptHTTP2:     true,
				MaxIdleConns:          100,
				IdleConnTimeout:       90 * time.Second,
				TLSHandshakeTimeout:   10 * time.Second,
				ExpectContinueTimeout: 1 * time.Second,
			},
			Timeout: Timeout,
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	req, _ := http.NewRequestWithContext(ctx, "GET", TestURL, nil)

	do, err := p.Cli.Do(req)
	if err != nil {
		return nil, err
	}
	if do.StatusCode != 200 {
		return nil, errors.New("invalid proxy")
	}

	return p, nil
}

func makeSocks5Proxy(c *Config) (*Proxy, error) {

	auth := proxy.Auth{
		User:     c.Login,
		Password: c.Password,
	}

	dialer, err := proxy.SOCKS5("tcp", c.Host, &auth, &net.Dialer{
		Timeout:   Timeout,
		KeepAlive: KeepAlive,
	})
	if err != nil {
		return nil, err
	}

	p := &Proxy{
		Config: c,
		Cli: &http.Client{
			Transport: &http.Transport{
				Proxy:                 nil,
				ForceAttemptHTTP2:     true,
				MaxIdleConns:          100,
				IdleConnTimeout:       90 * time.Second,
				TLSHandshakeTimeout:   10 * time.Second,
				ExpectContinueTimeout: 1 * time.Second,
				Dial:                  dialer.Dial,
			},
			Timeout: Timeout,
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	req, _ := http.NewRequestWithContext(ctx, "GET", TestURL, nil)

	do, err := p.Cli.Do(req)
	if err != nil {
		return nil, err
	}

	if err != nil || do.StatusCode != 200 {
		return nil, errors.New("proxy failed to connect")
	}

	return p, nil
}

func ParseProxy(s string) (*Config, error) {

	if s == "" {
		return &Config{
			Disable: true,
		}, nil
	}

	str := strings.Split(s, ":")
	if len(str) != 4 {
		return nil, errors.New("invalid proxy format: " + s + " format: " + mask)
	}

	if str[0] == "" || str[1] == "" {
		return nil, errors.New("invalid format. " + " format: " + mask)
	}

	return &Config{
		Host:     str[0] + ":" + str[1],
		Login:    str[2],
		Password: str[3],
		Disable:  false,
	}, nil
}
