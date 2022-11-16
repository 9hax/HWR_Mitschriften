BEGIN TRANSACTION;
CREATE TABLE IF NOT EXISTS "Software" (
	"id"	INTEGER NOT NULL,
	"name"	TEXT,
	"currentDownloadURL"	TEXT,
	"currentVersion"	INTEGER,
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "Device" (
	"id"	INTEGER NOT NULL,
	"installDate"	TEXT,
	"hasSoftware"	INTEGER,
	"Vendor"	TEXT,
	PRIMARY KEY("id"),
	FOREIGN KEY("hasSoftware") REFERENCES "Software"("id")
);
CREATE TABLE IF NOT EXISTS "User" (
	"id"	INTEGER,
	"name"	TEXT NOT NULL,
	"tickets"	INTEGER,
	"devices"	INTEGER,
	PRIMARY KEY("id"),
	FOREIGN KEY("devices") REFERENCES "Device"("id"),
	FOREIGN KEY("tickets") REFERENCES "Ticket"("id")
);
CREATE TABLE IF NOT EXISTS "Ticket" (
	"id"	INTEGER,
	"creationDate"	TEXT,
	"details"	TEXT,
	"supportNote"	TEXT,
	"conceriningDevice"	INTEGER,
	PRIMARY KEY("id")
);
COMMIT;
