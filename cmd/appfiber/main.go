package main

import (
	"github.com/gofiber/fiber/v2"
	"fmt"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	
	swagger  "github.com/arsmn/fiber-swagger/v2"
	_ "golearn/cmd/appfiber/docs"
)

type Person struct{
	Name string `json:"name" xml:"name" form:"name"`
}

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main(){

	// initial app
	//app := fiber.New()

    // initial app with config (example config)
	app := fiber.New(fiber.Config{
		Prefork: true,
		CaseSensitive: true,		// path casesensitive
		StrictRouting: true,		// strip / end path  > true mean /aa and /aa/ ard diferenct
		ServerHeader: "Fiber",
		AppName: "Demo",
	})



	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world")
	})

	// path /xxx
	// app.Get("/:value", func(c *fiber.Ctx) error {
	// 	return c.SendString("Value : " + c.Params("value"))
	// })

	// path /xxx in name variable paramm
	// app.Get("/:name?", func(c *fiber.Ctx) error{
	// 	return c.SendString("name : " + c.Params("name"))
	// })

	// path /xxx/yyyy in param 	
	// app.Get("/:name?/:name2?", func(c *fiber.Ctx) error{
	// 	n := c.Params("name")
	// 	n2 := c.Params("name2")
	// 	return c.SendString("name  : " + n +  " and " + n2)
	// })

	// path wildcard
	app.Get("/api/*", func(c *fiber.Ctx) error{
		return c.SendString("path : " + c.Params("*"))
	})

	app.Get("/ex1", customErr)
	// route handler all method suchas app.Post, app.Get, app.Put, app.Delete
	//
	// another way to custom handler

	// match any request
	// app.Use(func (c *fiber.Ctx) error {
	// 	fmt.Println("use")
	// 	return c.SendString("Use")
	// })

	// match request prefix /xxx
	app.Use("/xxx", func (c *fiber.Ctx) error {
		fmt.Println("use xxx")
		return c.Next()			// chain request function Next() for pass next handler that same route
	})
	app.Use("/xxx", func (c *fiber.Ctx) error {
		fmt.Println("use yyy")
		return c.SendString("Use yyy")
	})


	// Attach multiple handlers 
	// app.Use("/aaa",func(c *fiber.Ctx) error {
	// 	c.Set("X-Custom-Header",  random.String(32))
	//   		return c.Next()
	// }, func(c *fiber.Ctx) error {
	//   return c.Next()
	// })


	app.Static("/js","./static")		// serve static folder
	app.Static("/js2", "./static/t.js") // serve static file

	// app.Static("/js3", "./static", fiber.Static{  // serve static with config
	// 	CacheDuration: 10 * time.Second,
	// })



	// group api
    demo := app.Group("/demo", handler01)   // path /demo
	v1 := demo.Group("/v1", handler01)		// path /demo/v1
	v1.Get("/test",handler01)				// path /demo/v1/test
	
 
	app.Get("/aa",handlerContext)

    /// monitor
	app.Get("/dashboard", monitor.New())


	//swagger
	app.Get("/swagger/*", swagger.Handler) // default
	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:         "/swagger/doc.json",
		DeepLinking: false,
	}))

	//listen port
	app.Listen(":3000")

}

// func handler(c *fiber.Ctx) error {
// 	p1 := c.Params("foo")
// 	_ = p1
// 	return nil
// }

// custom error
func customErr(c *fiber.Ctx) error {
	return fiber.NewError(782, "Custom err")  // status code and error message
}


// demo handler
func handler01(c *fiber.Ctx) error {
	return c.SendString("handler01")
}

// fiber => app.Test() for test httptest



// ShowAccount godoc
// @Summary Show a account
// @Description get string by ID
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param id path int true "Account ID"
// @Success 200 {string} string "ok"
// @Router /accounts/{id} [get]
func handlerContext(c *fiber.Ctx) error{

	fmt.Println(c.BaseURL())
	fmt.Println(c.Hostname())
	//c.IP()
	//c.IPs() -> IP from x-forward
	// c.Is("json")  -> for check content type
	// c.JSON(data)   -> convert to json
	// c.JSON(fiber.Map{
	//	"key": "value"
	// })
	// example /v1/*/shop/*
	// call /v1/xx/shop/yy
	// use c.Params("*2")  =>  yy

	// use c.Query("key") => from url such as /xxx?id=name
	// use c.QueryParser  -> get and convert
	// use c.Redirect("path")  => redirect path in app or url external

	// use c.Send()   => send response 
	// use c.SendFile(path)  => send file

	//use c.sendStatus(int)  -> for status
	// can set sendString and return sendStatus
   
	// c.Set(key,value) => set response header
    // c.Status(number)


	// c.Locals  => scope of request
	 

	// c.Body()  // request body for post
	fmt.Println(c.Locals("Host"))
	return c.SendString("OK")
}

func handlerPost(c *fiber.Ctx) error{

	p := new(Person)
	if err := c.BodyParser(p); err!=nil {		// pase request to struct
		return err
	}
	return c.SendString("OK")
}

// use c.Download for transfer files
// use c.FormFile() for get file from field 
// use c.SaveFile() for save file to server
// use c.FormValue("name")  from form
// use c.Get("key")  from req header
// use c.GetRespHeader("key") get res header


// ***
// use app.Use for setup header like a middleware before any request
// such as
// app.Use(func(c *fiber.Ctx) error {
	//// Set some security headers:
	//c.Set("X-XSS-Protection", "1; mode=block")

    //// Go to next middleware:
    // return c.Next()
   //})