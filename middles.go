package i18n

import (
	"gopkg.in/orivil/orivil.v1"
)

// DataSender send all languages and current language to view template.
func DataSender(a *orivil.App) {

	if filter, ok := a.Get(orivil.SvcI18nFilter).(*Filter); ok {

		currentLang := GetFullName(filter.currentLang)
		a.With("currentLang", currentLang)
		a.With("i18nlangs", Config.Languages)
	}
}
