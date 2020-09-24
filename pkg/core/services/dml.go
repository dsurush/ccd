package services

const getUserByIdDML = `Select id, name, surname, lastname, login, phone, role, status, position, status_line from users where id = ($1)`

const getUsersDML = `Select id, name, surname, lastname, login, phone, role, status, position, status_line from users`
//get Users
const getUsersWithWorkTimeDML = `SELECT us.id,
    us.name,
    us.surname,
    us.lastname,
    us.login,
    us.phone,
    us.role,
    us.status,
    us."position",
    us.status_line,
    ( SELECT COALESCE(sum(st.work_time), 0::bigint) AS sum
           FROM states st
          WHERE st.user_id = us.id AND st.status = false AND st.unix_date > ($1)) AS worked,
    ( SELECT COALESCE(sum(states.work_time), 0::bigint) AS sum
           FROM states
          WHERE states.user_id = us.id AND states.status = true AND states.unix_date > ($2)) AS rest
   FROM users us`

const userSaveDML= `Insert into "users"(name, surname, lastname, login, password, phone, position) values($1, $2, $3, $4, $5, $6, $7)`

const editUserDML =  `Update users set name = ($1), surname = ($2), 
lastname = ($3), login = ($4), password = ($5), phone = ($6), position = ($7) where id = ($8)`

const editUserWithoutPassDML =  `Update users set name = ($1), surname = ($2), 
lastname = ($3), login = ($4), phone = ($5), position = ($6) where id = ($7)`

const setStateAndTimeDML = `Insert into "states" (user_id, work_time, status, unix_date, time_date) values($1, $2, $3, $4, $5)`

const editUserStateDML =  `Update users set status = ($1) where id = ($2)`

const editUserStatusLineDML =  `Update users set status_line = ($1) where login = ($2)`

const editUserStatusLineByIdDML =  `Update users set status_line = ($1) where id = ($2)`

const editUserStatusByIdDML =  `Update users set status = ($1) where id = ($2)`


const getUserStatsDML = `Select *from states where user_id = ($1) and unix_date > ($2)`

const getUserStatsForAdminDML = `Select *from states where user_id = ($1) and unix_date >= ($2) and unix_date <= ($3)`

const geUsersStatsDML = `select us.name, us.surname, sum(st.work_time) as work_time, count(*) as count_interval
from users as us, states as st
where st.user_id = us.id and st.status = 'false' and unix_date >= ($1) and unix_date <= ($2)
group by us.name, us.surname`

const getUserPassByIdDML = `Select password from users where id = ($1)`

const setUserPassByIdDML =  `Update users set password = ($1) where id = ($2)`