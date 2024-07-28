import {
    getLastMonthDate,
    getLastWeekData,
    getLastYearDate,
    getThisMonthData,
    getThisWeekData,
    getThisYearDate
} from '@/utils/timeutils';

export const shortcuts = [
    {
        text: '上周',
        value: getLastWeekData(),
    },
    {
        text: '上个月',
        value: getLastMonthDate(),
    },
    {
        text: '去年',
        value: getLastYearDate(),
    },
    {
        text: '本周',
        value: getThisWeekData(),
    },
    {
        text: '本月',
        value: getThisMonthData(),
    },
    {
        text: '今年',
        value: getThisYearDate(),
    },
    {
        text: '今天',
        value: [],
    },
    {
        text: '全部',
        value: [null, null],
    },
];