CREATE TABLE IF NOT EXISTS todos (
  id UUID PRIMARY KEY,
  text TEXT,
  done BOOLEAN,
  user_id UUID,
);
