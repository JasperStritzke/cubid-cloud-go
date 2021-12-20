import {State} from "@/store/types";

export default function isAuthenticated(state: State): boolean {
    return !!state.auth
}
