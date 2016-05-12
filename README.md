# i18n

Orivil I18n Bundle

## Version

```
v1
```

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

## Generate I18n View File

Run Server, it will auto generate i18n view file under "view" dir in each bundle.

## Uninstall

We should uninstall i18n bundle in development environment.

```GO
server.RegisterBundle (

    // new(i18n.Register),
)
```