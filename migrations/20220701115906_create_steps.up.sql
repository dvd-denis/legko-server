CREATE TABLE steps (
    id bigserial not null primary key,
    article_id bigserial REFERENCES articles (id) on delete CASCADE NOT NULL,
    title text not null,
    content text not null,
    num int not null,
    wifi boolean not null,
    question boolean default false not null
);