CREATE TABLE IF NOT EXISTS client.session (
  id uuid PRIMARY KEY,
  sub uuid NOT NULL,
  refresh varchar NOT NULL,
  agent varchar NOT NULL,
  ip varchar NOT NULL,
  blocked boolean NOT NULL DEFAULT false,
  expires timestamptz NOT NULL,
  created timestamptz DEFAULT Now() NOT NULL 
);

ALTER TABLE client.session ADD FOREIGN KEY (sub) REFERENCES client.profile (id);
