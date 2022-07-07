CREATE TABLE articles (
    id bigserial not null primary key,
    title text not null,
    icon_name text not null,
    icon text not null,
    url text not null,
    color text not null,
    wifi boolean not null,
    question boolean default false not null
);