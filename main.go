package main

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	/*err := Config.Connect()

	config := cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,DELETE",
		AllowHeaders:     "Origin,Content-Type,Accept",
		AllowCredentials: true,
	}

	app.Use(cors.New(config))
	*/
	app.Listen(":8000")

}

func getModuleRam() {
	cmd := exec.Command("sh", "-c", "cat /proc/ram_202004734")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	out := string(output[:])
	ram_Data := strings.Split(out, ",")

	ram_total, err1 := strconv.Atoi(ram_Data[0])
	ram_libre, err2 := strconv.Atoi(ram_Data[1])

	if err1 == nil && err2 == nil {
		porcentaje := float64(ram_total-ram_libre) / float64(ram_total) * 100
		fmt.Println("{")
		fmt.Println("Total RAM:", ram_total)
		fmt.Println("Total RAM in Use:", ram_total-ram_libre)
		fmt.Println("Total Ram Free:", ram_libre)
		fmt.Println("Porcentaje RAM in Use:", porcentaje)
		fmt.Println("}")
	}
}
