package database

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/wfrs-dev/hefesto/app/config"
)

func New(d config.Database) (*pgxpool.Pool, error) {
	url := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		d.User,
		d.Password,
		d.Host,
		d.Port,
		d.Database,
		d.SSLMode,
	)

	pool, err := pgxpool.New(context.Background(), url)
	if err != nil {
		return nil, err
	}

	// Control de conexion
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = pool.Ping(ctx)
	if err != nil {
		return nil, err
	}

	return pool, nil
}

// QueryAsMaps Ejecuta una consulta y devuelve los resultados como mapas
func QueryAsMaps(ctx context.Context, pool *pgxpool.Pool, query string) ([]map[string]any, error) {
	rows, err := pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	fields := rows.FieldDescriptions()

	result := make([]map[string]any, 0)

	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			return nil, err
		}

		row := make(map[string]any, len(fields))

		for i, field := range fields {
			row[string(field.Name)] = values[i]
		}

		result = append(result, row)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
