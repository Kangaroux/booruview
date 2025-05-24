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
    unknown: tagWithState[];
};

defineEmits<{ click: [tag: Tag] }>();

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
        unknown: [],
    };

    function sortTag(tag: Tag, state: TagState) {
        switch (tag.type) {
            case "artist":
                ret.artist = ret.artist.concat({ tag, state });
                break;
            case "character":
                ret.character = ret.character.concat({
                    tag,
                    state,
                });
                break;
            case "copyright":
                ret.copyright = ret.copyright.concat({
                    tag,
                    state,
                });
                break;
            case "tag":
                ret.tag = ret.tag.concat({ tag, state });
                break;
            case "metadata":
                ret.metadata = ret.metadata.concat({
                    tag,
                    state,
                });
                break;
            case "unknown":
                ret.unknown = ret.unknown.concat({
                    tag,
                    state,
                });
                break;
        }
    }

    includeTags.forEach((tag) => sortTag(tag, "include"));
    excludeTags.forEach((tag) => sortTag(tag, "exclude"));

    ret.artist.sort((a, b) => a.tag.name.localeCompare(b.tag.name));
    ret.character.sort((a, b) => a.tag.name.localeCompare(b.tag.name));
    ret.copyright.sort((a, b) => a.tag.name.localeCompare(b.tag.name));
    ret.tag.sort((a, b) => a.tag.name.localeCompare(b.tag.name));
    ret.metadata.sort((a, b) => a.tag.name.localeCompare(b.tag.name));
    ret.unknown.sort((a, b) => a.tag.name.localeCompare(b.tag.name));

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

    <h3 v-if="categories.unknown.length > 0">raw</h3>
    <TagChip
        v-for="t in categories.unknown"
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
