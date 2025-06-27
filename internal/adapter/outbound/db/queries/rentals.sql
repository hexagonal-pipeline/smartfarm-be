-- name: CreateRental :one
INSERT INTO rentals (renter_nickname, plot_id, start_date, end_date, monthly_rent)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetNicknameRentals :many
SELECT r.*, fp.name as plot_name, fp.location, fp.crop_type, fp.size_sqm
FROM rentals r
JOIN farm_plots fp ON r.plot_id = fp.id
WHERE r.renter_nickname = $1 AND r.status = 'active'
ORDER BY r.start_date DESC;

-- name: GetRental :one
SELECT r.*, fp.name as plot_name, fp.location, fp.crop_type, fp.size_sqm
FROM rentals r
JOIN farm_plots fp ON r.plot_id = fp.id
WHERE r.id = $1; 