create table public.event (
                              id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
                              name varchar(100) NOT NULL,
                              description varchar(255) NOT NULL,
                              date_and_time timestamp NOT NULL
);

create table public.event_author (
                                     id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
                                     name varchar(100) NOT NULL,
                                     author_id uuid NOT NULL,
                                     constraint author_fk FOREIGN KEY (author_id) REFERENCES public.event(id)
);

INSERT INTO event (name, description, date_and_time)
VALUES ('Hamlet',
        'Hamlet is not only one of Shakespeares finest plays, but also one of the most gripping tragedies in world literature.',
        '2022-06-08 12:05:00');
INSERT INTO event (name, description, date_and_time)
VALUES ('JULIUS CAESAR',
        'В этой знаменитой трагедии Шекспир обращается к истории Римской империи и описывает весьма известное преступление против Юлия Цезаря.',
        '2022-06-08 15:05:00');
INSERT INTO event (name, description, date_and_time)
VALUES ('LUCES DE BOHEMIA',
        'La obra narra las últimas horas de la vida de Max Estrella.',
        '2022-06-08 18:05:00');
INSERT INTO event (name, description, date_and_time)
VALUES ('Великая Екатерина',
        'Пьеса описывает приключения чопорного английского джентльмена при дворе любвеобильной российской императрицы Екатерины II.',
        '2022-06-08 21:05:00');
INSERT INTO event (name, description, date_and_time)
VALUES ('Индийская тушь',
        'Пьеса британского драматурга Тома Стоппарда, написанная в 1995 году.',
        '2022-06-08 23:00:00');

insert into event_author (name, author_id) VALUES ('William Shakespeare', '4390f70e-4657-45d4-b56d-f1c8b4d49b29');
insert into event_author (name, author_id) VALUES ('William Shakespeare', '292ef6a0-e605-42ba-bf06-c6f1f4666128');
insert into event_author (name, author_id) VALUES ('Ramon Maria Del Valle-Inclan', '2e5af243-13f0-4e5a-a7f0-4d6f4030488a');
insert into event_author (name, author_id) VALUES ('Бернард Шоу', '3492cad5-8a38-4ad5-85ea-bc8dd88f58d7');
insert into event_author (name, author_id) VALUES ('Том Стоппард', '8cfde0c8-fefe-4b09-b67b-cf5eb1117d23');