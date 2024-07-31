import EasyDB from "../electron/easyDB";
import * as path from "node:path";
import * as fs from "node:fs";

const DB_FILE = path.join(__dirname, 'test.db');
const DB_LOG_FILE = path.join(__dirname, 'test.db.log');

if (fs.existsSync(DB_FILE)) {
    fs.unlink(DB_FILE, function (error) {
        if (error) {
            console.log(error);
            return false;
        }
        console.log('删除文件成功');
    })
}

const CREATE_TEST_TABLE = `CREATE TABLE IF NOT EXISTS test (
    key TEXT,
    value TEXT
);`

const INSERT_KEY_VALUE = `INSERT INTO test (key, value) VALUES (?, ?);`;

const QUERY_KEY_VALUE = `select * from test;`;
const DELETE_KEY = `delete from test where key=?`;

let easyDB = new EasyDB(DB_FILE, DB_LOG_FILE);

const main = async () => {
    try {
        await easyDB.init();
        await easyDB.runSql(CREATE_TEST_TABLE);
        await easyDB.runSql(INSERT_KEY_VALUE, ['testkey1', 'testvalue1'])
        await easyDB.runSql(INSERT_KEY_VALUE, ['testkey2', 'testvalue2'])
        const rows = await easyDB.querySql(QUERY_KEY_VALUE);
        console.log(rows);
        console.log(JSON.stringify(rows))
        await easyDB.runSql(DELETE_KEY, ['testkey1'])
        const rows2 = await easyDB.querySql(QUERY_KEY_VALUE);
        console.log(rows2);
        console.log(JSON.stringify(rows2))
    } catch (err) {
        console.log(err);
    }
};

main();