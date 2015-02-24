CREATE TABLE locations (
	id        integer PRIMARY KEY DEFAULT nextval('serial'),
	target    GEOGRAPHY(POINT, 4326)
);
