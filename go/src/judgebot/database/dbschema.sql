CREATE USER judgebot WITH PASSWORD 'judgebot';

CREATE DATABASE judgebot OWNER judgebot;

\connect judgebot judgebot;

CREATE SCHEMA judgebot AUTHORIZATION judgebot;

CREATE TABLE judgebot.users (
    id serial NOT NULL,
    telegram_id integer UNIQUE NOT NULL,

    CONSTRAINT users_pkey PRIMARY KEY (id)
);

CREATE TABLE judgebot.judge_phrases (
    id serial NOT NULL,
    phrase text UNIQUE,

    CONSTRAINT judge_phrases_pkey PRIMARY KEY (id)
);

CREATE TABLE judgebot.votes (
    user_id integer NOT NULL,
    judge_phrase_id integer NOT NULL,

    CONSTRAINT vote_pkey PRIMARY KEY (user_id, judge_phrase_id),
    CONSTRAINT user_id_fkey FOREIGN KEY (user_id) REFERENCES judgebot.users (id),
    CONSTRAINT judge_phrase_id_fkey FOREIGN KEY (judge_phrase_id) REFERENCES judgebot.judge_phrases (id)
);
