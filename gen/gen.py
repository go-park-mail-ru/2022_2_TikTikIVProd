import sys
from faker import Faker
from random import randint 
import datetime


COUNT_USERS = 30
COUNT_POSTS = 20
COUNT_IMAGES = 30
COUNT_FILES = 10
COUNT_COMMUNITIES = 30

def gen_users():
    faker = Faker()
    def _gen_users_string():
        first_name = faker.first_name()
        last_name = faker.last_name()
        nick_name = faker.unique.word()
        avatar_att_id = faker.pyint(1, COUNT_IMAGES)
        email = faker.email()
        password = "hash"
        created_at = datetime.datetime.now()

        return f"{first_name};{last_name};{nick_name};{avatar_att_id};{email};{password};{created_at}"

    with open("users.csv", "w") as f: 
        f.write("first_name;last_name;nick_name;avatar_att_id;email;password;created_at\n")
        for _ in range(COUNT_USERS):
            f.write(_gen_users_string() + "\n")

def gen_posts():
    date = Faker().date_this_year()
    def _gen_post_string():
        faker = Faker()
        user_id = randint(1, COUNT_USERS)
        description = str(faker.text()).replace('\n', ' ')
        created_at = date
        return f"{user_id};;{description};{created_at}"
    def _gen_post_string_communities():
        faker = Faker()
        user_id = randint(1, COUNT_USERS)
        community_id = randint(1, COUNT_COMMUNITIES)
        description = str(faker.text()).replace('\n', ' ')
        created_at = date
        return f"{user_id};{community_id};{description};{created_at}"

    with open("user_posts.csv", "w") as f: 
        f.write("user_id;community_id;description;created_at\n")
        for i in range(COUNT_POSTS):
            if i % 2:
                f.write(_gen_post_string() + "\n")
            else:
                f.write(_gen_post_string_communities() + "\n")

def gen_attachments():
    with open("attachments.csv", "w") as f: 
        f.write("link;type\n")
        for i in range(COUNT_IMAGES):
            f.write(f"{i + 1}.png;0" + "\n")

def gen_files():
    with open("files.csv", "w") as f: 
        f.write("link\n")
        for i in range(COUNT_IMAGES):
            f.write(f"{i + 1}.html" + "\n")


def gen_stickers():
    with open("stickers.csv", "w") as f: 
        f.write("link\n")
        for i in range(10):
            f.write(f"{i + 1}.png" + "\n")

def gen_posts_attachments_relation():
    relations = []
    def _gen_posts_attachments_relation_string():
        post_id = randint(1, COUNT_POSTS)
        att_id = randint(1, COUNT_IMAGES)
        if (post_id, att_id) not in relations:
            relations.append((post_id, att_id))
            return f"{post_id};{att_id}"
        else:
            return ""

    with open("user_posts_attachments.csv", "w") as f: 
        f.write("user_post_id;att_id\n")
        for _ in range(COUNT_POSTS * 2):
            row = _gen_posts_attachments_relation_string()
            f.write(row + "\n" if row != "" else "")

def gen_posts_files_relation():
    relations = []
    def _gen_posts_files_relation_string():
        post_id = randint(1, COUNT_POSTS)
        file_id = randint(1, COUNT_FILES)
        if (post_id, file_id) not in relations:
            relations.append((post_id, file_id))
            return f"{post_id};{file_id}"
        else:
            return ""

    with open("user_posts_files.csv", "w") as f: 
        f.write("user_post_id;file_id\n")
        for _ in range(COUNT_POSTS // 2):
            row = _gen_posts_files_relation_string()
            f.write(row + "\n" if row != "" else "")

COUNT_CHATS = COUNT_USERS
COUNT_MESSAGES = 100

def gen_messages_for_user(user_id: int):
    def _gen_cur_user_chat(chat_id):
        return f"{user_id};{chat_id}"
    def _gen_user_chat(i: int):
        return f"{i};{i}"
    def _gen_message():
        faker = Faker()
        body = str(faker.text()).replace('\n', ' ')
        sender_id = randint(1, COUNT_USERS)
        chat_id = randint(1, COUNT_CHATS)
        created_at = datetime.datetime.now()
        return f"{body};{sender_id};{chat_id};{created_at}"
    def _gen_chats():
        created_at = datetime.datetime.now()
        return f";{created_at}"

    with open("chat.csv", "w") as f: 
        f.write("name;created_at\n")
        for _ in range(COUNT_CHATS):
            row = _gen_chats()
            f.write(row + "\n" if row != "" else "")

    with open("user_chat.csv", "w") as f: 
        f.write("user_id;chat_id\n")
        for i in range(COUNT_MESSAGES):
            row_1 = _gen_cur_user_chat(i + 1)
            row_2 = _gen_user_chat(i + 1)
            f.write(row_1 + "\n" if row_1 != "" else "")
            f.write(row_2 + "\n" if row_2 != "" else "")

    with open("message.csv", "w") as f: 
        f.write("body;sender_id;chat_id;created_at\n")
        for i in range(COUNT_MESSAGES):
            row = _gen_message()
            f.write(row + "\n" if row != "" else "")

def gen_communities():
    faker = Faker()
    def _gen_communities_string():
        name = faker.first_name()
        owner_id = randint(1, COUNT_USERS)
        avatar_att_id = faker.pyint(1, COUNT_IMAGES)
        description = str(faker.text()).replace('\n', ' ')
        created_at = datetime.datetime.now()

        return f"{name};{owner_id};{avatar_att_id};{description};{created_at}"

    with open("communities.csv", "w") as f: 
        f.write("name;owner_id;nickavatar_att_id_name;description;created_at\n")
        for _ in range(COUNT_COMMUNITIES):
            f.write(_gen_communities_string() + "\n")

def gen_likes():
    relations = []
    def _gen_likes_string():
        user_post_id = randint(1, COUNT_POSTS)
        user_id = randint(1, COUNT_USERS)
        if (user_post_id, user_id) not in relations:
            relations.append((user_post_id, user_id))
            return f"{user_post_id};{user_id}"
        else:
            return ""

    with open("like_post.csv", "w") as f: 
        f.write("user_post_id;user_id\n")
        for _ in range(COUNT_POSTS * 100):
            row = _gen_likes_string()
            f.write(row + "\n" if row != "" else "")

if __name__ == '__main__':
    gen_attachments()
    gen_files()
    gen_stickers()
    gen_posts()
    gen_users()
    gen_posts_attachments_relation()
    gen_posts_files_relation()
    gen_communities()
    gen_likes()
    
    if len(sys.argv) > 1:
        gen_messages_for_user(sys.argv[1])

