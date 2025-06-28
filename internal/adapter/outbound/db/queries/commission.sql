-- name: CreateCommissionWork :one
INSERT INTO commission_works (
  requester_nickname,
  plot_id,
  task_type,
  task_description,
  credit_cost
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetCommissionWork :one
SELECT * FROM commission_works
WHERE id = $1;

-- name: ListCommissionWorksByRequester :many
SELECT * FROM commission_works
WHERE requester_nickname = $1
ORDER BY requested_at DESC;

-- name: ListCommissionWorksByStatus :many
SELECT * FROM commission_works
WHERE status = $1
ORDER BY requested_at DESC;

-- name: UpdateCommissionWorkStatus :one
UPDATE commission_works
SET status = $2
WHERE id = $1
RETURNING *;

-- name: DeleteCommissionWork :exec
DELETE FROM commission_works
WHERE id = $1; 