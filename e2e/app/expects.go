package app

import "testing"

func ExpectDatabaseHas(t *testing.T, table string, fields map[string]any) {
	t.Helper()

	count, err := databaseRowsCount(t, table, fields)
	if err != nil {
		t.Errorf("ExpectDatabaseHas: %v", err)
	} else if count == 0 {
		t.Errorf("ExpectDatabaseHas: not found row in \"%s\" table", table)
	}
}

func ExpectDatabaseMissing(t *testing.T, table string, fields map[string]any) {
	t.Helper()

	count, err := databaseRowsCount(t, table, fields)
	if err != nil {
		t.Errorf("ExpectDatabaseHas: %v", err)
	} else if count > 0 {
		t.Errorf("ExpectDatabaseHas: found %d rows in \"%s\" table", count, table)
	}
}

func databaseRowsCount(t *testing.T, table string, fields map[string]any) (int, error) {
	t.Helper()

	sql := DBQueryBuilder().Table(table)
	for key, val := range fields {
		sql.Where(key, "=", val)
	}

	count, err := sql.Count()
	return int(count), err
}
