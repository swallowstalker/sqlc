-- name: SelectFundSummary :many
SELECT name,
       extent
FROM auth_user
WHERE user_type = 'manager'
ORDER BY name;

-- name: GetFundSummary :one
SELECT id,
       name,
       extent
FROM auth_user
WHERE id = $1
    LIMIT 1;