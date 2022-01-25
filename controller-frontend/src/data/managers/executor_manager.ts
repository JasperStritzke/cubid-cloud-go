import axios from "axios";

export interface Executor {
    name: string,
    macCpuUsage: number,
    maxMemory: number
}

export async function createExecutor(executor: Executor) {
    return axios.put("/executor/create", executor)
}

export async function getExecutors(): Promise<Executor[]> {
    return axios.get("/executor/")
        .then(r => r.data)
        .then(({executors}) => executors)
}

export async function deleteExecutor(name: string): Promise<void> {
    return axios.delete("/executor/" + name)
}
