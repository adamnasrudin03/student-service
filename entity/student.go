package entity

// student - Our struct for all students
type Student struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Age  int8   `json:"age"`
}
