<script setup lang="ts">
import { computed } from "vue";
import Chip from "./tag-chip/Chip.vue";
import store from "@/store";

type listCategories = {
    artist: Tag[];
    character: Tag[];
    copyright: Tag[];
    tag: Tag[];
    metadata: Tag[];
    unknown: Tag[];
};

const { jiggle, tags, showCheckmark } = defineProps<{
    jiggle: boolean;
    tags: Tag[];
    showCheckmark: boolean;
}>();

const categories = computed(() => {
    const ret: listCategories = {
        artist: [],
        character: [],
        copyright: [],
        tag: [],
        metadata: [],
        unknown: [],
    };

    function sortTag(tag: Tag) {
        switch (tag.type) {
            case "artist":
                ret.artist = ret.artist.concat(tag);
                break;
            case "character":
                ret.character = ret.character.concat(tag);
                break;
            case "copyright":
                ret.copyright = ret.copyright.concat(tag);
                break;
            case "deprecated":
            case "tag":
                ret.tag = ret.tag.concat(tag);
                break;
            case "metadata":
                ret.metadata = ret.metadata.concat(tag);
                break;
            case "unknown":
                ret.unknown = ret.unknown.concat(tag);
                break;
        }
    }

    tags.forEach(sortTag);

    ret.artist.sort((a, b) => a.name.localeCompare(b.name));
    ret.character.sort((a, b) => a.name.localeCompare(b.name));
    ret.copyright.sort((a, b) => a.name.localeCompare(b.name));
    ret.tag.sort((a, b) => a.name.localeCompare(b.name));
    ret.metadata.sort((a, b) => a.name.localeCompare(b.name));
    ret.unknown.sort((a, b) => a.name.localeCompare(b.name));

    return ret;
});

const tagState = (tag: Tag) =>
    computed<TagState>(() => {
        const included = store.query.include.has(tag.name);
        if (included) {
            return showCheckmark ? "include" : "none";
        }

        const excluded = store.query.exclude.has(tag.name);
        if (excluded) {
            return "exclude";
        }

        return "none";
    });
</script>

<template>
    <h3 v-if="categories.artist.length > 0">artist</h3>
    <Chip
        v-for="t in categories.artist"
        :key="t.name"
        :tag="t"
        :state="tagState(t).value"
        :jiggle="jiggle"
    />

    <h3 v-if="categories.character.length > 0">character</h3>
    <Chip
        v-for="t in categories.character"
        :key="t.name"
        :tag="t"
        :state="tagState(t).value"
        :jiggle="jiggle"
    />

    <h3 v-if="categories.copyright.length > 0">copyright</h3>
    <Chip
        v-for="t in categories.copyright"
        :key="t.name"
        :tag="t"
        :state="tagState(t).value"
        :jiggle="jiggle"
    />

    <h3 v-if="categories.tag.length > 0">tags</h3>
    <Chip
        v-for="t in categories.tag"
        :key="t.name"
        :tag="t"
        :state="tagState(t).value"
        :jiggle="jiggle"
    />

    <h3 v-if="categories.metadata.length > 0">metadata</h3>
    <Chip
        v-for="t in categories.metadata"
        :key="t.name"
        :tag="t"
        :state="tagState(t).value"
        :jiggle="jiggle"
    />

    <h3 v-if="categories.unknown.length > 0">raw</h3>
    <Chip
        v-for="t in categories.unknown"
        :key="t.name"
        :tag="t"
        :state="tagState(t).value"
        :jiggle="jiggle"
    />
</template>

<style lang="scss" scoped>
h3 {
    margin: 15px 0 10px;
    font-size: 18px;
}
</style>
