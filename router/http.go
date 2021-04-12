//Package router serve service controllers
package router

import (
	"fmt"
	"stt-service/conf"
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
	r.Use(sessions.Sessions("jwt_session", store))

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

//save authenticated sessions to
func saveAuthSession(c *gin.Context, user_id uint, token string) {
	user_id_str := fmt.Sprint(user_id)
	c.SetCookie("gin_user", user_id_str, conf.Config.WEB_SESSION_TIMEOUT, "/", "", false, true)
	session := sessions.Default(c)
	session.Set(user_id_str, token)
	session.Save()
}

//check if the session is authenticated
func isSessionAuthenticated(c *gin.Context, token string) bool {
	user_id, err := getUserId(c)
	return err == nil && sessions.Default(c).Get(user_id) == token && utils.VerifyJWT(token)
}

// retrive email from the cookie;
// do not use this without calling isSessionAuthenticated method first
func getUserId(c *gin.Context) (string, error) {
	return c.Cookie("gin_user")
}

//setup only API specific endpoints
func setupApiRouter(r *gin.Engine) {

	//Login API endpoint
	r.POST("/login", login)

	//Registration API endpoint
	r.POST("/register", register)

	//File Upload and Audio transcription API endpoint
	r.POST("/transcribe", transcribe)

	// Get all the uploaded transcribed data from current logged-in user
	r.POST("/all-data", search_all)

	// Get filtered data from search term from current logged-in user
	r.POST("/filter", filter)
}

//setup only WEB specific endpoints
func setupWebRouters(r *gin.Engine) {

	//set static assets to load directly
	r.Static("/web/", "./web/static")
}
