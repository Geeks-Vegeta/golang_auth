package controllers

import (
	"app/models"
	"app/utils"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection = models.GetCollection(models.DB, "user")
var validate = validator.New()

func LoginUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var user models.User
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		var candidatePassword = user.Password

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&user); validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": validationErr.Error(),
			})
			return
		}

		err := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&user)
		if err != nil {
			c.JSON(http.StatusConflict, gin.H{
				"message": "user with that email does not exists",
			})
			return
		}

		verify := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(candidatePassword))
		if verify != nil {
			c.JSON(http.StatusConflict, gin.H{
				"message": "Incorrect Password",
			})
			return
		}

		var userid = user.Id.Hex()
		token, _ := utils.GenerateToken(userid)

		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})

	}
}
