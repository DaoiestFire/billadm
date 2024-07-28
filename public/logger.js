const fs = require('fs');

const INFO = 'INFO';
const ERROR = 'ERROR';

const writeLog = (out, level, logFile) => {
    console.log(out);
    let log = "";
    const maxLogLines = 1024;
    try {
        if (fs.existsSync(logFile)) {
            log = fs.readFileSync(logFile).toString();
            let lines = log.split("\n");
            if (maxLogLines < lines.length) {
                log = lines.slice(maxLogLines / 2, maxLogLines).join("\n") + "\n";
            }
        }
        out = out.toString();
        out = new Date().toISOString().replace(/T/, " ").replace(/\..+/, "") + " " + level + " " + out;
        log += out + "\n";
        fs.writeFileSync(logFile, log);
    } catch (e) {
        console.error(e);
    }
};


exports.init_logger = (f) => {
    return {
        logFile: f,
        info: function (out) {
            writeLog(out, INFO, this.logFile);
        },
        error: function (out) {
            writeLog(out, ERROR, this.logFile);
        },
    }
};