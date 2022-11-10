CREATE TABLE IF NOT EXISTS images (
	id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	img_link TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS users (
	id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	first_name VARCHAR(32) NOT NULL,
	last_name VARCHAR(32) NOT NULL,
	nick_name VARCHAR(30) NOT NULL UNIQUE,
	avatar_img_id INT REFERENCES images(id),
	email VARCHAR(50) NOT NULL UNIQUE,
	password VARCHAR(128) NOT NULL
);

CREATE TABLE IF NOT EXISTS friends (
	id1 INT REFERENCES users(id) on delete cascade,
	id2 INT REFERENCES users(id) on delete cascade,
	PRIMARY KEY (id1, id2)
);

CREATE TABLE IF NOT EXISTS user_posts (
	id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	user_id INT REFERENCES users(id) on delete cascade,
	message TEXT NOT NULL,
	create_date DATE NOT NULL
);

CREATE TABLE IF NOT EXISTS user_posts_images (
	user_post_id INT REFERENCES user_posts(id) on delete cascade,
	img_id INT REFERENCES images(id) on delete cascade,
	PRIMARY KEY (user_post_id, img_id)
);

CREATE TABLE IF NOT EXISTS chat (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id1 INT REFERENCES users(id),
	user_id2 INT REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS message (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    body text DEFAULT '',
    sender_id INT REFERENCES users(id),
	receiver_id INT REFERENCES users(id),
	chat_id INT REFERENCES chat(id),
    created_at date NOT NULL
);
