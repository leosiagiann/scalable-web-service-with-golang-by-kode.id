package routers

import (
	"encoding/json"
	"lat-http-request/controllers"
	"lat-http-request/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.GET("/students", controllers.GetStudents)
	router.POST("/student/:studentID", controllers.GetStudent)
	router.POST("/student", controllers.CreateStudent)
	router.PUT("/student/:studentID", controllers.UpdateStudent)
	router.DELETE("/student/:studentID", controllers.DeleteStudent)

	return router
}

func GetMidStudent(w http.ResponseWriter, r *http.Request) {
	if !middleware.Auth(w, r) {
		return
	}

	if id := r.URL.Query().Get("id"); id != "" {
		OutputJSON(w, SelectStudent(id))
		return
	}

	OutputJSON(w, controllers.GetStudents())
}

func OutputJSON(w http.ResponseWriter, data interface{}) {
	res, err := json.Marshal(data)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
