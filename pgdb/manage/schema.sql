CREATE TABLE captains (
  uid	bigserial,
  name	text,
  PRIMARY KEY(uid)
);

CREATE TABLE stars (
  x	bigint NOT NULL,
  y	bigint NOT NULL,
  z	bigint NOT NULL,
  name text,
  PRIMARY KEY(x, y, z)
);

CREATE TABLE paths (
  x1	bigint NOT NULL,
  y1	bigint NOT NULL,
  z1	bigint NOT NULL,
  x2	bigint NOT NULL,
  y2	bigint NOT NULL,
  z2	bigint NOT NULL,
  FOREIGN KEY(x1, y1, z1) REFERENCES stars ON DELETE CASCADE,
  FOREIGN KEY(x2, y2, z2) REFERENCES stars ON DELETE CASCADE,
  PRIMARY KEY(x1, y1, z1, x2, y2, z2)
);
