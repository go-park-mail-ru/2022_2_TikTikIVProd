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
	passhash VARCHAR(128) NOT NULL
);

CREATE TABLE IF NOT EXISTS cookies (
	value varchar(64) PRIMARY KEY,
	user_id INT REFERENCES users(id),
	expires DATE
);


CREATE TABLE IF NOT EXISTS user_posts (
	id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	user_id INT REFERENCES users(id),
	message TEXT,
	create_date DATE NOT NULL
);


CREATE TABLE IF NOT EXISTS user_posts_images (
	user_post_id INT REFERENCES user_posts(id),
	img_id INT REFERENCES images(id),
	PRIMARY KEY (user_post_id, img_id)
);