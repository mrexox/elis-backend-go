use elis_test;
INSERT INTO image (url) VALUES ('/image/lala.png');
INSERT INTO post (name, content, cover, visible, created_at, permalink) VALUES ('Post One', 'TEXT', 1, TRUE, NOW(), 'post-one');
INSERT INTO post (name, content, created_at, permalink) VALUES ('Post Two', 'TEXT', NOW(), 'post-two');
INSERT INTO post (name, content, created_at, permalink) VALUES ('Post Three', 'TEXT', NOW(), 'post-three');
INSERT INTO post (name, content, created_at, permalink) VALUES ('Post Four', 'TEXT', NOW(), 'post-four');
INSERT INTO tag (name) VALUES ('birds');
INSERT INTO tag (name) VALUES ('sheeps');
INSERT INTO tag (name) VALUES ('mountains');
INSERT INTO tag (name) VALUES ('big-clouds');
INSERT INTO post_tag (post_id, tag_id) VALUES (1, 1);
INSERT INTO post_tag (post_id, tag_id) VALUES (1, 2);
INSERT INTO post_tag (post_id, tag_id) VALUES (2, 1);
INSERT INTO post_tag (post_id, tag_id) VALUES (2, 3);
INSERT INTO post_tag (post_id, tag_id) VALUES (2, 4);
INSERT INTO post_tag (post_id, tag_id) VALUES (3, 3);
INSERT INTO post_tag (post_id, tag_id) VALUES (4, 4);
INSERT INTO liker (post_id, ip) VALUES (1, '192.168.1.14');
INSERT INTO liker (post_id, ip) VALUES (1, '192.168.122.4');
INSERT INTO liker (post_id, ip) VALUES (2, '192.168.0.243');
INSERT INTO liker (post_id, ip) VALUES (3, '192.168.9.146');
INSERT INTO message (phone, email, content, name, theme, created_at) VALUES ('+7900121311', 'aa@bb.pu', 'Content', 'Anonimus', 'Theme 1', NOW());
INSERT INTO message (phone, email, content, name, theme, created_at) VALUES ('+7900121311', 'aa@bb.pu', 'Content', 'Anonimus', 'Theme 2', NOW());
INSERT INTO message (phone, email, content, name, theme, created_at) VALUES ('+7900121311', 'aa@bb.pu', 'Content', 'Anonimus', 'Theme 3', NOW());
