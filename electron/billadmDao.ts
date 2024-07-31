import EasyDB from "./easyDB";

class BilladmDao {
    private readonly easyDB: EasyDB;

    constructor(dbFile: string, logFile: string) {
        this.easyDB = new EasyDB(dbFile, logFile);
    }


}

export default BilladmDao;