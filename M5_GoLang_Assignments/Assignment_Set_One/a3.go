package main

import (
	"errors"
	"fmt"
	"sort"
)

// Define the Product struct
type Product struct {
	ID    int
	Name  string
	Price float64
	Stock int
}

// Slice to store products
var inventory []Product

// Function to add a product to the inventory
func addProduct(id int, name string, price interface{}, stock int) error {
	// Type casting the price to float64
	priceFloat, ok := price.(float64)
	if !ok {
		return errors.New("price must be a valid float64")
	}

	// Create a new product
	product := Product{ID: id, Name: name, Price: priceFloat, Stock: stock}

	// Add the product to inventory
	inventory = append(inventory, product)
	return nil
}

// Function to update the stock of a product
func updateStock(id int, stock int) error {
	if stock < 0 {
		return errors.New("stock cannot be negative")
	}

	// Search for the product by ID
	for i := range inventory {
		if inventory[i].ID == id {
			// Update the stock
			inventory[i].Stock = stock
			return nil
		}
	}
	return errors.New("product not found")
}

// Function to search for a product by name or ID
func searchProduct(query string) (*Product, error) {
	// Search by ID
	var id int
	if _, err := fmt.Sscanf(query, "%d", &id); err == nil {
		for i := range inventory {
			if inventory[i].ID == id {
				return &inventory[i], nil
			}
		}
	}

	// Search by Name
	for i := range inventory {
		if inventory[i].Name == query {
			return &inventory[i], nil
		}
	}
	return nil, errors.New("product not found")
}

// Function to display the inventory
func displayInventory() {
	if len(inventory) == 0 {
		fmt.Println("No products available in the inventory.")
		return
	}

	// Display table header
	fmt.Printf("%-5s %-20s %-10s %-10s\n", "ID", "Name", "Price", "Stock")
	fmt.Println("--------------------------------------------")

	// Display each product in the inventory
	for _, product := range inventory {
		fmt.Printf("%-5d %-20s ₹%-9.2f %-10d\n", product.ID, product.Name, product.Price, product.Stock)
	}
}

// Function to sort inventory by price or stock
func sortInventory(criteria string) {
	switch criteria {
	case "price":
		sort.SliceStable(inventory, func(i, j int) bool {
			return inventory[i].Price < inventory[j].Price
		})
	case "stock":
		sort.SliceStable(inventory, func(i, j int) bool {
			return inventory[i].Stock < inventory[j].Stock
		})
	default:
		fmt.Println("Invalid sort criteria. Use 'price' or 'stock'.")
	}
}

// Main function to implement the menu system
func main() {
	// Sample products with an Indian context
	addProduct(1, "Samsung Galaxy S24", 55000.50, 20)
	addProduct(2, "Dell Inspiron Laptop", 60000.75, 10)
	addProduct(3, "OnePlus 11", 50000.30, 15)
	addProduct(4, "Mi Smart TV", 25000.40, 30)
	addProduct(5, "Redmi Note 12", 15000.80, 50)

	for {
		// Display menu options
		fmt.Println("\nIndian Store Inventory Management System")
		fmt.Println("1. Add Product")
		fmt.Println("2. Update Stock")
		fmt.Println("3. Search Product")
		fmt.Println("4. Display Inventory")
		fmt.Println("5. Sort Inventory")
		fmt.Println("6. Exit")
		fmt.Print("Enter your choice: ")

		// Get user input
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			// Add Product
			var id, stock int
			var name, priceStr string
			var price float64
			fmt.Print("Enter Product ID: ")
			fmt.Scan(&id)
			fmt.Print("Enter Product Name: ")
			fmt.Scan(&name)
			fmt.Print("Enter Product Price (in ₹): ")
			fmt.Scan(&priceStr)
			fmt.Print("Enter Product Stock: ")
			fmt.Scan(&stock)

			// Type casting price to float64
			_, err := fmt.Sscanf(priceStr, "%f", &price)
			if err != nil {
				fmt.Println("Invalid price format.")
				continue
			}

			err = addProduct(id, name, price, stock)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Product added successfully!")
			}

		case 2:
			// Update Stock
			var id, stock int
			fmt.Print("Enter Product ID to Update Stock: ")
			fmt.Scan(&id)
			fmt.Print("Enter New Stock: ")
			fmt.Scan(&stock)

			err := updateStock(id, stock)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Stock updated successfully!")
			}

		case 3:
			// Search Product
			var query string
			fmt.Print("Enter Product Name or ID to Search: ")
			fmt.Scan(&query)

			product, err := searchProduct(query)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("Found Product: ID=%d, Name=%s, Price=₹%.2f, Stock=%d\n", product.ID, product.Name, product.Price, product.Stock)
			}

		case 4:
			// Display Inventory
			displayInventory()

		case 5:
			// Sort Inventory
			var criteria string
			fmt.Print("Enter sort criteria (price/stock): ")
			fmt.Scan(&criteria)
			sortInventory(criteria)
			fmt.Println("Inventory sorted successfully!")

		case 6:
			// Exit the program
			fmt.Println("Exiting... Goodbye!")
			return

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
