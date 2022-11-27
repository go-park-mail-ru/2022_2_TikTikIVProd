CREATE TABLE IF NOT EXISTS images (
	id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	img_link VARCHAR(260) NOT NULL
);

CREATE TABLE IF NOT EXISTS users (
	id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	first_name VARCHAR(35) NOT NULL,
	last_name VARCHAR(35) NOT NULL,
	nick_name VARCHAR(30) NOT NULL UNIQUE,
	avatar_img_id INT REFERENCES images(id),
	email VARCHAR(254) NOT NULL UNIQUE,
	password VARCHAR(128) NOT NULL,
	created_at date NOT NULL
);

CREATE TABLE IF NOT EXISTS friends (
	user_id1 INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
	user_id2 INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
	PRIMARY KEY (user_id1, user_id2)
);

CREATE TABLE IF NOT EXISTS communities (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    owner_id INT REFERENCES users(id),
    avatar_img_id INT REFERENCES images(id),
    name VARCHAR(30) NOT NULL,
    description TEXT DEFAULT '',
    created_at date NOT NULL
);

CREATE TABLE IF NOT EXISTS user_posts (
	id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
	community_id INT REFERENCES communities(id),
	description TEXT NOT NULL DEFAULT '',
	created_at DATE NOT NULL
);

CREATE TABLE IF NOT EXISTS user_posts_images (
	user_post_id INT NOT NULL REFERENCES user_posts(id) ON DELETE CASCADE,
	img_id INT NOT NULL REFERENCES images(id) ON DELETE CASCADE,
	PRIMARY KEY (user_post_id, img_id)
);

CREATE TABLE IF NOT EXISTS chat (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id1 INT NOT NULL REFERENCES users(id),
	user_id2 INT NOT NULL REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS message (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    text TEXT NOT NULL,
    sender_id INT NOT NULL REFERENCES users(id),
	receiver_id INT NOT NULL REFERENCES users(id),
	chat_id INT NOT NULL REFERENCES chat(id) ON DELETE CASCADE,
    created_at date NOT NULL
);

CREATE TABLE IF NOT EXISTS like_post (
	user_post_id INT NOT NULL REFERENCES user_posts(id) ON DELETE CASCADE,
	user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
	PRIMARY KEY (user_post_id, user_id)
);
