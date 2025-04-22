<script setup lang="ts">
import {defineProps } from "vue";
const { tag } = defineProps<{tag: Tag}>();
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
    background-color: #FFF;
}

.artist {
    background-color: #A00;
    color: #FFF;
}

.copyright {
    background-color: #A0A;
    color: #FFF;
}

.character {
    background-color: #0A0;
}

.metadata {
    background-color: #F80;
}

.deprecated, .unknown {
    background-color: #666;
    color: #FFF;
}
</style>
