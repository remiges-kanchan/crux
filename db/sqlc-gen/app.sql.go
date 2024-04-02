// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: app.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const appDelete = `-- name: AppDelete :exec
DELETE FROM app
WHERE shortnamelc = $1 AND realm = $2
`

type AppDeleteParams struct {
	Shortnamelc string `json:"shortnamelc"`
	Realm       string `json:"realm"`
}

func (q *Queries) AppDelete(ctx context.Context, arg AppDeleteParams) error {
	_, err := q.db.Exec(ctx, appDelete, arg.Shortnamelc, arg.Realm)
	return err
}

const appExist = `-- name: AppExist :one
SELECT
    CASE
        WHEN EXISTS (SELECT 1 FROM schema WHERE schema.app = $1) OR
             EXISTS (SELECT 1 FROM ruleset WHERE ruleset.app = $1)
        THEN 1
        ELSE 0
    END AS value_exists
`

func (q *Queries) AppExist(ctx context.Context, app string) (int32, error) {
	row := q.db.QueryRow(ctx, appExist, app)
	var value_exists int32
	err := row.Scan(&value_exists)
	return value_exists, err
}

const appNew = `-- name: AppNew :many
INSERT INTO
    app (
        realm, shortname, shortnamelc, longname, setby
    )
VALUES (
        $1, $2, $3, $4, $5
    )
RETURNING
    id, realm, shortname, shortnamelc, longname, setby, setat
`

type AppNewParams struct {
	Realm       string `json:"realm"`
	Shortname   string `json:"shortname"`
	Shortnamelc string `json:"shortnamelc"`
	Longname    string `json:"longname"`
	Setby       string `json:"setby"`
}

func (q *Queries) AppNew(ctx context.Context, arg AppNewParams) ([]App, error) {
	rows, err := q.db.Query(ctx, appNew,
		arg.Realm,
		arg.Shortname,
		arg.Shortnamelc,
		arg.Longname,
		arg.Setby,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []App
	for rows.Next() {
		var i App
		if err := rows.Scan(
			&i.ID,
			&i.Realm,
			&i.Shortname,
			&i.Shortnamelc,
			&i.Longname,
			&i.Setby,
			&i.Setat,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const appUpdate = `-- name: AppUpdate :exec
UPDATE app
set
    longname = $1,
    setat = NOW(),
    setby = $2
WHERE
    shortnamelc = $3
    AND realm = $4
`

type AppUpdateParams struct {
	Longname    string `json:"longname"`
	Setby       string `json:"setby"`
	Shortnamelc string `json:"shortnamelc"`
	Realm       string `json:"realm"`
}

func (q *Queries) AppUpdate(ctx context.Context, arg AppUpdateParams) error {
	_, err := q.db.Exec(ctx, appUpdate,
		arg.Longname,
		arg.Setby,
		arg.Shortnamelc,
		arg.Realm,
	)
	return err
}

const getAppList = `-- name: GetAppList :many
SELECT
    a.shortnamelc AS name,
    a.longname AS descr,
    a.setat AS createdat,
    a.setby AS createdby,
    ( SELECT COUNT(DISTINCT "user")
        FROM capgrant
        WHERE app = a.shortnamelc
    ) AS nusers,
    ( SELECT COUNT(*)
        FROM ruleset
        WHERE app = a.shortnamelc AND brwf = 'B'
    ) AS nrulesetsbre,
    ( SELECT COUNT(*)
        FROM ruleset
        WHERE app = a.shortnamelc AND brwf = 'W'
    ) AS nrulesetswfe
FROM
    app a
WHERE
a.realm= $1
`

type GetAppListRow struct {
	Name         string           `json:"name"`
	Descr        string           `json:"descr"`
	Createdat    pgtype.Timestamp `json:"createdat"`
	Createdby    string           `json:"createdby"`
	Nusers       int64            `json:"nusers"`
	Nrulesetsbre int64            `json:"nrulesetsbre"`
	Nrulesetswfe int64            `json:"nrulesetswfe"`
}

func (q *Queries) GetAppList(ctx context.Context, realm string) ([]GetAppListRow, error) {
	rows, err := q.db.Query(ctx, getAppList, realm)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAppListRow
	for rows.Next() {
		var i GetAppListRow
		if err := rows.Scan(
			&i.Name,
			&i.Descr,
			&i.Createdat,
			&i.Createdby,
			&i.Nusers,
			&i.Nrulesetsbre,
			&i.Nrulesetswfe,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAppName = `-- name: GetAppName :many
select id, realm, shortname, shortnamelc, longname, setby, setat FROM app WHERE shortnamelc = $1 AND realm = $2
`

type GetAppNameParams struct {
	Shortnamelc string `json:"shortnamelc"`
	Realm       string `json:"realm"`
}

func (q *Queries) GetAppName(ctx context.Context, arg GetAppNameParams) ([]App, error) {
	rows, err := q.db.Query(ctx, getAppName, arg.Shortnamelc, arg.Realm)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []App
	for rows.Next() {
		var i App
		if err := rows.Scan(
			&i.ID,
			&i.Realm,
			&i.Shortname,
			&i.Shortnamelc,
			&i.Longname,
			&i.Setby,
			&i.Setat,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAppNames = `-- name: GetAppNames :many
SELECT shortname FROM app WHERE realm = $1
`

func (q *Queries) GetAppNames(ctx context.Context, realm string) ([]string, error) {
	rows, err := q.db.Query(ctx, getAppNames, realm)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var shortname string
		if err := rows.Scan(&shortname); err != nil {
			return nil, err
		}
		items = append(items, shortname)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
