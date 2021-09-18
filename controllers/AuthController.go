package controllers

import (
   "fmt"
   "github.com/gofiber/fiber/v2"
   "github.com/google/uuid"
   "net/http"
   "ooni/backend/core/database"
   "ooni/backend/models"
)

type AuthController interface {
   Register(c *fiber.Ctx)
   Login (c *fiber.Ctx) error
}

type Response struct {
   status int
   data fiber.Map
}

type authController struct {
    // nous servira dans le cas de l'injection de dependance  
}

func (a authController) Register(c *fiber.Ctx) error {
   // get the data that will be send
   // check if the user exists
   // if  not create the users
   // if users exists throw an error and connect him after some seconds after

   db := database.InitDatabase()
   response := new(Response)

   newUser := new(models.User)
   jsonUser := c.BodyParser(newUser)
   fmt.Print(newUser.Password)
   if jsonUser != nil {
      return c.Status(503).SendString(jsonUser.Error())
   }
   // check if existing user
   isExistingUser := findExistingUserEmail(newUser)

   if isExistingUser.Email == "" {
      id := uuid.New()
        // create the new user
      user := models.User{ID: id.String() ,Email: newUser.Email,Password: newUser.Password}
      // save the user to database
      creationResult := db.Create(user)

      if (creationResult.Error != nil) && (user.ID != "") {
          response.status = http.StatusOK
          response.data = fiber.Map{
             "userId": user.ID,
             "message": "user successfully registered",
          }
      } else {
         // error on inserting
         response.status = http.StatusInternalServerError
         response.data = fiber.Map{
            "userId": "",
            "message": "Internal server error",
         }
      }
   }else {
      // user existing
      response.status = http.StatusOK
      response.data = fiber.Map{
         "userId": isExistingUser.ID,
         "message": "user already exists , you'll be redirect to the account",
      }
   }

   return c.Status(response.status).JSON(response.data)

}

func (a authController) Login(c *fiber.Ctx) error {

   user := new(models.User)
   jsonUser := c.BodyParser(user)
   if jsonUser != nil {
      return c.Status(503).SendString(jsonUser.Error())
   }
   existingUser := findExistingUserEmail(user)
   fmt.Print(existingUser.Email)
   return nil
}

//findExistingUser query the database to find if the new user exist
func findExistingUserEmail(_user *models.User) models.User {
   db := database.InitDatabase()
   var user models.User
   db.Find(&user,"email = ? " ,_user.Email)
   return user
}

func AuthenticationController() *authController {
   return &authController{}
}


