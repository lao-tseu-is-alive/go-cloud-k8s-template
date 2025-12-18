package template_4_your_project_name

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lao-tseu-is-alive/go-cloud-k8s-common-libs/pkg/database"
)

type PGX struct {
	Conn *pgxpool.Pool
	dbi  database.DB
	log  *slog.Logger
}

// NewPgxDB will instantiate a new storage of type postgres and ensure schema exist
func NewPgxDB(ctx context.Context, db database.DB, log *slog.Logger) (Storage, error) {
	var psql PGX
	pgConn, err := db.GetPGConn()
	if err != nil {
		return nil, err
	}
	psql.Conn = pgConn
	psql.dbi = db
	psql.log = log
	var numberOfTypetemplate4YourProjectNames int
	errTypetemplate4YourProjectNameTable := pgConn.QueryRow(ctx, typetemplate4YourProjectNameCount).Scan(&numberOfTypetemplate4YourProjectNames)
	if errTypetemplate4YourProjectNameTable != nil {
		log.Error("Unable to retrieve the number of typetemplate4YourProjectName", "error", err)
		return nil, errTypetemplate4YourProjectNameTable
	}

	if numberOfTypetemplate4YourProjectNames > 0 {
		log.Info("database contains records in template_4_your_project_name_db_schema.type_template_4_your_project_name", "count", numberOfTypetemplate4YourProjectNames)
	} else {
		log.Warn("template_4_your_project_name_db_schema.type_template_4_your_project_name is empty - it should contain at least one row")
		return nil, fmt.Errorf("«template_4_your_project_name_db_schema.type_template_4_your_project_name» contains %w should not be empty", numberOfTypetemplate4YourProjectNames)
	}

	return &psql, err
}

func (db *PGX) GeoJson(ctx context.Context, offset, limit int, params GeoJsonParams) (string, error) {
	db.log.Debug("trace: entering GeoJson", "offset", offset, "limit", limit)
	if params.Type != nil {
		db.log.Info("param type", "type", *params.Type)
	}
	if params.CreatedBy != nil {
		db.log.Info("params.CreatedBy", "createdBy", *params.CreatedBy)
	}
	var (
		mayBeResultIsNull *string
		err               error
	)
	isInactive := false
	if params.Inactivated != nil {
		isInactive = *params.Inactivated
	}
	listtemplate4YourProjectNames := baseGeoJsontemplate4YourProjectNameSearch + listtemplate4YourProjectNamesConditions
	if params.Validated != nil {
		db.log.Debug("params.Validated is not nil ")
		isValidated := *params.Validated
		listtemplate4YourProjectNames += " AND validated = coalesce($6, validated) " + geoJsonListEndOfQuery
		err = pgxscan.Select(ctx, db.Conn, &mayBeResultIsNull, listtemplate4YourProjectNames,
			limit, offset, &params.Type, &params.CreatedBy, isInactive, isValidated)
	} else {
		listtemplate4YourProjectNames += geoJsonListEndOfQuery
		err = pgxscan.Select(ctx, db.Conn, &mayBeResultIsNull, listtemplate4YourProjectNames,
			limit, offset, &params.Type, &params.CreatedBy, isInactive)
	}
	if err != nil {
		db.log.Error(SelectFailedInNWithErrorE, "List", err)
		return "", err
	}
	if mayBeResultIsNull == nil {
		db.log.Info("List returned no results")
		return "", pgx.ErrNoRows
	}
	return *mayBeResultIsNull, nil
}

