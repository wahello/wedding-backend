package main

import (
	"bufio"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/ssh/terminal"

	"wedding/ent"
	"wedding/pkg/server"
	"wedding/pkg/util"
)

var (
	pgHost     string
	pgPort     int
	pgUsername string
	pgDatabase string
	pgPassword string
	pgSSLMode  string
)

var (
	authZone   string
	authSecret string
)

var (
	hydraAdminURL  string
	hydraPublicURL string
)

func main() {
	app := cli.App{
		Name:      "wedding",
		ArgsUsage: "[[hostname]:port]",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "pg-host",
				Destination: &pgHost,
				EnvVars:     []string{"PG_HOST"},
				Value:       "localhost",
			},
			&cli.IntFlag{
				Name:        "pg-port",
				Destination: &pgPort,
				EnvVars:     []string{"PG_PORT"},
				Value:       5432,
			},
			&cli.StringFlag{
				Name:        "pg-username",
				Destination: &pgUsername,
				EnvVars:     []string{"PG_USERNAME"},
				Value:       "wedding",
			},
			&cli.StringFlag{
				Name:        "pg-database",
				Destination: &pgDatabase,
				EnvVars:     []string{"PG_DATABASE"},
				Value:       "wedding",
			},
			&cli.StringFlag{
				Name:        "pg-password",
				Destination: &pgPassword,
				EnvVars:     []string{"PG_PASSWORD"},
				Value:       "",
			},
			&cli.StringFlag{
				Name:        "pg-sslmode",
				Destination: &pgSSLMode,
				EnvVars:     []string{"PG_SSL_MODE"},
				Value:       "disable",
			},
			&cli.BoolFlag{
				Name:  "generate",
				Value: false,
			},
			&cli.StringFlag{
				Name:        "auth-zone",
				Destination: &authZone,
				EnvVars:     []string{"AUTH_ZONE"},
				Value:       "default",
			},
			&cli.StringFlag{
				Name:        "auth-secret",
				Destination: &authSecret,
				EnvVars:     []string{"AUTH_SECRET"},
				Value:       "default",
			},
			&cli.StringFlag{
				Name:        "hydra-admin-url",
				Destination: &hydraAdminURL,
				EnvVars:     []string{"HYDRA_ADMIN_URL"},
				Value:       "https://admin.hydra.imle.io",
			},
			&cli.StringFlag{
				Name:        "hydra-public-url",
				Destination: &hydraPublicURL,
				EnvVars:     []string{"HYDRA_PUBLIC_URL"},
				Value:       "https://hydra.imle.io",
			},
		},
		Commands: []*cli.Command{
			{
				Name: "add-backroom-user",
				Action: func(ctx *cli.Context) error {
					reader := bufio.NewReader(os.Stdin)

					fmt.Print("Enter Username: ")
					username, err := reader.ReadString('\n')
					if err != nil {
						return err
					}

					fmt.Print("Enter Password: ")
					bytePassword, err := terminal.ReadPassword(syscall.Stdin)
					if err != nil {
						return err
					}
					fmt.Println()

					username = strings.TrimSpace(username)
					password := strings.TrimSpace(string(bytePassword))

					hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
					if err != nil {
						return err
					}

					// Connect to db.
					client, err := ent.Open("postgres", getPgConnectionString(), ent.Log(log.Println))
					if err != nil {
						return err
					}
					defer client.Close()

					save, err := client.BackroomUser.Create().
						SetUsername(username).
						SetPassword(string(hash)).
						Save(ctx.Context)
					if err != nil {
						return err
					}

					log.Printf("user created: (%s)\n", save.String())

					return nil
				},
			},
		},
		Action: func(ctx *cli.Context) error {
			// Connect to db.
			client, err := ent.Open("postgres", getPgConnectionString(), ent.Log(log.Println))
			if err != nil {
				return err
			}
			defer client.Close()

			// TODO: Remove auto migrate.
			if err := client.Schema.Create(ctx.Context); err != nil {
				log.Fatalf("failed creating schema resources: %v", err)
			}

			// Gen some fake data if asked for.
			if ctx.Bool("generate") {
				generateFakeData(ctx.Context, client)
			}

			// For non-release mode we want to see the queries.
			if gin.Mode() != gin.ReleaseMode {
				client = client.Debug()
			}

			// Default port and listen on all.
			serverAddress := ctx.Args().First()
			if serverAddress == "" {
				serverAddress = ":8080"
			}

			// Setup gin.
			var engine = gin.Default()

			// Setup sessions
			store, err := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
			if err != nil {
				return err
			}
			engine.Use(sessions.Sessions("wedding", store))

			// Setup CORS.
			engine.Use(cors.New(cors.Config{
				AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
				AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
				AllowCredentials: true,
				AllowOrigins:     []string{"*"},
				MaxAge:           12 * time.Hour,
			}))

			// Register data handlers.
			router := engine.Group("/api")
			auth := server.RegisterAuth(client, router, store)
			server.RegisterAdminAPIv1(client, router.Group("/admin/v1/", auth.Middleware()))
			server.RegisterAPIv1(client, router.Group("/v1/invitees"))

			// Create server.
			srv := &http.Server{
				Addr:    serverAddress,
				Handler: engine,
			}

			// Initializing the server in a goroutine to enable graceful handling of shutdown.
			go func() {
				if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					log.Fatalf("listen: %s\n", err)
				}
			}()

			quit := make(chan os.Signal)
			signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
			<-quit
			log.Println("Shutting down server...")

			shutdownCtx, cancel := context.WithTimeout(ctx.Context, 5*time.Second)
			defer cancel()
			if err := srv.Shutdown(shutdownCtx); err != nil {
				log.Fatal("Server forced to shutdown:", err)
			}

			log.Println("Server exiting")

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func getPgConnectionString() string {
	connData := map[string]string{
		"host":     pgHost,
		"port":     strconv.Itoa(pgPort),
		"user":     pgUsername,
		"dbname":   pgDatabase,
		"password": pgPassword,
		"sslmode":  pgSSLMode,
	}
	var connectionString string
	for key, value := range connData {
		if value == "" {
			continue
		}
		connectionString += key + "=" + value + " "
	}

	return connectionString
}

func generateFakeData(ctx context.Context, client *ent.Client) {
	imleParty := client.InviteeParty.Create().
		SetName("Steven Immediate Family").
		SetCode(util.RandomString(10)).
		SaveX(ctx)

	client.Invitee.Create().
		SetName("Susan Hixson").
		SetParty(imleParty).
		SaveX(ctx)
	client.Invitee.Create().
		SetName("Todd Hixson").
		SetParty(imleParty).
		SaveX(ctx)
	client.Invitee.Create().
		SetName("Ryan Hixson").
		SetParty(imleParty).
		SaveX(ctx)

	smithParty := client.InviteeParty.Create().
		SetName("Savannah Immediate Family").
		SetCode(util.RandomString(10)).
		SaveX(ctx)

	client.Invitee.Create().
		SetName("Harold Smith").
		SetParty(smithParty).
		SaveX(ctx)
	client.Invitee.Create().
		SetName("Kimberly Smith").
		SetParty(smithParty).
		SaveX(ctx)
	client.Invitee.Create().
		SetName("Joseph Smith").
		SetParty(smithParty).
		SaveX(ctx)
	client.Invitee.Create().
		SetName("Chandler Smith").
		SetParty(smithParty).
		SaveX(ctx)
}
