package entity

// student - Our struct for all students
type Student struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Age  int8   `json:"age"`
}

type Datastore struct {
	Students []Student
}

type Page struct {
	Size      int64 `json:"size"`
	TotalData int64 `json:"total_data"`
	TotalPage int64 `json:"total_page"`
	Current   int64 `json:"current"`
}