// List returns the list of existing template_4_your_project_names with the given offset and limit.
func (db *PGX) List(ctx context.Context, offset, limit int, params ListParams) ([]*template4YourProjectNameList, error) {
	db.log.Debug("trace: entering List", "offset", offset, "limit", limit)
	if params.Type != nil {
		db.log.Info("param type", "type", *params.Type)
	}
	if params.CreatedBy != nil {
		db.log.Info("params.CreatedBy", "createdBy", *params.CreatedBy)
	}
	var (
		res []*template4YourProjectNameList
		err error
	)
	isInactive := false
	if params.Inactivated != nil {
		isInactive = *params.Inactivated
	}
	listtemplate4YourProjectNames := basetemplate4YourProjectNameListQuery + listtemplate4YourProjectNamesConditions
	if params.Validated != nil {
		db.log.Debug("params.Validated is not nil ")
		isValidated := *params.Validated
		listtemplate4YourProjectNames += " AND validated = coalesce($6, validated) " + template_4_your_project_nameListOrderBy
		err = pgxscan.Select(ctx, db.Conn, &res, listtemplate4YourProjectNames,
			limit, offset, &params.Type, &params.CreatedBy, isInactive, isValidated)
	} else {
		listtemplate4YourProjectNames += template_4_your_project_nameListOrderBy
		err = pgxscan.Select(ctx, db.Conn, &res, listtemplate4YourProjectNames,
			limit, offset, &params.Type, &params.CreatedBy, isInactive)
	}
	if err != nil {
		db.log.Error(SelectFailedInNWithErrorE, "List", err)
		return nil, err
	}
	if res == nil {
		db.log.Info("List returned no results")
		return nil, pgx.ErrNoRows
	}
	return res, nil
}

// ListByExternalId returns the list of existing template_4_your_project_names having given externalId with the given offset and limit.
func (db *PGX) ListByExternalId(ctx context.Context, offset, limit int, externalId int) ([]*template4YourProjectNameList, error) {
	db.log.Debug("trace: entering ListByExternalId", "externalId", externalId)
	var res []*template4YourProjectNameList
	listByExternalIdtemplate4YourProjectNames := basetemplate4YourProjectNameListQuery + listByExternalIdtemplate4YourProjectNamesCondition + template_4_your_project_nameListOrderBy
	err := pgxscan.Select(ctx, db.Conn, &res, listByExternalIdtemplate4YourProjectNames, limit, offset, externalId)
	if err != nil {
		db.log.Error("ListByExternalId failed", "error", err)
		return nil, err
	}
	if res == nil {
		db.log.Info("ListByExternalId returned no results")
		return nil, pgx.ErrNoRows
	}
	return res, nil
}

func (db *PGX) Search(ctx context.Context, offset, limit int, params SearchParams) ([]*template4YourProjectNameList, error) {
	db.log.Debug("trace: entering Search", "offset", offset, "limit", limit)
	var (
		res []*template4YourProjectNameList
		err error
	)
	searchtemplate4YourProjectNames := basetemplate4YourProjectNameListQuery + listtemplate4YourProjectNamesConditions
	if params.Keywords != nil {
		searchtemplate4YourProjectNames += " AND text_search @@ plainto_tsquery('french', unaccent($6))"
		if params.Validated != nil {
			searchtemplate4YourProjectNames += " AND validated = coalesce($7, validated) " + template_4_your_project_nameListOrderBy
			err = pgxscan.Select(ctx, db.Conn, &res, searchtemplate4YourProjectNames,
				limit, offset, &params.Type, &params.CreatedBy, &params.Inactivated, &params.Keywords, &params.Validated)
		} else {
			searchtemplate4YourProjectNames += template_4_your_project_nameListOrderBy
			err = pgxscan.Select(ctx, db.Conn, &res, searchtemplate4YourProjectNames,
				limit, offset, &params.Type, &params.CreatedBy, &params.Inactivated, &params.Keywords)
		}
	} else {
		if params.Validated != nil {
			searchtemplate4YourProjectNames += " AND validated = coalesce($6, validated) " + template_4_your_project_nameListOrderBy
			err = pgxscan.Select(ctx, db.Conn, &res, searchtemplate4YourProjectNames,
				limit, offset, &params.Type, &params.CreatedBy, &params.Inactivated, &params.Validated)
		} else {
			searchtemplate4YourProjectNames += template_4_your_project_nameListOrderBy
			err = pgxscan.Select(ctx, db.Conn, &res, searchtemplate4YourProjectNames,
				limit, offset, &params.Type, &params.CreatedBy, &params.Inactivated)
		}
	}

	if err != nil {
		db.log.Error("Search failed", "error", err)
		return nil, err
	}
	if res == nil {
		db.log.Info("Search returned no results")
		return nil, pgx.ErrNoRows
	}
	return res, nil
}

