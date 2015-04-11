CREATE TABLE locations (
	id        serial PRIMARY KEY,
	ds_id     integer REFERENCES dataset (id),
	target    GEOGRAPHY(POINT, 4326)
);
