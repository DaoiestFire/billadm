import {ElNotification} from "element-plus";

export const notify = (type, msg) => {
    ElNotification({
        type: type,
        message: msg,
        position: 'bottom-right',
        duration: 2000,
        offset: 40,
    });
}