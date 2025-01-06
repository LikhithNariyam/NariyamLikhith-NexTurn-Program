from flask import Flask, request, jsonify
import sqlite3
from sqlite3 import Error

app = Flask(__name__)

DATABASE = 'bookbuddy.db'

# Utility function to connect to the SQLite database
def get_db():
    conn = sqlite3.connect(DATABASE)
    return conn

# Helper function to handle errors
def handle_db_error(error):
    return jsonify({"error": "Database error", "message": str(error)}), 500

# POST /books: Add a new book
@app.route('/books', methods=['POST'])
def add_book():
    try:
        data = request.get_json()

        title = data.get('title')
        author = data.get('author')
        published_year = data.get('published_year')
        genre = data.get('genre')

        if not title or not author or not published_year or not genre:
            return jsonify({"error": "Invalid data", "message": "All fields are required"}), 400

        # Insert book into the database
        conn = get_db()
        c = conn.cursor()
        c.execute('''
            INSERT INTO books (title, author, published_year, genre)
            VALUES (?, ?, ?, ?)
        ''', (title, author, published_year, genre))
        conn.commit()

        book_id = c.lastrowid
        conn.close()

        return jsonify({"message": "Book added successfully", "book_id": book_id}), 201

    except Error as e:
        return handle_db_error(e)

# GET /books: Retrieve all books
@app.route('/books', methods=['GET'])
def get_books():
    try:
        conn = get_db()
        c = conn.cursor()
        c.execute('SELECT * FROM books')
        books = c.fetchall()
        conn.close()

        books_list = [{"id": book[0], "title": book[1], "author": book[2], "published_year": book[3], "genre": book[4]} for book in books]
        return jsonify(books_list), 200

    except Error as e:
        return handle_db_error(e)

# GET /books/<id>: Retrieve a specific book by id
@app.route('/books/<int:id>', methods=['GET'])
def get_book(id):
    try:
        conn = get_db()
        c = conn.cursor()
        c.execute('SELECT * FROM books WHERE id = ?', (id,))
        book = c.fetchone()
        conn.close()

        if not book:
            return jsonify({"error": "Book not found", "message": "No book exists with the provided ID"}), 404

        book_data = {"id": book[0], "title": book[1], "author": book[2], "published_year": book[3], "genre": book[4]}
        return jsonify(book_data), 200

    except Error as e:
        return handle_db_error(e)

# PUT /books/<id>: Update an existing book by id
@app.route('/books/<int:id>', methods=['PUT'])
def update_book(id):
    try:
        data = request.get_json()

        title = data.get('title')
        author = data.get('author')
        published_year = data.get('published_year')
        genre = data.get('genre')

        if not title or not author or not published_year or not genre:
            return jsonify({"error": "Invalid data", "message": "All fields are required"}), 400

        conn = get_db()
        c = conn.cursor()
        c.execute('''
            UPDATE books SET title = ?, author = ?, published_year = ?, genre = ?
            WHERE id = ?
        ''', (title, author, published_year, genre, id))
        conn.commit()

        if c.rowcount == 0:
            conn.close()
            return jsonify({"error": "Book not found", "message": "No book exists with the provided ID"}), 404

        conn.close()
        return jsonify({"message": "Book updated successfully"}), 200

    except Error as e:
        return handle_db_error(e)

# DELETE /books/<id>: Delete a book by id
@app.route('/books/<int:id>', methods=['DELETE'])
def delete_book(id):
    try:
        conn = get_db()
        c = conn.cursor()
        c.execute('DELETE FROM books WHERE id = ?', (id,))
        conn.commit()

        if c.rowcount == 0:
            conn.close()
            return jsonify({"error": "Book not found", "message": "No book exists with the provided ID"}), 404

        conn.close()
        return jsonify({"message": "Book deleted successfully"}), 200

    except Error as e:
        return handle_db_error(e)

if __name__ == '__main__':
    app.run(debug=True)
