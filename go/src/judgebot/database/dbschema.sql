CREATE DATABASE judgebot;

CREATE TABLE users (
    id serial NOT NULL,
    telegram_id integer UNIQUE NOT NULL,

    CONSTRAINT users_pkey PRIMARY KEY (id)
);

CREATE TABLE judge_phrases (
    id serial NOT NULL,
    phrase text UNIQUE,

    CONSTRAINT judge_phrases_pkey PRIMARY KEY (id)
);

CREATE TABLE votes (
    user_id integer NOT NULL,
    judge_phrase_id integer NOT NULL,

    CONSTRAINT vote_pkey PRIMARY KEY (user_id, judge_phrase_id),
    CONSTRAINT user_id_fkey FOREIGN KEY (user_id) REFERENCES users (id),
    CONSTRAINT judge_phrase_id_fkey FOREIGN KEY (judge_phrase_id) REFERENCES judge_phrases (id)
);
