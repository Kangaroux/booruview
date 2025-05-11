<script setup lang="ts">
import { computed } from "vue";
import StatefulTagChip from "./StatefulTagChip.vue";

type listCategories = {
    artist: StatefulTag[];
    character: StatefulTag[];
    copyright: StatefulTag[];
    tag: StatefulTag[];
    metadata: StatefulTag[];
};

const { tags } = defineProps<{ tags: StatefulTag[] }>();

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

function cycleState(tag: StatefulTag) {
    if (tag.state === "default") {
        tag.state = "include";
    } else if (tag.state === "include") {
        tag.state = "exclude";
    } else {
        tag.state = "default";
    }
}
</script>

<template>
    <h3 v-if="categories.artist.length > 0">artist</h3>
    <StatefulTagChip
        v-for="t in categories.artist"
        :tag="t"
        :key="t.name"
        @click="cycleState(t)"
    />

    <h3 v-if="categories.character.length > 0">character</h3>
    <StatefulTagChip
        v-for="t in categories.character"
        :tag="t"
        :key="t.name"
        @click="cycleState(t)"
    />

    <h3 v-if="categories.copyright.length > 0">copyright</h3>
    <StatefulTagChip
        v-for="t in categories.copyright"
        :tag="t"
        :key="t.name"
        @click="cycleState(t)"
    />

    <h3 v-if="categories.tag.length > 0">tags</h3>
    <StatefulTagChip
        v-for="t in categories.tag"
        :tag="t"
        :key="t.name"
        @click="cycleState(t)"
    />

    <h3 v-if="categories.metadata.length > 0">metadata</h3>
    <StatefulTagChip
        v-for="t in categories.metadata"
        :tag="t"
        :key="t.name"
        @click="cycleState(t)"
    />
</template>

<style lang="scss" scoped>
h3 {
    margin: 15px 0 10px;
    font-size: 18px;
}
</style>
