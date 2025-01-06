import os
from book_management import add_book, view_books, search_books
from customer_management import add_customer, view_customers
from sales_management import sell_book, view_sales

def main():
    book_list = []
    customer_list = []
    sales_list = []
    
    while True:
        print("\nWelcome to BookMart!")
        print("1. Book Management")
        print("2. Customer Management")
        print("3. Sales Management")
        print("4. Exit")
        
        choice = input("Enter your choice: ")

        if choice == '1':
            print("\n1. Add Book")
            print("2. View Books")
            print("3. Search Books")
            book_choice = input("Enter your choice: ")
            
            if book_choice == '1':
                title = input("Enter book title: ")
                author = input("Enter book author: ")
                try:
                    price = float(input("Enter book price: "))
                    quantity = int(input("Enter book quantity: "))
                    if price <= 0 or quantity <= 0:
                        raise ValueError("Price and quantity must be positive.")
                except ValueError as e:
                    print(f"Invalid input: {e}")
                    continue
                add_book(book_list, title, author, price, quantity)

            elif book_choice == '2':
                view_books(book_list)

            elif book_choice == '3':
                search_term = input("Enter title or author to search: ")
                search_books(book_list, search_term)

        elif choice == '2':
            print("\n1. Add Customer")
            print("2. View Customers")
            customer_choice = input("Enter your choice: ")

            if customer_choice == '1':
                name = input("Enter customer name: ")
                email = input("Enter customer email: ")
                phone = input("Enter customer phone: ")
                add_customer(customer_list, name, email, phone)

            elif customer_choice == '2':
                view_customers(customer_list)

        elif choice == '3':
            print("\n1. Sell Book")
            print("2. View Sales Records")
            sales_choice = input("Enter your choice: ")

            if sales_choice == '1':
                customer_name = input("Enter customer name: ")
                book_title = input("Enter book title: ")
                try:
                    quantity = int(input("Enter quantity: "))
                    if quantity <= 0:
                        raise ValueError("Quantity must be positive.")
                except ValueError as e:
                    print(f"Invalid input: {e}")
                    continue
                sell_book(book_list, customer_list, book_title, quantity, customer_name)

            elif sales_choice == '2':
                view_sales(sales_list)

        elif choice == '4':
            print("Exiting BookMart...")
            break

if __name__ == "__main__":
    main()
