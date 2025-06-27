-- name: CreateUserStats :one
INSERT INTO user_stats (nickname)
VALUES ($1)
RETURNING *;

-- name: UpdateUserExperience :one
UPDATE user_stats
SET experience = experience + $2,
    level = CASE 
        WHEN (experience + $2) >= level * 100 THEN level + 1
        ELSE level
    END,
    updated_at = CURRENT_TIMESTAMP
WHERE nickname = $1
RETURNING *;

-- name: UpdateUserRevenue :one
UPDATE user_stats
SET total_revenue = total_revenue + $2,
    updated_at = CURRENT_TIMESTAMP
WHERE nickname = $1
RETURNING *;

-- name: IncrementSuccessfulRaids :one
UPDATE user_stats
SET successful_raids = successful_raids + 1,
    updated_at = CURRENT_TIMESTAMP
WHERE nickname = $1
RETURNING *;

-- name: GetLeaderboard :many
SELECT nickname, level, experience, total_revenue, successful_raids
FROM user_stats
ORDER BY total_revenue DESC, level DESC
LIMIT 10;

-- name: GetUserStats :one
SELECT * FROM user_stats
WHERE nickname = $1; 