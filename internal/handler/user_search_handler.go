package handler

import (
	"firstGoProject/internal/domain/entity"
	"firstGoProject/internal/domain/enum"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

// SearchHandler gets url parameters and sends them to DTO
func (h *UserHandler) SearchHandler(c *gin.Context) {
	info := &entity.SearchUserDTO{
		Name:      strings.TrimSpace(c.Query("name")),
		Surname:   strings.TrimSpace(c.Query("surname")),
		Education: strings.TrimSpace(c.Query("education")),
		Age:       -1,
		Status:    enum.NotInitialized,
		Gender:    "",
	}

	// gets age from the url and if it's not null, turns it into an int
	if ageStr := c.Query("age"); ageStr != "" {
		age, err := strconv.Atoi(ageStr)
		if err == nil {
			info.Age = age
		} else {
			c.JSON(400, gin.H{"error": "invalid age"})
			return
		}
	}

	// gets status from url and if it's not null, turns it into an enum
	if statusStr := c.Query("status"); statusStr != "" {
		status, err := strconv.Atoi(statusStr)
		if err == nil {
			info.Status = enum.UserStatus(status)
		} else {
			c.JSON(400, gin.H{"error": "invalid status"})
			return
		}
	}

	// gets gender from url and if it's not null, turns it into an enum
	if genderStr := strings.ToLower(c.Query("gender")); genderStr != "" {
		switch genderStr {
		case "male", "female", "other":
			info.Gender = enum.UserGender(genderStr)
		default:
			c.JSON(400, gin.H{"error": "invalid gender"})
			return
		}
	}

	users, err := h.s.SearchUser(info)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	if users == nil || len(*users) == 0 {
		c.JSON(404, gin.H{"message": "no user found"})
		return
	}

	c.JSON(200, users)
}
