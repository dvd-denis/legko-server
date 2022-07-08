CREATE TABLE articles (
    id bigserial not null primary key,
    title text not null,
    tags text not null,
    group_id bigserial REFERENCES groups (id) on delete CASCADE NOT NULL
);