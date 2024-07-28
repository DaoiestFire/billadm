const { init_logger } = require('../public/logger');

logger = init_logger("./test/test.log");
logger.info("start to test");
logger.error("test error");

logger2 = init_logger("./test/test1.log");
logger2.info("start to test");
logger2.error("test error");