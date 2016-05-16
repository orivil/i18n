# I18n Bundle

## Introduction

**I18n Bundle** provides a full-featured i18n management system. Automatically generate template files and message files according to the language configuration. Make your project become international immediately, even if your have finish your project.

## Step By Step

### Install

* make dir "i18n" under your bundle dir
* cd "i18n" dir
* git init
* git remote add i18n https://github.com/orivil/i18n
* git pull i18n master

### Configure Languages

Under `i18n/config` directory has a `i18n.yml` file:
```yaml
# all languages
languages:
  简体中文: "zh-CN"
  English: "en"
  Französisch: "fr"

# set default language
defaultlang: "English"
```

### Display Languages

Use `DataSender` function to send data to template:

```GO
i18n.DataSender(app) 
```

Display data:

```html
{{if $.currentLang}}
<span>{{$.currentLang}}</span>
<ul>
    {{/* range all languages */}}
    {{range $lang, $shortName := .i18nlangs}}
        {{/* except current lang */}}
        {{if ne $lang $.currentLang}}
            <li><a class="lang-set" href="/setlang/{{$lang}}">{{$lang}}</a></li>
        {{end}}
    {{end}}
</ul>
{{end}}
```

### Sent Ajax Request To Set Language

```html
<script>
  $(document).ready(function() {
      $(".lang-set").on("click", function (e) {
    
          e.preventDefault();
    
          $.ajax({
              type: "GET",
    
              // send request to i18n controller
              url: $(this).attr("href"),
    
              // server will return a one-year i18n cookie
              success: function () {
    
                  window.location.reload(true)
              }
          })
      })
  })
</script>
```

### Register I18n Bundle

```GO
server.RegisterBundle (

    new(i18n.Register),
)
```

### Generate I18n View Files

Before the server starts:
```
└── view
    └── index.tmpl
```

After the server started:

```
└── view
    ├── i18n
    │   ├── en
    │   │   └── index.tmpl
    │   ├── fr
    │   │   └── index.tmpl
    │   └── zh-CN
    │       └── index.tmpl
    └── index.tmpl
```
> **Note:** now we can translate them.

### Generate I18n Messages

Suppose we have messages:
```GO
msgUnSignIn := "Please sign in"
msgIncorrectPass := "Incorrect password"
```

Add messages to configuration generator:

```GO
i18n.AddMsgs (
    msgUnSignIn,
    msgIncorrectPass,
)
```

After the server started, it will generate the default language configuration file `en.yml` under `/i18n/config/msgs` directory:

```
├── i18n
│   ├── config
│   │   ├── i18n.yml
│   │   └── msgs
│   │       └── en.yml

```
Here is the generated file `en.yml`:
```yaml
Incorrect password:
  fr: ""
  zh-CN: ""
Please sign in:
  fr: ""
  zh-CN: ""
```

Translate all messages:
```yaml
Incorrect password:
  fr: "Mot de passe incorrect"
  zh-CN: "密码错误"
Please sign in:
  fr: "S'il vous plaît vous connecter"
  zh-CN: "请登陆"
```

### Get I18n Message

```GO
msg := app.FilterI18n(msgUnSignIn)

// if customer language is "zh-CN", msg == "请登陆".
// if customer language is "en",    msg == "Please sign in"
// if customer language is "fr",    msg == "S'il vous plaît vous connecter"
```

## Uninstall

We should uninstall i18n bundle in development environment.

```GO
server.RegisterBundle (

    // new(i18n.Register),
)
```