class Book:
    def __init__(self, title, author, price, quantity):
        self.title = title
        self.author = author
        self.price = price
        self.quantity = quantity

    def display_book(self):
        return f"Title: {self.title}, Author: {self.author}, Price: {self.price}, Quantity: {self.quantity}"

# Function to add a book to the inventory
def add_book(book_list, title, author, price, quantity):
    new_book = Book(title, author, price, quantity)
    book_list.append(new_book)
    print("Book added successfully!")

# Function to view all books
def view_books(book_list):
    if not book_list:
        print("No books available.")
    for book in book_list:
        print(book.display_book())

# Function to search for a book by title or author
def search_books(book_list, search_term):
    found_books = [book for book in book_list if search_term.lower() in book.title.lower() or search_term.lower() in book.author.lower()]
    if found_books:
        for book in found_books:
            print(book.display_book())
    else:
        print("No books found.")