<script setup lang="ts">
import store from "@/store";
import {
    computed,
    onActivated,
    onDeactivated,
    onMounted,
    onUnmounted,
    useTemplateRef,
} from "vue";
import createPanZoom, { type PanZoom } from "panzoom";

const imgRef = useTemplateRef("imgRef");
let pz: PanZoom;

const htmlRoot = document.body.parentElement as HTMLElement;
const overscrollCssClass = "prevent-overscroll";

onMounted(() => {
    pz = createPanZoom(imgRef.value!, {
        autocenter: true,
        bounds: true,
        maxZoom: 3,
        minZoom: 0.1,
        onTouch() {
            // Don't block the touch event so the user can right click
            return false;
        },
    });

    // Since the touch event isn't being blocked we need to prevent the user from
    // overscrolling the page (refresh by pulling down)
    htmlRoot.classList.add(overscrollCssClass);
});

onUnmounted(() => {
    pz.dispose();
    htmlRoot.classList.remove(overscrollCssClass);
});

onDeactivated(() => pz.pause());
onActivated(() => pz.resume());

const img = computed(() => {
    const post = store.fullscreenPost!;
    const hasHighRes = post.image_url.length > 0;
    const hasLowRes = post.lowres_url.length > 0;

    if (!hasLowRes || (hasHighRes && store.settings.highResImages)) {
        return {
            url: post.image_url,
            height: post.height,
            width: post.width,
        };
    }

    return {
        url: post.lowres_url,
        height: post.lowres_height,
        width: post.lowres_width,
    };
});
</script>

<template>
    <img
        :src="img.url"
        :width="img.width"
        :height="img.height"
        loading="lazy"
        ref="imgRef"
    />
</template>
