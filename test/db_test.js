const dbController = require('../public/sqlite');

const CREATE_TEST_TABLE = `CREATE TABLE test (
    name TEXT,
    value TEXT
)`;

const INSERT_TEST_VALUE = 'INSERT INTO test (name,value) VALUES (?,?)';

const QUERY_TEST_VALUE = 'SELECT * FROM test';

dbInstance = dbController.get_db("./test/test.db", "./test/test.db.log");
dbController.db_run_sql_no_value(dbInstance, CREATE_TEST_TABLE);
dbController.db_run_sql_wth_value(dbInstance, INSERT_TEST_VALUE, ['testkey', 'testvalue']);
dbController.db_query_data_no_value(dbInstance, QUERY_TEST_VALUE, (rows) => {
    console.log(rows);
})
dbController.close_db(dbInstance);