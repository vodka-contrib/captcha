// Copyright 2014 The Macaron Authors
// Copyright 2016 The Vodka Authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package captcha

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/insionng/vodka"
	"github.com/insionng/vodka/engine/standard"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/vodka-contrib/cache"
	"gopkg.in/macaron.v1"
)

func Test_Version(t *testing.T) {
	Convey("Get version", t, func() {
		So(Version(), ShouldEqual, _VERSION)
	})
}

func Test_Captcha(t *testing.T) {
	Convey("Captch service", t, func() {

		v := vodka.New()
		v.Use(cache.Cacher(cache.Options{Adapter: "memory"}))
		v.Use(Captchaer())
		v.Get("/", func(c vodka.Context) error {
			return c.String(http.StatusOK, "")
		})
		go v.Run(standard.New(":7891"))

		resp := httptest.NewRecorder()
		req, err := http.NewRequest("GET", "http://localhost:7891/", nil)
		So(err, ShouldBeNil)

		m := macaron.New()
		m.ServeHTTP(resp, req)

	})

}
