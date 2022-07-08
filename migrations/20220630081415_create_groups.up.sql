CREATE TABLE groups (
    id bigserial not null primary key,
    title text not null,
    icon_name text not null,
    icon text not null,
    color text not null,
    wifi boolean not null,
    model text default('default') not null
);