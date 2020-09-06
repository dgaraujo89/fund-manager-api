
CREATE TABLE IF NOT EXISTS stock_order (
    id CHAR(36) NOT NULL,
    stock VARCHAR(10) NOT NULL,
    purchase_price INT NOT NULL,
    target_price INT NOT NULL,
    purchase_date DATETIME NOT NULL,
    sale_price INT NULL,
    sale_date DATETIME NULL,
    stop_percentage NUMERIC(10,2) NULL,
    amount INT NOT NULL,
    finished CHAR(1) NOT NULL,
    order_type VARCHAR(50) NOT NULL,
    PRIMARY KEY (id)
);

CREATE INDEX stock_order_stock ON stock_order (stock);
CREATE INDEX stock_order_purchase_date ON stock_order (purchase_date);
CREATE INDEX stock_order_stock_purchase_date ON stock_order (stock, purchase_date);
CREATE INDEX stock_order_sale_date ON stock_order (sale_date);
CREATE INDEX stock_order_stock_sale_date ON stock_order (stock, sale_date);
CREATE INDEX stock_order_finished ON stock_order (finished);
CREATE INDEX stock_order_stock_finished ON stock_order (stock, finished);