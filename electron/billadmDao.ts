import EasyDB from "./easyDB";
import {getCurrentUTCTime, UUID} from "./utils";
import {BUILT_IN_BILLBOOK, BUILT_IN_TYPES} from "./constants";


/**
 * 创建表格
 * creation_time: 至UTC时间的毫秒数
 */
const CREATE_TABLE_BILLBOOKS: string = `CREATE TABLE IF NOT EXISTS t_billbooks (
    id            TEXT PRIMARY KEY NOT NULL,
    name          TEXT NOT NULL,
    description   TEXT DEFAULT '',
    creation_time INTEGER NOT NULL
)`;

const CREATE_TABLE_BILLS: string = `CREATE TABLE IF NOT EXISTS t_bills (
    id            TEXT PRIMARY KEY NOT NULL,
    money         REAL NOT NULL,
    type          TEXT NOT NULL,
    income        TEXT NOT NULL DEFAULT 'false',
    book_id       TEXT NOT NULL,
    description   TEXT DEFAULT '',
    tags          TEXT NOT NULL DEFAULT '[]',
    creation_time INTEGER NOT NULL
)`;

const CREATE_TABLE_TYPES: string = `CREATE TABLE IF NOT EXISTS t_types (
    id   TEXT PRIMARY KEY NOT NULL,
    name TEXT NOT NULL
)`;

/** t_types*/
const INSERT_ONE_TYPE: string = 'INSERT INTO t_types (id,name) VALUES (?,?)';
const QUERY_ALL_TYPE: string = 'SELECT * FROM t_types';
const DELETE_ONE_TYPE_BY_ID: string = 'DELETE FROM t_types WHERE id=?';

/** t_billbooks*/
const INSERT_ONE_BILLBOOK: string = 'INSERT INTO t_billbooks (id,name,description,creation_time) VALUES (?,?,?,?)';
const QUERY_ALL_BILLBOOK: string = 'SELECT * FROM t_billbooks';
const DELETE_ONE_BILLBOOK_BY_ID: string = 'DELETE FROM t_billbooks WHERE id=?';

/** t_bills*/
const INSERT_ONE_BILL: string = 'INSERT INTO t_bills (id,money,type,income,book_id,description,tags,creation_time) VALUES (?,?,?,?,?,?,?,?);';
const QUERY_ALL_BILL_BY_BOOK_ID: string = 'SELECT * FROM t_bills WHERE book_id=?';
const DELETE_ONE_BILL_BY_ID: string = 'DELETE FROM t_bills WHERE id=?';

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
        await this.easyDB.runSql(INSERT_ONE_BILLBOOK, [BUILT_IN_BILLBOOK.id, BUILT_IN_BILLBOOK.name, BUILT_IN_BILLBOOK.description, getCurrentUTCTime()]);
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
    async deleteOneTypeByID(id: string) {
        await this.easyDB.runSql(DELETE_ONE_TYPE_BY_ID, [id]);
    }

    /** 创建一个账本*/
    async insertOneBillbook(name: string, description: string) {
        await this.easyDB.runSql(INSERT_ONE_BILLBOOK, [UUID(), name, description, getCurrentUTCTime()]);
    }

    /** 查询所有账本*/
    async queryAllBillbook() {
        return await this.easyDB.querySql(QUERY_ALL_BILLBOOK);
    }

    /** 删除一个账本*/
    async deleteOneBillbookByID(id: string) {
        await this.easyDB.runSql(DELETE_ONE_BILLBOOK_BY_ID, [id]);
    }

    /** 创建一条消费记录*/
    async insertOneBill(money: number, type: string, income: string, book_id: string, description: string, tags: string) {
        await this.easyDB.runSql(INSERT_ONE_BILL, [UUID(), money, type, income, book_id, description, tags, getCurrentUTCTime()]);
    }

    /** 查询一个账本中的所有消费记录*/
    async queryAllBillByBookID(book_id: string) {
        console.log(book_id);
        return await this.easyDB.querySql(QUERY_ALL_BILL_BY_BOOK_ID, [book_id]);
    }

    /** 删除一条消费记录*/
    async deleteOneBillByID(id: string) {
        await this.easyDB.runSql(DELETE_ONE_BILL_BY_ID, [id]);
    }
}

export default BilladmDao;