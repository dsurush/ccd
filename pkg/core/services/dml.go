package services

const getUserByIdDML = `Select id, name, surname, lastname, login, phone, role, status from users where id = ($1)`

const getUsersDML = `Select id, name, surname, lastname, login, phone, role, status from users`

const userSaveDML= `Insert into "users"(name, surname, lastname, login, password, phone) values($1, $2, $3, $4, $5, $6)`

const editUserDML =  `Update users set name = ($1), surname = ($2), 
lastname = ($3), login = ($4), password = ($5), phone = ($6) where id = ($7)`