<script setup lang="ts">
import { computed } from "vue";

defineEmits(["click"]);

const { tag } = defineProps<{ tag: StatefulTag }>();

const cls = computed(() => ({
    [`state-${tag.state}`]: true,
    [tag.type]: true,
}));
</script>

<template>
    <div class="chip" :class="cls" @click="$emit('click')">
        {{ tag.name }}
        <i class="bi bi-check-lg" v-if="tag.state == 'include'"></i>
    </div>
</template>

<style lang="scss" scoped>
.chip {
    padding: 8px;
    margin: 0 4px 4px 0;
    border: none;
    border-radius: 8px;
    display: inline-block;
    font-size: 16px;

    &.state-exclude {
        filter: brightness(0.8);
        text-decoration: line-through;
    }
}

.tag {
    background-color: #303030;
    color: hsl(208, 56%, 75%);

    &.active {
        background-color: hsl(208, 56%, 75%);
        color: #303030;
    }
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
    background-color: #6275ae;
    color: #0b1227;
}
</style>
