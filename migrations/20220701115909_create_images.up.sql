CREATE TABLE images (
    id bigserial not null primary key,
    step_id bigserial REFERENCES steps (id) on delete CASCADE NOT NULL,
    image_name text not null,
    image text not null
);