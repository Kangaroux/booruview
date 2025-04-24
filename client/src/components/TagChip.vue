<script setup lang="ts">
import { defineProps } from "vue";
const { tag } = defineProps<{ tag: Tag }>();
</script>

<template>
    <button class="chip" :class="tag.type">
        {{ tag.name }} ({{ tag.count }})
    </button>
</template>

<style scoped>
.chip {
    padding: 8px;
    border: none;
    border-radius: 8px;
}

.tag {
    background-color: #fff;
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
    background-color: #666;
    color: #fff;
}
</style>
