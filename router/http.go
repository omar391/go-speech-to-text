//Package router serve service controllers
package router

import (
	"net/http"
	"stt-service/models"
	"stt-service/service"
	"stt-service/utils"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/secure"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

//Start start http web server + API endpoints
func Start(addr string) error {
	return setupRouter().Run(addr)
}

//setupRouter setup global route for api server
func setupRouter() *gin.Engine {
	r := gin.Default()

	//add security middleware
	attachSecurityLayers(r)

	//setup cookies based sessions
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	//setup API specific routers
	setupApiRouter(r)

	//setup WEB specific routers
	setupWebRouters(r)

	return r
}

//add security middleware
func attachSecurityLayers(r *gin.Engine) {
	//setting-up cors headers
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	//setting-up security headers
	r.Use(secure.New(secure.Config{
		FrameDeny:             true,
		ContentTypeNosniff:    true,
		BrowserXssFilter:      true,
		ContentSecurityPolicy: "default-src 'self'",
		IENoOpen:              true,
		ReferrerPolicy:        "strict-origin-when-cross-origin",
	}))
}

//setup only API specific endpoints
func setupApiRouter(r *gin.Engine) {
	r.POST("/login", func(c *gin.Context) {
		email, fe := c.GetPostForm("email")
		pass, fp := c.GetPostForm("password")

		response := models.ApiBooleanResponse{}
		if !fe || !fp {
			response.IsScuess = false
			response.Msg = "Please provide ALL login credentials!"

		} else {
			response, id := service.Login(email, pass)
			if response.IsScuess {
				response.Token, _ = utils.GenerateJWT(id)
			}
		}
		c.JSON(http.StatusOK, response)
	})

	r.POST("/register", func(c *gin.Context) {

		email, fe := c.GetPostForm("email")
		pass, fp := c.GetPostForm("password")
		name, fn := c.GetPostForm("name")

		response := models.ApiBooleanResponse{}
		if !fe || !fp || !fn {
			response.IsScuess = false
			response.Msg = "Please provide ALL registration credentials!"

		} else {
			var id uint
			response, id = service.AddNewUser(&models.User{Email: email, Password: pass, Name: name})
			if response.IsScuess {
				response.Token, _ = utils.GenerateJWT(id)
				response.Msg = "Registration successfull!"
			}
		}
		c.JSON(http.StatusOK, response)
	})

	r.GET("/search", func(c *gin.Context) {
		// // text, _ := c.GetQuery("text")
		// // list := modules.Search(strings.ToLower(text))

		// c.JSON(http.StatusOK, list)
	})

	r.GET("/all-data", func(c *gin.Context) {
		// text, _ := c.GetQuery("text")
		// list := modules.Search(strings.ToLower(text))

		// c.JSON(http.StatusOK, list)
	})
}

//setup only WEB specific endpoints
func setupWebRouters(r *gin.Engine) {

	//set static assets to load directly
	r.Static("/web/static", "./static")
}
