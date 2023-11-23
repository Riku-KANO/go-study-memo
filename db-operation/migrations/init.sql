DROP TABLE IF EXISTS users CASCADE;
CREATE TABLE users (
   id bigserial PRIMARY KEY,
   created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
   name TEXT NOT NULL,
   updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
);