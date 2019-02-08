package api_server

import (
	"app-version-manager/api_server/admins"
	"app-version-manager/api_server/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func New() *APIServer {
	return &APIServer{
		engine: gin.Default(),
	}
}

type APIServer struct {
	engine *gin.Engine
}

func (a *APIServer) Start() {
	//ports := []string{":25000", ":25001"}
	//for _, v := range ports {
	//	go func(port string) { //每个端口都扔进一个goroutine中去监听
	//		mux := http.NewServeMux()
	//		mux.HandleFunc("/", handler1)
	//		http.ListenAndServe(port, mux)
	//	}(v)
	//}
	//select {}

	go func() {
		http.ListenAndServe(":8080", http.FileServer(http.Dir("./dist")))
	}()
	go func() {
		a.registry()
		a.engine.Run(":5000")
	}()
	select {}
}

func (a *APIServer) registry() {
	APIServerInit(a.engine)

}

func (a *APIServer) init() {

}

func APIServerInit(r *gin.Engine) {

	r.Use(controller.ConfigHeader())
	v1 := r.Group("/v1/api")
	adminsRegistry(v1)
	v1.Use(controller.AuthRequired())
	checkToken(v1)

}

func checkToken(r *gin.RouterGroup) {
	r.GET("/test_token", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code": 0,
			"msg": "token is ok",
			"data": "",
		})
	})
}

func adminsRegistry(r *gin.RouterGroup) {
	r.POST("/admins/sign_in", admins.SignIn)
	r.POST("/admins/sign_up", admins.SignUp)
}






