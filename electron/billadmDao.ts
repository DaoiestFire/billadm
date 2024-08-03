import EasyDB from "./easyDB";
import {UUID} from "./utils";
import {BUILT_IN_TYPES} from "./constans";

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

/** t_types*/
const INSERT_ONE_TYPE: string = 'INSERT INTO t_types (id,name) VALUES (?,?);'
const QUERY_ALL_TYPE: string = 'SELECT * FROM t_types';
const DELETE_ONE_TYPE_BY_ID: string = 'DELETE FROM t_types WHERE id=?';

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
        for (let value of BUILT_IN_TYPES) {
            await this.insertOneType(value);
        }
    }

    /** 插入一条消费类型*/
    async insertOneType(typeName: string) {
        await this.easyDB.runSql(INSERT_ONE_TYPE, [UUID(), typeName]);
    }

    /** 查询所有消费类型*/
    async queryAllType() {
        return await this.easyDB.querySql(QUERY_ALL_TYPE);
    }

    /** 删除一条消费类型*/
    async deleteOneType(id: string) {
        await this.easyDB.runSql(DELETE_ONE_TYPE_BY_ID, [id]);
    }
}

export default BilladmDao;