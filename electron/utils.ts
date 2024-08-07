import * as uuid from "node-uuid";

export const UUID = () => uuid.v1().replace(/-/g, "");

export const getCurrentUTCTimeString = () => {
    let date = new Date();
    let Y = date.getUTCFullYear() + '-';
    let M = (date.getUTCMonth() + 1 < 10 ? '0' + (date.getUTCMonth() + 1) : date.getUTCMonth() + 1) + '-';
    let D = (date.getUTCDate() < 10 ? '0' + date.getUTCDate() : date.getUTCDate()) + ' ';
    let h = (date.getUTCHours() < 10 ? '0' + date.getUTCHours() : date.getUTCHours()) + ':';
    let m = (date.getUTCMinutes() < 10 ? '0' + date.getUTCMinutes() : date.getUTCMinutes()) + ':';
    let s = date.getUTCSeconds() < 10 ? '0' + date.getUTCSeconds() : date.getUTCSeconds();
    return Y + M + D + h + m + s;
};