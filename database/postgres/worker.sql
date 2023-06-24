-- Создание роли worker_role
CREATE ROLE worker_role;

-- Предоставление разрешений на запись в таблицы user_order и item
GRANT INSERT,SELECT ON TABLE user_order TO worker_role;
-- this privilege allows the use of the currval and nextval functions
GRANT USAGE ON ALL SEQUENCES IN SCHEMA public TO worker_role;
GRANT INSERT ON TABLE item TO worker_role;

-- Создание пользователя worker
CREATE USER worker WITH PASSWORD 'worker';

-- Присвоение роли worker_role пользователю worker
GRANT worker_role TO worker;
