package routes

import (
  "net/http"
  "strconv"
  "github.com/gin-gonic/gin"
  "example.com/rest-api/models"
)

func createEvent(context *gin.Context){
  var event models.Event
  err := context.ShouldBindJSON(&event)

  if err != nil {
    context.JSON(http.StatusBadRequest, gin.H{"message":"Could not parse request data"})
    return
  }

  event.ID = 1
  event.UserID = 1

  err = event.Save()

  if err != nil {
    context.JSON(http.StatusInternalServerError, gin.H{"message":"Could not create events. Try again later."})
    return
  }

  context.JSON(http.StatusCreated, gin.H{"message":"Event created", "event":event})
}


func getEvents(context *gin.Context){
  events, err := models.GetAllEvents()

  if err != nil {
    context.JSON(http.StatusInternalServerError, gin.H{"message":"Could not fetch events. Try again later."})
    return
  }

  context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context) {
  eventId, err := strconv.ParseInt(context.Param("id"), 10, 64) // base 10, 64 bit

  if err != nil {
    context.JSON(http.StatusBadRequest, gin.H{"message":"Could not parse event id."})
    return
  }

  event, err := models.GetEventByID(eventId)

  if err != nil {
    context.JSON(http.StatusInternalServerError, gin.H{"message":"Could not fetch event."})
    return
  }

  context.JSON(http.StatusOK, event)

}

func updateEvent(context *gin.Context){
  eventId, err := strconv.ParseInt(context.Param("id"), 10, 64) // base 10, 64 bit

  if err != nil {
    context.JSON(http.StatusBadRequest, gin.H{"message":"Could not parse event id."})
    return
  }

   _, err = models.GetEventByID(eventId)

  if err != nil {
    context.JSON(http.StatusInternalServerError, gin.H{"message":"Could not fetch event."})
    return
  }

  var updatedEvent models.Event

  err = context.ShouldBindJSON(&updatedEvent)

  if err != nil {
    context.JSON(http.StatusBadRequest, gin.H{"message":"Could not parse request data"})
    return
  }

  updatedEvent.ID = eventId

  err = updatedEvent.Update()

  if err != nil {
    context.JSON(http.StatusInternalServerError, gin.H{"message":"Could not update event."})
    return
  }

  context.JSON(http.StatusOK, gin.H{"message":"Event updated successfully"})
}

func deleteEvent(context *gin.Context){
  eventId, err := strconv.ParseInt(context.Param("id"), 10, 64) // base 10, 64 bit

  if err != nil {
    context.JSON(http.StatusBadRequest, gin.H{"message":"Could not parse event id."})
    return
  }

  event, err := models.GetEventByID(eventId)

  if err != nil {
    context.JSON(http.StatusInternalServerError, gin.H{"message":"Could not fetch event."})
    return
  }

  err = event.Delete()

  if err != nil {
    context.JSON(http.StatusInternalServerError, gin.H{"message":"Could not delete event."})
    return
  }

  context.JSON(http.StatusOK, gin.H{"message":"Event deleted successfully"})


 }









