package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Student struct {
	StudentID string `json:"student_id"`
	Name      string `json:"name"`
}

var Students = []Student{}

func init() {
	Students = append(Students, Student{
		StudentID: "s-1",
		Name:      "John Doe",
	})
	Students = append(Students, Student{
		StudentID: "s-2",
		Name:      "Jane Doe",
	})
}

func CreateStudent(c *gin.Context) {
	var newStudent Student

	if err := c.ShouldBindJSON(&newStudent); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newStudent.StudentID = fmt.Sprintf("s-%d", len(Students)+1)
	Students = append(Students, newStudent)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Student created successfully",
		"student": newStudent,
	})
}

func GetStudents(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"students": Students,
	})
}

func GetStudent(c *gin.Context) {
	StudentID := c.Param("studentID")
	isFInd := false
	var student Student

	for i, s := range Students {
		if s.StudentID == StudentID {
			student = Students[i]
			isFInd = true
			break
		}
	}

	if !isFInd {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "Student not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"student": student,
	})
}

func UpdateStudent(c *gin.Context) {
	StudentID := c.Param("studentID")
	isFInd := false
	var updateStudent Student

	if err := c.ShouldBindJSON(&updateStudent); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for i, s := range Students {
		if s.StudentID == StudentID {
			Students[i] = updateStudent
			Students[i].StudentID = StudentID
			isFInd = true
			break
		}
	}

	if !isFInd {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "Student not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Student updated successfully",
	})
}

func DeleteStudent(c *gin.Context) {
	StudentID := c.Param("studentID")
	isFInd := false
	var deleteStudent Student

	for i, s := range Students {
		if s.StudentID == StudentID {
			deleteStudent = Students[i]
			Students = append(Students[:i], Students[i+1:]...)
			isFInd = true
			break
		}
	}

	if !isFInd {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "Student not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Student deleted successfully",
		"student": deleteStudent,
	})
}
