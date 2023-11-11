package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type fact struct {
	ID   string `json:"id"`
	Fact string `json:"fact"`
}

var facts = []fact{
	{ID: "1", Fact: "A group of flamingos is called a 'flamboyance."},
	{ID: "2", Fact: "The heart of a shrimp is located in its head"},
	{ID: "3", Fact: "The peregrine falcon is the fastest animal in the world, reaching speeds of up to 240 mph when diving for prey"},
	{ID: "4", Fact: "The mantis shrimp has one of the most complex visual systems in the animal kingdom"},
	{ID: "5", Fact: "A giraffe's tongue can be up to 21 inches long."},
	{ID: "6", Fact: "The tongue of a blue whale can weigh as much as an elephant."},
	{ID: "7", Fact: "A platypus is one of the few mammals that lay eggs."},
	{ID: "8", Fact: "The longest recorded lifespan of a goldfish is 43 years."},
	{ID: "9", Fact: "A newborn kangaroo is about 1 inch long."},
	{ID: "10", Fact: "Polar bears are excellent swimmers and can swim for long distances in the Arctic Ocean."},
	{ID: "11", Fact: "The mimic octopus can imitate the appearance and behavior of other marine animals."},
	{ID: "12", Fact: "A group of crows is called a 'murder."},
	{ID: "13", Fact: "The electric eel can generate electric shocks of up to 600 volts to stun prey."},
	{ID: "14", Fact: "Pigs are highly intelligent animals and can even learn to play video games."},
	{ID: "15", Fact: "Sloths move so slowly that algae can grow on their fur."},
	{ID: "16", Fact: "A group of owls is called a 'parliament."},
	{ID: "17", Fact: "Honey never spoils. Archaeologists have found pots of honey in ancient Egyptian tombs that are over 3,000 years old and still perfectly edible."},
	{ID: "18", Fact: "The only mammal capable of flight is the bat."},
	{ID: "19", Fact: "Crocodiles can go through extended periods of fasting and can survive without food for months."},
	{ID: "20", Fact: "A chameleon's tongue can be twice the length of its body."},
}

func getFacts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, facts)
}

func postFact(c *gin.Context) {
	var newFact fact

	if err := c.BindJSON(&newFact); err != nil {
		return
	}

	facts = append(facts, newFact)
	c.IndentedJSON(http.StatusCreated, newFact)
}

func getFactById(c *gin.Context) {
	id := c.Param("id")

	for _, a := range facts {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "fact not found"})
}

func main() {
	router := gin.Default()

	router.GET("/facts", getFacts)
	router.GET("/facts/:id", getFactById)
	router.POST("/facts", postFact)

	router.Run("localhost:8080")
}
