import { type ShallowRef } from "vue";

export function useDeepFocusOut(
    el: Readonly<ShallowRef<HTMLElement | null>>,
    onFocusOut: () => void,
) {
    function handler(e: FocusEvent) {
        if (el.value === null) {
            return;
        }

        if (
            e.relatedTarget === null ||
            !el.value.contains(e.relatedTarget as Node)
        ) {
            onFocusOut();
        }
    }

    return handler;
}
