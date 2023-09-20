package config

var (
	createDriverTable = `CREATE TABLE IF NOT EXISTS driver(
		id BIGSERIAL PRIMARY KEY,
		phone varchar(50),
		firstName varchar(50),
		lastName varchar(50),
		inn varchar(50),
		ava varchar(200),
		carNumber varchar(200),
		carColor varchar(200),
		carModel varchar(200),
		docsfront varchar(200),
		docsback varchar(200),
		carType varchar(50),
		token varchar(500)
	);`
	createUserTable = `CREATE TABLE IF NOT EXISTS customer(
		id BIGSERIAL PRIMARY KEY,
		phone varchar(50),
		firstName varchar(50),
		lastName varchar(50),
		ava varchar(200),
		token varchar(500)
	);`
	createHistoryTable = `CREATE TABLE IF NOT EXISTS history(
		id BIGSERIAL PRIMARY KEY,
		orderId BIGINT,
		driverId BIGINT,
		userId BIGINT,
		startDate varchar(100),
		finishedDate varchar(100)
	);`
	createOfferDriverTable = `CREATE TABLE IF NOT EXISTS offer_driver(
		id BIGSERIAL PRIMARY KEY,
		comment varchar(500),
		locationFrom varchar(100),
		locationTo varchar(100),
		price INTEGER, 
		type varchar(100),
		driver BIGINT
	);`
	createOfferUserTable = `CREATE TABLE IF NOT EXISTS offer_user(
		id BIGSERIAL PRIMARY KEY,
		comment VARCHAR(500),
		locationFrom VARCHAR(100),
		locationTo VARCHAR(100),
		price INTEGER,
		type VARCHAR(100),
		customer BIGINT
	);`
	createOrderProcessTable = `CREATE TABLE IF NOT EXISTS order_process(
		id BIGSERIAL PRIMARY KEY,
		userId BIGINT,
		latitudeFrom varchar(100),
		longitudeFrom varchar(100),
		latitudeTo varchar(100),
		longitudeTo varchar(100),
		comments varchar(500),
		price INTEGER,
		type varchar(100),
		orderStatus INTEGER
	);`
	createTableSMS = `CREATE TABLE IF NOT EXISTS sms_cache(
		id BIGSERIAL PRIMARY KEY,
		contact varchar(50),
		code INTEGER
	);`

	createTableSecurity = `CREATE TABLE IF NOT EXISTS security(
		id BIGSERIAL PRIMARY KEY,
		userId       BIGINT,
		firstName    varchar(50),
		lastName     varchar(50), 
		A            varchar(50),
		B            varchar(50),
		fiod         varchar(50),
		phone        varchar(50), 
		carNumber    varchar(200),
		timeStart    varchar(200),
		timeFinish   varchar(200),
		status       boolean
	);`
)
