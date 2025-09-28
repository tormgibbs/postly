CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  whatsapp_id TEXT UNIQUE NOT NULL,
  access_token TEXT,
  refresh_token TEXT,
  token_expiry TIMESTAMPTZ,
  is_logged_in BOOLEAN NOT NULL DEFAULT false,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);
