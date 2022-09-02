CREATE DATABASE product_db;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE products (
	product_id serial4 NOT NULL,
	product_item_id UUID NOT NULL DEFAULT uuid_generate_v1(),
    product_tile VARCHAR(100) NOT NULL,
	product_description VARCHAR(100) NOT NULL,
	product_information VARCHAR(100) NOT NULL,
	product_datetime TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT product_page_pkey PRIMARY KEY (product_id),
	UNIQUE(product_item_id)
);

CREATE TABLE products_sku (
	product_sku_id serial4 NOT NULL,
	product_id serial4 NOT NULL,
    product_sku_code VARCHAR(20) NOT NULL,
	product_sku_option VARCHAR(100) NOT NULL,
	product_sku_option_value VARCHAR(100) NOT NULL,
	product_sku_details VARCHAR(100) NOT NULL,
	product_sku_qtd INTEGER DEFAULT 0 NOT NULL,
	product_sku_price NUMERIC (5,2) DEFAULT 0.0 NOT NULL,
	CONSTRAINT products_sku_pkey PRIMARY KEY (product_sku_id),
	CONSTRAINT fk_products FOREIGN key (product_id) REFERENCES products (product_id)
);