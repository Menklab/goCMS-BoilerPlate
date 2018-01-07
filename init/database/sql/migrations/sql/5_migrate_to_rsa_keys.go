package migrations

import "github.com/rubenv/sql-migrate"

func MigrateToRSAKeys() *migrate.Migration {
	migrateToRSAKeys := migrate.Migration{
		Id: "5",
		Up: []string{`
			DELETE FROM gocms_settings WHERE name = 'AUTHENTICATION_KEY';
			`, `
			INSERT INTO gocms_settings (name, value, description) VALUES('RSA_PRIV', '', 'RSA private key used for authentication');
			`, `
			INSERT INTO gocms_settings (name, value, description) VALUES('RSA_PUB', '', 'RSA public key used for authentication');
			`,`
			ALTER TABLE gocms_settings MODIFY value TEXT NOT NULL;
			`,
		},
		Down: []string{
		},
	}

	return &migrateToRSAKeys
}
