CREATE TYPE user_role AS ENUM ('buyer', 'seller');

CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    role user_role NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE products (
    product_id SERIAL PRIMARY KEY,
    seller_id_fk INTEGER NOT NULL,
    name VARCHAR(255) NOT NULL,
    quantity INTEGER NOT NULL,
    price INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT seller_fk FOREIGN KEY (seller_id_fk) REFERENCES users(user_id)
);

CREATE TABLE orders (
    order_id SERIAL PRIMARY KEY,
    user_id_fk INTEGER NOT NULL,
    product_id_fk INTEGER NOT NULL,
    name VARCHAR(255) NOT NULL,
    quantity INTEGER NOT NULL,
    total_price INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT user_fk FOREIGN KEY (user_id_fk) REFERENCES users(user_id),
    CONSTRAINT product_fk FOREIGN KEY (product_id_fk) REFERENCES products(product_id)
)