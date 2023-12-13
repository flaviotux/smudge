CREATE TABLE IF NOT EXISTS users (
  id UUID PRIMARY KEY,
  name TEXT,
) WITH compaction = { 'class' : 'LeveledCompactionStrategy' };
