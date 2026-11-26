// Using the gin framework, we'll build a basic api for creating, updating, and retrieval
// In vscode, requires the gin framework to be imported within the terminal
package main

import (
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
)

type phones struct {
	ID              string `json:"id"`
	PhoneName       string `json:"name"`
	PhoneGeneration string `json:"generation"`
	ReleaseDate     string `json:"date"`
}

var phoneModels = []phones{
	{ID: "1", PhoneName: "Iphone 8", PhoneGeneration: "5th Gen", ReleaseDate: "2012"},
	{ID: "2", PhoneName: "Iphone 12", PhoneGeneration: "3rd Gen", ReleaseDate: "2015"},
	{ID: "3", PhoneName: "Iphone 16", PhoneGeneration: "2nd Gen", ReleaseDate: "2018"},
}

// implementing endpoints for customers
// We'll use handler functions to do this
// 1. create a function that will retrieve the phone models from the json at the request of the user.
func getPhoneModels(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, phoneModels)
}

//Command in cmd: curl http://localhost:8080/phoneModels

// 2. create a POST request function that adds a phone to the json directly from the body request
func addPhoneModel(c *gin.Context) {
	var newPhone phones

	// call BINDJSON so the newly received json is binded to your newphone variable. If an error occurs, return and quit the function
	if err := c.BindJSON(&newPhone); err != nil {
		return
	}

	// once valid, assign the newPhone into the slice
	if newPhone.ID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Message": "You cannot have an empty ID value."})
		return
	}
	phoneModels = append(phoneModels, newPhone)

	c.IndentedJSON(http.StatusOK, newPhone)
}

// Command: curl http://localhost:8080/phoneModels --include --header "Content-Type: applications/json" --request POST --data "{\"id\": \"4\", \"name\": \"Iphone 21\", \"generation\": \"9th\", \"date\": \"2020\"}"

// 3. Create a handler that allows you to retrieve phonemodels by ID
func getPhoneById(c *gin.Context) {
	id := c.Param("id") // retrieves the id path parameter from URL, allows us to compare id values

	for _, a := range phoneModels {
		if a.ID == id { // if item is found, return to user
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Phone model not found"}) // returns 404 if not found
}

// command: curl http://localhost:8080/phoneModels/x (X == ID value of item)

// 4. Create a handler that allows you to delete an item via id
func removePhoneById(c *gin.Context) {
	idParam := c.Param("id") // references the id within the json

	for i, phones := range phoneModels {
		if phones.ID == idParam {
			phoneModels = slices.Delete(phoneModels, i, i+1)
			c.JSON(http.StatusOK, gin.H{"Message": "Item Successfully Deleted"})
			return
		}
	}

	// if operation failed
	c.JSON(http.StatusNotFound, gin.H{"Message": "Could not find entry with this ID"})

}

// creating a router within go
func main() {
	router := gin.Default()
	router.GET("/phoneModels", getPhoneModels) // Before, you wrote the folder of Restful_Go_Api. Instead, provide the slice of phones
	router.POST("/phoneModels", addPhoneModel)
	router.GET("/phoneModels/:id", getPhoneById)
	router.DELETE("/phoneModels/:id", removePhoneById)
	router.Run("localhost:8080")
}

// TBA: Implement a delete handler function that deletes an entry by ID (Done - 11/25/26)
// TBA: Implement a function that will catch if an empty string is passed as the id for a new phone entry (Done - 11/25/2026)
