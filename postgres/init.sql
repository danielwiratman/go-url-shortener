CREATE DATABASE projects
    WITH
    OWNER = postgres
    ENCODING = 'UTF8'
    LC_COLLATE = 'en_US.utf8'
    LC_CTYPE = 'en_US.utf8'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1
    IS_TEMPLATE = False;
    
CREATE SCHEMA IF NOT EXISTS public
    AUTHORIZATION pg_database_owner;

COMMENT ON SCHEMA public
    IS 'standard public schema';

GRANT USAGE ON SCHEMA public TO PUBLIC;

GRANT ALL ON SCHEMA public TO pg_database_owner;

CREATE TABLE IF NOT EXISTS public.go_url_shortener
(
    id integer NOT NULL DEFAULT 'nextval('go_url_shortener_id_seq'::regclass)',
    url text COLLATE pg_catalog."default" NOT NULL,
    short_url text COLLATE pg_catalog."default" NOT NULL,
    created_date timestamp with time zone NOT NULL DEFAULT 'now()',
    CONSTRAINT go_url_shortener_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.go_url_shortener
    OWNER to postgres;