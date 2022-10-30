from faker import Faker
from random import randint 

COUNT_USERS = 30
COUNT_POSTS = 30
COUNT_IMAGES = 30

def gen_users():
    faker = Faker()
    def _gen_users_string():
        first_name = faker.first_name()
        last_name = faker.last_name()
        nick_name = faker.unique.word()
        avatar_img_id = faker.pyint(1, COUNT_IMAGES)
        email = faker.email()
        password = "hash"

        return f"{first_name};{last_name};{nick_name};{avatar_img_id};{email};{password}"

    with open("users.csv", "w") as f: 
        f.write("first_name;last_name;nick_name;avatar_img_id;email;password\n")
        for _ in range(COUNT_USERS):
            f.write(_gen_users_string() + "\n")

def gen_posts():
    def _gen_post_string():
        faker = Faker()
        user_id = randint(1, COUNT_USERS)
        message = str(faker.text()).replace('\n', ' ')
        create_date = faker.date_this_year()
        return f"{user_id};{message};{create_date}"

    with open("user_posts.csv", "w") as f: 
        f.write("user_id;message;create_date\n")
        for _ in range(COUNT_POSTS):
            f.write(_gen_post_string() + "\n")

def gen_images():
    with open("images.csv", "w") as f: 
        f.write("link\n")
        for i in range(COUNT_IMAGES):
            f.write(f"{i + 1}.png" + "\n")

def gen_posts_images_relation():
    relations = []
    def _gen_posts_images_relation_string():
        user_id = randint(1, COUNT_USERS)
        img_id = randint(1, COUNT_IMAGES)
        if (user_id, img_id) not in relations:
            relations.append((user_id, img_id))
            return f"{user_id};{img_id}"
        else:
            return ""

    with open("user_posts_images.csv", "w") as f: 
        f.write("user_post_id;img_id\n")
        for _ in range(COUNT_POSTS * 2):
            row = _gen_posts_images_relation_string()
            f.write(row + "\n" if row != "" else "")

if __name__ == '__main__':
    gen_images()
    gen_posts()
    gen_users()
    gen_posts_images_relation()

