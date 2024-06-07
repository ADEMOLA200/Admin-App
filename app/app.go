package app

import (
    "context"
    "log"

    "github.com/joho/godotenv"

    "github.com/ADEMOLA200/Admin-App.git/database"
    "github.com/ADEMOLA200/Admin-App.git/logger"
    "github.com/ADEMOLA200/Admin-App.git/routes"
    "github.com/gofiber/fiber/v2"
    "go.uber.org/fx"
)

func init() {
    if err := godotenv.Load(); err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }
}

var Module = fx.Options(
    fx.Provide(
        NewApp,
        logger.NewLogger,
    ),
    fx.Invoke(
        RegisterHooks,
    ),
)

func NewApp() *fiber.App {
    database.Connect()
    return fiber.New()
}

func RegisterHooks(
    lifecycle fx.Lifecycle,
    app *fiber.App,
    logger *logger.Logger,
) {
    lifecycle.Append(
        fx.Hook{
            OnStart: func(ctx context.Context) error {
                go func() {
                    logger.Info("Starting Application")
                    logger.Info("================================")
                    logger.Info("====== ADMIN APPLICATION ======")
                    logger.Info("================================")

                    routes.Setup(app)
                    routes.UserSetup(app)
                    routes.RolesSetup(app)
                    routes.PermissionsSetup(app)
                    if err := app.Listen(":9000"); err != nil {
                        logger.Fatalf("Error starting server: %v", err)
                    }
                }()
                return nil
            },
            OnStop: func(ctx context.Context) error {
                logger.Info("Stopping Application")
                return app.Shutdown()
            },
        },
    )
}
