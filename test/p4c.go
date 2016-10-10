package main

import (
	"fmt"

	"github.com/insionng/vodka"
	"github.com/insionng/vodka/engine/standard"
	//"github.com/insionng/vodka/engine/fasthttp"
	"github.com/insionng/vodka/middleware"
	"github.com/vodka-contrib/cache"
	"github.com/vodka-contrib/captcha"
	"github.com/vodka-contrib/pongor"
)

func main() {
	v := vodka.New()
	v.Use(middleware.Logger())
	v.Use(middleware.Recover())
	v.Use(cache.VodkaCacher(cache.Options{Adapter: "memory"}))
	v.Use(captcha.Captchaer())

	r := pongor.Renderor()
	v.SetRenderer(r)

	v.Get("/", func(self vodka.Context) error {
		if cpt := self.Get("Captcha"); cpt != nil {
			fmt.Println("Got:")
			fmt.Println(cpt)
		} else {
			fmt.Println("CPT IS NIL!")
		}

		return self.Render(200, "index.html", map[string]interface{}{
			"title":   "你好，世界",
			"Captcha": self.Get("Captcha"),
		})
	})

	//v.Run(fasthttp.New(":7891"))
	v.Run(standard.New(":1987"))
}
