CREATE TABLE IF NOT EXISTS categories(
    category_id SMALLINT GENERATED ALWAYS AS IDENTITY,
    category_name VARCHAR(50) NOT NULL UNIQUE,
    description TEXT NOT NULL,
    image TEXT,
    CONSTRAINT pk_categories_category_id PRIMARY KEY (category_id)
);

CREATE TABLE IF NOT EXISTS suppliers(
    supplier_id SMALLINT GENERATED ALWAYS AS IDENTITY,
    company_name VARCHAR(70) NOT NULL,
    contact_name VARCHAR(70),
    contact_title VARCHAR(70),
    address VARCHAR(100),
    phone VARCHAR(24),
    contract TEXT,
    CONSTRAINT pk_suppliers_supplier_id PRIMARY KEY (supplier_id)
);

CREATE TABLE IF NOT EXISTS products(
    product_id INTEGER GENERATED ALWAYS AS IDENTITY,
    product_name VARCHAR(50) NOT NULL UNIQUE,
    supplier_id SMALLINT NOT NULL,
    category_id SMALLINT NOT NULL,
    unit_price REAL NOT NULL,
    in_stock BOOLEAN NOT NULL DEFAULT TRUE,
    discount REAL DEFAULT 0.0,
    quantity_per_init VARCHAR(50),
    weight VARCHAR(10),
    image TEXT,
    CONSTRAINT pk_products_product_id PRIMARY KEY (product_id),
    CONSTRAINT fk_products_supplier_id FOREIGN KEY (supplier_id) REFERENCES suppliers(supplier_id),
    CONSTRAINT chk_unit_price CHECK (unit_price >= 0),
    CONSTRAINT chk_discount CHECK (discount >= 0)
);

CREATE TABLE IF NOT EXISTS promocodes(
    promocode_id VARCHAR(25) NOT NULL UNIQUE,
    discount REAL DEFAULT 0.0,
    start_at DATE NOT NULL,
    finish_at DATE NOT NULL,
    pk_promocodes_promocode_id PRIMARY KEY(promocode_id),
    CONSTRAINT chk_discount CHECK (discount >= 0)
);

CREATE TABLE IF NOT EXISTS orders (
    order_id BIGINT GENERATED ALWAYS AS IDENTITY,
    customer_id VARCHAR(24) NOT NULL,
    order_date TIMESTAMPTZ DEFAULT NOW(),
    shipper_id VARCHAR(24),
    accepted_at TIMESTAMPTZ,
    delivered_at TIMESTAMPTZ,
    ship_region VARCHAR(45) NOT NULL,
    ship_city VARCHAR(35) NOT NULL,
    ship_address VARCHAR(55) NOT NULL,
    porch SMALLINT,
    floor SMALLINT,
    apartment SMALLINT,
    intercom VARCHAR(15),
    description TEXT,
    CONSTRAINT pk_orders_order_id PRIMARY KEY (order_id)
);

CREATE TABLE IF NOT EXISTS orders_details(
    order_id BIGINT NOT NULL,
    product_id INTEGER NOT NULL,
    quantity SMALLINT NOT NULL,
    promocode_id VARCHAR(25),
    CONSTRAINT pk_order_details PRIMARY KEY (order_id, product_id),
    CONSTRAINT fk_order_details_order_id FOREIGN KEY(order_id) REFERENCES orders(order_id),
    CONSTRAINT fk_order_details_product_id FOREIGN KEY(product_id) REFERENCES products(product_id),
    CONSTRAINT fk_order_details_promocode_id FOREIGN KEY(promocode_id) REFERENCES promocodes(promocode_id),
    CONSTRAINT chk_quantity CHECK (quantity >= 1)
);

CREATE TABLE IF NOT EXISTS receipts (
    receipt_id BIGINT GENERATED ALWAYS AS IDENTITY,
    total_amount REAL NOT NULL,
    receipt_date TIMESTAMPTZ DEFAULT NOW(),
    CONSTRAINT pk_receipts_receipt_id PRIMARY KEY (receipt_id)
);

CREATE TABLE IF NOT EXISTS receipt_details (
    receipt_id BIGINT NOT NULL,
    order_id BIGINT NOT NULL,
    CONSTRAINT pk_receipt_details PRIMARY KEY (receipt_id, order_id),
    CONSTRAINT fk_receipt_details_receipt_id FOREIGN KEY (receipt_id) REFERENCES receipts(receipt_id),
    CONSTRAINT fk_receipt_details_order_id FOREIGN KEY (order_id) REFERENCES orders(order_id)
);
