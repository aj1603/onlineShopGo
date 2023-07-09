INSERT INTO languages ("language")
VALUES ('tm'), ('en'), ('ru');

INSERT INTO admins ("name", surname, "password", email, phone_num)
VALUES
    ('Jora', 'Abdullayev', 'sowgat', 'abdullayevj.1603@gmail.com', '64806141'),
    ('Omar', 'Abdullayev', 'omar2003', 'abdullayevo.@gmail.com', '64453960');

INSERT INTO categories ("name", img_url)
VALUES
    ('Electronika we robot', 'http://192.168.0.103:8091/public/category/1.png'),
    ('Hojalyk harytlary', 'http://192.168.0.103:8091/public/category/2.png'),
    ('Bezeg harytlary', 'http://192.168.0.103:8091/public/category/3.png'),
    ('Arassacylyk harytlary', 'http://192.168.0.103:8091/public/category/4.png'),
    ('Azyk harytlary', 'http://192.168.0.103:8091/public/category/5.png');

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
    ('ESP-8266 WiFi module', 'WiFi module', 140.0, '0321657', 10, 1, 1, 1),
    ('Palta', 'Orginal made in turkmenistan', 32.0, '0321658', 150, 2, 1, 1),
    ('Orak', 'Orginal made in turkmenistan', 25.0, '0321659', 20, 2, 1, 1),
    ('Relay modul', '12V relay modul', 23.0, '0321660', 25, 1, 1, 1),
    ('Pulse sensor', 'Yurek urgusyny anyklayan sensor', 140.0, '0321661', 10, 1, 1, 1),
    ('Oliva sabyn', 'Arzan bahadan', 8.0, '0321664', 25, 4, 1, 1),
    ('Ariel', '2in1 soda mashyn ucin', 105.0, '0321665', 10, 4, 1, 1),
    ('Elidor', 'Shampun beti Uniliver', 56.0, '0321666', 150, 4, 1, 1),
    ('Fa sabyny', '150gr aramy', 12.0, '0321667', 20, 4, 1, 1),
    ('Dorgo stanok', '5in1 paketli', 7.0, '0321668', 25, 4, 1, 1),
    ('Arko', 'Arko krem posle britva', 43.0, '0321669', 10, 4, 1, 1);

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
    ('http://192.168.0.103:8091/public/products/2.jpg', 1),
    ('http://192.168.0.103:8091/public/products/4.jpg', 2),
    ('http://192.168.0.103:8091/public/products/3.jpg', 3),
    ('http://192.168.0.103:8091/public/products/5.jpg', 5),
    ('http://192.168.0.103:8091/public/products/6.jpg', 6),
    ('http://192.168.0.103:8091/public/products/7.jpg', 9),
    ('http://192.168.0.103:8091/public/products/8.jpg', 9),
    ('http://192.168.0.103:8091/public/products/9.jpg', 12),
    ('http://192.168.0.103:8091/public/products/10.jpg', 12),
    ('http://192.168.0.103:8091/public/products/11.jpg', 10),
    ('http://192.168.0.103:8091/public/products/12.jpg', 9),
    ('http://192.168.0.103:8091/public/products/13.jpg', 9),
    ('http://192.168.0.103:8091/public/products/14.jpg', 11),
    ('http://192.168.0.103:8091/public/products/15.jpg', 11),
    ('http://192.168.0.103:8091/public/products/16.jpg', 11),
    ('http://192.168.0.103:8091/public/products/17.jpg', 8),
    ('http://192.168.0.103:8091/public/products/18.jpg', 13),
    ('http://192.168.0.103:8091/public/products/19.jpg', 13),
    ('http://192.168.0.103:8091/public/products/20.jpg', 7),
    ('http://192.168.0.103:8091/public/products/20.jpg', 14),
    ('http://192.168.0.103:8091/public/products/1.jpg', 4);