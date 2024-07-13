package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
    "github.com/gin-contrib/cors"
)

type Character struct {
	ID		 string    `json:"id"`
    Name       string `json:"name"`
    Class      string `json:"class"`
    Race       string `json:"race"`
    Strength   int    `json:"strength"`
    Dexterity  int    `json:"dexterity"`
    Constitution int  `json:"constitution"`
    Intelligence int  `json:"intelligence"`
    Wisdom      int   `json:"wisdom"`
    Charisma    int   `json:"charisma"`
}

// create dummy for database
var characters = []Character{
	{ID: "1", Name: "Aragorn", Class: "Ranger", Race: "Human", Strength: 15, Dexterity: 14, Constitution: 13, Intelligence: 12, Wisdom: 10, Charisma: 8},
	{ID: "2", Name: "Legolas", Class: "Archer", Race: "Elf", Strength: 14, Dexterity: 18, Constitution: 12, Intelligence: 11, Wisdom: 13, Charisma: 9},
	{ID: "3", Name: "Gimli", Class: "Warrior", Race: "Dwarf", Strength: 16, Dexterity: 12, Constitution: 15, Intelligence: 10, Wisdom: 9, Charisma: 7},
}

func init() {
		err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	fmt.Println("Server is successfully running on port:", os.Getenv("PORT"))
}

func getCharacters(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, characters)
}

func getCharactersBy(c *gin.Context) {
	// get id from url param
	id := c.Param("id")
	// grab the character by id
	char, err := getCharacterByIdHelper(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Character not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, char)
}

func getCharacterByIdHelper(id string) (*Character, error) {
	for i, char := range characters {
		if char.ID == id {
			return &characters[i], nil
		}
	}
	return nil, errors.New("Character not found")
}

func createCharacter(c *gin.Context) {
	var newCharacter Character
	// bind the request body to newCharacter
	if err := c.BindJSON(&newCharacter); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}

	// append newCharacter to characters
	characters = append(characters, newCharacter)
	c.IndentedJSON(http.StatusCreated, newCharacter)
}

func main() {
	// setup server
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/characters", getCharacters)
	router.GET("/characters/:id", getCharactersBy)
	router.POST("/characters", createCharacter)

	router.Run("localhost:8080")
}
