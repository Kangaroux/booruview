<script setup lang="ts">
import { computed } from "vue";
import TagChip from "./TagChip.vue";

type tagWithState = { tag: Tag; state: TagState };
type listCategories = {
    artist: tagWithState[];
    character: tagWithState[];
    copyright: tagWithState[];
    tag: tagWithState[];
    metadata: tagWithState[];
};

defineEmits<{click: [tag: Tag]}>();

const { jiggle, excludeTags, includeTags } = defineProps<{
    jiggle: boolean;
    includeTags: Tag[];
    excludeTags: Tag[];
}>();

const categories = computed(() => {
    const ret: listCategories = {
        artist: [],
        character: [],
        copyright: [],
        tag: [],
        metadata: [],
    };

    for (const t of includeTags) {
        switch (t.type) {
            case "artist":
                ret.artist = ret.artist.concat({ tag: t, state: "include" });
                break;
            case "character":
                ret.character = ret.character.concat({
                    tag: t,
                    state: "include",
                });
                break;
            case "copyright":
                ret.copyright = ret.copyright.concat({
                    tag: t,
                    state: "include",
                });
                break;
            case "tag":
                ret.tag = ret.tag.concat({ tag: t, state: "include" });
                break;
            case "metadata":
            case "unknown":
                ret.metadata = ret.metadata.concat({
                    tag: t,
                    state: "include",
                });
                break;
        }
    }

    for (const t of excludeTags) {
        switch (t.type) {
            case "artist":
                ret.artist = ret.artist.concat({ tag: t, state: "exclude" });
                break;
            case "character":
                ret.character = ret.character.concat({
                    tag: t,
                    state: "exclude",
                });
                break;
            case "copyright":
                ret.copyright = ret.copyright.concat({
                    tag: t,
                    state: "exclude",
                });
                break;
            case "tag":
                ret.tag = ret.tag.concat({ tag: t, state: "exclude" });
                break;
            case "metadata":
            case "unknown":
                ret.metadata = ret.metadata.concat({
                    tag: t,
                    state: "exclude",
                });
                break;
        }
    }

    ret.artist.sort((a, b) => a.tag.name.localeCompare(b.tag.name));
    ret.character.sort((a, b) => a.tag.name.localeCompare(b.tag.name));
    ret.copyright.sort((a, b) => a.tag.name.localeCompare(b.tag.name));
    ret.tag.sort((a, b) => a.tag.name.localeCompare(b.tag.name));
    ret.metadata.sort((a, b) => a.tag.name.localeCompare(b.tag.name));

    return ret;
});
</script>

<template>
    <h3 v-if="categories.artist.length > 0">artist</h3>
    <TagChip
        v-for="t in categories.artist"
        :tag="t.tag"
        :key="t.tag.name"
        :state="t.state"
        :jiggle="jiggle"
        @click="$emit('click', t.tag)"
    />

    <h3 v-if="categories.character.length > 0">character</h3>
    <TagChip
        v-for="t in categories.character"
        :tag="t.tag"
        :key="t.tag.name"
        :state="t.state"
        :jiggle="jiggle"
        @click="$emit('click', t.tag)"
    />

    <h3 v-if="categories.copyright.length > 0">copyright</h3>
    <TagChip
        v-for="t in categories.copyright"
        :tag="t.tag"
        :key="t.tag.name"
        :state="t.state"
        :jiggle="jiggle"
        @click="$emit('click', t.tag)"
    />
    <h3 v-if="categories.tag.length > 0">tags</h3>
    <TagChip
        v-for="t in categories.tag"
        :tag="t.tag"
        :key="t.tag.name"
        :state="t.state"
        :jiggle="jiggle"
        @click="$emit('click', t.tag)"
    />

    <h3 v-if="categories.metadata.length > 0">metadata</h3>
    <TagChip
        v-for="t in categories.metadata"
        :tag="t.tag"
        :key="t.tag.name"
        :state="t.state"
        :jiggle="jiggle"
        @click="$emit('click', t.tag)"
    />
</template>

<style lang="scss" scoped>
h3 {
    margin: 15px 0 10px;
    font-size: 18px;
}
</style>
