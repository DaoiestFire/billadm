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

export function timestampToLocalTimeString(timestamp) {
    timestamp = timestamp ? timestamp : null;
    let date = new Date(timestamp);
    let Y = date.getFullYear() + '-';
    let M = (date.getMonth() + 1 < 10 ? '0' + (date.getMonth() + 1) : date.getMonth() + 1) + '-';
    let D = (date.getDate() < 10 ? '0' + date.getDate() : date.getDate());
    let h = (date.getHours() < 10 ? '0' + date.getHours() : date.getHours()) + ':';
    let m = (date.getMinutes() < 10 ? '0' + date.getMinutes() : date.getMinutes()) + ':';
    let s = date.getSeconds() < 10 ? '0' + date.getSeconds() : date.getSeconds();
    return Y + M + D;
}

export function localTimeStringToTimeStamp(timeString) {

}