CREATE TABLE captains (
  uid	bigserial,
  name	text,
  PRIMARY KEY(uid)
);

CREATE TABLE stars (
  x	numeric NOT NULL,
  y	numeric NOT NULL,
  z	numeric NOT NULL,
  name text,
  PRIMARY KEY(x, y, z)
);

CREATE TABLE paths (
  x1	numeric NOT NULL,
  y1	numeric NOT NULL,
  z1	numeric NOT NULL,
  x2	numeric NOT NULL,
  y2	numeric NOT NULL,
  z2	numeric NOT NULL,
  FOREIGN KEY(x1, y1, z1) REFERENCES stars ON DELETE CASCADE,
  FOREIGN KEY(x2, y2, z2) REFERENCES stars ON DELETE CASCADE,
  PRIMARY KEY(x1, y1, z1, x2, y2, z2)
);
