CREATE USER worker WITH PASSWORD 'password'; -- create user
GRANT SELECT, INSERT ON service.public.user_order, service.public. TO read_write_user;
; -- granting write access to the write database
