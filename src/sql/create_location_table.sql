CREATE TABLE locations (
	id        serial PRIMARY KEY,
	ds_id     integer REFERENCES dataset (id),
	lat       float,
	lon       float
);
