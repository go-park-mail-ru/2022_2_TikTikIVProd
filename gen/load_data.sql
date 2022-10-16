COPY public.images (img_link) FROM '/home/gen/images.csv' DELIMITER ';' CSV HEADER;
COPY public.users (first_name, last_name, nick_name, avatar_img_id, email, passhash) FROM '/home/gen/users.csv' DELIMITER ';' CSV HEADER;
COPY public.user_posts (user_id, message, create_date) FROM '/home/gen/user_posts.csv' DELIMITER ';' CSV HEADER;
COPY public.user_posts_images (user_post_id, img_id) FROM '/home/gen/user_posts_images.csv' DELIMITER ';' CSV HEADER;