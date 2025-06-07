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
import { useIsVideo } from "@/composable";

const imgRef = useTemplateRef("imgRef");
let pz: PanZoom | undefined;

const post = store.fullscreenPost!;
const htmlRoot = document.body.parentElement as HTMLElement;
const overscrollCssClass = "prevent-overscroll";

const isVideo = useIsVideo(post);

onMounted(() => {
    if (isVideo.value) {
        return;
    }

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
    if (isVideo.value) {
        return;
    }

    pz?.dispose();
    htmlRoot.classList.remove(overscrollCssClass);
});

onDeactivated(() => pz?.pause());
onActivated(() => pz?.resume());

const content = computed(() => {
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
    <video
        v-if="isVideo"
        :poster="post.thumbnail_url || post.lowres_url"
        controls
        loop
    >
        <source
            :src="content.url"
            type="video/mp4"
            v-if="content.url.endsWith('.mp4')"
        />
        <source
            :src="content.url"
            type="video/webm"
            v-if="content.url.endsWith('.webm')"
        />
    </video>

    <img
        v-else
        :src="content.url"
        :width="content.width"
        :height="content.height"
        ref="imgRef"
    />
</template>

<style scoped>
video {
    max-width: 100%;
    max-height: 100%;
}
</style>
