import * as sqlite3 from "sqlite3";
import Logger from "./logger";

class EasyDB {
    private readonly dbFile: string;
    private readonly logger: Logger;
    private dbInstance: sqlite3.Database;

    constructor(dbFile: string, logFile: string) {
        this.dbFile = dbFile;
        this.logger = new Logger(logFile);
    }

    init() {
        return new Promise((resolve, reject) => {
            this.dbInstance = new sqlite3.Database(this.dbFile, sqlite3.OPEN_READWRITE | sqlite3.OPEN_CREATE, (err) => {
                if (err) {
                    this.logger.error(`connect database [${this.dbFile}] failed, err: ${err.message}`);
                    reject(err);
                    return;
                }
                this.logger.info(`connect database [${this.dbFile}] success`);
                resolve(true)
            });
        })
    }

    runSql(sql: string, value?: any) {
        if (value) {
            return new Promise(
                (resolve, reject) => {
                    this.dbInstance.run(sql, value, (err) => {
                        if (err) {
                            this.logger.error(`run sql [${sql}] failed, err: ${err.message}`);
                            reject(err);
                            return;
                        }
                        this.logger.info(`run sql [${sql}] success`);
                        resolve(true);
                    })
                }
            )
        } else {
            return new Promise(
                (resolve, reject) => {
                    this.dbInstance.run(sql, (err) => {
                        if (err) {
                            this.logger.error(`run sql [${sql}] failed, err: ${err.message}`);
                            reject(err);
                            return;
                        }
                        this.logger.info(`run sql [${sql}] success`);
                        resolve(true);
                    })
                }
            )
        }
    }

    querySql(sql: string, value?: any) {
        if (value) {
            return new Promise(
                (resolve, reject) => {
                    this.dbInstance.all(sql, value, (err, rows) => {
                        if (err) {
                            this.logger.error(`query data with sql [${sql}] failed, err: ${err.message}`);
                            reject(err);
                            return;
                        }
                        this.logger.info(`query data with sql [${sql}] success`);
                        resolve(rows);
                    });
                }
            )
        } else {
            return new Promise(
                (resolve, reject) => {
                    this.dbInstance.all(sql, (err, rows) => {
                        if (err) {
                            this.logger.error(`query data with sql [${sql}] failed, err: ${err.message}`);
                            reject(err);
                            return;
                        }
                        this.logger.info(`query data with sql [${sql}] success`);
                        resolve(rows);
                    });
                }
            )
        }
    }

    closeDB() {
        this.dbInstance.close((err) => {
            if (err) {
                this.logger.error(`close db failed, err: ${err}`);
                return;
            }
            this.logger.info('close db success');
        });
    }
}

export default EasyDB;