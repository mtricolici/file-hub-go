package web

import (
	"fmt"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/mtricolici/file-hub-go/config"
	"github.com/mtricolici/file-hub-go/web/controllers"

	"github.com/kataras/iris/v12/sessions"
)

const (
	sessionCookieName = "zuzu"
)

func registerControllers(app *iris.Application) {
	mvc.New(app.Party("/")).Handle(new(controllers.HomeController))
	mvc.New(app.Party("/login")).Handle(new(controllers.LoginController))
}

func startWebServer(app *iris.Application) {
	listen := fmt.Sprintf("%s:%d",
		config.Get().ListenInterface(), config.Get().ListenPort())
	fmt.Printf("Start web app on '%s' ...\n", listen)

	err := app.Listen(listen)
	if err != nil {
		panic(err)
	}
}

func InitAndStart() {
	app := iris.New()

	// configure session
	sess := sessions.New(sessions.Config{
		Cookie: sessionCookieName,
		// CookieSecureTLS: true,
		AllowReclaim: true,
		//Expires:      -1,
	})

	app.Use(sess.Handler())

	// Register HTML templates
	app.RegisterView(iris.HTML(config.Get().TemplatesDir(), ".html"))

	registerControllers(app)
	startWebServer(app)
}
