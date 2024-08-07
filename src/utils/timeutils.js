export function getThisWeekData() {
    let today = new Date();
    let dayOfWeek = today.getDay() - 1;
    let startDate = new Date(today.getTime() - dayOfWeek * 86400000);
    let endDate = new Date(startDate.getTime() + 6 * 86400000);
    return [startDate, endDate];
}

export function getLastWeekData() {
    let thisWeek = getThisWeekData();
    let startDate = new Date(new Date(thisWeek[0]).getTime() - 7 * 86400000);
    let endDate = new Date(new Date(startDate).getTime() + 6 * 86400000);
    return [startDate, endDate];
}

export function getThisMonthData() {
    let today = new Date();
    let year = today.getFullYear();
    let month = today.getMonth() + 1;
    let daysInMonth = new Date(year, month, 0).getDate();
    let startDate = new Date(year, month - 1, 1);
    let endDate = new Date(year, month - 1, daysInMonth);
    return [startDate, endDate];
}

export function getLastMonthDate() {
    let thisMonth = getThisMonthData();
    let startDate = new Date(new Date(thisMonth[0]).setMonth(new Date(thisMonth[0]).getMonth() - 1));
    let endDate = new Date(new Date(thisMonth[1]).setMonth(new Date(thisMonth[1]).getMonth(), 0));
    return [startDate, endDate];
}

export function getThisYearDate() {
    let startDate = new Date();
    startDate.setMonth(0);
    startDate.setDate(1);
    let endDate = new Date(startDate);
    endDate.setMonth(11);
    endDate.setDate(31);
    return [startDate, endDate];
}


export function getLastYearDate() {
    let thisYear = getThisYearDate();
    let startDate = new Date(thisYear[0]);
    let endDate = new Date(thisYear[1]);
    startDate.setFullYear(startDate.getFullYear() - 1);
    endDate.setFullYear(endDate.getFullYear() - 1);
    return [startDate, endDate];
}

export function dateObjectToUTCTimeString(date) {
    let Y = date.getUTCFullYear() + '-';
    let M = (date.getUTCMonth() + 1 < 10 ? '0' + (date.getUTCMonth() + 1) : date.getUTCMonth() + 1) + '-';
    let D = (date.getUTCDate() < 10 ? '0' + date.getUTCDate() : date.getUTCDate()) + ' ';
    let h = (date.getUTCHours() < 10 ? '0' + date.getUTCHours() : date.getUTCHours()) + ':';
    let m = (date.getUTCMinutes() < 10 ? '0' + date.getUTCMinutes() : date.getUTCMinutes()) + ':';
    let s = date.getUTCSeconds() < 10 ? '0' + date.getUTCSeconds() : date.getUTCSeconds();
    return Y + M + D + h + m + s;
}

export function utcTimeStringToDateObject(utcDateString) {
    let dateString = utcDateString.trim().split(' ')[0].trim().split('-');
    let timeString = utcDateString.trim().split(' ')[1].trim().split(':');
    let Y = dateString[0];
    let M = dateString[1];
    let D = dateString[2];
    let h = timeString[0];
    let m = timeString[1];
    let s = timeString[2];
    let date = new Date();
    date.setUTCFullYear(Number(Y));
    date.setUTCMonth(Number(M) - 1);
    date.setUTCDate(Number(D));
    date.setUTCHours(Number(h), Number(m), Number(s));
    return date;
}

export function dateObjectToLocalTimeString(date) {
    let Y = date.getFullYear() + '-';
    let M = (date.getMonth() + 1 < 10 ? '0' + (date.getMonth() + 1) : date.getMonth() + 1) + '-';
    let D = (date.getDate() < 10 ? '0' + date.getDate() : date.getDate()) + ' ';
    let h = (date.getHours() < 10 ? '0' + date.getHours() : date.getHours()) + ':';
    let m = (date.getMinutes() < 10 ? '0' + date.getMinutes() : date.getMinutes()) + ':';
    let s = date.getSeconds() < 10 ? '0' + date.getSeconds() : date.getSeconds();
    return Y + M + D + h + m + s;
}

export function isValidDate(date) {
    return date instanceof Date && !isNaN(date.getTime())
}
