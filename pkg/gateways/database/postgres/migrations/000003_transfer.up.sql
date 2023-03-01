BEGIN;

CREATE TABLE IF NOT EXISTS transfer
(
	id serial PRIMARY KEY,
	account_origin_id serial REFERENCES account,
	account_destination_id serial REFERENCES account,
	amount int NOT NULL,
	created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

COMMIT;