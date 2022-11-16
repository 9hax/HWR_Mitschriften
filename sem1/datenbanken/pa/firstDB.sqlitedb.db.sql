BEGIN TRANSACTION;
CREATE TABLE IF NOT EXISTS "Software" (
	"id"					INTEGER,
	"name"					TEXT,
	"currentDownloadURL"	TEXT,
	"currentVersion"		INTEGER,
);
CREATE TABLE IF NOT EXISTS "Device" (
	"id"			INTEGER,
	"installDate"	TEXT,
	"hasSoftware"	INTEGER,
	"Vendor"		TEXT,
	FOREIGN KEY("hasSoftware") REFERENCES "Software"("id")
);
CREATE TABLE IF NOT EXISTS "User" (
	"id"		INTEGER,
	"name"		TEXT,
	"tickets"	INTEGER,
	"devices"	INTEGER,
	FOREIGN KEY("devices") REFERENCES "Device"("id"),
	FOREIGN KEY("tickets") REFERENCES "Ticket"("id")
);
CREATE TABLE IF NOT EXISTS "Ticket" (
	"id"				INTEGER,
	"creationDate"		TEXT,
	"details"			TEXT,
	"supportNote"		TEXT,
	"conceriningDevice"	INTEGER,
	FOREIGN KEY("conceriningDevice") REFERENCES "Device" ("id")
);
COMMIT;
