-- name: GetSetting :one
SELECT id, category, "key", value, created_at, updated_at
FROM system_settings
WHERE category = ? AND "key" = ?;

-- name: GetSettingsByCategory :many
SELECT id, category, "key", value, created_at, updated_at
FROM system_settings
WHERE category = ?
ORDER BY "key";

-- name: UpsertSetting :exec
INSERT INTO system_settings (category, "key", value, updated_at)
VALUES (?, ?, ?, CURRENT_TIMESTAMP)
ON CONFLICT (category, "key") DO UPDATE SET value = excluded.value, updated_at = CURRENT_TIMESTAMP;

-- name: DeleteSetting :exec
DELETE FROM system_settings WHERE category = ? AND "key" = ?;