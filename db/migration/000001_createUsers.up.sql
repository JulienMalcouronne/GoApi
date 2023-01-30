CREATE TABLE users
(
  id                     SERIAL PRIMARY KEY,
  email                  character varying NOT NULL UNIQUE,
  password_no_hash       character varying NOT NULL,
  first_name             character varying,
  last_name              character varying,
  created_at             timestamp with time zone DEFAULT now() NOT NULL,
  updated_at             timestamp with time zone DEFAULT now() NOT NULL
);
