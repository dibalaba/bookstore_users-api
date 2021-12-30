package users

import (
	//"encoding/json"
	//"fmt"
	//"io/ioutil"
	//"fmt"
	"net/http"

	"github.com/dibalaba/bookstore_users-api/domain/users"
	"github.com/dibalaba/bookstore_users-api/services"
	"github.com/dibalaba/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user users.User

	/* fmt.Println(user)

	bytes, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		//TODO: handle error
		return
	}
	if err := json.Unmarshal(bytes, &user); err != nil {
		fmt.Println(err.Error())
		//TODO: handle json error
		return
	} */

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		//TODO: return bad request to the caller
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		// TODO: handle user creation error
		return
	}

	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}
