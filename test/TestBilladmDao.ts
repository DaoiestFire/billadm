import BilladmDao from "../electron/billadmDao";
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

let billadmDao = new BilladmDao(DB_FILE, DB_LOG_FILE);

const main = async () => {
    try {
        await billadmDao.init();
        await billadmDao.initDB();
    } catch (err) {
        console.log(err);
    }
};

main();