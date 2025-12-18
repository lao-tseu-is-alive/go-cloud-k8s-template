package template_4_your_project_name

const (
	basetemplate4YourProjectNameListQuery = `
SELECT 
       id,
       type_id,
       name,
       description,
       external_id,
       inactivated,
       validated,
       status, 
       _created_by as created_by,
       _created_at as created_at,
	   st_x(position) as pos_x,
       st_y(position) as pos_y
FROM template_4_your_project_name_db_schema.template_4_your_project_name
WHERE _deleted = false AND position IS NOT NULL
`
	template_4_your_project_nameListOrderBy = " ORDER BY _created_at DESC LIMIT $1 OFFSET $2;"
	listtemplate4YourProjectNamesConditions = `
 AND type_id = coalesce($3, type_id)
 AND _created_by = coalesce($4, _created_by)
 AND inactivated = coalesce($5, inactivated) 
`
	listByExternalIdtemplate4YourProjectNamesCondition = " AND external_id = $3 "
	searchtemplate4YourProjectNamesConditions          = `
 AND type_id = coalesce($3, type_id)
 AND _created_by = coalesce($4, _created_by)
 AND inactivated = coalesce($5, inactivated)
 AND text_search @@ plainto_tsquery('french', unaccent($6))
`
	createtemplate4YourProjectName = `
INSERT INTO template_4_your_project_name_db_schema.template_4_your_project_name
(id, type_id, name, description, comment, external_id, external_ref,
 build_at, status, contained_by, contained_by_old,validated, validated_time, validated_by,
 managed_by, _created_at, _created_by, more_data, text_search, position)
VALUES ($1, $2, $3, $4, $5, $6, $7,
        $8, $9, $10, $11, $12, $13, $14,
        $15, CURRENT_TIMESTAMP, $16, $17,
        to_tsvector('french', unaccent($3) ||
                              ' ' || coalesce(unaccent($4), ' ') ||
                              ' ' || coalesce(unaccent($5), ' ') ),
        ST_SetSRID(ST_MakePoint($18,$19), 2056));
`

	gettemplate4YourProjectName = `SELECT id,
       type_id,
       name,
       description,
       comment,
       external_id,
       external_ref,
       build_at,
       status,
       contained_by,
       contained_by_old,
       inactivated,
       inactivated_time,
       inactivated_by,
       inactivated_reason,
       validated,
       validated_time,
       validated_by,
       managed_by,
       _created_at as created_at,
       _created_by as created_by,
       _last_modified_at as last_modified_at,
       _last_modified_by as last_modified_by,
       _deleted as deleted,
       _deleted_at as deleted_at,
       _deleted_by as deleted_by,
       more_data, 
       round(st_x(ST_Centroid(position))::numeric, 2) AS pos_x,
       round(st_y(ST_Centroid(position))::numeric, 2) AS pos_y
FROM template_4_your_project_name_db_schema.template_4_your_project_name
WHERE id = $1;
`
	existtemplate4YourProjectName        = `SELECT COUNT(*) FROM template_4_your_project_name_db_schema.template_4_your_project_name WHERE id = $1;`
	isActivetemplate4YourProjectName     = `SELECT COUNT(*) FROM template_4_your_project_name_db_schema.template_4_your_project_name WHERE inactivated=false AND id = $1;`
	existtemplate4YourProjectNameOwnedBy = `SELECT COUNT(*) FROM template_4_your_project_name_db_schema.template_4_your_project_name WHERE id = $1 AND _created_by = $2;`
	counttemplate4YourProjectName        = `SELECT COUNT(*) FROM template_4_your_project_name_db_schema.template_4_your_project_name `
	deletetemplate4YourProjectName       = `
UPDATE template_4_your_project_name_db_schema.template_4_your_project_name
SET
    _deleted = true,
    _deleted_by = $1,
    _deleted_at = CURRENT_TIMESTAMP
WHERE id = $2;`
	updatetemplate4YourProjectName = `
UPDATE template_4_your_project_name_db_schema.template_4_your_project_name SET
       type_id = $2,
       name = $3,
       description = $4,
       comment = $5,
       external_id = $6,
       external_ref = $7,
       build_at = $8,
       status = $9,
       contained_by = $10,
       contained_by_old = $11,
       inactivated = $12,
       inactivated_time = $13,
       inactivated_by = $14,
       inactivated_reason = $15,
       validated = $16,
       validated_time = $17,
       validated_by = $18,
       managed_by = $19,
       _last_modified_at = CURRENT_TIMESTAMP,
       _last_modified_by =$20,
       more_data =$21,
       position = ST_SetSRID(ST_MakePoint($22,$23), 2056),
       text_search = to_tsvector('french', unaccent($3) ||
                             ' ' || coalesce(unaccent($4), ' ') ||
                             ' ' || coalesce(unaccent($5), ' ') )
WHERE id = $1;
`

	baseGeoJsontemplate4YourProjectNameSearch = `
SELECT row_to_json(fc)
FROM (SELECT 'FeatureCollection'                         AS type,
             coalesce(array_to_json(array_agg(f)), '[]') AS features
      FROM (SELECT 'Feature'                             AS TYPE,
                   ST_AsGeoJSON(t.position, 6)::JSON     AS GEOMETRY,
                   row_to_json((SELECT l
                                FROM (SELECT id,
                                             type_id,
                                             name,
                                             description,
                                             external_id,
                                             inactivated,
                                             validated,
                                             status,
										     (SELECT icon_path FROM template_4_your_project_name_db_schema.type_template_4_your_project_name tt WHERE tt.id = t.type_id) as icon_path,
                                             _created_by    as created_by,
                                             _created_at    as created_at,
                                             st_x(position) as pos_x,
                                             st_y(position) as pos_y) AS l)) AS properties
            FROM template_4_your_project_name_db_schema.template_4_your_project_name t
            WHERE _deleted = false AND position IS NOT NULL
               
`
	geoJsonListEndOfQuery = `
        ORDER BY _created_at DESC
        LIMIT $1 OFFSET $2) AS f) AS fc
`
)
