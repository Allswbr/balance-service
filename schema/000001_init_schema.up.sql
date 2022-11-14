CREATE TYPE TRANSENUM AS ENUM ('ADD', 'RESERVE', 'TAKE', 'FREE', 'TRANSFER');

CREATE TABLE IF NOT EXISTS users
(
    user_id          INTEGER 		 PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    deposit_account  NUMERIC 		 NOT NULL,
    reserve_account  NUMERIC 		 NOT NULL,
    CONSTRAINT       account_check CHECK (deposit_account >= 0 AND
                                          reserve_account >= 0)
);

CREATE TABLE IF NOT EXISTS transaction
(
    trans_id      INTEGER   	  PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    user_id       INTEGER   	  NOT NULL,
    service_id    INTEGER   	  NULL,
    order_id      INTEGER   	  NULL,
    datetime      TIMESTAMP 	  NOT NULL,
    amount        NUMERIC   	  NOT NULL,
    start_deposit NUMERIC   	  NOT NULL,
    end_deposit   NUMERIC   	  NOT NULL,
    event         TRANSENUM 	  NOT NULL,
    message       TEXT      	  NULL,
    FOREIGN KEY   (user_id) 	  REFERENCES users (user_id),
    CONSTRAINT    deposit_check CHECK (start_deposit >= 0 AND
                                       end_deposit >= 0 AND
                                       amount >= 0)
);