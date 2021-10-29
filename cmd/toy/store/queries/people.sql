-- name: ListPeople :many
SELECT * FROM people;

-- name: CreatePerson :exec
INSERT INTO people (id,name,email,pb_data)
VALUES ($1,$2,$3,$4);