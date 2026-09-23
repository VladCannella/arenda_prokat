package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"rental/internal/app"
	"rental/internal/domain"
	"rental/internal/infra"
	"strconv"
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
		if fields[0] == "exit" {
			return
		}
		func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("internal error", r)
				}
			}()
			switch fields[0] {
			case "rent":
				if len(fields) != 4 {
					fmt.Println("usage: rent <itemID> <customerID> <days>")
					return
				}
				days, err := strconv.Atoi(fields[3])
				if err != nil {
					fmt.Println("invalid days: ", err)
					return
				}
				itemID := domain.ID(fields[1])
				customerID := domain.ID(fields[2])
				start := time.Now()
				end := start.AddDate(0, 0, days)

				rental, err := service.RentItem(itemID, customerID, start, end)
				if err != nil {
					printError(err)
					return
				}
				fmt.Println("rented, rental id: ", rental.ID)
			case "return":
				if len(fields) != 2 {
					fmt.Println("usage: rent <rentalID>")
					return
				}
				rentalID := domain.ID(fields[1])
				fine, err := service.ReturnItem(rentalID, time.Now())
				if err != nil {
					printError(err)
					return
				}

				fmt.Println("returned, rental id: ", rentalID, "fine: ", fine)
			case "list":
				items, err := itemRepo.List()
				if err != nil {
					printError(err)
					return
				}

				for _, it := range items {
					fmt.Println(it.ID, it.Name, "-", itemStatusLabel(it.Status))
				}
				rentals, err := rentalRepo.List()
				if err != nil {
					printError(err)
					return
				}

				for _, r := range rentals {
					fmt.Println(r.ID, "item: ", r.ItemID, "customer: ", r.CustomerID)
				}
			default:
				fmt.Println("unknown command: ", fields[0])
			}
		}()

	}
}

func itemStatusLabel(status domain.ItemStatus) string {
	switch status {
	case domain.ItemAvailable:
		return "available"
	case domain.ItemRented:
		return "rented"
	default:
		panic(fmt.Sprintf("invariant violated: unknown item stats %q", status))
	}
}

func printError(err error) {
	switch {
	case errors.Is(err, domain.ErrItemAlreadyRented):
		fmt.Println("item already rented")
	case errors.Is(err, domain.ErrEntityNotFound):
		fmt.Println("not found - check ID")
	case errors.Is(err, domain.ErrRentalClosed):
		fmt.Println("rental already closed")
	default:
		var ve domain.ValidationError
		if errors.As(err, &ve) {
			fmt.Println("validation error:", ve.Error())
		} else {
			fmt.Println("error: ", err)
		}
	}
}
