const path = require('node:path');
const sqlite3 = require('sqlite3').verbose();
const { init_logger } = require('./logger');

let logger;

const db_run_sql_no_value = (db, sql) => {
    db.run(sql, function (err) {
        if (err) {
            logger.error(`run sql [${sql}] failed, err: ${err}`);
            return;
        }
        logger.info(`run sql [${sql}] success`);
    });
};

const db_run_sql_wth_value = (db, sql, value) => {
    db.run(sql, value, function (err) {
        if (err) {
            logger.error(`run sql [${sql}] failed, err: ${err}`);
            return;
        }
        logger.info(`run sql [${sql}] success`);
    });
};

const db_query_data_no_value = (db, sql, callback) => {
    db.all(sql, function (err, rows) {
        if (err) {
            logger.error(`query data with sql [${sql}] failed, err: ${err}`);
            return;
        }
        logger.info(`query data with sql [${sql}] success`);
        callback(rows);
    });
}

const db_query_data_with_value = (db, sql, value, callback) => {
    db.all(sql, value, function (err, rows) {
        if (err) {
            logger.error(`query data with sql [${sql}] failed, err: ${err}`);
            return;
        }
        logger.info(`query data with sql [${sql}] success`);
        callback(rows);
    });
}

exports.get_db = (db_path, log_path) => {
    logger = init_logger(log_path);
    let db = new sqlite3.Database(
        db_path,
        sqlite3.OPEN_READWRITE | sqlite3.OPEN_CREATE,
        function (err) {
            if (err) {
                logger.error(`connect database [${db_path}] failed, err: ${err}`);
                return;
            }
            logger.info(`connect database [${db_path}] success`);
        }
    );
    return db ? db : null;
};

exports.close_db = (db) => {
    db.close(function (err) {
        if (err) {
            logger.error(`close db failed, err: ${err}`);
            return;
        }
        logger.info('close db success');
    });
};

exports.db_run_sql_no_value = db_run_sql_no_value;
exports.db_run_sql_wth_value = db_run_sql_wth_value;
exports.db_query_data_no_value = db_query_data_no_value;
exports.db_query_data_with_value = db_query_data_with_value;