package main

import (
	"errors"
	"fmt"
)

// Define constants for menu options
const (
	DepositOption     = 1
	WithdrawOption    = 2
	ViewBalanceOption = 3
	ViewHistoryOption = 4
	ExitOption        = 5
)

// Define the Account struct
type Account struct {
	ID                 int
	Name               string
	Balance            float64
	TransactionHistory []string
}

// Slice to store accounts
var accounts []Account

// Function to find an account by ID
func findAccountByID(id int) (*Account, error) {
	for i := range accounts {
		if accounts[i].ID == id {
			return &accounts[i], nil
		}
	}
	return nil, errors.New("account not found")
}

// Function to deposit money
func deposit(id int, amount float64) error {
	// Validate deposit amount
	if amount <= 0 {
		return errors.New("deposit amount must be greater than zero")
	}

	// Find the account
	account, err := findAccountByID(id)
	if err != nil {
		return err
	}

	// Update balance and transaction history
	account.Balance += amount
	account.TransactionHistory = append(account.TransactionHistory, fmt.Sprintf("Deposited: %.2f", amount))
	return nil
}

// Function to withdraw money
func withdraw(id int, amount float64) error {
	// Validate withdraw amount
	if amount <= 0 {
		return errors.New("withdraw amount must be greater than zero")
	}

	// Find the account
	account, err := findAccountByID(id)
	if err != nil {
		return err
	}

	// Ensure sufficient balance
	if account.Balance < amount {
		return errors.New("insufficient balance")
	}

	// Update balance and transaction history
	account.Balance -= amount
	account.TransactionHistory = append(account.TransactionHistory, fmt.Sprintf("Withdrew: %.2f", amount))
	return nil
}

// Function to view balance
func viewBalance(id int) (float64, error) {
	// Find the account
	account, err := findAccountByID(id)
	if err != nil {
		return 0, err
	}

	// Return the balance
	return account.Balance, nil
}

// Function to view transaction history
func viewTransactionHistory(id int) ([]string, error) {
	// Find the account
	account, err := findAccountByID(id)
	if err != nil {
		return nil, err
	}

	// Return the transaction history
	return account.TransactionHistory, nil
}

// Main function to implement the menu system
func main() {
	// Add some sample accounts
	accounts = append(accounts, Account{ID: 1, Name: "Likhith", Balance: 1000, TransactionHistory: []string{}})
	accounts = append(accounts, Account{ID: 2, Name: "Pranav", Balance: 2000, TransactionHistory: []string{}})
	accounts = append(accounts, Account{ID: 3, Name: "Maurya", Balance: 500, TransactionHistory: []string{}})

	for {
		// Display menu options
		fmt.Println("\nBank Transaction System")
		fmt.Println("1. Deposit")
		fmt.Println("2. Withdraw")
		fmt.Println("3. View Balance")
		fmt.Println("4. View Transaction History")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")

		// Get user input
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case DepositOption:
			// Handle deposit
			var id int
			var amount float64
			fmt.Print("Enter Account ID: ")
			fmt.Scan(&id)
			fmt.Print("Enter Amount to Deposit: ")
			fmt.Scan(&amount)

			err := deposit(id, amount)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Deposit successful!")
			}

		case WithdrawOption:
			// Handle withdrawal
			var id int
			var amount float64
			fmt.Print("Enter Account ID: ")
			fmt.Scan(&id)
			fmt.Print("Enter Amount to Withdraw: ")
			fmt.Scan(&amount)

			err := withdraw(id, amount)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Withdrawal successful!")
			}

		case ViewBalanceOption:
			// Handle viewing balance
			var id int
			fmt.Print("Enter Account ID: ")
			fmt.Scan(&id)

			balance, err := viewBalance(id)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("Account Balance: %.2f\n", balance)
			}

		case ViewHistoryOption:
			// Handle viewing transaction history
			var id int
			fmt.Print("Enter Account ID: ")
			fmt.Scan(&id)

			history, err := viewTransactionHistory(id)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Transaction History:")
				for _, transaction := range history {
					fmt.Println(transaction)
				}
			}

		case ExitOption:
			// Exit the program
			fmt.Println("Exiting... Goodbye!")
			return

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
