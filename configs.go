package i18n

import (
	"gopkg.in/orivil/orivil.v1"
	"gopkg.in/orivil/config.v0"
	"path/filepath"
)

var Config = &struct {
	Languages           map[string]string
	DefaultLang         string
	CookieKey           string
	Auto_generate_files bool
}{}

// sortName: fullName
var langs map[string]string

// use shortName for get fullName
func GetFullName(shortName string) string {

	return langs[shortName]
}

func init() {

	cfg := config.NewConfig(filepath.Join(orivil.DirBundle, "i18n", "config"))

	cfg.ReadStruct("i18n.yml", Config)

	langs = make(map[string]string, len(Config.Languages))

	for fullName, shortName := range Config.Languages {

		langs[shortName] = fullName
	}
}
