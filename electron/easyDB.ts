import * as sqlite3 from "sqlite3";
import Logger from "./logger";

class EasyDB {
    private readonly logger: Logger;
    private readonly dbInstance: sqlite3.Database;

    constructor(dbFile: string, logFile: string) {
        this.logger = new Logger(logFile);
        this.dbInstance = new sqlite3.Database(dbFile, sqlite3.OPEN_READWRITE | sqlite3.OPEN_CREATE, (err) => {
            if (err) {
                this.logger.error(`connect database [${dbFile}] failed, err: ${err}`);
                return;
            }
            this.logger.info(`connect database [${dbFile}] success`);
        });
    }

    runSql(sql: string, value?: any) {
        if (value) {
            return new Promise(
                (resolve, reject) => {
                    this.dbInstance.run(sql, value, (err) => {
                        if (err) {
                            this.logger.error(`run sql [${sql}] failed, err: ${err}`);
                            reject(err);
                            return;
                        }
                        this.logger.info(`run sql [${sql}] success`);
                        resolve("");
                    })
                }
            )
        } else {
            return new Promise(
                (resolve, reject) => {
                    this.dbInstance.run(sql, (err) => {
                        if (err) {
                            this.logger.error(`run sql [${sql}] failed, err: ${err}`);
                            reject(err);
                            return;
                        }
                        this.logger.info(`run sql [${sql}] success`);
                        resolve("");
                    })
                }
            )
        }
    }

    querySql(sql: string, value?: any) {
        if (value) {
            return new Promise(
                (resolve, reject) => {
                    this.dbInstance.all(sql, (err, rows) => {
                        if (err) {
                            this.logger.error(`query data with sql [${sql}] failed, err: ${err}`);
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
                    this.dbInstance.all(sql, value, (err, rows) => {
                        if (err) {
                            this.logger.error(`query data with sql [${sql}] failed, err: ${err}`);
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