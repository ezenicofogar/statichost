package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
)

var version_major uint8 = 0
var version_minor uint8 = 1

var multithreaded bool
var connection_host string
var connection_port uint16
var folder string

var statichost = &cobra.Command{
	Use:   "statichost",
	Short: "Serve files from a directory over localhost.",
	Long:  "Statichost serve files from a directory. It uses Go Fiber's Static functionality.",
	Run: func(cmd *cobra.Command, args []string) {
		app := fiber.New(fiber.Config{
			Prefork:               multithreaded,
			CaseSensitive:         true,
			DisableStartupMessage: true,
		})
		app.Static("/", folder, fiber.Static{
			CacheDuration: -1,
		})
		fmt.Printf("Statichost v%d.%d\n", version_major, version_minor)
		fmt.Printf("Serving %s\n", folder)
		fmt.Printf("at %s:%d\n", connection_host, connection_port)
		app.Listen(fmt.Sprintf("%s:%d", connection_host, connection_port))
	},
}

func main() {
	statichost.PersistentFlags().BoolVarP(&multithreaded,
		"multithreaded",
		"M",
		false,
		"Uses Fiber's \"Prefork\" option.")
	statichost.PersistentFlags().Uint16VarP(&connection_port,
		"port",
		"P",
		5500,
		"The port in which the server will run.")
	statichost.PersistentFlags().StringVar(&connection_host,
		"host",
		"localhost",
		"The host in which the server will run.")
	statichost.PersistentFlags().StringVarP(&folder,
		"location",
		"L",
		"./",
		"The location that the server will serve.")
	if err := statichost.Execute(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
