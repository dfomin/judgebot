CREATE USER judgebot WITH PASSWORD 'judgebot';

CREATE DATABASE judgebot OWNER judgebot;

\connect judgebot judgebot;

CREATE SCHEMA judgebot AUTHORIZATION judgebot;

CREATE TABLE judgebot.chat_users (
    id SERIAL NOT NULL,
    user_id INTEGER NOT NULL,
    chat_id BIGINT NOT NULL,

    CONSTRAINT users_pkey PRIMARY KEY (id),
    CONSTRAINT chat_users_unique UNIQUE (user_id, chat_id)
);

CREATE TABLE judgebot.judge_phrases (
    id SERIAL NOT NULL,
    phrase TEXT UNIQUE,

    CONSTRAINT judge_phrases_pkey PRIMARY KEY (id)
);

CREATE TABLE judgebot.votes (
    vote BOOLEAN NOT NULL,
    chat_user_id INTEGER NOT NULL,
    judge_phrase_id INTEGER NOT NULL,

    CONSTRAINT vote_pkey PRIMARY KEY (chat_user_id, judge_phrase_id),
    CONSTRAINT user_id_fkey FOREIGN KEY (chat_user_id) REFERENCES judgebot.chat_users (id),
    CONSTRAINT judge_phrase_id_fkey FOREIGN KEY (judge_phrase_id) REFERENCES judgebot.judge_phrases (id)
);

GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA judgebot TO judgebot;
GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA judgebot TO judgebot;