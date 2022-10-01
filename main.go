package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/adamnasrudin03/student-service/configs"
	"github.com/adamnasrudin03/student-service/controller"
	"github.com/adamnasrudin03/student-service/entity"
	"github.com/adamnasrudin03/student-service/utils"
)

var (
	studentController controller.StudentController = controller.NewStudentController(&entity.Datastore{})
)

func main() {
	config := configs.GetInstance()
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		switch r.URL.Path {
		case "/":
			utils.APIResponseSuccess(w, r, "Welcome my application", http.StatusOK, "success", nil)
			fmt.Println(time.Now().Format("02-Jan-2006 15:04:05"), "Endpoint Hit: [Home Page]")
		case "/healty-check":
			utils.APIResponseSuccess(w, r, "server is healty", http.StatusOK, "success", nil)
			fmt.Println(time.Now().Format("02-Jan-2006 15:04:05"), "Endpoint Hit: [Healty Check]")
		default:
			fmt.Println(time.Now().Format("02-Jan-2006 15:04:05"), "Endpoint Hit: [Not Found] "+r.URL.Path)
			utils.NotFound(w, r)
		}
	})

	mux.HandleFunc("/api/v1/students", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		switch r.Method {
		case "GET":
			studentController.GetAll(w, r)
		default:
			utils.NotFound(w, r)
		}
	})

	mux.HandleFunc("/api/v1/student", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		switch r.Method {
		case "POST":
			studentController.Create(w, r)
		default:
			utils.NotFound(w, r)
		}
	})

	mux.HandleFunc("/api/v1/student/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		switch r.Method {
		case "GET":
			studentController.GetByID(w, r)
		case "PUT":
			studentController.UpdateByID(w, r)
		case "DELETE":
			studentController.DeleteByID(w, r)
		default:
			utils.NotFound(w, r)
		}
	})

	fmt.Printf("Server is running on port %v\n", config.Appconfig.Port)
	err := http.ListenAndServe(":"+config.Appconfig.Port, mux)
	if err != nil {
		log.Fatal(err)
	}

}
