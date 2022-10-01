package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/adamnasrudin03/student-service/entity"
	"github.com/adamnasrudin03/student-service/utils"
)

//StudentController interface is a contract what this controller can do
type StudentController interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	GetByID(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	DeleteByID(w http.ResponseWriter, r *http.Request)
	UpdateByID(w http.ResponseWriter, r *http.Request)
}

type studentController struct {
	Store *entity.Datastore
}

//NewStudentController creates a new instance of StudentController
func NewStudentController(s *entity.Datastore) StudentController {
	return &studentController{
		Store: s,
	}
}

func (h *studentController) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println(time.Now().Format("02-Jan-2006 15:04:05"), "Endpoint Hit: [Student] Create")
	mutex := sync.RWMutex{}
	student := entity.Student{}
	sliceID := []int64{}

	mutex.Lock()

	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &student)
	student.Name = strings.Trim(student.Name, " ")

	if strings.Compare(student.Name, "") == 0 {
		mutex.Unlock()
		utils.APIResponseError(w, r, "request name is required", http.StatusBadRequest)
		return
	}

	if student.Age <= 0 {
		mutex.Unlock()
		utils.APIResponseError(w, r, "request name is required, number type, and more than 0", http.StatusBadRequest)
		return
	}

	for _, student := range h.Store.Students {
		sliceID = append(sliceID, student.Id)
	}

	_, max := utils.GetMinAndMax(sliceID)
	student.Id = max + 1

	h.Store.Students = append(h.Store.Students, student)

	mutex.Unlock()
	utils.APIResponseSuccess(w, r, "success create student", http.StatusOK, "success", student)
}

func (h *studentController) GetAll(w http.ResponseWriter, r *http.Request) {
	fmt.Println(time.Now().Format("02-Jan-2006 15:04:05"), "Endpoint Hit: [Student] GetAll")
	sizeQuery := 10
	pageQuery := 1

	sizeParam := r.URL.Query().Get("size")
	if strings.Compare(sizeParam, "") != 0 {
		value, err := strconv.Atoi(sizeParam)
		if err != nil {
			utils.APIResponseError(w, r, "query param size must be number", http.StatusBadRequest)
			return
		}
		sizeQuery = value
	}

	pageParam := r.URL.Query().Get("page")
	if strings.Compare(pageParam, "") != 0 {
		value, err := strconv.Atoi(pageParam)
		if err != nil {
			utils.APIResponseError(w, r, "query param page must be number", http.StatusBadRequest)
			return
		}
		pageQuery = value
	}

	mutex := sync.RWMutex{}

	totalPage := int64(len(h.Store.Students)) / int64(sizeQuery)
	if totalPage == 0 {
		totalPage = 1
	}

	page := entity.Page{
		Size:      int64(sizeQuery),
		TotalData: int64(len(h.Store.Students)),
		TotalPage: totalPage,
		Current:   int64(pageQuery),
	}

	mutex.Lock()
	data := h.Store.Students
	if data == nil {
		utils.APIResponseListSuccess(w, r, "success get list students", http.StatusOK, "success", []entity.Student{}, page)
		mutex.Unlock()
		return
	}

	start := int((pageQuery - 1) * sizeQuery)
	if start > len(h.Store.Students) {
		start = len(h.Store.Students)
	}

	end := start + int(sizeQuery)
	if end > len(h.Store.Students) {
		end = len(h.Store.Students)
	}

	data = h.Store.Students[start:end]

	mutex.Unlock()
	utils.APIResponseListSuccess(w, r, "success get list students", http.StatusOK, "success", data, page)
}

func (h *studentController) GetByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println(time.Now().Format("02-Jan-2006 15:04:05"), "Endpoint Hit: [Student] GetByID")
	mutex := sync.RWMutex{}
	getID := regexp.MustCompile(`^\/api/v1/students\/(\d+)$`)
	matches := getID.FindStringSubmatch(r.URL.Path)
	if len(matches) < 2 {
		utils.NotFound(w, r)
		return
	}

	mutex.Lock()
	ID, err := strconv.ParseInt(matches[1], 10, 64)
	if err != nil {
		panic(err)
	}

	for _, student := range h.Store.Students {
		if student.Id == ID {
			utils.APIResponseSuccess(w, r, "success get student", http.StatusOK, "success", student)
			mutex.Unlock()
			return
		}
	}

	mutex.Unlock()
	utils.APIResponseError(w, r, "data not found", http.StatusNotFound)
}

func (h *studentController) UpdateByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println(time.Now().Format("02-Jan-2006 15:04:05"), "Endpoint Hit: [Student] UpdateByID")
	mutex := sync.RWMutex{}
	getID := regexp.MustCompile(`^\/api/v1/students\/(\d+)$`)
	matches := getID.FindStringSubmatch(r.URL.Path)
	if len(matches) < 2 {
		utils.NotFound(w, r)
		return
	}

	mutex.Lock()
	ID, err := strconv.ParseInt(matches[1], 10, 64)
	if err != nil {
		panic(err)
	}

	reqBody, _ := ioutil.ReadAll(r.Body)
	var studentReq entity.Student
	json.Unmarshal(reqBody, &studentReq)
	studentReq.Name = strings.Trim(studentReq.Name, " ")

	for index, student := range h.Store.Students {
		if student.Id == ID {

			if strings.Compare(studentReq.Name, "") != 0 {
				student.Name = studentReq.Name
			}

			if studentReq.Age > 0 {
				student.Age = studentReq.Age
			}

			h.Store.Students[index] = student
			mutex.Unlock()
			utils.APIResponseSuccess(w, r, "success update student", http.StatusOK, "success", student)
			return
		}
	}

	mutex.Unlock()
	utils.APIResponseError(w, r, "data not found", http.StatusNotFound)
}

func (h *studentController) DeleteByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println(time.Now().Format("02-Jan-2006 15:04:05"), "Endpoint Hit: [Student] DeleteByID ")
	mutex := sync.RWMutex{}
	getID := regexp.MustCompile(`^\/api/v1/students\/(\d+)$`)
	matches := getID.FindStringSubmatch(r.URL.Path)
	if len(matches) < 2 {
		utils.NotFound(w, r)
		return
	}

	mutex.Lock()
	ID, err := strconv.ParseInt(matches[1], 10, 64)
	if err != nil {
		panic(err)
	}

	for index, student := range h.Store.Students {
		if student.Id == ID {
			h.Store.Students = append(h.Store.Students[:index], h.Store.Students[index+1:]...)
			mutex.Unlock()
			utils.APIResponseSuccess(w, r, "success delete student", http.StatusOK, "success", nil)
			return
		}
	}

	mutex.Unlock()
	utils.APIResponseError(w, r, "data not found", http.StatusNotFound)

}
