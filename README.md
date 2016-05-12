# i18n

Orivil I18n bundle

## Version

```
v1
```

## Install

* make dir "i18n" under your project bundle dir
* cd "i18n" dir
* git init
* git remote add i18n https://github.com/orivil/i18n
* git pull i18n v1

## Config "DataSender" Middleware

"DataSender" middleware will send languages data to template:

```GO
func (this *Controller)SetMiddle(bag *middle.Bag) {

    bag.Set(i18n.MidDataSender).OnlyActions("Index")
} 
```

Display languages in template:

```html
<span>{{$.currentLang}}</span>

<ul class="dropdown-menu">

    {{/* range all languages */}}
    {{range $lang, $shortName := .i18nlangs}}
        {{/* except current lang */}}
        {{if ne $lang $.currentLang}}
            <li><a class="lang-set" href="/setlang/{{$lang}}">{{$lang}}</a></li>
        {{end}}
    {{end}}
</ul>
```

## Sent Ajax Request

```html
<script>
  $(document).ready(function() {
      $(".lang-set").on("click", function (e) {
    
          e.preventDefault()
    
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