package main

import (
	"database/sql"
)

// Up is executed when this migration is applied
func Up_20140816162235(txn *sql.Tx) {
	txn.Exec("CREATE EXTENSION \"uuid-ossp\";")

	txn.Exec(`CREATE TABLE entities (
	  id uuid primary key default uuid_generate_v4()
	);`)

	txn.Exec(`CREATE TABLE positions (
	  id uuid references entities(id) ON DELETE CASCADE,
	  z integer NOT NULL,
	  x integer NOT NULL,
	  y integer NOT NULL,
	  cx integer NOT NULL,
	  cz integer NOT NULL,
	  cy integer NOT NULL
	);`)

	txn.Exec(`CREATE INDEX positions_id_fkey_idx ON positions (id);`)
	txn.Exec(`CREATE INDEX positions_chunk_idx ON positions (cz, cx, cy);`)

	txn.Exec(`CREATE TABLE materials (
	  id uuid references entities(id) ON DELETE CASCADE,
	  material_type integer NOT NULL
	);`)

	txn.Exec(`CREATE INDEX materials_id_fkey_idx ON materials (id);`)

	txn.Exec(`CREATE TABLE brains (
	  id uuid references entities(id) ON DELETE CASCADE,
	  strategy integer NOT NULL
	);`)

	txn.Exec(`CREATE INDEX brains_id_fkey_idx ON brains (id);`)

	txn.Exec(`CREATE TABLE controlled (
	  id uuid references entities(id) ON DELETE CASCADE
	);`)

	txn.Exec(`CREATE INDEX controlled_id_fkey_idx ON controlled (id);`)
}

// Down is executed when this migration is rolled back
func Down_20140816162235(txn *sql.Tx) {
	txn.Exec("DROP TABLE controlled;")
	txn.Exec("DROP TABLE brains;")
	txn.Exec("DROP TABLE positions;")
	txn.Exec("DROP TABLE materials;")
	txn.Exec("DROP TABLE entities;")
	txn.Exec("DROP EXTENSION \"uuid-ossp\";")
}
