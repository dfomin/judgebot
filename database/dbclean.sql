DELETE FROM judgebot.votes;
DELETE FROM judgebot.judge_phrases;
DELETE FROM judgebot.chat_users;

DROP TABLE judgebot.votes;
DROP TABLE judgebot.judge_phrases;
DROP TABLE judgebot.chat_users;

DROP SCHEMA IF EXISTS judgebot;

\DISCONNECT judgebot;

DROP DATABASE IF EXISTS judgebot;
DROP USER IF EXISTS judgebot;
