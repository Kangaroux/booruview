<script setup lang="ts">
import { computed, defineProps } from "vue";
const { tag } = defineProps<{ tag: Tag }>();
</script>

<template>
    <div class="chip" :class="tag.type">{{ tag.name }}</div>
</template>

<style scoped>
.chip {
    padding: 8px;
    border: none;
    border-radius: 8px;
    display: inline-block;
    font-size: 16px;
    color: #000;
}

.tag {
    background-color: #303030;
    color: hsl(208, 56%, 75%);
}

.artist {
    background-color: #a00;
    color: #fff;
}

.copyright {
    background-color: #a0a;
    color: #fff;
}

.character {
    background-color: #0a0;
}

.metadata {
    background-color: #f80;
}

.deprecated,
.unknown {
    background-color: #DDD;
    color: #000;
}
</style>
