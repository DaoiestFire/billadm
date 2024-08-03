import * as uuid from "node-uuid";

export const UUID = () => uuid.v1().replace(/-/g, "");