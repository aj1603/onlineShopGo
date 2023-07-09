CREATE TABLE languages (
    id SERIAL PRIMARY KEY,
    "language" VARCHAR
);

CREATE TABLE admins (
    id SERIAL PRIMARY KEY,
    "name" VARCHAR,
    surname VARCHAR,
    "password" VARCHAR,
    email VARCHAR UNIQUE,
    phone_num INTEGER UNIQUE
);

CREATE TABLE customers (
    id SERIAL PRIMARY KEY,
    "name" VARCHAR,
    "password" VARCHAR,
    email VARCHAR UNIQUE,
    phone_num INTEGER UNIQUE
);

CREATE TABLE addresss (
    id SERIAL PRIMARY KEY,
    region VARCHAR,
    city VARCHAR,
    address_line VARCHAR,
    customers_id INTEGER,

    CONSTRAINT customers_id
        FOREIGN KEY (customers_id)
            REFERENCES customers (id)
);

CREATE TABLE addresstrans (
    id SERIAL PRIMARY KEY,
    region VARCHAR,
    city VARCHAR,
    addresss_id INTEGER,
    languages_id INTEGER,

    CONSTRAINT addresss_id
        FOREIGN KEY (addresss_id)
            REFERENCES addresss (id),

    CONSTRAINT languages_id
        FOREIGN KEY (languages_id)
            REFERENCES languages (id)
);

CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    "name" VARCHAR,
    img_url VARCHAR
);

CREATE TABLE discounts (
    id SERIAL PRIMARY KEY,
    "name" VARCHAR,
    "description" VARCHAR,
    discount_rate FLOAT,
    start_time DATE,
    end_time DATE
);

CREATE TABLE brands (
    id SERIAL PRIMARY KEY,
    "name" VARCHAR,
    img_url VARCHAR
);


CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    "name" VARCHAR,
    "description" VARCHAR,
    price FLOAT NOT NULL,
    product_sku VARCHAR UNIQUE,
    quantity INTEGER,
    categories_id INTEGER,
    discounts_id INTEGER,
    brands_id INTEGER,

    CONSTRAINT categories_id
        FOREIGN KEY (categories_id)
            REFERENCES categories (id),

    CONSTRAINT discounts_id
        FOREIGN KEY (discounts_id)
            REFERENCES discounts (id),

    CONSTRAINT brands_id
        FOREIGN KEY (brands_id)
            REFERENCES brands (id)
);

CREATE TABLE favorities (
    id SERIAL PRIMARY KEY,
    customers_id INTEGER,
    products_id INTEGER,

    CONSTRAINT customers_id
        FOREIGN KEY (customers_id)
            REFERENCES customers (id),

    CONSTRAINT products_id
        FOREIGN KEY (products_id)
            REFERENCES products (id)
);

CREATE TABLE order_status (
    id SERIAL PRIMARY KEY,
    "status" VARCHAR
);

CREATE TABLE payment_methods (
    id SERIAL PRIMARY KEY,
    method VARCHAR
);

CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    total FLOAT NOT NULL,
    customers_id INTEGER,
    addresss_id INTEGER,
    order_status_id INTEGER,
    payment_methods_id INTEGER,

    CONSTRAINT customers_id
        FOREIGN KEY (customers_id)
            REFERENCES customers (id),

    CONSTRAINT addresss_id
        FOREIGN KEY (addresss_id)
            REFERENCES addresss (id),

    CONSTRAINT order_status_id
        FOREIGN KEY (order_status_id)
            REFERENCES order_status (id),

    CONSTRAINT payment_methods_id
        FOREIGN KEY (payment_methods_id)
            REFERENCES payment_methods (id)
);

CREATE TABLE order_items (
    id SERIAL PRIMARY KEY,
    quantity INTEGER,
    products_id INTEGER,
    orders_id INTEGER,

    CONSTRAINT products_id
        FOREIGN KEY (products_id)
            REFERENCES products (id),

    CONSTRAINT orders_id
        FOREIGN KEY (orders_id)
            REFERENCES orders (id)
);

CREATE TABLE payment_method_trans (
    id SERIAL PRIMARY KEY,
    method VARCHAR,
    languages_id INTEGER,
    payment_methods_id INTEGER,

    CONSTRAINT languages_id
        FOREIGN KEY (languages_id)
            REFERENCES languages (id),

    CONSTRAINT payment_methods_id
        FOREIGN KEY (payment_methods_id)
            REFERENCES payment_methods (id)
);

