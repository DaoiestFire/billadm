import EasyDB from "./easyDB";
import {getCurrentUTCTimeString, UUID} from "./utils";
import {BUILT_IN_BILLBOOK, BUILT_IN_TYPES} from "./constants";


/**
 * 创建表格
 * creation_time: YYYY-MM-DD hh:mm:ss 格式的UTC时间字符串
 */
const CREATE_TABLE_BILLBOOKS: string = `CREATE TABLE IF NOT EXISTS t_billbooks (
    id            TEXT PRIMARY KEY NOT NULL,
    name          TEXT NOT NULL,
    description   TEXT DEFAULT '',
    creation_time TEXT NOT NULL
)`;

const CREATE_TABLE_BILLS: string = `CREATE TABLE IF NOT EXISTS t_bills (
    id            TEXT PRIMARY KEY NOT NULL,
    money         REAL NOT NULL,
    type          TEXT NOT NULL,
    income        TEXT NOT NULL DEFAULT 'false',
    book_id       TEXT NOT NULL,
    description   TEXT DEFAULT '',
    tags          TEXT NOT NULL DEFAULT '[]',
    creation_time TEXT NOT NULL
)`;

const CREATE_TABLE_TYPES: string = `CREATE TABLE IF NOT EXISTS t_types (
    id   TEXT PRIMARY KEY NOT NULL,
    name TEXT NOT NULL
)`;

const CREATE_TABLE_BILLADM: string = `CREATE TABLE IF NOT EXISTS t_billadm (
    version       TEXT NOT NULL,
    creation_time TEXT NOT NULL
)`;

/** t_billadm*/
const INSERT_BILLADM_INFO: string = 'INSERT INTO t_billadm (version,creation_time) VALUES (?,?)';

/** t_types*/
const INSERT_ONE_TYPE: string = 'INSERT INTO t_types (id,name) VALUES (?,?)';
const QUERY_ALL_TYPE: string = 'SELECT * FROM t_types';
const DELETE_ONE_TYPE_BY_ID: string = 'DELETE FROM t_types WHERE id=?';

/** t_billbooks*/
const INSERT_ONE_BILLBOOK: string = 'INSERT INTO t_billbooks (id,name,description,creation_time) VALUES (?,?,?,?)';
const QUERY_ALL_BILLBOOK: string = 'SELECT * FROM t_billbooks';
const DELETE_ONE_BILLBOOK_BY_ID: string = 'DELETE FROM t_billbooks WHERE id=?';

/** t_bills*/
const INSERT_ONE_BILL: string = 'INSERT INTO t_bills (id,money,type,income,book_id,description,tags,creation_time) VALUES (?,?,?,?,?,?,?,?)';
const UPDATE_ONE_BILL: string = 'UPDATE t_bills set money=?,type=?,income=?,description=?,tags=?,creation_time=? WHERE id=?';
const QUERY_ALL_BILL_BY_BOOK_ID: string = 'SELECT * FROM t_bills WHERE book_id=?';
const DELETE_ONE_BILL_BY_ID: string = 'DELETE FROM t_bills WHERE id=?';
const DELETE_BILLS_BY_BOOK_ID: string = 'DELETE FROM t_bills WHERE book_id=?';

class BilladmDao {
    private easyDB: EasyDB;

    constructor(dbFile: string, logFile: string) {
        this.easyDB = new EasyDB(dbFile, logFile);
    }

    init() {
        return this.easyDB.init();
    }

    /** 初始化数据库，创建表格*/
    async initDB(version: string) {
        await this.easyDB.runSql(CREATE_TABLE_BILLADM);
        await this.easyDB.runSql(CREATE_TABLE_BILLBOOKS);
        await this.easyDB.runSql(CREATE_TABLE_BILLS);
        await this.easyDB.runSql(CREATE_TABLE_TYPES);
        await this.easyDB.runSql(INSERT_BILLADM_INFO, [version, getCurrentUTCTimeString()]);
        await this.easyDB.runSql(INSERT_ONE_BILLBOOK, [BUILT_IN_BILLBOOK.id, BUILT_IN_BILLBOOK.name, BUILT_IN_BILLBOOK.description, getCurrentUTCTimeString()]);
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
        await this.easyDB.runSql(INSERT_ONE_BILLBOOK, [UUID(), name, description, getCurrentUTCTimeString()]);
    }

    /** 查询所有账本*/
    async queryAllBillbook() {
        return await this.easyDB.querySql(QUERY_ALL_BILLBOOK);
    }

    /** 删除一个账本*/
    async deleteOneBillbookByID(id: string) {
        await this.easyDB.runSql(DELETE_BILLS_BY_BOOK_ID, [id]);
        await this.easyDB.runSql(DELETE_ONE_BILLBOOK_BY_ID, [id]);
    }

    /** 创建一条消费记录*/
    async insertOneBill(money: number, type: string, income: string, bookId: string, description: string, tags: string, creationTime: string) {
        await this.easyDB.runSql(INSERT_ONE_BILL, [UUID(), money, type, income, bookId, description, tags, creationTime]);
    }

    /** 更新一条消费记录*/
    async updateOneBill(id: string, money: number, type: string, income: string, description: string, tags: string, creationTime: string) {
        await this.easyDB.runSql(UPDATE_ONE_BILL, [money, type, income, description, tags, creationTime, id]);
    }

    /** 查询一个账本中的所有消费记录*/
    async queryAllBillByBookID(bookId: string) {
        return await this.easyDB.querySql(QUERY_ALL_BILL_BY_BOOK_ID, [bookId]);
    }

    /** 查询一个账本中的过滤后的消费记录*/
    async queryAllBillByBookIDWithFilters(bookId: string, filters: object) {
        let newSql: string;
        let values: any[] = [];
        values.push(bookId);
        if ('start_time' in filters && filters.start_time && 'end_time' in filters && filters.end_time) {
            newSql = QUERY_ALL_BILL_BY_BOOK_ID + " AND datetime(creation_time) BETWEEN datetime(?) AND datetime(?)";
            values.push(filters.start_time);
            values.push(filters.end_time);
        }
        return await this.easyDB.querySql(newSql, values);
    }

    /** 删除一条消费记录*/
    async deleteOneBillByID(id: string) {
        await this.easyDB.runSql(DELETE_ONE_BILL_BY_ID, [id]);
    }
}

export default BilladmDao;