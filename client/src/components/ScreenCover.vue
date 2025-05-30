<script setup lang="ts">
defineEmits(["click"]);
const { blackout } = defineProps<{ blackout?: boolean }>();
</script>

<template>
    <div
        class="screen-cover"
        :class="{ blackout }"
        @click="$emit('click')"
    ></div>
</template>

<style lang="scss" scoped>
.screen-cover {
    position: absolute;
    z-index: 1;
    width: 100%;
    height: 100%;
    background-color: black;
    opacity: 0.9;

    &.blackout {
        opacity: 1;
    }
}
</style>
