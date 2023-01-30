CREATE TABLE users (
   id BIGSERIAL primary key,
   first_name TEXT not null,
   last_name TEXT,
   email TEXT,
   password TEXT(60),
   created_at TIMESTAMP default now(),
   updated_at TIMESTAMP default now(),
);