CREATE TABLE order_status_trans (
    id SERIAL PRIMARY KEY,
    "status" VARCHAR,
    languages_id INTEGER,
    order_status_id INTEGER,

    CONSTRAINT languages_id
        FOREIGN KEY (languages_id)
            REFERENCES languages (id),

    CONSTRAINT order_status_id
        FOREIGN KEY (order_status_id)
            REFERENCES order_status (id)
);

CREATE TABLE carts (
    id SERIAL PRIMARY KEY,
    carts_sku VARCHAR UNIQUE,
    customers_id INTEGER,

    CONSTRAINT customers_id
        FOREIGN KEY (customers_id)
            REFERENCES customers (id)
);

CREATE TABLE cart_item (
    id SERIAL PRIMARY KEY,
    quantity INTEGER,
    products_id INTEGER,
    carts_id INTEGER,

    CONSTRAINT products_id
        FOREIGN KEY (products_id)
            REFERENCES products (id),

    CONSTRAINT carts_id
        FOREIGN KEY (carts_id)
            REFERENCES carts (id)
);

CREATE TABLE products_trans (
    id SERIAL PRIMARY KEY,
    "name" VARCHAR,
    "description" VARCHAR,
    products_id INTEGER,
    languages_id INTEGER,

    CONSTRAINT products_id
        FOREIGN KEY (products_id)
            REFERENCES products (id),

    CONSTRAINT languages_id
        FOREIGN KEY (languages_id)
            REFERENCES languages (id)
);

CREATE TABLE products_images (
    id SERIAL PRIMARY KEY,
    img_url VARCHAR,
    products_id INTEGER,

    CONSTRAINT products_id
        FOREIGN KEY (products_id)
            REFERENCES products (id)
);

CREATE TABLE categories_trans (
    id SERIAL PRIMARY KEY,
    "name" VARCHAR,
    categories_id INTEGER,
    languages_id INTEGER,

    CONSTRAINT categories_id
        FOREIGN KEY (categories_id)
            REFERENCES categories (id),

    CONSTRAINT languages_id
        FOREIGN KEY (languages_id)
            REFERENCES languages (id)
);

CREATE TABLE discounts_trans (
    id SERIAL PRIMARY KEY,
    "name" VARCHAR,
    "description" VARCHAR,
    discounts_id INTEGER,
    languages_id INTEGER,

    CONSTRAINT discounts_id
        FOREIGN KEY (discounts_id)
            REFERENCES discounts (id),

    CONSTRAINT languages_id
        FOREIGN KEY (languages_id)
            REFERENCES languages (id)
);

CREATE TABLE products_spec (
    id SERIAL PRIMARY KEY,
    "name" VARCHAR,
    categories_id INTEGER,

    CONSTRAINT categories_id
        FOREIGN KEY (categories_id)
            REFERENCES categories (id)
);

CREATE TABLE spec_value (
    id SERIAL PRIMARY KEY,
    "value" VARCHAR,
    products_spec_id INTEGER,

    CONSTRAINT products_spec_id
        FOREIGN KEY (products_spec_id)
            REFERENCES products_spec (id)
);

CREATE TABLE products_config (
    id SERIAL PRIMARY KEY,
    products_id INTEGER,
    spec_value_id INTEGER,

    CONSTRAINT products_id
        FOREIGN KEY (products_id)
            REFERENCES products (id),
            
    CONSTRAINT spec_value_id
        FOREIGN KEY (spec_value_id)
            REFERENCES spec_value (id)
);

CREATE TABLE products_spec_trans (
    id SERIAL PRIMARY KEY,
    "name" VARCHAR,
    products_spec_id INTEGER,
    languages_id INTEGER,

    CONSTRAINT products_spec_id
        FOREIGN KEY (products_spec_id)
            REFERENCES products_spec (id),

    CONSTRAINT languages_id
        FOREIGN KEY (languages_id)
            REFERENCES languages (id)
);

CREATE TABLE hytay_turk (
    id SERIAL PRIMARY KEY,
    customers_id INTEGER,

    CONSTRAINT customers_id
        FOREIGN KEY (customers_id)
            REFERENCES customers (id)
);

CREATE TABLE hytay_turk_item (
    id SERIAL PRIMARY KEY,
    link_for_product VARCHAR,
    title VARCHAR,
    "description" VARCHAR,
    hytay_turk_id INTEGER,

    CONSTRAINT hytay_turk_id
        FOREIGN KEY (hytay_turk_id)
            REFERENCES hytay_turk (id)
);