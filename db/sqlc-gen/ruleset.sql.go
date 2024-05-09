// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: ruleset.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
)

const activateBRERuleSet = `-- name: ActivateBRERuleSet :exec
UPDATE ruleset
SET is_active = true
WHERE realm = $1
AND slice = $2
AND app = $3
AND class = $4
AND setname = $5
AND brwf = $6
`

type ActivateBRERuleSetParams struct {
	Realm   string   `json:"realm"`
	Slice   int32    `json:"slice"`
	App     string   `json:"app"`
	Class   string   `json:"class"`
	Setname string   `json:"setname"`
	Brwf    BrwfEnum `json:"brwf"`
}

func (q *Queries) ActivateBRERuleSet(ctx context.Context, arg ActivateBRERuleSetParams) error {
	_, err := q.db.Exec(ctx, activateBRERuleSet,
		arg.Realm,
		arg.Slice,
		arg.App,
		arg.Class,
		arg.Setname,
		arg.Brwf,
	)
	return err
}

const allRuleset = `-- name: AllRuleset :many
SELECT
    id, realm, slice, app, brwf, class, setname, schemaid, is_active, is_internal, ruleset, createdat, createdby, editedat, editedby
FROM
    public.ruleset
`

func (q *Queries) AllRuleset(ctx context.Context) ([]Ruleset, error) {
	rows, err := q.db.Query(ctx, allRuleset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Ruleset
	for rows.Next() {
		var i Ruleset
		if err := rows.Scan(
			&i.ID,
			&i.Realm,
			&i.Slice,
			&i.App,
			&i.Brwf,
			&i.Class,
			&i.Setname,
			&i.Schemaid,
			&i.IsActive,
			&i.IsInternal,
			&i.Ruleset,
			&i.Createdat,
			&i.Createdby,
			&i.Editedat,
			&i.Editedby,
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

const deActivateBRERuleSet = `-- name: DeActivateBRERuleSet :exec
UPDATE ruleset
SET is_active = false
WHERE realm = $1
AND slice = $2
AND app = $3
AND class = $4
AND setname = $5
AND brwf = $6
`

type DeActivateBRERuleSetParams struct {
	Realm   string   `json:"realm"`
	Slice   int32    `json:"slice"`
	App     string   `json:"app"`
	Class   string   `json:"class"`
	Setname string   `json:"setname"`
	Brwf    BrwfEnum `json:"brwf"`
}

func (q *Queries) DeActivateBRERuleSet(ctx context.Context, arg DeActivateBRERuleSetParams) error {
	_, err := q.db.Exec(ctx, deActivateBRERuleSet,
		arg.Realm,
		arg.Slice,
		arg.App,
		arg.Class,
		arg.Setname,
		arg.Brwf,
	)
	return err
}

const getApp = `-- name: GetApp :one
SELECT app
FROM ruleset
WHERE
    slice = $1
    AND app = $2
    AND class = $3
    AND realm = $4
    AND brwf = 'W'
`

type GetAppParams struct {
	Slice int32  `json:"slice"`
	App   string `json:"app"`
	Class string `json:"class"`
	Realm string `json:"realm"`
}

func (q *Queries) GetApp(ctx context.Context, arg GetAppParams) (string, error) {
	row := q.db.QueryRow(ctx, getApp,
		arg.Slice,
		arg.App,
		arg.Class,
		arg.Realm,
	)
	var app string
	err := row.Scan(&app)
	return app, err
}

const getBRERuleSetActiveStatus = `-- name: GetBRERuleSetActiveStatus :one
SELECT ruleset, is_active ,setname FROM ruleset
WHERE realm = $1
AND slice = $2
AND app = $3
AND class = $4
AND setname = $5
AND brwf = $6
`

type GetBRERuleSetActiveStatusParams struct {
	Realm   string   `json:"realm"`
	Slice   int32    `json:"slice"`
	App     string   `json:"app"`
	Class   string   `json:"class"`
	Setname string   `json:"setname"`
	Brwf    BrwfEnum `json:"brwf"`
}

type GetBRERuleSetActiveStatusRow struct {
	Ruleset  []byte      `json:"ruleset"`
	IsActive pgtype.Bool `json:"is_active"`
	Setname  string      `json:"setname"`
}

func (q *Queries) GetBRERuleSetActiveStatus(ctx context.Context, arg GetBRERuleSetActiveStatusParams) (GetBRERuleSetActiveStatusRow, error) {
	row := q.db.QueryRow(ctx, getBRERuleSetActiveStatus,
		arg.Realm,
		arg.Slice,
		arg.App,
		arg.Class,
		arg.Setname,
		arg.Brwf,
	)
	var i GetBRERuleSetActiveStatusRow
	err := row.Scan(&i.Ruleset, &i.IsActive, &i.Setname)
	return i, err
}

const getBRERuleSetCount = `-- name: GetBRERuleSetCount :one
SELECT count(*) FROM ruleset
WHERE realm = $1
AND slice = $2
AND app = $3
AND class = $4
AND setname = $5
AND brwf = $6
`

type GetBRERuleSetCountParams struct {
	Realm   string   `json:"realm"`
	Slice   int32    `json:"slice"`
	App     string   `json:"app"`
	Class   string   `json:"class"`
	Setname string   `json:"setname"`
	Brwf    BrwfEnum `json:"brwf"`
}

func (q *Queries) GetBRERuleSetCount(ctx context.Context, arg GetBRERuleSetCountParams) (int64, error) {
	row := q.db.QueryRow(ctx, getBRERuleSetCount,
		arg.Realm,
		arg.Slice,
		arg.App,
		arg.Class,
		arg.Setname,
		arg.Brwf,
	)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getClass = `-- name: GetClass :one
SELECT class
FROM ruleset
WHERE
    slice = $1
    AND app = $2
    AND class = $3
    AND realm = $4
    AND brwf = 'W'
`

type GetClassParams struct {
	Slice int32  `json:"slice"`
	App   string `json:"app"`
	Class string `json:"class"`
	Realm string `json:"realm"`
}

func (q *Queries) GetClass(ctx context.Context, arg GetClassParams) (string, error) {
	row := q.db.QueryRow(ctx, getClass,
		arg.Slice,
		arg.App,
		arg.Class,
		arg.Realm,
	)
	var class string
	err := row.Scan(&class)
	return class, err
}

const getWFActiveStatus = `-- name: GetWFActiveStatus :one
SELECT is_active
FROM ruleset
WHERE
    slice = $1
    AND app = $2
    AND class = $3
    AND realm = $4
    AND brwf = 'W'
    AND setname = $5
`

type GetWFActiveStatusParams struct {
	Slice   int32  `json:"slice"`
	App     string `json:"app"`
	Class   string `json:"class"`
	Realm   string `json:"realm"`
	Setname string `json:"setname"`
}

func (q *Queries) GetWFActiveStatus(ctx context.Context, arg GetWFActiveStatusParams) (pgtype.Bool, error) {
	row := q.db.QueryRow(ctx, getWFActiveStatus,
		arg.Slice,
		arg.App,
		arg.Class,
		arg.Realm,
		arg.Setname,
	)
	var is_active pgtype.Bool
	err := row.Scan(&is_active)
	return is_active, err
}

const getWFInternalStatus = `-- name: GetWFInternalStatus :one
SELECT is_internal
FROM ruleset
WHERE
    slice = $1
    AND app = $2
    AND class = $3
    AND realm = $4
    AND brwf = 'W'
    AND setname = $5
`

type GetWFInternalStatusParams struct {
	Slice   int32  `json:"slice"`
	App     string `json:"app"`
	Class   string `json:"class"`
	Realm   string `json:"realm"`
	Setname string `json:"setname"`
}

func (q *Queries) GetWFInternalStatus(ctx context.Context, arg GetWFInternalStatusParams) (bool, error) {
	row := q.db.QueryRow(ctx, getWFInternalStatus,
		arg.Slice,
		arg.App,
		arg.Class,
		arg.Realm,
		arg.Setname,
	)
	var is_internal bool
	err := row.Scan(&is_internal)
	return is_internal, err
}

const isWorkflowReferringSchema = `-- name: IsWorkflowReferringSchema :one
select count(*)
From ruleset
Where realm = $1
AND slice = $2
AND app = $3
AND class = $4
AND is_active = true
`

type IsWorkflowReferringSchemaParams struct {
	Realm string `json:"realm"`
	Slice int32  `json:"slice"`
	App   string `json:"app"`
	Class string `json:"class"`
}

func (q *Queries) IsWorkflowReferringSchema(ctx context.Context, arg IsWorkflowReferringSchemaParams) (int64, error) {
	row := q.db.QueryRow(ctx, isWorkflowReferringSchema,
		arg.Realm,
		arg.Slice,
		arg.App,
		arg.Class,
	)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const loadRuleSet = `-- name: LoadRuleSet :one
SELECT id, realm, slice, app, brwf, class, setname, schemaid, is_active, is_internal, ruleset, createdat, createdby, editedat, editedby 
FROM ruleset 
WHERE
     realm = $3::varchar
    AND slice = (SELECT realmslice.id FROM realmslice WHERE realmslice.id= $4 AND realmslice.realm = $3 )
    AND class = $1
    AND app = (SELECT app.shortnamelc FROM app WHERE app.shortnamelc= $5 AND app.realm = $3)
    AND setname = $2
`

type LoadRuleSetParams struct {
	Class     string `json:"class"`
	Setname   string `json:"setname"`
	RealmName string `json:"realm_name"`
	Slice     int32  `json:"slice"`
	App       string `json:"app"`
}

func (q *Queries) LoadRuleSet(ctx context.Context, arg LoadRuleSetParams) (Ruleset, error) {
	row := q.db.QueryRow(ctx, loadRuleSet,
		arg.Class,
		arg.Setname,
		arg.RealmName,
		arg.Slice,
		arg.App,
	)
	var i Ruleset
	err := row.Scan(
		&i.ID,
		&i.Realm,
		&i.Slice,
		&i.App,
		&i.Brwf,
		&i.Class,
		&i.Setname,
		&i.Schemaid,
		&i.IsActive,
		&i.IsInternal,
		&i.Ruleset,
		&i.Createdat,
		&i.Createdby,
		&i.Editedat,
		&i.Editedby,
	)
	return i, err
}

const rulesetRowLock = `-- name: RulesetRowLock :one
SELECT id, realm, slice, app, brwf, class, setname, schemaid, is_active, is_internal, ruleset, createdat, createdby, editedat, editedby 
FROM ruleset 
WHERE
     realm = $2::varchar
    AND slice = (SELECT realmslice.id FROM realmslice WHERE realmslice.id= $3 AND realmslice.realm = $2 )
    AND class = $1
    AND app = (SELECT app.shortnamelc FROM app WHERE app.shortnamelc= $4 AND app.realm = $2)
FOR UPDATE
`

type RulesetRowLockParams struct {
	Class     string `json:"class"`
	RealmName string `json:"realm_name"`
	Slice     int32  `json:"slice"`
	App       string `json:"app"`
}

func (q *Queries) RulesetRowLock(ctx context.Context, arg RulesetRowLockParams) (Ruleset, error) {
	row := q.db.QueryRow(ctx, rulesetRowLock,
		arg.Class,
		arg.RealmName,
		arg.Slice,
		arg.App,
	)
	var i Ruleset
	err := row.Scan(
		&i.ID,
		&i.Realm,
		&i.Slice,
		&i.App,
		&i.Brwf,
		&i.Class,
		&i.Setname,
		&i.Schemaid,
		&i.IsActive,
		&i.IsInternal,
		&i.Ruleset,
		&i.Createdat,
		&i.Createdby,
		&i.Editedat,
		&i.Editedby,
	)
	return i, err
}

const workFlowNew = `-- name: WorkFlowNew :one
INSERT INTO
    ruleset (
        realm, slice, app, brwf, class, setname, schemaid, is_active, is_internal, ruleset, createdat, createdby
    )
VALUES (
        $8::varchar,
        (SELECT realmslice.id FROM realmslice WHERE realmslice.id= $9 AND realmslice.realm = $8 ),
        (SELECT app.shortnamelc FROM app WHERE app.shortnamelc= $10 AND app.realm = $8), 
        $1, $2, $3, $4, false, $5, $6, CURRENT_TIMESTAMP, $7
    )RETURNING id
`

type WorkFlowNewParams struct {
	Brwf       BrwfEnum `json:"brwf"`
	Class      string   `json:"class"`
	Setname    string   `json:"setname"`
	Schemaid   int32    `json:"schemaid"`
	IsInternal bool     `json:"is_internal"`
	Ruleset    []byte   `json:"ruleset"`
	Createdby  string   `json:"createdby"`
	RealmName  string   `json:"realm_name"`
	Slice      int32    `json:"slice"`
	App        string   `json:"app"`
}

func (q *Queries) WorkFlowNew(ctx context.Context, arg WorkFlowNewParams) (int32, error) {
	row := q.db.QueryRow(ctx, workFlowNew,
		arg.Brwf,
		arg.Class,
		arg.Setname,
		arg.Schemaid,
		arg.IsInternal,
		arg.Ruleset,
		arg.Createdby,
		arg.RealmName,
		arg.Slice,
		arg.App,
	)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const workFlowUpdate = `-- name: WorkFlowUpdate :execresult
UPDATE ruleset
SET
    ruleset = $4,
    editedat = CURRENT_TIMESTAMP,
    editedby = $5
WHERE
    realm = $6::varchar
    AND slice = (SELECT realmslice.id FROM realmslice WHERE realmslice.id= $7 AND realmslice.realm = $6 )
    AND class = $1
    AND app = (SELECT app.shortnamelc FROM app WHERE app.shortnamelc= $8 AND app.realm = $6)
    AND brwf = $2
    AND setname = $3
    AND is_active = false
`

type WorkFlowUpdateParams struct {
	Class     string      `json:"class"`
	Brwf      BrwfEnum    `json:"brwf"`
	Setname   string      `json:"setname"`
	Ruleset   []byte      `json:"ruleset"`
	Editedby  pgtype.Text `json:"editedby"`
	RealmName string      `json:"realm_name"`
	Slice     int32       `json:"slice"`
	App       string      `json:"app"`
}

func (q *Queries) WorkFlowUpdate(ctx context.Context, arg WorkFlowUpdateParams) (pgconn.CommandTag, error) {
	return q.db.Exec(ctx, workFlowUpdate,
		arg.Class,
		arg.Brwf,
		arg.Setname,
		arg.Ruleset,
		arg.Editedby,
		arg.RealmName,
		arg.Slice,
		arg.App,
	)
}

const workflowDelete = `-- name: WorkflowDelete :execresult
DELETE from ruleset
where
    brwf = $6
    AND is_active = false
    and slice = $1
    and app = $2
    and class = $3
    and setname = $4
    AND realm = $5
`

type WorkflowDeleteParams struct {
	Slice   int32    `json:"slice"`
	App     string   `json:"app"`
	Class   string   `json:"class"`
	Setname string   `json:"setname"`
	Realm   string   `json:"realm"`
	Brwf    BrwfEnum `json:"brwf"`
}

func (q *Queries) WorkflowDelete(ctx context.Context, arg WorkflowDeleteParams) (pgconn.CommandTag, error) {
	return q.db.Exec(ctx, workflowDelete,
		arg.Slice,
		arg.App,
		arg.Class,
		arg.Setname,
		arg.Realm,
		arg.Brwf,
	)
}

const workflowList = `-- name: WorkflowList :many
select
    id,
    slice,
    app,
    class,
    setname as name,
    is_active,
    is_internal,
    createdat,
    createdby,
    editedat,
    editedby
from ruleset
where
    realm = $1
    AND ($2::INTEGER is null OR slice = $2::INTEGER)
    AND ($3::text[] is null OR app = any( $3::text[]))
    AND ($4::text is null OR class = $4::text)
    AND ($5::text is null OR setname = $5::text)
    AND ($6::BOOLEAN is null OR is_active = $6::BOOLEAN)
    AND ($7::BOOLEAN is null OR is_internal = $7::BOOLEAN)
    and brwf =  $8
`

type WorkflowListParams struct {
	Realm      string      `json:"realm"`
	Slice      pgtype.Int4 `json:"slice"`
	App        []string    `json:"app"`
	Class      pgtype.Text `json:"class"`
	Setname    pgtype.Text `json:"setname"`
	IsActive   pgtype.Bool `json:"is_active"`
	IsInternal pgtype.Bool `json:"is_internal"`
	Brwf       BrwfEnum    `json:"brwf"`
}

type WorkflowListRow struct {
	ID         int32            `json:"id"`
	Slice      int32            `json:"slice"`
	App        string           `json:"app"`
	Class      string           `json:"class"`
	Name       string           `json:"name"`
	IsActive   pgtype.Bool      `json:"is_active"`
	IsInternal bool             `json:"is_internal"`
	Createdat  pgtype.Timestamp `json:"createdat"`
	Createdby  string           `json:"createdby"`
	Editedat   pgtype.Timestamp `json:"editedat"`
	Editedby   pgtype.Text      `json:"editedby"`
}

func (q *Queries) WorkflowList(ctx context.Context, arg WorkflowListParams) ([]WorkflowListRow, error) {
	rows, err := q.db.Query(ctx, workflowList,
		arg.Realm,
		arg.Slice,
		arg.App,
		arg.Class,
		arg.Setname,
		arg.IsActive,
		arg.IsInternal,
		arg.Brwf,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []WorkflowListRow
	for rows.Next() {
		var i WorkflowListRow
		if err := rows.Scan(
			&i.ID,
			&i.Slice,
			&i.App,
			&i.Class,
			&i.Name,
			&i.IsActive,
			&i.IsInternal,
			&i.Createdat,
			&i.Createdby,
			&i.Editedat,
			&i.Editedby,
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

const workflowget = `-- name: Workflowget :one
select
    id,
    slice,
    app,
    class,
    setname as name,
    is_active,
    is_internal,
    ruleset as flowrules,
    createdat,
    createdby,
    editedat,
    editedby
from ruleset
where
    slice = $1
    and app = $2
    and class = $3
    and setname = $4
    and realm = $5
    AND brwf = $6
`

type WorkflowgetParams struct {
	Slice   int32    `json:"slice"`
	App     string   `json:"app"`
	Class   string   `json:"class"`
	Setname string   `json:"setname"`
	Realm   string   `json:"realm"`
	Brwf    BrwfEnum `json:"brwf"`
}

type WorkflowgetRow struct {
	ID         int32            `json:"id"`
	Slice      int32            `json:"slice"`
	App        string           `json:"app"`
	Class      string           `json:"class"`
	Name       string           `json:"name"`
	IsActive   pgtype.Bool      `json:"is_active"`
	IsInternal bool             `json:"is_internal"`
	Flowrules  []byte           `json:"flowrules"`
	Createdat  pgtype.Timestamp `json:"createdat"`
	Createdby  string           `json:"createdby"`
	Editedat   pgtype.Timestamp `json:"editedat"`
	Editedby   pgtype.Text      `json:"editedby"`
}

func (q *Queries) Workflowget(ctx context.Context, arg WorkflowgetParams) (WorkflowgetRow, error) {
	row := q.db.QueryRow(ctx, workflowget,
		arg.Slice,
		arg.App,
		arg.Class,
		arg.Setname,
		arg.Realm,
		arg.Brwf,
	)
	var i WorkflowgetRow
	err := row.Scan(
		&i.ID,
		&i.Slice,
		&i.App,
		&i.Class,
		&i.Name,
		&i.IsActive,
		&i.IsInternal,
		&i.Flowrules,
		&i.Createdat,
		&i.Createdby,
		&i.Editedat,
		&i.Editedby,
	)
	return i, err
}

const ruleExists = `-- name: ruleExists :one
select 1 
from ruleset 
where realm = $1
AND app= $2
AND slice= $3 
AND class = $4
`

type ruleExistsParams struct {
	Realm string `json:"realm"`
	App   string `json:"app"`
	Slice int32  `json:"slice"`
	Class string `json:"class"`
}

func (q *Queries) ruleExists(ctx context.Context, arg ruleExistsParams) (int32, error) {
	row := q.db.QueryRow(ctx, ruleExists,
		arg.Realm,
		arg.App,
		arg.Slice,
		arg.Class,
	)
	var column_1 int32
	err := row.Scan(&column_1)
	return column_1, err
}
