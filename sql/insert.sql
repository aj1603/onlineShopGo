INSERT INTO languages ("language")
VALUES ('tm'), ('en'), ('ru');

INSERT INTO admins ("name", surname, "password", email, phone_num)
VALUES
    ('Jora', 'Abdullayev', 'sowgat', 'abdullayevj.1603@gmail.com', '64806141'),
    ('Omar', 'Abdullayev', 'omar2003', 'abdullayevo.@gmail.com', '64453960');

INSERT INTO categories ("name")
VALUES
    ('Electronika we robot'),
    ('Hojalyk harytlary'),
    ('Bezeg harytlary'),
    ('Arassacylyk harytlary'),
    ('Azyk harytlary');

INSERT INTO discounts ("name", "description", discount_rate, start_time, end_time)
VALUES
    ('default', 'default', 0.0, '03/03/2024', '03/03/2025');

INSERT INTO discounts_trans ("name", "description", discounts_id, languages_id)
VALUES
    ('default', 'default', 1, 2),
    ('default', 'default', 1, 3);

INSERT INTO brands ("name", img_url)
VALUES
    ('Coco_cola', '/public/products/1.jpg'),
    ('Salam', '/public/products/2.jpg');

INSERT INTO products ("name", "description", price, product_sku, quantity, categories_id, discounts_id, brands_id)
VALUES
    ('Arduino', 'Orginal made in italy', 378.0, '0321654', 150, 1, 1, 1),
    ('Arduino LCD display', 'LCD display', 99.0, '0321655', 20, 1, 1, 1),
    ('Lion battery charger', 'DC plug charger', 441.0, '0321656', 25, 1, 1, 1),
    ('ESP-8266 WiFi module', 'WiFi module', 140.0, '0321657', 10, 1, 1, 1);

INSERT INTO order_status ("status")
VALUES
    ('Kabul edildi'),
    ('Garaşylýar'),
    ('Sargyt ýatyryldy');

INSERT INTO order_status_trans ("status", languages_id, order_status_id)
VALUES
    ('Delivered', 2, 1),
    ('Waiting', 2, 2),
    ('Canceled', 2, 3),
    ('Доставлено', 3, 1),
    ('В ожидании', 3, 2),
    ('Отмена', 3, 3);

INSERT INTO payment_methods (method)
VALUES
    ('Kartdan töleg'),
    ('Nagt töleg');

INSERT INTO payment_method_trans (method, languages_id, payment_methods_id)
VALUES
    ('Payment terminal', 2, 1),
    ('Cash', 2, 2),
    ('Банковская карта', 3, 1),
    ('Наличными', 3, 2);

INSERT INTO products_images (img_url, products_id)
VALUES
    ('/public/products/2.jpg', 1),
    ('/public/products/4.jpg', 2),
    ('/public/products/3.jpg', 3),
    ('/public/products/1.jpg', 4);