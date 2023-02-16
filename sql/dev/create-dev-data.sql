-- CATEGORY DATA
INSERT INTO categories(name) VALUES('cat 1');
INSERT INTO categories(name) VALUES('cat 2');

-- PRODUCTS DATA
INSERT INTO products(name, price, article, brand, country, weight, sale)
    VALUES('test1', 3.25, 26157, 'ok', 'kek', 180, 5);
INSERT INTO products(name, price, article, brand, country, weight, sale)
    VALUES('test2', 4.80, 421949, 'lol', 'kek', 240, 10);
INSERT INTO products(name, price, article, brand, country, weight, sale)
    VALUES('test3', 12.50, 10142, 'hah', 'kek', 65, 15);
INSERT INTO products(name, price, article, brand, country, weight, sale)
    VALUES('lol14', 13.90, 46183, 'kek', 'kek', 31, 43);

-- REVIEWS DATA
INSERT INTO reviews(name, description, date, rating, product_id)
    VALUES('rev 1', 'des 1', '2023-02-14', 3, 1);
INSERT INTO reviews(name, description, date, rating, product_id)
    VALUES('rev 2', 'des 2', '2023-02-14', 5, 1);
INSERT INTO reviews(name, description, date, rating, product_id)
    VALUES('rev 3', 'des 3', '2023-02-14', 2, 2);

-- ARTICLES DATA
INSERT INTO articles(title, description, date, image_url)
    VALUES('art 1', 'jwhefjwjjlkqjdevwnvnownwonfo', DATE(NOW()), 'qjkwcejv');
INSERT INTO articles(title, description, date, image_url)
    VALUES('art 2', 'wnevubbdiqbcwiub', DATE(NOW()), 'wfbewbevrnweur');