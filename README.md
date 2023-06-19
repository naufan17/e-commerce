# e-commerce
A basic online store API written with Golang Programming language

This API is basic implementation of an e-commerce(online store)
- You can diplay product and category data
- Also can display product by category
- Authentication is based JWT(JSON Web Token)
- Register and login to get token
- User can display, add and delete product from cart
- Built using golang programming language and MySQL database

## API Documentation and Usage

### 1. View Category
- **Method** - `GET` <br>
- **URL Pattern** - `/categories` <br>
- **Authentication** - `false` <br>
- **Usage**
```
curl -X GET BASE_URL/categories
```
- **Example**
![Screenshot](/screenshots/GetCategory.png)

### 2. View Product
- **Method** - `GET` <br>
- **URL Pattern** - `/products` <br>
- **Authentication** - `false` <br>
- **Usage**
```
curl -X GET BASE_URL/products 
```
- **Example**
![Screenshot](/screenshots/GetProduct.png)

### 3. View Product by Category
- **Method** - `GET` <br>
- **URL Pattern** - `/products?category={category}` <br>
- **Authentication** - `false` <br>
- **Usage**
```
curl -X GET BASE_URL/products?=category={category}
```
- **Example**
![Screenshot](/screenshots/GetProductbyCategory.png)

### 4. Register User
- **Method** - `POST` <br>
- **URL Pattern** - `/register` <br>
- **Authentication** - `false` <br>
- **Usage**
```
curl -X POST \
-d '{ "username": "username", 
    "password": "password"}' \
BASE_URL/register
```
- **Example**
![Screenshot](/screenshots/Register.png)

### 5. Login User
- **Method** - `POST` <br>
- **URL Pattern** - `/login` <br>
- **Authentication** - `false` <br>
- **Usage**
```
curl -X POST \
-d '{ "username": "username", 
    "password": "password"}' \
BASE_URL/login
```
- **Example**
![Screenshot](/screenshots/Login.png)

### 6. Add Product to Cart
- **Method** - `POST` <br>
- **URL Pattern** - `/cart` <br>
- **Authentication** - `true` <br>
- **Usage**
```
curl -X POST \
-H "Authorization: Bearer <ACCESS_TOKEN>" \
-d '{ "product_id": "product_id", 
    "count": "count"}' \
BASE_URL/cart
```
- **Example**
![Screenshot](/screenshots/PostCart.png)

### 7. View Product in Cart
- **Method** - `GET` <br>
- **URL Pattern** - `/cart` <br>
- **Authentication** - `true` <br>
- **Usage**
```
curl -X GET \
-H "Authorization: Bearer <ACCESS_TOKEN>" \
BASE_URL/cart
```
- **Example**
![Screenshot](/screenshots/GetCart.png)

### 8. Update Product in Cart
- **Method** - `PUT` <br>
- **URL Pattern** - `/cart` <br>
- **Authentication** - `true` <br>
- **Usage**
```
curl -X PUT \
-H "Authorization: Bearer <ACCESS_TOKEN>" \
BASE_URL/cart
```
- **Example**
![Screenshot](/screenshots/PutCart.png)

### 9. Delete Product in Cart
- **Method** - `DELETE` <br>
- **URL Pattern** - `/cart/{cart_id}` <br>
- **Authentication** - `true` <br>
- **Usage**
```
curl -X DELETE \
-H "Authorization: Bearer <ACCESS_TOKEN>" \
BASE_URL/cart/{id}
```
- **Example**
![Screenshot](/screenshots/DeleteCart.png)
