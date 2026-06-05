-- name: GetGroupByID :one
SELECT id, name, description, created_at, updated_at
FROM `groups`
WHERE id = ?;

-- name: ListGroups :many
SELECT id, name, description, created_at, updated_at
FROM `groups`
ORDER BY id;

-- name: CreateGroup :execresult
INSERT INTO `groups` (name, description)
VALUES (?, ?);

-- name: UpdateGroup :exec
UPDATE `groups`
SET name = ?, description = ?, updated_at = CURRENT_TIMESTAMP
WHERE id = ?;

-- name: DeleteGroup :exec
DELETE FROM `groups` WHERE id = ?;

-- name: AddUserToGroup :exec
INSERT INTO user_groups (user_id, group_id)
VALUES (?, ?);

-- name: RemoveUserFromGroup :exec
DELETE FROM user_groups WHERE user_id = ? AND group_id = ?;

-- name: GetUserGroups :many
SELECT g.id, g.name, g.description, g.created_at, g.updated_at
FROM `groups` g
JOIN user_groups ug ON ug.group_id = g.id
WHERE ug.user_id = ?;

-- name: AddGroupRole :exec
INSERT INTO group_roles (group_id, role_id)
VALUES (?, ?);

-- name: RemoveGroupRole :exec
DELETE FROM group_roles WHERE group_id = ? AND role_id = ?;

-- name: GetGroupRoles :many
SELECT r.id, r.name, r.description, r.created_at, r.updated_at
FROM roles r
JOIN group_roles gr ON gr.role_id = r.id
WHERE gr.group_id = ?;
