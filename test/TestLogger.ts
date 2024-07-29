import Logger from "../electron/logger";
import * as path from "node:path";

let logger: Logger = new Logger(path.join(__dirname, 'test.log'));

logger.info("Test Logger");
logger.error("Test Logger");

let logger1: Logger = new Logger();

logger1.info("Test Logger");
logger1.error("Test Logger");