// Get will retrieve the template_4_your_project_name with given id
func (db *PGX) Get(ctx context.Context, id uuid.UUID) (*template4YourProjectName, error) {
	db.log.Debug("trace: entering Get", "id", id)
	res := &template4YourProjectName{}
	err := pgxscan.Get(ctx, db.Conn, res, gettemplate4YourProjectName, id)
	if err != nil {
		db.log.Error("Get failed", "error", err)
		return nil, err
	}
	if res == nil {
		db.log.Info("Get returned no results")
		return nil, pgx.ErrNoRows
	}
	return res, nil
}

// Exist returns true only if a template_4_your_project_name with the specified id exists in store.
func (db *PGX) Exist(ctx context.Context, id uuid.UUID) bool {
	db.log.Debug("trace: entering Exist", "id", id)
	count, err := db.dbi.GetQueryInt(ctx, existtemplate4YourProjectName, id)
	if err != nil {
		db.log.Error("Exist could not be retrieved from DB", "id", id, "error", err)
		return false
	}
	if count > 0 {
		db.log.Info("Exist: id does exist", "id", id, "count", count)
		return true
	} else {
		db.log.Info("Exist: id does not exist", "id", id, "count", count)
		return false
	}
}

// Count returns the number of template_4_your_project_name stored in DB
func (db *PGX) Count(ctx context.Context, params CountParams) (int32, error) {
	db.log.Debug("trace : entering Count()")
	var (
		count int
		err   error
	)
	queryCount := counttemplate4YourProjectName + " WHERE _deleted = false AND position IS NOT NULL "
	withoutSearchParameters := true
	if params.Keywords != nil {
		withoutSearchParameters = false
		queryCount += `AND text_search @@ plainto_tsquery('french', unaccent($1))
		AND type_id = coalesce($2, type_id)
		AND _created_by = coalesce($3, _created_by)
		AND inactivated = coalesce($4, inactivated)
`
		if params.Validated != nil {
			db.log.Debug("params.Validated is not nil ")
			isValidated := *params.Validated
			queryCount += " AND validated = coalesce($4, validated) "
			count, err = db.dbi.GetQueryInt(ctx, queryCount, &params.Keywords, &params.Type, &params.CreatedBy, &params.Inactivated, isValidated)

		} else {
			count, err = db.dbi.GetQueryInt(ctx, queryCount, &params.Keywords, &params.Type, &params.CreatedBy, &params.Inactivated)
		}
	}
	if withoutSearchParameters {
		queryCount += `
		AND type_id = coalesce($1, type_id)
		AND _created_by = coalesce($2, _created_by)
		AND inactivated = coalesce($3, inactivated)
`
		if params.Validated != nil {
			db.log.Debug("params.Validated is not nil ")
			isValidated := *params.Validated
			queryCount += " AND validated = coalesce($4, validated) "
			count, err = db.dbi.GetQueryInt(ctx, queryCount, &params.Type, &params.CreatedBy, &params.Inactivated, isValidated)

		} else {
			count, err = db.dbi.GetQueryInt(ctx, queryCount, &params.Type, &params.CreatedBy, &params.Inactivated)
		}

	}

	if err != nil {
		db.log.Error("Count failed", "error", err)
		return 0, err
	}
	return int32(count), nil
}

// Create will store the new template4YourProjectName in the database
func (db *PGX) Create(ctx context.Context, t template4YourProjectName) (*template4YourProjectName, error) {
	db.log.Debug("trace: entering Create", "name", t.Name, "id", t.Id)

	rowsAffected, err := db.dbi.ExecActionQuery(ctx, createtemplate4YourProjectName,
		t.Id, t.TypeId, t.Name, &t.Description, &t.Comment, &t.ExternalId, &t.ExternalRef, //$7
		&t.BuildAt, &t.Status, &t.ContainedBy, &t.ContainedByOld, t.Validated, &t.ValidatedTime, &t.ValidatedBy, //$14
		&t.ManagedBy, t.CreatedBy, &t.MoreData, t.PosX, t.PosY)
	if err != nil {
		db.log.Error("Create unexpectedly failed", "name", t.Name, "error", err)
		return nil, err
	}
	if rowsAffected < 1 {
		db.log.Error("Create no row was created", "name", t.Name)
		return nil, err
	}
	db.log.Info("Create success", "name", t.Name, "id", t.Id)

	// if we get to here all is good, so let's retrieve a fresh copy to send it back
	createdtemplate4YourProjectName, err := db.Get(ctx, t.Id)
	if err != nil {
		return nil, fmt.Errorf("error %w: template_4_your_project_name was created, but can not be retrieved", err)
	}
	return createdtemplate4YourProjectName, nil
}

