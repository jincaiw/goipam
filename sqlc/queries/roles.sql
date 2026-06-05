-- name: GetRoleByID :one
SELECT id, name, description, created_at, updated_at
FROM roles
WHERE id = ?;

-- name: ListRoles :many
SELECT id, name, description, created_at, updated_at
FROM roles
ORDER BY id;

-- name: CreateRole :execresult
INSERT INTO roles (name, description)
VALUES (?, ?);

-- name: UpdateRole :exec
UPDATE roles
SET name = ?, description = ?, updated_at = CURRENT_TIMESTAMP
WHERE id = ?;

-- name: DeleteRole :exec
DELETE FROM roles WHERE id = ?;

-- name: GetPermissionByCode :one
SELECT id, code, description, created_at
FROM permissions
WHERE code = ?;

-- name: ListPermissions :many
SELECT id, code, description, created_at
FROM permissions
ORDER BY id;

-- name: CreatePermission :execresult
INSERT INTO permissions (code, description)
VALUES (?, ?);

-- name: GetRolePermissions :many
SELECT p.id, p.code, p.description, p.created_at
FROM permissions p
JOIN role_permissions rp ON rp.permission_id = p.id
WHERE rp.role_id = ?;

-- name: AddRolePermission :exec
INSERT INTO role_permissions (role_id, permission_id)
VALUES (?, ?);

-- name: RemoveRolePermission :exec
DELETE FROM role_permissions WHERE role_id = ? AND permission_id = ?;

-- name: GetUserRoles :many
SELECT r.id, r.name, r.description, r.created_at, r.updated_at
FROM roles r
JOIN user_roles ur ON ur.role_id = r.id
WHERE ur.user_id = ?;

-- name: AddUserRole :exec
INSERT INTO user_roles (user_id, role_id)
VALUES (?, ?);

-- name: RemoveUserRole :exec
DELETE FROM user_roles WHERE user_id = ? AND role_id = ?;
