package web

import "github.com/gofiber/fiber/v2"

type Server struct {
	App *fiber.App
}

func CreateServer() *Server {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			// Status code defaults to 500
			code := fiber.StatusInternalServerError
			var msg string
			// Retrieve the custom status code if it's a fiber.*Error
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
				msg = e.Message
			}

			if msg == "" {
				msg = "cannot process the http call"
			}

			// Send custom error page
			err = ctx.Status(code).JSON(InternalError{
				Msg: msg,
			})
			return nil
		},
	})

	setUpRoutes(app)

	return &Server{
		App: app,
	}
}

func setUpRoutes(app *fiber.App) {
	setupUserRoutes(app)
}

func (srv *Server) Run(port string) error {
	return srv.App.Listen(":" + port)
}
