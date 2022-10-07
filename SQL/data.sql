INSERT INTO images (img_link) VALUES ('bucket1/1.png'), ('bucket2/2.png'), ('bucket3/3.png');

INSERT INTO users (first_name, last_name, nick_name, avatar_img_id, email, passhash) 
VALUES 
('Valera', 'Vin', 'valera', 1, 'vkV@vk.team', 'pweuga[dfnalf'), 
('Dima', 'Neu', 'p1xel', 2, 'vkD@vk.team', 'fkgms'),
('Nastya', 'Kuz', 'KuzKus', 3, 'vkN@vk.team', 'sdfmnsdmf');

INSERT INTO user_posts (user_id, message, create_date)
VALUES 
(1, 'My post1!', '2022-12-12'),
(2, 'My post2!', '2022-12-13'),
(3, 'My post2!', '2022-12-13');

INSERT INTO user_posts_images (user_post_id, img_id) 
VALUES
(1, 1),
(2, 2),
(3, 3);

