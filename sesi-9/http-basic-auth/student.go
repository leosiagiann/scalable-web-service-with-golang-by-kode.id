package main

var stundents = []*Student{}

type Student struct {
	Id    string
	Name  string
	Grade string
}

func init() {
	stundents = append(stundents, &Student{Id: "1", Name: "John", Grade: "A"})
	stundents = append(stundents, &Student{Id: "2", Name: "Mary", Grade: "B"})
	stundents = append(stundents, &Student{Id: "3", Name: "Mike", Grade: "C"})
}

func GetStudents() []*Student {
	return stundents
}

func SelectStudent(id string) *Student {
	for _, student := range stundents {
		if student.Id == id {
			return student
		}
	}
	return nil
}
