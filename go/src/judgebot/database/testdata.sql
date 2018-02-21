\connect judgebot judgebot;

INSERT INTO judgebot.users (id, telegram_id) VALUES (1, 1);
INSERT INTO judgebot.users (id, telegram_id) VALUES (2, 2);
INSERT INTO judgebot.users (id, telegram_id) VALUES (3, 3);
INSERT INTO judgebot.users (id, telegram_id) VALUES (4, 4);
INSERT INTO judgebot.users (id, telegram_id) VALUES (5, 5);

INSERT INTO judgebot.judge_phrases (id, phrase) VALUES (1, 'aaa');
INSERT INTO judgebot.judge_phrases (id, phrase) VALUES (2, 'bbb');
INSERT INTO judgebot.judge_phrases (id, phrase) VALUES (3, 'ccc');
INSERT INTO judgebot.judge_phrases (id, phrase) VALUES (4, 'ddd');
INSERT INTO judgebot.judge_phrases (id, phrase) VALUES (5, 'eee');

INSERT INTO judgebot.votes (user_id, judge_phrase_id) VALUES (1, 1);
INSERT INTO judgebot.votes (user_id, judge_phrase_id) VALUES (2, 1);
INSERT INTO judgebot.votes (user_id, judge_phrase_id) VALUES (3, 1);
INSERT INTO judgebot.votes (user_id, judge_phrase_id) VALUES (4, 1);
INSERT INTO judgebot.votes (user_id, judge_phrase_id) VALUES (5, 1);
INSERT INTO judgebot.votes (user_id, judge_phrase_id) VALUES (1, 2);
INSERT INTO judgebot.votes (user_id, judge_phrase_id) VALUES (2, 2);
INSERT INTO judgebot.votes (user_id, judge_phrase_id) VALUES (3, 2);
INSERT INTO judgebot.votes (user_id, judge_phrase_id) VALUES (4, 2);
INSERT INTO judgebot.votes (user_id, judge_phrase_id) VALUES (1, 4);
INSERT INTO judgebot.votes (user_id, judge_phrase_id) VALUES (3, 4);
INSERT INTO judgebot.votes (user_id, judge_phrase_id) VALUES (1, 5);
INSERT INTO judgebot.votes (user_id, judge_phrase_id) VALUES (5, 5);
