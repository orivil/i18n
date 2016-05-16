# i18n

Orivil I18n Bundle

## Version

```
v1
```

## Introduce

**I18n Bundle** provide a full-featured i18n management system. Automatically generate template files and message files according to the language configuration. Make your project become international immediately, even if your project has been completed.

## Install

* make dir "i18n" under your bundle dir
* cd "i18n" dir
* git init
* git remote add i18n https://github.com/orivil/i18n
* git pull i18n v1

## Display Languages

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

## Sent Ajax Request To Set Language

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

## Register I18n Bundle

```GO
server.RegisterBundle (

    new(i18n.Register),
)
```

## Generate I18n View Files

Run Server, it will auto generate i18n view files under "view" dir in each bundle.

## Uninstall

We should uninstall i18n bundle in development environment.

```GO
server.RegisterBundle (

    // new(i18n.Register),
)
```