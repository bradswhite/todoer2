CREATE TABLE TODOS (
  ID UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  Username TEXT NOT NULL,
  Text TEXT NOT NULL,
  Completed BOOLEAN NOT NULL DEFAULT FALSE,
  Timestamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO TODOS (USERNAME, TEXT) VALUES
  ('Brad', 'First Todo'),
  ('Brad', 'Second Todo'),
  ('Brad', 'Third Todo');

INSERT INTO TODOS VALUES
  (DEFAULT, 'Brad', 'First Todo', DEFAULT, DEFAULT),
  (DEFAULT, 'Brad', 'Second Todo', DEFAULT, DEFAULT),
  (DEFAULT, 'Brad', 'Third Todo', DEFAULT, DEFAULT);
