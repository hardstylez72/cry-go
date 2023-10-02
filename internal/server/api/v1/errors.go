package v1

import (
	"github.com/hardstylez72/cry/internal/exchange/okex"
	"github.com/pkg/errors"
)

var (
	ErrIpNotInWhiteList         = errors.New("IP не в белом списке на okex. проверьте настройки API-KEY для okex и сравните их с аккаунтом в разделе 'Биржи'")
	ErrNetworkOkex              = errors.New("Okex отклонил запрос. Возможно прокси - нестабилен, возможно okex недоступен")
	ErrWalletNotInWhiteListOkex = errors.New("Адрес на который вы пытаетесь вывести монету не в белом списке адресов на которые доступен вывод на okex")
)

func Rus(err error) error {
	switch true {
	case err == nil:
		return nil
	case errors.Is(err, okex.ErrIpNotInWhiteList):
		return ErrIpNotInWhiteList
	case errors.Is(err, okex.ErrNetworkError):
		return ErrNetworkOkex
	case errors.Is(err, okex.ErrAddressNotInWhiteList):
		return ErrWalletNotInWhiteListOkex

	default:
		return err
	}

}
