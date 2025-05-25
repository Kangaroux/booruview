<script setup lang="ts">
import store from "@/store";
import {
    onActivated,
    onDeactivated,
    onUnmounted,
    useTemplateRef,
    watchPostEffect,
} from "vue";
import createPanZoom, { type PanZoom } from "panzoom";

const imgRef = useTemplateRef("img");
let pz: PanZoom | undefined;

watchPostEffect(() => {
    if (imgRef.value == null) {
        return;
    }

    pz = createPanZoom(imgRef.value, {
        autocenter: true,
        bounds: true,
        maxZoom: 3,
        minZoom: 0.1,
    });
});

onUnmounted(() => {
    pz?.dispose();
});

onDeactivated(() => pz?.pause());
onActivated(() => pz?.resume());

// const urlLowRes = computed(() => {
//     if (store.fullscreenPost === null) {
//         return "";
//     }

//     return store.fullscreenPost.lowres_url || store.fullscreenPost.image_url;
// });
// const urlOriginal = computed(() => {
//     if (store.fullscreenPost === null) {
//         return "";
//     }

//     return store.fullscreenPost.image_url;
// });
</script>

<template>
    <img
        class="content high-res"
        :src="store.fullscreenPost?.image_url"
        :width="store.fullscreenPost?.width"
        :height="store.fullscreenPost?.height"
        loading="lazy"
        ref="img"
    />
</template>
