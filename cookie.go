package cdpcookies

import (
	"context"

	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
)

// CookieParams represent cookies to be set.
type CookiesParams struct {
	Cookies []*network.CookieParam
}

// SetCookies sets given cookies against the provided context and return chromdedp.Action.
func (s *CookiesParams) SetCookies(ctx context.Context) chromedp.Action {
	return chromedp.ActionFunc(func(ctx context.Context) error {
		err := network.SetCookies(s.Cookies).Do(ctx)
		if err != nil {
			return err
		}

		return nil
	})
}
