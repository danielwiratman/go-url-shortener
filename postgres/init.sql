create table if not exists go_url_shortener (
    id serial not null,
    url text,
    short_url text,
    created_date timestamp with time zone not null default now(),
    primary key (id)
);