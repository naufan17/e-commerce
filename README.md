# e-commerce
A basic inline store API written with Golang Programming language

This API is basic implementation of an e-commerce(online store)
- You can diplay product and category data
- Also can display product bt category
- Authentication is based JWT(JSON Web Token)
- Register and login to get Token
- User can display, add and delete product from cart
- Built using golang programming language and MySQL database

## API Docementation and Usage

### 1. View Category
- **Method** - `GET` <br>
- **URL Pattern** - `/categories` <br>
- **Authentication** - `false` <br>
- **Example**
![Screenshot](/screenshot/GetCategory.png)
### 2. View Product
- **Method** - `GET` <br>
- **URL Pattern** - `/products` <br>
- **Authentication** - `false` <br>
- **Example**
![Screenshot](/screenshot/GetProduct.png)
### 3. View Product by Category
- **Method** - `GET` <br>
- **URL Pattern** - `/products?category={category}` <br>
- **Authentication** - `false` <br>
- **Example**
![Screenshot](/screenshot/GetProductbyCategory.png)
### 4. Register User
- **Method** - `POST` <br>
- **URL Pattern** - `/register` <br>
- **Authentication** - `false` <br>
- **Example**
![Screenshot](/screenshot/Register.png)
### 5. Login User
- **Method** - `POST` <br>
- **URL Pattern** - `/login` <br>
- **Authentication** - `false` <br>
- **Example**
![Screenshot](/screenshot/Login.png)
### 6. Add Product to Cart
- **Method** - `POST` <br>
- **URL Pattern** - `/cart` <br>
- **Authentication** - `true` <br>
- **Example**
![Screenshot](/screenshot/PostCart.png)
### 7. View Product in Cart
- **Method** - `GET` <br>
- **URL Pattern** - `/cart` <br>
- **Authentication** - `true` <br>
- **Example**
![Screenshot](/screenshot/GetCart.png)
### 8. Delete Product in Cart
- **Method** - `DELETE` <br>
- **URL Pattern** - `/cart/{cart_id}` <br>
- **Authentication** - `true` <br>
- **Example**
![Screenshot](/screenshot/DeleteCart.png)
