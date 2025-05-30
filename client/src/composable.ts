import { onMounted, onUnmounted, type ShallowRef } from "vue";

export function useDismiss(
    el: Readonly<ShallowRef<HTMLElement | null>>,
    onDismiss: () => void,
) {
    function handler(e: MouseEvent) {
        if (el.value === null) {
            return;
        }

        if (e.target === null || !el.value.contains(e.target as Node)) {
            onDismiss();
        }
    }

    onMounted(() => document.addEventListener("click", handler));
    onUnmounted(() => document.removeEventListener("click", handler));
}
