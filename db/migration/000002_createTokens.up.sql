CREATE TABLE tokens
(
  id                     SERIAL PRIMARY KEY,
  user_id                integer REFERENCES users (id),
  email                  character varying NOT NULL UNIQUE,
  token                  character varying,
  token_hash             BYTEA,
  created_at             timestamp with time zone DEFAULT now() NOT NULL,
  updated_at             timestamp with time zone DEFAULT now() NOT NULL,
  expiry                 timestamp without time zone
);