// Update the template_4_your_project_name stored in DB with given id and other information in struct
func (db *PGX) Update(ctx context.Context, id uuid.UUID, t template4YourProjectName) (*template4YourProjectName, error) {
	db.log.Debug("trace: entering Update", "id", t.Id)

	rowsAffected, err := db.dbi.ExecActionQuery(ctx, updatetemplate4YourProjectName,
		t.Id, t.TypeId, t.Name, &t.Description, &t.Comment, &t.ExternalId, &t.ExternalRef, //$7
		&t.BuildAt, &t.Status, &t.ContainedBy, &t.ContainedByOld, t.Inactivated, &t.InactivatedTime, &t.InactivatedBy, &t.InactivatedReason, //$15
		t.Validated, &t.ValidatedTime, &t.ValidatedBy, //$18
		&t.ManagedBy, &t.LastModifiedBy, &t.MoreData, t.PosX, t.PosY) //$23
	if err != nil {

		db.log.Error("Update unexpectedly failed", "id", t.Id, "error", err)
		return nil, err
	}
	if rowsAffected < 1 {
		db.log.Error("Update no row was updated", "id", t.Id)
		return nil, err
	}

	// if we get to here all is good, so let's retrieve a fresh copy to send it back
	updatedtemplate4YourProjectName, err := db.Get(ctx, t.Id)
	if err != nil {
		return nil, fmt.Errorf("error %w: template_4_your_project_name was updated, but can not be retrieved", err)
	}
	return updatedtemplate4YourProjectName, nil
}

// Delete the template_4_your_project_name stored in DB with given id
func (db *PGX) Delete(ctx context.Context, id uuid.UUID, userId int32) error {
	db.log.Debug("trace: entering Delete", "id", id)
	rowsAffected, err := db.dbi.ExecActionQuery(ctx, deletetemplate4YourProjectName, userId, id)
	if err != nil {
		db.log.Error("template_4_your_project_name could not be deleted", "id", id, "error", err)
		return fmt.Errorf("template_4_your_project_name could not be deleted: %w", err)
	}
	if rowsAffected < 1 {
		db.log.Error("template_4_your_project_name was not deleted", "id", id)
		return fmt.Errorf("template_4_your_project_name was not marked for deletetion")
	}
	return nil
}

// Istemplate4YourProjectNameActive returns true if the template_4_your_project_name with the specified id has the inactivated attribute set to false
func (db *PGX) Istemplate4YourProjectNameActive(ctx context.Context, id uuid.UUID) bool {
	db.log.Debug("trace: entering Istemplate4YourProjectNameActive", "id", id)
	count, err := db.dbi.GetQueryInt(ctx, isActivetemplate4YourProjectName, id)
	if err != nil {
		db.log.Error("Istemplate4YourProjectNameActive could not be retrieved from DB", "id", id, "error", err)
		return false
	}
	if count > 0 {
		db.log.Info("Istemplate4YourProjectNameActive is true", "id", id, "count", count)
		return true
	} else {
		db.log.Info("Istemplate4YourProjectNameActive is false", "id", id, "count", count)
		return false
	}
}

// IsUserOwner returns true only if userId is the creator of the record (owner) of this template_4_your_project_name in store.
func (db *PGX) IsUserOwner(ctx context.Context, id uuid.UUID, userId int32) bool {
	db.log.Debug("trace: entering IsUserOwner", "id", id, "userId", userId)
	count, err := db.dbi.GetQueryInt(ctx, existtemplate4YourProjectNameOwnedBy, id, userId)
	if err != nil {
		db.log.Error("IsUserOwner could not be retrieved from DB", "id", id, "userId", userId, "error", err)
		return false
	}
	if count > 0 {
		db.log.Info("IsUserOwner is true", "id", id, "userId", userId, "count", count)
		return true
	} else {
		db.log.Info("IsUserOwner is false", "id", id, "userId", userId, "count", count)
		return false
	}
}

