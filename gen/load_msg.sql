COPY public.chat (name, created_at) FROM '/home/gen/chat.csv' DELIMITER ';' CSV HEADER;
COPY public.message (body, sender_id, chat_id, created_at) FROM '/home/gen/message.csv' DELIMITER ';' CSV HEADER;
COPY public.user_chat (user_id, chat_id) FROM '/home/gen/user_chat.csv' DELIMITER ';' CSV HEADER;
