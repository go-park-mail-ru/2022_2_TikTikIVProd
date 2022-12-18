COPY public.images (img_link) FROM '/home/gen/images.csv' DELIMITER ';' CSV HEADER;
COPY public.files (file_link) FROM '/home/gen/files.csv' DELIMITER ';' CSV HEADER;
COPY public.stickers (link) FROM '/home/gen/stickers.csv' DELIMITER ';' CSV HEADER;
COPY public.users (first_name, last_name, nick_name, avatar_img_id, email, password, created_at) FROM '/home/gen/users.csv' DELIMITER ';' CSV HEADER;
COPY public.communities (name, owner_id, avatar_img_id, description, created_at) FROM '/home/gen/communities.csv' DELIMITER ';' CSV HEADER;
COPY public.user_posts (user_id, community_id, description, created_at) FROM '/home/gen/user_posts.csv' DELIMITER ';' CSV HEADER;
COPY public.user_posts_images (user_post_id, img_id) FROM '/home/gen/user_posts_images.csv' DELIMITER ';' CSV HEADER;
COPY public.user_posts_files (user_post_id, file_id) FROM '/home/gen/user_posts_files.csv' DELIMITER ';' CSV HEADER;
COPY public.like_post (user_post_id, user_id) FROM '/home/gen/like_post.csv' DELIMITER ';' CSV HEADER;
