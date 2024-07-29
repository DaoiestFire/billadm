import * as fs from 'fs';

const INFO = 'INFO ';
const ERROR = 'ERROR';

function formatDateTime(date: Date): string {
    const year = date.getFullYear();
    const month = date.getMonth() + 1;
    const day = date.getDate();
    const hour = date.getHours();
    const minute = date.getMinutes();
    const second = date.getSeconds();
    return `${year}-${pad(month)}-${pad(day)} ${pad(hour)}:${pad(minute)}:${pad(second)}`;
}

function pad(num: number) {
    return num.toString().padStart(2, '0');
}

const writeLog = (out: string, level: string, logFile: string, maxLogLines: number) => {
    out = out.toString();
    out = formatDateTime(new Date()) + " " + level + " " + out;
    console.log(out);
    let log = "";
    if (logFile.length == 0) {
        return;
    }
    try {
        if (fs.existsSync(logFile)) {
            log = fs.readFileSync(logFile).toString();
            let lines = log.split("\n");
            if (maxLogLines < lines.length) {
                log = lines.slice(maxLogLines / 2, maxLogLines).join("\n") + "\n";
            }
        }
        log += out + "\n";
        fs.writeFileSync(logFile, log);
    } catch (e) {
        console.error(e);
    }
};

class Logger {
    private readonly logFile: string;
    private readonly maxLogLines: number;

    constructor(logFile?: string, maxLogLines: number = 1024) {
        this.logFile = logFile ? logFile : "";
        this.maxLogLines = maxLogLines;
    }

    info(msg: string) {
        writeLog(msg, INFO, this.logFile, this.maxLogLines);
    }

    error(msg: string) {
        writeLog(msg, ERROR, this.logFile, this.maxLogLines);
    }
}

export default Logger;