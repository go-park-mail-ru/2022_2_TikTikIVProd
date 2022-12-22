CREATE TABLE IF NOT EXISTS attachments (
	id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	att_link VARCHAR(260) NOT NULL,
	ttype int NOT NULL
);

CREATE TABLE IF NOT EXISTS stickers (
	id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	link VARCHAR(260) NOT NULL
);

CREATE TABLE IF NOT EXISTS users (
	id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	first_name VARCHAR(35) NOT NULL,
	last_name VARCHAR(35) NOT NULL,
	nick_name VARCHAR(30) NOT NULL UNIQUE,
	avatar_att_id INT REFERENCES attachments(id),
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
    avatar_att_id INT REFERENCES attachments(id),
    name VARCHAR(30) NOT NULL,
    description TEXT DEFAULT '',
    created_at date NOT NULL
);

CREATE TABLE IF NOT EXISTS user_posts (
	id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
	community_id INT REFERENCES communities(id),
	description TEXT NOT NULL DEFAULT '',
	created_at TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS comments (
	id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	user_id INT NOT NULL REFERENCES users(id),
	post_id INT NOT NULL REFERENCES user_posts(id) ON DELETE CASCADE,
	text TEXT NOT NULL,
	created_at TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS user_posts_attachments (
	user_post_id INT NOT NULL REFERENCES user_posts(id) ON DELETE CASCADE,
	att_id INT NOT NULL REFERENCES attachments(id) ON DELETE CASCADE,
	PRIMARY KEY (user_post_id, att_id)
);

CREATE TABLE IF NOT EXISTS chat (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id1 INT NOT NULL REFERENCES users(id),
	user_id2 INT NOT NULL REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS message (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    text TEXT,
    sender_id INT NOT NULL REFERENCES users(id),
	receiver_id INT NOT NULL REFERENCES users(id),
	chat_id INT NOT NULL REFERENCES chat(id) ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL,
	sticker_id int REFERENCES stickers(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS like_post (
	user_post_id INT NOT NULL REFERENCES user_posts(id) ON DELETE CASCADE,
	user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
	PRIMARY KEY (user_post_id, user_id)
);

CREATE TABLE IF NOT EXISTS message_attachments (
	message_id INT NOT NULL REFERENCES message(id) ON DELETE CASCADE,
	att_id INT NOT NULL REFERENCES attachments(id) ON DELETE CASCADE,
	PRIMARY KEY (message_id, att_id)
);
