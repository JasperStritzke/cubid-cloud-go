import {Alert} from "@/store/types"

let idCounter = 0;

export default function trigger_alert(context: any, payload: Alert) {
    idCounter++;
    context.commit('trigger_alert', {...payload, id: idCounter})
}
