-- +migrate Up

CREATE TABLE IF NOT EXISTS orders
(
    id         VARCHAR(255) PRIMARY KEY,
    name       VARCHAR(10) NOT NULL,
    show_count INTEGER   DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +migrate Down

DROP TABLE IF EXISTS orders;