import {
    computed,
    onMounted,
    onUnmounted,
    ref,
    toValue,
    type ComponentPublicInstance,
    type ComputedRef,
    type MaybeRefOrGetter,
    type ShallowRef,
} from "vue";

export function useDismiss(el: (HTMLElement | null)[], onDismiss: () => void) {
    function handler(e: MouseEvent) {
        const clickedOutside =
            el.findIndex((v) => {
                if (v === null) {
                    return false;
                }

                return v.contains(e.target as Node);
            }) === -1; // Clicked element was not found in any of `el`

        if (clickedOutside) {
            onDismiss();
        }
    }

    onMounted(() => document.addEventListener("click", handler));
    onUnmounted(() => document.removeEventListener("click", handler));
}

export function useIsVideo(post: MaybeRefOrGetter<Post>): ComputedRef<boolean> {
    return computed(() => {
        const val = toValue(post);

        return (
            val.image_url.endsWith(".mp4") || val.image_url.endsWith(".webm")
        );
    });
}

type StationaryClickReturn = {
    mouseDown: (e: MouseEvent) => void;
    mouseUp: (e: MouseEvent) => void;
};

// Fires a click event only if the cursor didn't move.
export function useStationaryClick(
    onClick: (e: MouseEvent) => void,
): StationaryClickReturn {
    const originX = ref(0);
    const originY = ref(0);

    // Number of pixels the cursor can move and still be considered a stationary click
    const allowedDistance = 10;

    function mouseDown(e: MouseEvent) {
        originX.value = e.x;
        originY.value = e.y;
    }

    function mouseUp(e: MouseEvent) {
        const dist = Math.sqrt(
            Math.pow(e.x - originX.value, 2) + Math.pow(e.y - originY.value, 2),
        );

        // Not a stationary click
        if (dist > allowedDistance) {
            return;
        }

        onClick(e);
    }

    return { mouseDown, mouseUp };
}
