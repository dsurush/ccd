package services

const getUserByIdDML = `Select id, name, surname, lastname, login, phone, role from users where id = ($1)`

const getUsersDML = `Select id, name, surname, lastname, login, phone, role from users`