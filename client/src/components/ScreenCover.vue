<script setup lang="ts">
defineEmits(["click"]);
</script>

<template>
    <div class="screen-cover" @click="$emit('click')"></div>
</template>

<style scoped>
.screen-cover {
    position: absolute;
    z-index: 1;
    background-color: rgba(0, 0, 0, 0.95);
    width: 100%;
    height: 100%;
}
</style>
