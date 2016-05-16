package i18n

import (
	"gopkg.in/orivil/event.v0"
	"gopkg.in/orivil/middle.v0"
	"gopkg.in/orivil/router.v0"
	"gopkg.in/orivil/service.v0"
	"gopkg.in/orivil/orivil.v1"
)

type Register struct {}

func (*Register) RegisterRoute(c *router.Container) {
	c.Add("{get}/", func() interface{} { return new(Controller) })
}

func (*Register) SetMiddle(bag *middle.Bag) {}

func (*Register) RegisterService(c *service.Container) {

	c.Add(orivil.SvcI18nFilter, func(c *service.Container) interface{} {

		a := c.Get(orivil.SvcApp).(*orivil.App)

		return NewFilter(a)
	})
}

func (*Register) RegisterMiddle(c *middle.Container) {}

func (*Register) Boot(c *service.Container) {}

func (*Register) AddServerListener(d *event.Dispatcher) {

	// auto generate I18n view files
	d.AddListener(new(Listener))
}
