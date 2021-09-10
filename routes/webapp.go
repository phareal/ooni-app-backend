package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"net/http"
	"ooni/backend/core/database"
	"ooni/backend/models"
	"regexp"
)

// configure the session
var rxEmail = regexp.MustCompile(".+@.+\\..+")
var store = session.New()


func SetWebappRoutes(app *fiber.App) {

	app.Get("/", func(ctx *fiber.Ctx) error {
		sess, err := store.Get(ctx)
		if err != nil {
			panic(err)
		}
		// get the user id
		userId := sess.Get("userId")
		if userId == nil {
			err := ctx.Redirect("/login", http.StatusSeeOther)
			if err != nil {
				panic(err)
			}
		}
		return ctx.Render("index", fiber.Map{},"templates/base")
	})
	// TODO create middleware for handling session
	app.Get("/login", func(ctx *fiber.Ctx) error {
		sess, err := store.Get(ctx)
		if err != nil {
			panic(err)
		}
		userId := sess.Get("userId")
		if userId != nil {
			ctx.Status(http.StatusOK).Redirect("/")
		}else {
			return ctx.Render("login", fiber.Map{})
		}
	    return nil
	})

	app.Get("/users", func(ctx *fiber.Ctx) error {
		// get all the users

		var db = database.InitDatabase()
		var users [] models.User

		results := db.Find(&users);
		if results.Error != nil {
			panic(results.Error)
		}
	   return  ctx.Render("users",fiber.Map{
	   	 "users":users,
	   },"templates/base")

	})






	app.Post("/login", func(ctx *fiber.Ctx) error {
		sess, err := store.Get(ctx)
		if err != nil {
			panic(err)
		}

		errorsMessage := make(map[string]string)
		if ctx.FormValue("email") == "" || ctx.FormValue("password") == ""{
			errorsMessage["empty_fields"] = "Please fill all the mandatory field"
		} else {
			match := rxEmail.Match([]byte(ctx.FormValue("email")))
			if match == false {
				errorsMessage["email"] = "Please enter a valid email address"
			}
		}

		if len(errorsMessage) != 0{
			return ctx.Render("login", fiber.Map{
				"errors":errorsMessage,
			})
		}else {
			 // do the research here to find and admin that have the credentials as other
			var admin models.Admin
			//var db = database.InitDatabase()
			if ctx.FormValue("email") == "admin@ooni.app" && ctx.FormValue("password") == "admin" {
				sess.Set("userId",admin.ID)
				err := sess.Save()
				if err != nil {
					panic(err)
				}else {
					ctx.Status(http.StatusOK).Redirect("/")
				}
			}else {
				errorsMessage["noUser"] = "No user found with these credentials"
				return ctx.Render("login", fiber.Map{
					"errors":errorsMessage,
				})
			}
			//db.Find(&admin,"username = ? and password = ?" ,ctx.FormValue("email"), ctx.FormValue("password"))

			/*if admin.Username == "" {
				errorsMessage["noUser"] = "No user found with these credentials"
				return ctx.Render("login", fiber.Map{
					"errors":errorsMessage,
				})
			} else {
				 // set the session here

			}*/
		}
		return nil
	})

}
