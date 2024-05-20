package okex

import (
	"net/url"

	"github.com/hardstylez72/cry/internal/exchange/okex/driver/responses"
	"github.com/pkg/errors"
)

var (
	ErrIpNotInWhiteList      = errors.New("IP is not in white list")
	ErrNetworkError          = errors.New("network error")
	ErrAddressNotInWhiteList = errors.New("wallet address is not in white list")
)

func WrapErr(err error) error {
	if err == nil {
		return nil
	}
	switch err.(type) {
	case *url.Error:
		return errors.Wrap(ErrNetworkError, err.Error())
	case *responses.Basic:
		uw := err.(*responses.Basic)
		switch uw.Code {
		case 50110:
			return errors.Wrap(ErrIpNotInWhiteList, err.Error())
		case 58207:
			return errors.Wrap(ErrAddressNotInWhiteList, err.Error())
		default:
			return errors.Wrap(err, "unknown error")
		}
	default:
		return err
	}
}
