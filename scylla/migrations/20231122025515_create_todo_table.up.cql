CREATE TABLE IF NOT EXISTS todos (
  id UUID,
  text TEXT,
  done BOOLEAN,
  user_id UUID,
  PRIMARY KEY ((id, user_id), done)
) WITH compaction = { 'class' : 'LeveledCompactionStrategy' };
