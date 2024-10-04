CREATE TABLE IF NOT EXISTS categories (
    category_id SMALLINT GENERATED ALWAYS AS IDENTITY,
    category_name VARCHAR(50) NOT NULL UNIQUE,
    description TEXT NOT NULL,
    image TEXT,
    CONSTRAINT pk_categories_category_id PRIMARY KEY (category_id)
);

CREATE TABLE IF NOT EXISTS suppliers (
    supplier_id SMALLINT GENERATED ALWAYS AS IDENTITY,
    company_name VARCHAR(70) NOT NULL,
    contact_name VARCHAR(70),
    contact_title VARCHAR(70),
    address VARCHAR(100),
    phone VARCHAR(50),
    contract TEXT,
    CONSTRAINT pk_suppliers_supplier_id PRIMARY KEY (supplier_id)
);

CREATE TABLE IF NOT EXISTS products (
    product_id INTEGER GENERATED ALWAYS AS IDENTITY,
    product_name VARCHAR(50) NOT NULL UNIQUE,
    supplier_id SMALLINT NOT NULL,
    category_id SMALLINT NOT NULL,
    unit_price NUMERIC(10, 2) NOT NULL,
    in_stock BOOLEAN NOT NULL DEFAULT TRUE,
    discount NUMERIC(5, 2) DEFAULT 0.0,
    quantity_per_unit VARCHAR(50),
    weight VARCHAR(10),
    image TEXT,
    CONSTRAINT pk_products_product_id PRIMARY KEY (product_id),
    CONSTRAINT fk_products_supplier_id FOREIGN KEY (supplier_id) REFERENCES suppliers(supplier_id),
    CONSTRAINT chk_unit_price CHECK (unit_price >= 0),
    CONSTRAINT chk_discount CHECK (discount >= 0)
);

CREATE TABLE IF NOT EXISTS baskets (
    customer_id VARCHAR(24) NOT NULL,
    product_id INTEGER NOT NULL,
    quantity SMALLINT DEFAULT 1,
    CONSTRAINT pk_baskets_customer_id_product_id PRIMARY KEY (customer_id, product_id)
);

CREATE TABLE IF NOT EXISTS promocodes (
    promocode_id VARCHAR(25) NOT NULL UNIQUE,
    discount NUMERIC(5, 2) DEFAULT 0.0,
    start_at DATE NOT NULL,
    finish_at DATE NOT NULL,
    CONSTRAINT pk_promocodes_promocode_id PRIMARY KEY(promocode_id),
    CONSTRAINT chk_discount CHECK (discount >= 0)
);

CREATE TABLE IF NOT EXISTS customers (
    customer_id BIGINT GENERATED ALWAYS AS IDENTITY,
    user_id VARCHAR(24),
    customer_full_name VARCHAR(50) NOT NULL,
    customer_phone VARCHAR(50) NOT NULL,
    CONSTRAINT pk_customers_customer_id PRIMARY KEY(customer_id)
);

CREATE TABLE IF NOT EXISTS orders (
    order_id BIGINT GENERATED ALWAYS AS IDENTITY,
    customer_id BIGINT NOT NULL,
    transaction_id VARCHAR(255) NOT NULL UNIQUE,
    order_date TIMESTAMPTZ DEFAULT NOW(),
    CONSTRAINT pk_orders_order_id PRIMARY KEY (order_id),
    CONSTRAINT fk_orders_customer_id FOREIGN KEY(customer_id) REFERENCES customers(customer_id)
);

CREATE TABLE IF NOT EXISTS orders_products (
    order_id BIGINT NOT NULL,
    product_id INTEGER NOT NULL,
    quantity NOT NULL,
    CONSTRAINT pk_orders_products_order_id_product_id PRIMARY KEY (order_id, product_id),
    CONSTRAINT chk_quantity CHECK (quantity >= 1)
);

CREATE TABLE IF NOT EXISTS shipping_status (
    shipping_status_id BIGINT GENERATED ALWAYS AS IDENTITY,
    shipper_id VARCHAR(24) NOT NULL,
    accepted_at TIMESTAMPTZ DEFAULT NOW(),
    delivered_at TIMESTAMPTZ,
    CONSTRAINT pk_shipping_status_shipping_status_id PRIMARY KEY(shipping_status_id)
);

CREATE TABLE IF NOT EXISTS shipping_details (
    shipping_details_id BIGINT GENERATED ALWAYS AS IDENTITY,
    shipping_status_id BIGINT,
    ship_region VARCHAR(45) NOT NULL,
    ship_city VARCHAR(35) NOT NULL,
    ship_address VARCHAR(55) NOT NULL,
    porch SMALLINT,
    floor SMALLINT,
    apartment SMALLINT,
    intercom VARCHAR(15),
    description TEXT,
    CONSTRAINT pk_shipping_details_shipping_details_id PRIMARY KEY(shipping_details_id),
    CONSTRAINT fk_shipping_details_shipping_status_id FOREIGN KEY(shipping_status_id) REFERENCES shipping_status(shipping_status_id)
);

CREATE TABLE IF NOT EXISTS orders_details (
    orders_details_id BIGINT GENERATED ALWAYS AS IDENTITY,
    order_id BIGINT NOT NULL,
    shipping_details_id BIGINT NOT NULL,
    promocode_id VARCHAR(25),
    CONSTRAINT pk_orders_details_orders_details_id PRIMARY KEY(orders_details_id),
    CONSTRAINT fk_orders_details_order_id FOREIGN KEY(order_id) REFERENCES orders(order_id),
    CONSTRAINT fk_orders_details_shipping_details_id FOREIGN KEY(shipping_details_id) REFERENCES shipping_details(shipping_details_id),
    CONSTRAINT fk_order_details_promocode_id FOREIGN KEY(promocode_id) REFERENCES promocodes(promocode_id)
);
