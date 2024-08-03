import BilladmDao from "../electron/billadmDao";
import * as path from "node:path";
import * as fs from "node:fs";
import {BUILT_IN_BILLBOOK} from "../electron/constants";

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
        let res = await billadmDao.queryAllType();
        console.log(res);
        console.log(`delete type by id ${res[0].id} ${res[0].name}`);
        await billadmDao.deleteOneTypeByID(res[0].id);
        res = await billadmDao.queryAllType();
        console.log(res);
        await billadmDao.insertOneBillbook('test', 'this a test billbook');
        res = await billadmDao.queryAllBillbook();
        console.log(res)
        console.log(`delete billbook by id ${res[0].id} ${res[0].name}`);
        await billadmDao.deleteOneBillbookByID(res[0].id);
        res = await billadmDao.queryAllBillbook();
        console.log(res);
        await billadmDao.insertOneBill(10.5, 'test_type', 'false', BUILT_IN_BILLBOOK.id, 'test_des', '[test]');
        await billadmDao.insertOneBill(3.6, 'test_type1', 'true', BUILT_IN_BILLBOOK.id, 'test_des1', '[test1]');
        res = await billadmDao.queryAllBillByBookID(BUILT_IN_BILLBOOK.id);
        console.log(res);
        console.log(`delete bill by id ${res[0].id} ${res[0].money}`);
        await billadmDao.deleteOneBillByID(res[0].id);
        res = await billadmDao.queryAllBillByBookID(BUILT_IN_BILLBOOK.id);
        console.log(res);
    } catch (err) {
        console.log(err);
    }
};

main();