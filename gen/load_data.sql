COPY public.images (img_link) FROM '/home/gen/images.csv' DELIMITER ';' CSV HEADER;
COPY public.users (first_name, last_name, nick_name, avatar_img_id, email, password, created_at) FROM '/home/gen/users.csv' DELIMITER ';' CSV HEADER;
COPY public.communities (name, owner_id, avatar_img_id, description, created_at) FROM '/home/gen/communities.csv' DELIMITER ';' CSV HEADER;
COPY public.user_posts (user_id, community_id, description, created_at) FROM '/home/gen/user_posts.csv' DELIMITER ';' CSV HEADER;
COPY public.user_posts_images (user_post_id, img_id) FROM '/home/gen/user_posts_images.csv' DELIMITER ';' CSV HEADER;