// CreateTypetemplate4YourProjectName will store the new Typetemplate4YourProjectName in the database
func (db *PGX) CreateTypetemplate4YourProjectName(ctx context.Context, tt Typetemplate4YourProjectName) (*Typetemplate4YourProjectName, error) {
	db.log.Debug("trace: entering CreateTypetemplate4YourProjectName", "name", tt.Name, "createdBy", tt.CreatedBy)
	var lastInsertId int = 0
	err := db.Conn.QueryRow(ctx, createTypetemplate4YourProjectName,
		tt.Name, &tt.Description, &tt.Comment, &tt.ExternalId, &tt.TableName, &tt.GeometryType, //$6
		&tt.ManagedBy, tt.IconPath, tt.CreatedBy, &tt.MoreDataSchema).Scan(&lastInsertId)
	if err != nil {
		db.log.Error("CreateTypetemplate4YourProjectName unexpectedly failed", "name", tt.Name, "error", err)
		return nil, err
	}
	db.log.Info("CreateTypetemplate4YourProjectName success", "name", tt.Name, "id", lastInsertId)

	// if we get to here all is good, so let's retrieve a fresh copy to send it back
	createdTypetemplate4YourProjectName, err := db.GetTypetemplate4YourProjectName(ctx, int32(lastInsertId))
	if err != nil {
		return nil, fmt.Errorf("error %w: typetemplate4YourProjectName was created, but can not be retrieved", err)
	}
	return createdTypetemplate4YourProjectName, nil
}

// UpdateTypetemplate4YourProjectName updates the Typetemplate4YourProjectName stored in DB with given id and other information in struct
func (db *PGX) UpdateTypetemplate4YourProjectName(ctx context.Context, id int32, tt Typetemplate4YourProjectName) (*Typetemplate4YourProjectName, error) {
	db.log.Debug("trace: entering UpdateTypetemplate4YourProjectName", "id", id)

	rowsAffected, err := db.dbi.ExecActionQuery(ctx, updateTypeTing,
		id, tt.Name, &tt.Description, &tt.Comment, &tt.ExternalId, &tt.TableName, //$6
		&tt.GeometryType, tt.Inactivated, &tt.InactivatedTime, &tt.InactivatedBy, &tt.InactivatedReason, //$11
		&tt.ManagedBy, tt.IconPath, &tt.LastModifiedBy, &tt.MoreDataSchema) //$14
	if err != nil {

		db.log.Error("UpdateTypetemplate4YourProjectName unexpectedly failed", "id", id, "error", err)
		return nil, err
	}
	if rowsAffected < 1 {
		db.log.Error("UpdateTypetemplate4YourProjectName no row was updated", "id", id)
		return nil, err
	}

	// if we get to here all is good, so let's retrieve a fresh copy to send it back
	updatedTypetemplate4YourProjectName, err := db.GetTypetemplate4YourProjectName(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error %w: template_4_your_project_name was updated, but can not be retrieved", err)
	}
	return updatedTypetemplate4YourProjectName, nil
}

// DeleteTypetemplate4YourProjectName deletes the Typetemplate4YourProjectName stored in DB with given id
func (db *PGX) DeleteTypetemplate4YourProjectName(ctx context.Context, id int32, userId int32) error {
	db.log.Debug("trace: entering DeleteTypetemplate4YourProjectName", "id", id)
	rowsAffected, err := db.dbi.ExecActionQuery(ctx, deleteTypetemplate4YourProjectName, userId, id)
	if err != nil {
		db.log.Error("typetemplate_4_your_project_name could not be deleted", "id", id, "error", err)
		return fmt.Errorf("typetemplate_4_your_project_name could not be deleted: %w", err)
	}
	if rowsAffected < 1 {
		db.log.Error("typetemplate_4_your_project_name was not deleted", "id", id)
		return fmt.Errorf("typetemplate_4_your_project_name was not marked for deletion")
	}
	return nil
}

