
CREATE TABLE IF NOT EXISTS cities (
  id   SERIAL PRIMARY KEY,
  name varchar
);

CREATE TABLE IF NOT EXISTS circuits (
  id   SERIAL PRIMARY KEY,
  name varchar
);

CREATE TABLE IF NOT EXISTS national_parks (
  id   SERIAL PRIMARY KEY,
  name varchar
);

CREATE TABLE IF NOT EXISTS locations (
  id   SERIAL PRIMARY KEY,
  name varchar
);

CREATE TABLE IF NOT EXISTS sites (
  id                SERIAL PRIMARY KEY,
  name              varchar,
  has_picture       boolean,
  has_engraving     boolean,
  other             boolean,
  is_historic       boolean,
  year_of_discovery integer,

  latitude          varchar,
  longitude         varchar,

  city_id           integer REFERENCES cities(id),
  circuit_id        integer REFERENCES circuits(id),
  national_park_id  integer REFERENCES national_parks(id),
  location_id       integer REFERENCES locations(id)
);

CREATE EXTENSION unaccent;
