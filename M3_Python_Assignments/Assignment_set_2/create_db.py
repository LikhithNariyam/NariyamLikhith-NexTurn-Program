import sqlite3

def create_database():
    # Connect to SQLite database (it will be created if it doesn't exist)
    conn = sqlite3.connect('bookbuddy.db')
    c = conn.cursor()

    # Create the books table
    c.execute('''
        CREATE TABLE IF NOT EXISTS books (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            title TEXT NOT NULL,
            author TEXT NOT NULL,
            published_year INTEGER NOT NULL,
            genre TEXT NOT NULL
        )
    ''')

    # Add some sample data for testing
    sample_books = [
        ('The Great Gatsby', 'F. Scott Fitzgerald', 1925, 'Fiction'),
        ('To Kill a Mockingbird', 'Harper Lee', 1960, 'Fiction'),
        ('1984', 'George Orwell', 1949, 'Dystopian'),
        ('The Catcher in the Rye', 'J.D. Salinger', 1951, 'Fiction')
    ]
    
    c.executemany('''
        INSERT INTO books (title, author, published_year, genre) 
        VALUES (?, ?, ?, ?)
    ''', sample_books)

    # Commit changes and close the connection
    conn.commit()
    conn.close()

# Run the function to create the database
create_database()