// ListTypetemplate4YourProjectName returns the list of existing Typetemplate4YourProjectName with the given offset and limit.
func (db *PGX) ListTypetemplate4YourProjectName(ctx context.Context, offset, limit int, params Typetemplate4YourProjectNameListParams) ([]*Typetemplate4YourProjectNameList, error) {
	db.log.Debug("trace : entering ListTypetemplate4YourProjectName")
	var (
		res []*Typetemplate4YourProjectNameList
		err error
	)
	listTypetemplate4YourProjectNames := typetemplate4YourProjectNameListQuery
	if params.Keywords != nil {
		listTypetemplate4YourProjectNames += listTypetemplate4YourProjectNamesConditionsWithKeywords + typetemplate4YourProjectNameListOrderBy
		err = pgxscan.Select(ctx, db.Conn, &res, listTypetemplate4YourProjectNames,
			limit, offset, &params.Keywords, &params.CreatedBy, &params.ExternalId, &params.Inactivated)
	} else {
		listTypetemplate4YourProjectNames += listTypetemplate4YourProjectNamesConditionsWithoutKeywords + typetemplate4YourProjectNameListOrderBy
		err = pgxscan.Select(ctx, db.Conn, &res, listTypetemplate4YourProjectNames,
			limit, offset, &params.CreatedBy, &params.ExternalId, &params.Inactivated)
	}

	if err != nil {
		db.log.Error("ListTypetemplate4YourProjectName failed", "error", err)
		return nil, err
	}
	if res == nil {
		db.log.Info("ListTypetemplate4YourProjectName returned no results")
		return nil, pgx.ErrNoRows
	}
	return res, nil
}

// GetTypetemplate4YourProjectName will retrieve the Typetemplate4YourProjectName with given id
func (db *PGX) GetTypetemplate4YourProjectName(ctx context.Context, id int32) (*Typetemplate4YourProjectName, error) {
	db.log.Debug("trace: entering GetTypetemplate4YourProjectName", "id", id)
	res := &Typetemplate4YourProjectName{}
	err := pgxscan.Get(ctx, db.Conn, res, getTypetemplate4YourProjectName, id)
	if err != nil {
		db.log.Error("GetTypetemplate4YourProjectName failed", "error", err)
		return nil, err
	}
	if res == nil {
		db.log.Info("GetTypetemplate4YourProjectName returned no results", "id", id)
		return nil, pgx.ErrNoRows
	}
	return res, nil
}

// CountTypetemplate4YourProjectName returns the number of Typetemplate4YourProjectName based on search criteria
func (db *PGX) CountTypetemplate4YourProjectName(ctx context.Context, params Typetemplate4YourProjectNameCountParams) (int32, error) {
	db.log.Debug("trace : entering CountTypetemplate4YourProjectName()")
	var (
		count int
		err   error
	)
	queryCount := countTypetemplate4YourProjectName + " WHERE 1 = 1 "
	withoutSearchParameters := true
	if params.Keywords != nil {
		withoutSearchParameters = false
		queryCount += `AND text_search @@ plainto_tsquery('french', unaccent($1))
		AND _created_by = coalesce($2, _created_by)
		AND inactivated = coalesce($3, inactivated)
`
		count, err = db.dbi.GetQueryInt(ctx, queryCount, &params.Keywords, &params.CreatedBy, &params.Inactivated)
	}
	if withoutSearchParameters {
		queryCount += `
		AND _created_by = coalesce($1, _created_by)
		AND inactivated = coalesce($2, inactivated)
`
		count, err = db.dbi.GetQueryInt(ctx, queryCount, &params.CreatedBy, &params.Inactivated)

	}
	if err != nil {
		db.log.Error("CountTypetemplate4YourProjectName failed", "error", err)
		return 0, err
	}
	return int32(count), nil
}

// GetTypetemplate4YourProjectNameMaxId will retrieve maximum value of Typetemplate4YourProjectName id existing in store.
func (db *PGX) GetTypetemplate4YourProjectNameMaxId(ctx context.Context) (int32, error) {
	db.log.Debug("trace : entering GetTypetemplate4YourProjectNameMaxId")
	existingMaxId, err := db.dbi.GetQueryInt(ctx, typetemplate4YourProjectNameMaxId)
	if err != nil {
		db.log.Error("GetTypetemplate4YourProjectNameMaxId() failed", "error", err)
		return 0, err
	}
	return int32(existingMaxId), nil
}
