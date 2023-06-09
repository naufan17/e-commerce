# e-commerce
A basic online store API written with Golang Programming language

This API is basic implementation of an e-commerce(online store)
- You can diplay product and category data
- Also can display product by category
- Authentication is based JWT(JSON Web Token)
- Register for new user to get token
- Login for registered user to get token
- User can display, add and delete shipping address
- User can display, add and delete product in cart
- User can orders product in cart by entering shipping address
- Built using golang programming language and MySQL database
- Applications are bundled in containers using Docker

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

### 6. Add Shipping Address
- **Method** - `POST` <br>
- **URL Pattern** - `/address` <br>
- **Authentication** - `true` <br>
- **Usage**
```
curl -X POST \
-H "Authorization: Bearer <ACCESS_TOKEN>" \
-d '{ "shipping_address": "shipping_address"}' \
BASE_URL/address
```
- **Example**
![Screenshot](/screenshots/PostAddress.png)

### 7. View Shipping Address
- **Method** - `GET` <br>
- **URL Pattern** - `/address` <br>
- **Authentication** - `true` <br>
- **Usage**
```
curl -X GET \
-H "Authorization: Bearer <ACCESS_TOKEN>" \
BASE_URL/address
```
- **Example**
![Screenshot](/screenshots/GetAddress.png)

### 8. Update Shipping Address
- **Method** - `PUT` <br>
- **URL Pattern** - `/address` <br>
- **Authentication** - `true` <br>
- **Usage**
```
curl -X PUT \
-H "Authorization: Bearer <ACCESS_TOKEN>" \
-d '{ "address_id": "address_id",
    "shipping_address": "shipping_address"}' \
BASE_URL/address
```
- **Example**
![Screenshot](/screenshots/PutAddress.png)

### 9. Delete Shipping Address
- **Method** - `DELETE` <br>
- **URL Pattern** - `/address/{address_id}` <br>
- **Authentication** - `true` <br>
- **Usage**
```
curl -X DELETE \
-H "Authorization: Bearer <ACCESS_TOKEN>" \
BASE_URL/address/{address_id}
```
- **Example**
![Screenshot](/screenshots/DeleteAddress.png)

### 10. Add Product to Cart
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

### 11. View Product in Cart
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

### 12. Update Product in Cart
- **Method** - `PUT` <br>
- **URL Pattern** - `/cart` <br>
- **Authentication** - `true` <br>
- **Usage**
```
curl -X PUT \
-H "Authorization: Bearer <ACCESS_TOKEN>" \
-d '{ "cart_id": "cart_id", 
    "count": "count"}' \
BASE_URL/cart
```
- **Example**
![Screenshot](/screenshots/PutCart.png)

### 13. Delete Product in Cart
- **Method** - `DELETE` <br>
- **URL Pattern** - `/cart/{cart_id}` <br>
- **Authentication** - `true` <br>
- **Usage**
```
curl -X DELETE \
-H "Authorization: Bearer <ACCESS_TOKEN>" \
BASE_URL/cart/{cart_id}
```
- **Example**
![Screenshot](/screenshots/DeleteCart.png)


