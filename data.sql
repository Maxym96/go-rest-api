create table public.author (
                               id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
                               name varchar(100) NOT NULL
);
create table public.book (
                             id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
                             name varchar(100) NOT NULL,
                             author_id uuid NOT NULL,
                             constraint author_fk FOREIGN KEY (author_id) REFERENCES public.author(id)
);

INSERT INTO author (name) VALUES ('William Shakespeare');
INSERT INTO author (name) VALUES ('Sir Joseph Rudyard Kipling');
INSERT INTO author (name) VALUES ('Oscar Wills Wilde');

insert into book (name, author_id) VALUES ('Портрет Дориана Грея', '3e7c9288-5e82-4804-8e84-71ca81c520fe');
insert into book (name, author_id) VALUES ('Книга джунглей', 'e4bc5ea4-bb50-482a-a471-fcd432144c90');
insert into book (name, author_id) VALUES ('Ромео и Джульетта', 'ef9d9ef2-b6eb-4109-bf0e-340196d47d18');