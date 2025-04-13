package server

import (
	"github.com/Le-BlitzZz/onchain-auth-app/internal/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

// registerRoutes registers the routes for handling HTTP requests with the built-in web server.
func registerRoutes(router *gin.Engine, conf *config.Config) {
	registerStaticRoutes(router, conf)
	registerWebAppRoutes(router, conf)
}

func registerStaticRoutes(router *gin.Engine, conf *config.Config) {
	// Redirects to the login page.
	login := func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	router.Any("/", login)

	// Server static assets like js and css.
	if dir := conf.StaticPath(); dir != "" {
		group := router.Group(config.StaticUri)
		group.Static("", dir)
	}
}

func registerWebAppRoutes(router *gin.Engine, conf *config.Config) {
	ui := func(c *gin.Context) {
		// Get Client configuration.
		clientConfig := conf.GetClientConfig()

		// Set bootstrap template values.
		values := gin.H{
			"config": clientConfig,
		}

		// Render bootstrap template.
		c.HTML(http.StatusOK, conf.TemplateName(), values)
	}

	router.Any("/login", ui)
}
