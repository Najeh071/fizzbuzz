package db

import (
	"context"
	"database/sql"
	"fmt"

	// Enregistre silencieusement le driver SQLite dans database/sql
	_ "modernc.org/sqlite"

	"fizz-buzz/pkg/fizzbuzz"
)

type DB struct {
	conn *sql.DB
}

type StatRecord struct {
	Parameters fizzbuzz.Request `json:"parameters"`
	Hits       int64            `json:"hits"`
}

//initialisation de base de donnée
func InitDB(dbPath string) (*DB, error) {
	conn, err := sql.Open("sqlite", dbPath+"?_pragma=busy_timeout=5000&_pragma=journal_mode=WAL")
	if err != nil {
		return nil, fmt.Errorf("failed to open sqlite database: %w", err)
	}

	db := &DB{conn: conn}
	if err := db.migrate(); err != nil {
		return nil, err
	}

	return db, nil
}


func (d *DB) migrate() error {
	query := `
	CREATE TABLE IF NOT EXISTS request_stats (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		int1 INTEGER NOT NULL,
		int2 INTEGER NOT NULL,
		fizz_limit INTEGER NOT NULL,
		str1 TEXT NOT NULL,
		str2 TEXT NOT NULL,
		hits INTEGER NOT NULL DEFAULT 1,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		UNIQUE(int1, int2, fizz_limit, str1, str2)
	);`

	_, err := d.conn.Exec(query)
	if err != nil {
		return fmt.Errorf("migration failed: %w", err)
	}
	return nil
}

func (d *DB) TrackRequest(ctx context.Context, req fizzbuzz.Request) error {
	query := `
	INSERT INTO request_stats (int1, int2, fizz_limit, str1, str2, hits, updated_at)
	VALUES (?, ?, ?, ?, ?, 1, CURRENT_TIMESTAMP)
	ON CONFLICT(int1, int2, fizz_limit, str1, str2) DO UPDATE SET
		hits = hits + 1,
		updated_at = CURRENT_TIMESTAMP;
	`
	_, err := d.conn.ExecContext(ctx, query, req.Int1, req.Int2, req.Limit, req.Str1, req.Str2)
	if err != nil {
		return fmt.Errorf("failed to track request in sqlite: %w", err)
	}
	return nil
}

// fonction pour l'utilisation plus tard pour les statestiques.
func (d *DB) GetMostFrequent(ctx context.Context) (*StatRecord, error) {
	query := `
	SELECT int1, int2, fizz_limit, str1, str2, hits
	FROM request_stats
	ORDER BY hits DESC
	LIMIT 1;
	`

	var stat StatRecord
	err := d.conn.QueryRowContext(ctx, query).Scan(
		&stat.Parameters.Int1,
		&stat.Parameters.Int2,
		&stat.Parameters.Limit,
		&stat.Parameters.Str1,
		&stat.Parameters.Str2,
		&stat.Hits,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to fetch stats: %w", err)
	}

	return &stat, nil
}

func (d *DB) Close() error {
	return d.conn.Close()
}