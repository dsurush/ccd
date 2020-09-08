package services

const getUserByIdDML = `Select id, name, surname, lastname, login, phone, role, status, position from users where id = ($1)`

const getUsersDML = `Select id, name, surname, lastname, login, phone, role, status, position from users`

const userSaveDML= `Insert into "users"(name, surname, lastname, login, password, phone, position) values($1, $2, $3, $4, $5, $6, $7)`

const editUserDML =  `Update users set name = ($1), surname = ($2), 
lastname = ($3), login = ($4), password = ($5), phone = ($6), position = ($7) where id = ($8)`

const editUserWithoutPassDML =  `Update users set name = ($1), surname = ($2), 
lastname = ($3), login = ($4), phone = ($5), position = ($6) where id = ($7)`

const setStateAndTimeDML = `Insert into "states" (user_id, work_time, status, unix_date, time_date) values($1, $2, $3, $4, $5)`

const editUserStateDML =  `Update users set status = ($1) where id = ($2)`

const getUserStatsDML = `Select *from states where user_id = ($1) and unix_date > ($2)`

const getUserStatsForAdminDML = `Select *from states where user_id = ($1) and unix_date >= ($2) and unix_date <= ($3)`

const geUsersStatsDML = `select us.name, us.surname, sum(st.work_time) as work_time, count(*) as count_interval
from users as us, states as st
where st.user_id = us.id and st.status = 'false' and unix_date >= ($1) and unix_date <= ($2)
group by us.name, us.surname`