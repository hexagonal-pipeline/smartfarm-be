-- name: CreateRaid :one
INSERT INTO raids (title, description, crop_type, target_quantity, min_participation, max_participation, price_per_kg, deadline, creator_nickname)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING *;

-- name: ListOpenRaids :many
SELECT r.*,
       COALESCE(SUM(rp.quantity), 0) as current_quantity,
       COUNT(rp.id) as participant_count
FROM raids r
LEFT JOIN raid_participations rp ON r.id = rp.raid_id
WHERE r.status = 'open' AND r.deadline > NOW()
GROUP BY r.id
ORDER BY r.deadline ASC;

-- name: GetRaidDetails :one
SELECT r.*,
       COALESCE(SUM(rp.quantity), 0) as current_quantity,
       COUNT(rp.id) as participant_count
FROM raids r
LEFT JOIN raid_participations rp ON r.id = rp.raid_id
WHERE r.id = $1
GROUP BY r.id;

-- name: JoinRaid :one
INSERT INTO raid_participations (raid_id, participant_nickname, quantity, expected_revenue)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetRaidParticipations :many
SELECT rp.*, r.title as raid_title
FROM raid_participations rp
JOIN raids r ON rp.raid_id = r.id
WHERE rp.raid_id = $1
ORDER BY rp.created_at DESC;

-- name: GetNicknameRaidHistory :many
SELECT rp.*, r.title, r.crop_type, r.deadline, r.status as raid_status
FROM raid_participations rp
JOIN raids r ON rp.raid_id = r.id
WHERE rp.participant_nickname = $1
ORDER BY rp.created_at DESC;

-- name: UpdateRaidStatus :one
UPDATE raids
SET status = $2
WHERE id = $1
RETURNING *; 