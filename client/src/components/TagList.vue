<script setup lang="ts">
import { computed } from "vue";
import TagChip from "./TagChip.vue";

type listCategories = {
    artist: Tag[];
    character: Tag[];
    copyright: Tag[];
    tag: Tag[];
    metadata: Tag[];
};

const {
    jiggle,
    activeTags = [],
    tags,
} = defineProps<{ jiggle: boolean; activeTags?: Tag[]; tags: Tag[] }>();
const activeSet = computed(() => new Set(activeTags.map((t) => t.name)));

const categories = computed(() => {
    const ret: listCategories = {
        artist: [],
        character: [],
        copyright: [],
        tag: [],
        metadata: [],
    };

    const sorted = [...tags].sort((a, b) => a.name.localeCompare(b.name));

    for (const t of sorted) {
        switch (t.type) {
            case "artist":
                ret.artist = ret.artist.concat(t);
                break;
            case "character":
                ret.character = ret.character.concat(t);
                break;
            case "copyright":
                ret.copyright = ret.copyright.concat(t);
                break;
            case "tag":
                ret.tag = ret.tag.concat(t);
                break;
            case "metadata":
            case "unknown":
                ret.metadata = ret.metadata.concat(t);
                break;
        }
    }

    return ret;
});
</script>

<template>
    <h3 v-if="categories.artist.length > 0">artist</h3>
    <TagChip
        v-for="t in categories.artist"
        :tag="t"
        :key="t.name"
        :jiggle="jiggle"
        :active="activeSet.has(t.name)"
    />

    <h3 v-if="categories.character.length > 0">character</h3>
    <TagChip
        v-for="t in categories.character"
        :tag="t"
        :key="t.name"
        :jiggle="jiggle"
        :active="activeSet.has(t.name)"
    />

    <h3 v-if="categories.copyright.length > 0">copyright</h3>
    <TagChip
        v-for="t in categories.copyright"
        :tag="t"
        :key="t.name"
        :jiggle="jiggle"
        :active="activeSet.has(t.name)"
    />
    <h3 v-if="categories.tag.length > 0">tags</h3>
    <TagChip
        v-for="t in categories.tag"
        :tag="t"
        :key="t.name"
        :jiggle="jiggle"
        :active="activeSet.has(t.name)"
    />

    <h3 v-if="categories.metadata.length > 0">metadata</h3>
    <TagChip
        v-for="t in categories.metadata"
        :tag="t"
        :key="t.name"
        :jiggle="jiggle"
        :active="activeSet.has(t.name)"
    />
</template>

<style lang="scss" scoped>
h3 {
    margin: 15px 0 10px;
    font-size: 18px;
}
</style>
