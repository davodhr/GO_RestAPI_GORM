package main

import (
	"net/http"
	"time"

	"github.com/davodhr/GO_RestAPI_GORM/app/models"
	"github.com/gin-gonic/gin"
)

// getAlbums responds with the list of all albums as JSON.
func getData(c *gin.Context) {
	allData := []models.Rectangle{}
	allDataPure := []models.PureRectangle{}
	var tempRect models.Rectangle
	var tempRectPure models.PureRectangle

	models.DB.Find(&allData)

	for i := 0; i < len(allData); i++ {
		tempRect = allData[i]
		tempRectPure.X_rect = tempRect.X_rect
		tempRectPure.Y_rect = tempRect.Y_rect
		tempRectPure.W_rect = tempRect.W_rect
		tempRectPure.H_rect = tempRect.H_rect
		tempRectPure.Time = tempRect.Time
		allDataPure = append(allDataPure, tempRectPure)
	}
	c.IndentedJSON(http.StatusOK, &allDataPure) //200 status code
}

// postData to check validation and overlap of Rectangles and if true, save them in db
func postData(c *gin.Context) {
	var newData models.Data_in

	// Call BindJSON to bind the received JSON to newData
	if err := c.BindJSON(&newData); err != nil {

		return
	}

	t1 := time.Now()

	s := t1.Format("2006-01-02 15:04:05")

	inputs := newData.Inputs
	main_rect := newData.Main_Rectangle
	if !validateRectangle(main_rect) {
		// "Main rect is a line and line is not valid"
		return
	}

	var temp_rect models.Rectangle

	for i := 0; i < len(inputs); i++ {

		temp_rect = inputs[i]
		if validateRectangle(temp_rect) {
			if checkRectanglesOverlap(main_rect, temp_rect) {
				// add time to temp_rect then add it to db
				temp_rect.Time = s
				models.DB.Create(&temp_rect)
				// data_base = append(data_base, temp_rect)
			}

		}
	}

	c.IndentedJSON(http.StatusCreated, main_rect) //201 status code
}

func validateRectangle(rect models.Rectangle) bool {
	if (rect.W_rect == 0) || (rect.H_rect == 0) {
		return false
	} else {
		return true
	}
}

func checkRectanglesOverlap(rect1 models.Rectangle, rect2 models.Rectangle) bool {
	if (rect1.X_rect >= (rect2.X_rect + rect2.W_rect)) || (rect2.X_rect >= (rect1.X_rect + rect1.W_rect)) {
		return false
	} else if (rect1.Y_rect >= (rect2.Y_rect + rect2.H_rect)) || (rect2.Y_rect >= (rect1.Y_rect + rect1.H_rect)) {
		return false
	} else {
		return true
	}

}

func main() {
	// Connect Application To Postgres Database
	models.Connect()

	// Create Router
	router := gin.Default()
	router.GET("/", getData)
	router.POST("/", postData)

	router.Run(":3200")

}
