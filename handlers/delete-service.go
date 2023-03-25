package handlers

import (
	"net/http"

	"github.com/cassiusbessa/delete-service/logs"
	"github.com/cassiusbessa/delete-service/repositories"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func DeleteService(c *gin.Context) {
	db, id := c.Param("company"), c.Param("id")
	logrus.Warnf("Deleting a Service on %v", db)
	if _, err := repositories.DeleteService(db, id); err != nil {
		logrus.Errorf("Error Deleting Service to MongoDB: %v %v", db, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer logs.Elapsed("Delete Service")()
	c.JSON(http.StatusOK, gin.H{"message": "Service Deleted successfully"})
}
