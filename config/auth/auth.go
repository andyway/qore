package auth

import (
	"time"

	"github.com/andyway/qore/config"
	"github.com/andyway/qore/config/bindatafs"
	"github.com/andyway/qore/config/db"
	"github.com/andyway/qore/models/users"
	"github.com/qor/auth"
	"github.com/qor/auth/authority"
	"github.com/qor/auth/providers/facebook"
	"github.com/qor/auth/providers/github"
	"github.com/qor/auth/providers/google"
	"github.com/qor/auth/providers/twitter"
	"github.com/qor/auth_themes/clean"
	"github.com/qor/render"
)

var (
	// Auth initialize Auth for Authentication
	Auth = clean.New(&auth.Config{
		DB:         db.DB,
		Mailer:     config.Mailer,
		Render:     render.New(&render.Config{AssetFileSystem: bindatafs.AssetFS.NameSpace("auth")}),
		UserModel:  users.User{},
		Redirector: auth.Redirector{RedirectBack: config.RedirectBack},
	})

	// Authority initialize Authority for Authorization
	Authority = authority.New(&authority.Config{
		Auth: Auth,
	})
)

func init() {
	Auth.RegisterProvider(github.New(&config.Config.Github))
	Auth.RegisterProvider(google.New(&config.Config.Google))
	Auth.RegisterProvider(facebook.New(&config.Config.Facebook))
	Auth.RegisterProvider(twitter.New(&config.Config.Twitter))

	Authority.Register("logged_in_half_hour", authority.Rule{TimeoutSinceLastLogin: time.Minute * 30})
}
