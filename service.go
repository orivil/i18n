// Copyright 2016 orivil Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package i18n

import (
	"gopkg.in/orivil/cache.v0"
	"regexp"
	"gopkg.in/orivil/orivil.v1"
	"net/http"
	"path/filepath"
	"gopkg.in/orivil/validator.v0"
)

var I18n = &i18n{}

type Translator interface {
	Translate(src, srcLang, dstLang string) (dst string)
}

// Implement orivil.I18nFilter interface, register Filter structure as a service.
type Filter struct {
	currentLang string
	defaultLang string
}

func NewFilter(a *orivil.App) *Filter {

	var currentLang string

	var defaultLang = Config.Languages[Config.DefaultLang]

	// read language from cookie
	if c, err := a.Request.Cookie(Config.CookieKey); err != http.ErrNoCookie {

		currentLang = c.Value
	} else {

		currentLang = defaultLang
	}

	a.Response.Header().Set("Content-Language", currentLang)

	return &Filter {

		currentLang:currentLang,
		defaultLang:defaultLang,
	}
}

func (this *Filter) FilterMsg(src string) (dst string) {

	if this.currentLang != this.defaultLang {

		return I18n.Filter(src, this.currentLang)
	} else {

		return src
	}
}

func (this *Filter) ViewSubDir() (dir string) {


	if this.currentLang != this.defaultLang {

		return filepath.Join(i18nDir, this.currentLang)
	} else {

		return ""
	}
}

type i18n struct {
	cacheDir    string
	langs       []string
	defaultLang string
	currentLang string
	caches      map[string]map[string]string
	validators  []*validator.Validator
	msgs        []string
	translator  Translator
}

func (this *i18n) Init(cacheDir, defaultLang string, langs []string) {

	this.cacheDir = cacheDir
	this.defaultLang = defaultLang
	this.langs = langs
	this.translator = &defaultTranslator{}
	this.readConfig()
}

type defaultTranslator struct{}

// TODO: build a default translator
func (this *defaultTranslator) Translate(src, srcLang, dstLang string) (dst string) {
	return ""
}

// ReadConfig for read current language config
func (this *i18n) readConfig() {
	this.caches = map[string]map[string]string{}
	cacher, err := cache.NewJsonCache(this.cacheDir, this.defaultLang + ".yml")
	if err != nil {
		panic(err)
	}

	err = cacher.Read(this.caches)
	if err != nil {
		panic(err)
	}
}

func (this *i18n) Filter(msg, currentLang string) (i18nMsg string) {

	if m, ok := this.caches[msg][currentLang]; ok && m != "" {
		return m
	} else {
		return msg
	}
}

// SetTranslator set customer translator
func (this *i18n) SetTranslator(t Translator) {
	this.translator = t
}

func (this *i18n) AddMsgs(msg ...string) {
	this.msgs = append(this.msgs, msg...)
}

// AddValidator for store validator messages
func (this *i18n) AddValidators(v ...*validator.Validator) {

	this.validators = append(this.validators, v...)
}

func AddMsgs(msg ...string) {

	I18n.AddMsgs(msg...)
}

func AddValidators(v ...*validator.Validator) {

	I18n.AddValidators(v...)
}

// UpdateI18nConfig for updating i18n messages
func UpdateI18nConfig() {
	msgs := GetValidatorMsgs(I18n.validators...)
	compileMsgs(I18n.msgs, msgs)
	var caches = map[string]map[string]string{}
	cache, err := cache.NewJsonCache(I18n.cacheDir, I18n.defaultLang + ".yml")
	if err != nil {
		panic(err)
	}

	err = cache.Read(caches)
	if err != nil {
		panic(err)
	}

	for msg, _ := range msgs {
		for _, lang := range I18n.langs {
			if lang != I18n.defaultLang {
				if caches[msg] == nil {
					caches[msg] = map[string]string{lang: I18n.translator.Translate(msg, I18n.defaultLang, lang)}
				} else if caches[msg][lang] == "" {
					caches[msg][lang] = I18n.translator.Translate(msg, I18n.defaultLang, lang)
				}
			}
		}
	}

	err = cache.Write(caches)
	if err != nil {
		panic(err)
	}
}

func compileMsgs(srcms []string, msgs map[string]bool) {
	for _, msg := range srcms {
		msgs[msg] = true
	}
}

func GetValidatorMsgs(vs ...*validator.Validator) (msgs map[string]bool) {
	msgs = make(map[string]bool, 5)

	for _, v := range vs {
		getMsgs(v.Required, msgs)
		getMsgs(v.Email, msgs)
		getMsgs(v.Confirm, msgs)
		getMsgs(v.SliceRange, msgs)
		getMsgs(v.StringRange, msgs)
		getMsgs(v.NumRange, msgs)
		getMsgs(v.Min, msgs)
		getMsgs(v.Max, msgs)
		getMsgs(v.Regexp, msgs)
	}
	return
}

func getMsgs(data interface{}, msgs map[string]bool) {

	switch mp := data.(type) {
	case map[string]string:
		for _, msg := range mp {
			msgs[msg] = true
		}
	case map[string]map[string]string:
		for _, _msgs := range mp {
			for _, msg := range _msgs {
				msgs[msg] = true
			}
		}
	case map[string]map[string]*regexp.Regexp:
		for _, _msgs := range mp {
			for msg, _ := range _msgs {
				msgs[msg] = true
			}
		}
	default:
		panic("validata.getMsgs: unknown validate field type")
	}
}
