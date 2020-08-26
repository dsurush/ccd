package services

const getUserByIdDML = `Select id, name, surname, lastname, login, phone, role, status from users where id = ($1)`

const getUsersDML = `Select id, name, surname, lastname, login, phone, role, status from users`

const userSaveDML= `Insert into "users"(name, surname, lastname, login, password, phone) values($1, $2, $3, $4, $5, $6)`
