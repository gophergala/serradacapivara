CREATE TABLE sites (
  `id`                integer PRIMARY KEY,
  `name`              varchar,
  `has_picture`       boolean,
  `other`             boolean,
  `is_historic`       boolean,
  `year_of_discovery` integer,

  `city_id`          integer REFERENCES cities(id),
  `circuit_id`       integer REFERENCES circuits(id),
  `national_park_id` integer REFERENCES national_parks(id),
  `location_id`      integer REFERENCES locations(id)
);

CREATE TABLE cities (
  `id`   integer PRIMARY KEY,
  `name` varchar
);

CREATE TABLE circuits (
  `id`   integer PRIMARY KEY,
  `name` varchar
);

CREATE TABLE national_parks (
  `id`   integer PRIMARY KEY,
  `name` varchar
);

CREATE TABLE locations (
  `id`   integer PRIMARY KEY,
  `name` varchar
);
