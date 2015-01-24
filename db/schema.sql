CREATE TABLE IF NOT EXISTS Sites (
  Id              SERIAL PRIMARY KEY,
  Name            varchar,
  HasPicture      boolean,
  HasEngraving    boolean,
  Other           boolean,
  IsHistoric      boolean,
  YearOfDiscovery integer,
  Latitude        varchar,
  Longitude       varchar,
  City            varchar,
  Circuit         varchar,
  NationalPark    varchar,
  Location        varchar
);

CREATE EXTENSION unaccent;
