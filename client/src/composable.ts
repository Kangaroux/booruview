import {
    computed,
    onMounted,
    onUnmounted,
    ref,
    type ComputedRef,
    type ShallowRef,
} from "vue";

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

export function useIsVideo(post: Post): ComputedRef<boolean> {
    return computed(() => {
        return (
            post.image_url.endsWith(".mp4") || post.image_url.endsWith(".webm")
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
