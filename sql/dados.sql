insert into users (name, nick, email, password) 
values 
("Bidi", "bibi", "bibi@gmail.com", "$2a$10$1jhubBHlbh0Ai4CE5kdwTOrX1B6IVsNc6gCV5Xgm2nnrTAlWKjmnK"),
("Vanderleia", "wanda", "vanderleia@gmail.com", "$2a$10$1jhubBHlbh0Ai4CE5kdwTOrX1B6IVsNc6gCV5Xgm2nnrTAlWKjmnK"),
("Denis", "din", "din@gmail.com", "$2a$10$1jhubBHlbh0Ai4CE5kdwTOrX1B6IVsNc6gCV5Xgm2nnrTAlWKjmnK"),
("Aryosmar", "ary", "ary@gmail.com", "$2a$10$1jhubBHlbh0Ai4CE5kdwTOrX1B6IVsNc6gCV5Xgm2nnrTAlWKjmnK");

insert into followers (user_id, follower_id)
values
(1, 2),
(2, 1),
(3, 5),
(3, 4),
(2, 3),
(1, 3),
(5, 4),
(4, 1);