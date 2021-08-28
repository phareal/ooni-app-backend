package controllers

import (
   "fmt"
   "github.com/gofiber/fiber/v2"
   "ooni/backend/core/database"
   "ooni/backend/models"
)

type AuthController interface {
   Register(c *fiber.Ctx)
   Login (c *fiber.Ctx) error
}


type authController struct {
    // nous servira dans le cas de l'injection de dependance  
}

func (a authController) Register(c *fiber.Ctx) {
   panic("implement me")
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
   var user models.User;
   db.Find(&user,"email = ? " ,_user.Email)
   return user
}

func AuthenticationController() *authController {
   return &authController{}
}


