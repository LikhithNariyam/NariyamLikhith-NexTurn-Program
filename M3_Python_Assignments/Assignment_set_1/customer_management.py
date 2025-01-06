class Customer:
    def __init__(self, name, email, phone):
        self.name = name
        self.email = email
        self.phone = phone

    def display_customer(self):
        return f"Name: {self.name}, Email: {self.email}, Phone: {self.phone}"

# Function to add a customer
def add_customer(customer_list, name, email, phone):
    new_customer = Customer(name, email, phone)
    customer_list.append(new_customer)
    print("Customer added successfully!")

# Function to view all customers
def view_customers(customer_list):
    if not customer_list:
        print("No customers available.")
    for customer in customer_list:
        print(customer.display_customer())
