package postgres

import (
	"net/url"
	"sync"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // The database driver in use.
)

var once sync.Once
var instance *sqlx.DB

func GetInstance() (*sqlx.DB, error) {
	if instance != nil {
		return instance, nil
	}

	var err error

	once.Do(func() {
		config := NewConfig()

		sslMode := "require"
		if config.DisableTLS {
			sslMode = "disable"
		}

		q := make(url.Values)
		q.Set("sslmode", sslMode)
		q.Set("timezone", "utc")

		u := url.URL{
			Scheme:   "postgres",
			User:     url.UserPassword(config.Username, config.Password),
			Host:     config.Host,
			Path:     config.Database,
			RawQuery: q.Encode(),
		}

		instance, err = sqlx.Open("postgres", u.String())
	})

	return instance, err
}
