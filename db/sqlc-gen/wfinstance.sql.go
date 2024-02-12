// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: wfinstance.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const addWFNewInstace = `-- name: AddWFNewInstace :one
INSERT INTO 
wfinstance
(entityid,slice, app, class, workflow, step,loggedat, nextstep,parent)
VALUES ($1,$2,$3,$4,$5,$6,(NOW() :: timestamp),$7,$8) 
RETURNING id
`

type AddWFNewInstaceParams struct {
	Entityid string      `json:"entityid"`
	Slice    int32       `json:"slice"`
	App      string      `json:"app"`
	Class    string      `json:"class"`
	Workflow string      `json:"workflow"`
	Step     string      `json:"step"`
	Nextstep string      `json:"nextstep"`
	Parent   pgtype.Int4 `json:"parent"`
}

func (q *Queries) AddWFNewInstace(ctx context.Context, arg AddWFNewInstaceParams) (int32, error) {
	row := q.db.QueryRow(ctx, addWFNewInstace,
		arg.Entityid,
		arg.Slice,
		arg.App,
		arg.Class,
		arg.Workflow,
		arg.Step,
		arg.Nextstep,
		arg.Parent,
	)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const getLoggedate = `-- name: GetLoggedate :one
SELECT loggedat 
FROM wfinstance
WHERE id = $1
`

func (q *Queries) GetLoggedate(ctx context.Context, id int32) (pgtype.Timestamp, error) {
	row := q.db.QueryRow(ctx, getLoggedate, id)
	var loggedat pgtype.Timestamp
	err := row.Scan(&loggedat)
	return loggedat, err
}

const getWFINstance = `-- name: GetWFINstance :many

SELECT id, entityid, slice, app, class, workflow, step, loggedat, doneat, nextstep, parent 
FROM wfinstance
WHERE slice = $1 
AND app = $2
AND workflow = $3
AND entityid = $4
`

type GetWFINstanceParams struct {
	Slice    int32  `json:"slice"`
	App      string `json:"app"`
	Workflow string `json:"workflow"`
	Entityid string `json:"entityid"`
}

func (q *Queries) GetWFINstance(ctx context.Context, arg GetWFINstanceParams) ([]Wfinstance, error) {
	rows, err := q.db.Query(ctx, getWFINstance,
		arg.Slice,
		arg.App,
		arg.Workflow,
		arg.Entityid,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Wfinstance
	for rows.Next() {
		var i Wfinstance
		if err := rows.Scan(
			&i.ID,
			&i.Entityid,
			&i.Slice,
			&i.App,
			&i.Class,
			&i.Workflow,
			&i.Step,
			&i.Loggedat,
			&i.Doneat,
			&i.Nextstep,
			&i.Parent,
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