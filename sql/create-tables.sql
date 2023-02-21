CREATE TABLE products(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255),
    price DECIMAL,
    image_url VARCHAR(255),
    article INT,
    rating FLOAT DEFAULT 0.0,
    brand VARCHAR(50),
    country VARCHAR(50),
    weight INT,
    sale INT,
    category_id BIGINT
);
-- ADD CATEGORY_ID COLUMN

CREATE TABLE reviews(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(1000),
    description VARCHAR(255),
    date DATE,
    rating INT,
    product_id BIGINT,
    CONSTRAINT fk_products FOREIGN KEY(product_id) REFERENCES products(id)
);

CREATE TABLE categories(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(50),
    image_url VARCHAR(255)
);

CREATE TABLE articles(
    id BIGSERIAL PRIMARY KEY,
    title VARCHAR(100),
    description VARCHAR(1000),
    date DATE,
    image_url VARCHAR(255)
);

CREATE TABLE special_offers(
    id BIGSERIAL PRIMARY KEY,
    title VARCHAR(100),
    description VARCHAR(1000),
    image_url VARCHAR(255)
);

ALTER TABLE products
ADD CONSTRAINT FK_Products_Categories FOREIGN KEY(category_id)
    REFERENCES categories(id);