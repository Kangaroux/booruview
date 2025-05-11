<script setup lang="ts">
import store from "@/store";
import { computed, ref } from "vue";

const fitHeight = ref(true);
const url = computed(() => {
    if (store.fullscreenPost === null) {
        return "";
    }

    return store.fullscreenPost.image_url;
});
</script>

<template>
    <div class="content-container">
        <img
            class="content"
            :class="{
                'fit-height': fitHeight,
                'fit-width': !fitHeight,
            }"
            :src="url"
            @click="fitHeight = !fitHeight"
        />
    </div>
</template>

<style lang="scss" scoped>
.content-container {
    height: 100%;
    overflow-y: scroll;
    line-height: 0;

    scrollbar-width: none;

    &::-webkit-scrollbar {
        display: none;
    }
}

.content {
    max-width: 100%;

    &.fit-height {
        max-height: 100%;
        width: auto;

        cursor: zoom-in;
    }

    &.fit-width {
        height: auto;

        cursor: zoom-out;
    }
}
</style>
