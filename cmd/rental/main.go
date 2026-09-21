package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"rental/internal/app"
	"rental/internal/domain"
	"rental/internal/infra"
	"strings"
	"time"
)

func main() {

	itemRepo := infra.NewInMemoryRepo[domain.Item]()
	customerRepo := infra.NewInMemoryRepo[domain.Customer]()
	rentalRepo := infra.NewInMemoryRepo[domain.Rental]()
	notifier := &infra.ConsoleNotifier{}
	pricing := domain.DailyPricing{}

	service := app.NewRentalService(itemRepo, customerRepo, rentalRepo, notifier, pricing)
	dailyRate, err := domain.NewMoney(10000, "RUB")
	if err != nil {
		log.Fatal(err)
	}
	item := domain.Item{
		BaseEntity: domain.BaseEntity{ID: "item-1", CreatedAt: time.Now()},
		Name:       "Axe",
		DailyRate:  dailyRate,
		Status:     domain.ItemAvailable,
	}
	if err := itemRepo.Save(item); err != nil {
		log.Fatal(err)
	}

	customer := domain.Customer{
		BaseEntity:  domain.BaseEntity{ID: "customer-1", CreatedAt: time.Now()},
		Name:        "John",
		RentalCount: 0,
	}
	if err := customerRepo.Save(customer); err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		fields := strings.Fields(line)
		switch fields[0] {
		case "exit":
			return
		case "rent":
			fmt.Println("rent: not yet implemented")
		case "return":
			fmt.Println("return: not yet implemented")
		case "list":
			fmt.Println("list: not yet implemented")
		default:
			fmt.Println("unknown command: ", fields[0])
		}

	}
}
