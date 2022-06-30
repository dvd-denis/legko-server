CREATE TABLE articles (
    id bigserial not null primary key,
    title text not null,
    icon_name text not null,
    icon bytea not null,
    color text not null
);