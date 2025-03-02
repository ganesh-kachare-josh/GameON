package pkg 


// Request Queries
const GetRequestByIdQuery = "SELECT id , user_id , sport , location , time , court_price , status FROM requests WHERE id = $1"

const GetAllRequestsQuery = "SELECT requests.id, requests.user_id, requests.sport, requests.location, requests.time,requests.court_price, requests.status, users.name, users.email, users.phone_number FROM requests JOIN users ON requests.user_id = users.id"

const GetAllParticipantsQuery = "SELECT p.id, p.user_id, u.name , p.status FROM participants p JOIN users u ON p.user_id = u.id WHERE p.request_id = $1"

const AcceptRequestQuery = "INSERT INTO participants (request_id , user_id , status) VALUES($1 ,$2 ,$3) RETURNING *"

const ConfirmRequestQuery = "UPDATE participants SET status = 'Confirmed' WHERE request_id = $1 AND user_id = $2 RETURNING *"

const DeleteRequestQuery = "DELETE FROM requests WHERE id = $1"

const RejectParticipantQuery = "DELETE FROM participants WHERE id = $1"

const CreateRequestQuery = "INSERT INTO requests (user_id , sport , location , time , court_price , status , created_at) VALUES($1,$2,$3,$4,$5,'Open',NOW()) RETURNING id , user_id , sport , location , time , court_price , status"

const GetEmailById = "SELECT email from users WHERE id = $1"  

const GetUserIdByRequestId = "SELECT user_id FROM requests WHERE id = $1"

const GetNameByIdQuery = "SELECT name from users WHERE id = $1"

const GetSportByRequestId = "SELECT sport from requests WHERE id = $1"

const GetJoinedRequestByIdQuery = "SELECT request_id FROM participants WHERE user_id = $1"


// Profile Queries 
const GetUserByIdQuery = "SELECT id , name , email , sports , phone_number , created_at FROM users WHERE id = $1" 

const UpdateUserByIdQuery = "UPDATE users SET name = $2, email = $3, sports = $4, phone_number = $5 WHERE id = $1 RETURNING id, name, email, sports, phone_number"


// Authentication Queries 
const RegisterUserQuery = "INSERT INTO users (name , email , password , phone_number , sports , created_at) VALUES ($1 , $2 , $3 , $4 , $5 , NOW()) RETURNING id , name , email , password , phone_number , sports , created_at"

// Rating Queries 
const GiveRatingQuery = "INSERT into ratings(given_by , given_to , request_id , rating , feedback , created_at) VALUES ($1 , $2 , $3 , $4 , $5 , NOW())"

const GetRatingByUserIdQuery = "SELECT r.id, r.given_by, u.name, r.given_to, r.request_id, req.sport, r.rating, r.feedback, r.created_at FROM ratings r JOIN requests req ON r.request_id = req.id JOIN users u ON r.given_by = u.id WHERE r.given_to = $1"