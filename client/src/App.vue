<script setup lang="ts">
import { ref } from "vue";
import TagSearch from "./components/TagSearch.vue";
import TagChip from "./components/TagChip.vue";

const tags = ref<Tag[]>([]);
</script>

<template>
    <TagSearch @submit="t => tags = tags.concat(t)"/>

    <template v-for="t in tags">
        <TagChip :tag="t" />
    </template>
</template>
