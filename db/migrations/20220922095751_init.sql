-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "brands"
(
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    name TEXT  NOT NULL
);


CREATE TABLE IF NOT EXISTS "products"
(
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    name TEXT  NOT NULL,
    price INTEGER  NOT NULL,
    brand_id INTEGER NOT NULL,
    FOREIGN KEY (brand_id) REFERENCES "brands" (id)
    ON DELETE NO ACTION ON UPDATE NO ACTION
    );

CREATE TABLE IF NOT EXISTS "orders"
(
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    total INTEGER  NOT NULL
);

CREATE TABLE IF NOT EXISTS "orders_products"
(
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    product_id INTEGER NOT NULL,
    order_id INTEGER  NOT NULL,
    product_price INTEGER NOT NULL,
    amount INTEGER NOT NULL,
    total INTEGER NOT NULL,
    FOREIGN KEY (product_id) REFERENCES "products" (id)
    ON DELETE NO ACTION ON UPDATE NO ACTION,
    FOREIGN KEY (order_id) REFERENCES "orders" (id)
    ON DELETE NO ACTION ON UPDATE NO ACTION
    );


CREATE INDEX [IFK_ProductsBrandId] ON "products" (brand_id);
CREATE INDEX [IFK_OrdersProductsProductId] ON "orders_products" (product_id);
CREATE INDEX [IFK_OrdersProductsOrderId] ON "orders_products" (order_id);
CREATE INDEX [IFK_OrdersProductsOrderIdProductID] ON "orders_products" (order_id, product_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE brands;
DROP TABLE products;
DROP TABLE orders;
DROP TABLE orders_products;

-- +goose StatementEnd

