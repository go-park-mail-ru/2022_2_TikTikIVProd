import sys
from faker import Faker
from random import randint 
import datetime


COUNT_USERS = 30
COUNT_CHATS = COUNT_USERS
COUNT_MESSAGES = 100

def gen_messages_for_user(user_id: int):
    def _gen_cur_user_chat(chat_id):
        if chat_id != user_id:
            return f"{user_id};{chat_id}"
        else:
            return ""
    def _gen_user_chat(i: int):
        if i != user_id:
            return f"{i};{i}"
        else:
            return ""
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
        for i in range(COUNT_CHATS):
            row_1 = _gen_cur_user_chat(i + 1)
            row_2 = _gen_user_chat(i + 1)
            f.write(row_1 + "\n" if row_1 != "" else "")
            f.write(row_2 + "\n" if row_2 != "" else "")

    with open("message.csv", "w") as f: 
        f.write("body;sender_id;chat_id;created_at\n")
        for i in range(COUNT_MESSAGES):
            row = _gen_message()
            f.write(row + "\n" if row != "" else "")


if __name__ == '__main__':
    if len(sys.argv) > 1:
        gen_messages_for_user(sys.argv[1])

