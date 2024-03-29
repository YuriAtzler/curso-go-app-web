# My Go Store Application

This is a Go application that simulates a store, allowing users to create, edit, and remove products.

## Features

- Create products: Users can add new products to the store.
- Edit products: Users can update the details of existing products.
- Remove products: Users can delete products from the store.

## Installation

To run this application, you need to have Go installed on your machine and postgresql 16. Follow these steps:

1. Clone the repository: `git clone https://github.com/YuriAtzler/curso-go-app-web.git`
2. Navigate to the project directory: `cd curso-go-app-web.git`
3. Start the application: `go run main.go`
4. Run the application: `localhost:8000`

### Important

1. Configure your database in db/db.go
2. Create a table called `produtos` with columns `id serial primary key, nome varchar, descricao varchar, preco numeric, quantidade, integer`. 
