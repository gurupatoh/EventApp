package main

import (
	"EventApp/internal/database"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (app *application) createEvent(c *gin.Context) {
	var event database.Event
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return

	}


	err := app.models.Event.Insert(&event)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create event"})
	}	
	c.JSON(http.StatusCreated, event)
	
}

func (app *application) getAllEvents(c *gin.Context) {
	events, err := app.models.Event.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch events"})
		return
	}
	c.JSON(http.StatusOK, events)
}
func (app *application) getEventByID(c *gin.Context) {
	id ,err:= strconv.Atoi(c.Param("id")) 
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}
	event, err := app.models.Event.Get(id)

	if event == nil{
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch event"})
		return		
	}
	c.JSON(http.StatusOK, event)
}
func (app *application) updateEvent(c *gin.Context) {
	id ,err := strconv.Atoi(c.Param("id")) 
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	existingEvent, err := app.models.Event.Get(id)
 if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch event"})
		return		
	}	

	if existingEvent == nil{
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}	

	updateEvent := &database.Event{}
	if err := c.ShouldBindJSON(updateEvent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updateEvent.Id = int64(id)
if err := app.models.Event.Update(updateEvent); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update event"})
		return
	}	

}

func (app *application) deleteEvent(c *gin.Context) {
	id ,err := strconv.Atoi(c.Param("id")) 
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to delete event"})
	}
	 if err := app.models.Event.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete event"})
	 }
	 c.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})

}

