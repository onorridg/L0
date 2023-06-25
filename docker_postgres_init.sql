---------------------------------SERVICE_L0----------------------------------------------
CREATE TABLE user_order (
                            id BIGSERIAL PRIMARY KEY,
                            order_uid VARCHAR(255) DEFAULT '',
                            track_number VARCHAR(255) UNIQUE DEFAULT '',
                            entry VARCHAR(255) DEFAULT '',
                            locale VARCHAR(255) DEFAULT '',
                            internal_signature VARCHAR(255) DEFAULT '',
                            customer_id VARCHAR(255) DEFAULT '',
                            delivery_service VARCHAR(255) DEFAULT '',
                            shardkey VARCHAR(255) DEFAULT '',
                            sm_id INT DEFAULT 0,
                            date_created TIMESTAMP DEFAULT 'epoch',
                            oof_shard INT DEFAULT 0,
                            name VARCHAR(255) DEFAULT '',
                            phone VARCHAR(255) DEFAULT '',
                            zip VARCHAR(255) DEFAULT '',
                            city VARCHAR(255) DEFAULT '',
                            address VARCHAR(255) DEFAULT '',
                            region VARCHAR(255) DEFAULT '',
                            email VARCHAR(255) DEFAULT '',
                            transaction VARCHAR(255) DEFAULT '',
                            request_id VARCHAR(255) DEFAULT '',
                            currency VARCHAR(255) DEFAULT '',
                            provider VARCHAR(255) DEFAULT '',
                            amount FLOAT DEFAULT 0.0,
                            payment_dt BIGINT DEFAULT 0,
                            bank VARCHAR(255) DEFAULT '',
                            delivery_cost FLOAT DEFAULT 0.0,
                            goods_total FLOAT DEFAULT 0.0,
                            custom_fee INT DEFAULT 0

);

CREATE TABLE public.item (
                             id BIGSERIAL PRIMARY KEY,
                             user_order_id BIGSERIAL,
                             chrt_id INT DEFAULT 0,
                             track_number VARCHAR(255) DEFAULT '',
                             price FLOAT DEFAULT 0.0,
                             rid VARCHAR(255) DEFAULT '',
                             name VARCHAR(255) DEFAULT '',
                             sale FLOAT DEFAULT 0.0,
                             size FLOAT DEFAULT 0.0,
                             total_price FLOAT DEFAULT 0.0,
                             nm_id INT DEFAULT 0,
                             brand VARCHAR(255) DEFAULT '',
                             status INT DEFAULT 0,
                             FOREIGN KEY (user_order_id) REFERENCES user_order (id)
);

-- Создание роли worker_role
CREATE ROLE worker_role;

-- Предоставление разрешений на запись в таблицы user_order и item
GRANT INSERT,SELECT ON TABLE user_order TO worker_role;
GRANT INSERT,SELECT ON TABLE item TO worker_role;

-- This privilege allows the use of the currval and nextval functions
GRANT USAGE ON ALL SEQUENCES IN SCHEMA public TO worker_role;

-- Создание пользователя worker
CREATE USER worker WITH PASSWORD 'worker';

-- Присвоение роли worker_role пользователю worker
GRANT worker_role TO worker;


---------------------------NATS_STREAMING----------------------------------------------
CREATE TABLE IF NOT EXISTS ServerInfo (uniquerow INTEGER DEFAULT 1, id VARCHAR(1024), proto BYTEA, version INTEGER, PRIMARY KEY (uniquerow));
CREATE TABLE IF NOT EXISTS Clients (id VARCHAR(1024), hbinbox TEXT, PRIMARY KEY (id));
CREATE TABLE IF NOT EXISTS Channels (id INTEGER, name VARCHAR(1024) NOT NULL, maxseq BIGINT DEFAULT 0, maxmsgs INTEGER DEFAULT 0, maxbytes BIGINT DEFAULT 0, maxage BIGINT DEFAULT 0, deleted BOOL DEFAULT FALSE, PRIMARY KEY (id));
CREATE INDEX Idx_ChannelsName ON Channels (name(256));
CREATE TABLE IF NOT EXISTS Messages (id INTEGER, seq BIGINT, timestamp BIGINT, size INTEGER, data BYTEA, CONSTRAINT PK_MsgKey PRIMARY KEY(id, seq));
CREATE INDEX Idx_MsgsTimestamp ON Messages (timestamp);
CREATE TABLE IF NOT EXISTS Subscriptions (id INTEGER, subid BIGINT, lastsent BIGINT DEFAULT 0, proto BYTEA, deleted BOOL DEFAULT FALSE, CONSTRAINT PK_SubKey PRIMARY KEY(id, subid));
CREATE TABLE IF NOT EXISTS SubsPending (subid BIGINT, row BIGINT, seq BIGINT DEFAULT 0, lastsent BIGINT DEFAULT 0, pending BYTEA, acks BYTEA, CONSTRAINT PK_MsgPendingKey PRIMARY KEY(subid, row));
CREATE INDEX Idx_SubsPendingSeq ON SubsPending (seq);
CREATE TABLE IF NOT EXISTS StoreLock (id VARCHAR(30), tick BIGINT DEFAULT 0);

-- Updates for 0.10.0
ALTER TABLE Clients ADD proto BYTEA;

CREATE ROLE ns_role;
GRANT ALL PRIVILEGES ON TABLE ServerInfo,Clients,Channels,Messages,Subscriptions,SubsPending,StoreLock TO ns_role;
CREATE USER ns_user WITH PASSWORD 'ns_user';

GRANT ns_role TO ns_user;
