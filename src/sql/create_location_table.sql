CREATE TABLE locations (
	id        integer PRIMARY KEY DEFAULT nextval('serial'),
	ds_id     integer REFERENCES dataset (id),
	target    GEOGRAPHY(POINT, 4326)
);
