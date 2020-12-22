CREATE TABLE IF NOT EXISTS gadgets(
  gadget_uuid uuid PRIMARY KEY,
  user_id int NOT NULL,
  enabled BOOLEAN NOT NULL,
  FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);
