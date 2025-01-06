from customer_management import Customer
from book_management import Book

class Transaction(Customer):
    def __init__(self, name, email, phone, book_title, quantity_sold):
        super().__init__(name, email, phone)
        self.book_title = book_title
        self.quantity_sold = quantity_sold

    def display_transaction(self):
        return f"Customer: {self.name}, Book: {self.book_title}, Quantity Sold: {self.quantity_sold}"

# Function to sell a book
def sell_book(book_list, customer_list, book_title, quantity, customer_name):
    book = next((b for b in book_list if b.title.lower() == book_title.lower()), None)
    customer = next((c for c in customer_list if c.name.lower() == customer_name.lower()), None)
    
    if not book:
        print(f"Error: Book '{book_title}' not found.")
        return

    if not customer:
        print(f"Error: Customer '{customer_name}' not found.")
        return

    if book.quantity < quantity:
        print(f"Error: Only {book.quantity} copies available. Sale cannot be completed.")
        return

    book.quantity -= quantity
    transaction = Transaction(customer.name, customer.email, customer.phone, book.title, quantity)
    print(f"Sale successful! Remaining quantity: {book.quantity}")
    print(transaction.display_transaction())

# Function to view all sales records
def view_sales(sales_list):
    if not sales_list:
        print("No sales records available.")
    for sale in sales_list:
        print(sale.display_transaction())
