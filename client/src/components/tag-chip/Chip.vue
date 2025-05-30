<script setup lang="ts">
import store from "@/store";
import { computed, ref, useAttrs, useTemplateRef } from "vue";
import ChipStatic from "./ChipStatic.vue";
import { useDismiss } from "@/composable";

const props = defineProps<{ jiggle?: boolean; state: TagState; tag: Tag }>();
const { tag } = props;
const showOptions = ref(false);
const containerRef = useTemplateRef("container");

useDismiss(containerRef, () => (showOptions.value = false));

const isIncluded = computed(() => store.query.include.has(tag.name));
const isExcluded = computed(() => store.query.exclude.has(tag.name));

function onAdd() {
    store.query.includeTag(tag);
    showOptions.value = false;
}

function onExclude() {
    store.query.excludeTag(tag);
    showOptions.value = false;
}

function onRemove() {
    store.query.removeTag(tag);
    showOptions.value = false;
}
</script>

<template>
    <div class="chip-container" ref="container">
        <ChipStatic v-bind="props" @click="showOptions = !showOptions" />
        <div class="options" v-if="showOptions">
            <button
                class="btn-primary option-btn"
                v-if="isExcluded || !isIncluded"
                @click="onAdd"
            >
                <i class="bi bi-plus-lg"></i> include
            </button>
            <button
                class="btn-primary option-btn"
                v-if="!isExcluded || isIncluded"
                @click="onExclude"
            >
                <i class="bi bi-dash-lg"></i> exclude
            </button>
            <button
                class="btn-primary option-btn"
                v-if="isExcluded || isIncluded"
                @click="onRemove"
            >
                <i class="bi bi-x-lg"></i> remove
            </button>
            <button class="btn-primary option-btn blacklist" disabled>
                <i class="bi bi-ban"></i> blacklist
            </button>
        </div>
    </div>
</template>

<style lang="scss" scoped>
@import "@/assets/buttons";
@import "@/assets/colors";

.chip-container {
    display: inline-block;
    position: relative;
}

.options {
    position: absolute;
    left: 0;
    display: flex;
    flex-direction: column;
    z-index: 2;
    width: max-content;
    box-shadow: 0 0 10px black;
}

.option-btn {
    text-align: left;

    .bi {
        margin-right: 5px;
    }

    &:first-of-type {
        border-radius: 4px 4px 0 0;
    }

    &:last-of-type {
        border-radius: 0 0 4px 4px;
    }

    &:not(:last-of-type) {
        border-bottom: none;
    }
}

.blacklist {
    color: #ff5d5d;
}
</style>
