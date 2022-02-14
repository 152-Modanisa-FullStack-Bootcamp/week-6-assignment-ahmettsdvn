-------------------------------------------------------ORACLE-------------------------------------------------------
create table Wallet(
    id int generated always as IDENTITY not null,
    username varchar(50),
    balanceAmount int,
    CONSTRAINT pk_wallet PRIMARY KEY (id)
);

INSERT INTO
    Wallet (username, balanceamount)
values
('ahmettsdvn', 9000);

INSERT INTO
    Wallet (username, balanceamount)
values
('gophers', 19000);

INSERT INTO
    Wallet (username, balanceamount)
values
('java', 22333);

INSERT INTO
    Wallet (username, balanceamount)
values
('csharp', -20);

-------------------------------------------------------SQLITE3-------------------------------------------------------
CREATE TABLE "Wallet" (
	"id"	INTEGER NOT NULL,
	"username"	TEXT,
	"balanceAmount"	INTEGER,
	PRIMARY KEY("id" AUTOINCREMENT)
)

INSERT INTO
    Wallet (username, balanceamount)
values
('ahmettsdvn', 9000);

INSERT INTO
    Wallet (username, balanceamount)
values
('gophers', 19000);

INSERT INTO
    Wallet (username, balanceamount)
values
('java', 22333);

INSERT INTO
    Wallet (username, balanceamount)
values
('csharp', -20);