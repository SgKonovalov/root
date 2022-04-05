CREATE TABLE advertisement (
	adv_id serial NOT NULL,
	adv_name varchar,
	price bigint,
	description text,
	pub_date timestamp,
	PRIMARY KEY (adv_id)
);

CREATE TABLE photos (
	photo_id serial NOT NULL,
	adv_id integer NOT NULL,
	photo varchar,
	main_photo boolean,
	PRIMARY KEY (photo_id),
	FOREIGN KEY (adv_id)
	REFERENCES advertisement (adv_id)
	ON DELETE CASCADE
	ON UPDATE CASCADE
);