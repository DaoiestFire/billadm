import EasyDB from "./easyDB";

const CREATE_TABLE_BILLBOOKS: string = `CREATE TABLE IF NOT EXISTS t_billbooks (
    id            TEXT PRIMARY KEY NOT NULL,
    name          TEXT NOT NULL,
    description   TEXT DEFAULT '',
    creation_time TEXT NOT NULL
);`;

const CREATE_TABLE_BILLS: string = `CREATE TABLE IF NOT EXISTS t_bills (
    id            TEXT PRIMARY KEY NOT NULL,
    money         REAL NOT NULL,
    type          TEXT NOT NULL,
    income        TEXT NOT NULL DEFAULT 'false',
    book_id       TEXT NOT NULL,
    description   TEXT DEFAULT '',
    tags          TEXT NOT NULL DEFAULT '[]',
    creation_time TEXT NOT NULL
);`;

const CREATE_TABLE_TYPES: string = `CREATE TABLE IF NOT EXISTS t_types (
    id   TEXT PRIMARY KEY NOT NULL,
    name TEXT NOT NULL
);`;

class BilladmDao {
    private easyDB: EasyDB;

    constructor(dbFile: string, logFile: string) {
        this.easyDB = new EasyDB(dbFile, logFile);
    }

    init() {
        return this.easyDB.init();
    }

    /** 初始化数据库，创建表格*/
    async initDB() {
        await this.easyDB.runSql(CREATE_TABLE_BILLBOOKS);
        await this.easyDB.runSql(CREATE_TABLE_BILLS);
        await this.easyDB.runSql(CREATE_TABLE_TYPES);
    }
}

export default BilladmDao;