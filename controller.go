package i18n

import (
	"gopkg.in/orivil/orivil.v1"
)

const (

// cookie max age
	Year = 60 * 60 * 24 * 365
)

type Controller struct {
	*orivil.App
}

// set the cookie to client
//
// @route {get}/setlang/::language
func (c *Controller) Setlang() {

	name := c.Params["language"]

	if shortLang, ok := Config.Languages[name]; ok {

		if name != Config.DefaultLang {

			c.SetCookie(Config.CookieKey, shortLang, Year)
		} else {

			c.SetCookie(Config.CookieKey, shortLang, 0)
		}

		c.With("success", "success")
	}
}
