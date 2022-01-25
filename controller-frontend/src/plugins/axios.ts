import axios from "axios";

export default function setAuthToken(token: string | undefined) {
    if (!token) {
        delete axios.defaults.headers.common["Authorization"]
        return
    }

    axios.defaults.headers.common["Authorization"] = token
}
