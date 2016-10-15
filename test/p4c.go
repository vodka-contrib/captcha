package main

import (
	"fmt"

	"github.com/insionng/vodka"
	//"github.com/insionng/vodka/engine/standard"
	"github.com/insionng/vodka/engine/fasthttp"
	"github.com/insionng/vodka/middleware"
	"github.com/vodka-contrib/cache"
	"github.com/vodka-contrib/captcha"
	"github.com/vodka-contrib/pongor"
)

func main() {
	v := vodka.New()
	v.Use(middleware.Logger())
	v.Use(middleware.Recover())
	v.Use(cache.Cacher(cache.Options{Adapter: "memory"}))
	v.Use(captcha.Captchaer())
	v.SetRenderer(pongor.Renderor())

	v.Get("/", func(self vodka.Context) error {
		if cpt := self.Get("Captcha"); cpt != nil {
			fmt.Println("Got:", cpt)
		} else {
			fmt.Println("Captcha is nil!")
		}

		self.Set("title", "你好，世界")
		return self.Render(200, "index.html")
	})

	v.Run(fasthttp.New(":7891"))
	//v.Run(standard.New(":1987"))
}
