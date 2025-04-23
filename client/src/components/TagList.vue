<script setup lang="ts">
import { defineProps, ref, useTemplateRef, watch } from "vue";

const { tags } = defineProps<{tags: Tag[]}>();
const focusIndex = ref(0);

const listRef = useTemplateRef("list");

function doFocus(index: number) {
    if(!listRef.value || tags.length === 0)
        return;

    const item = listRef.value.children.item(index) as HTMLLIElement;
    focusIndex.value = index;
    item.focus();
}

function selectPrev() {
    if(tags.length === 0)
        return;

    let val = focusIndex.value - 1;
    if(val < 0)
        val = tags.length - 1;

    doFocus(val);
}

function selectNext() {
    if(tags.length === 0)
        return;

    let val = focusIndex.value + 1;
    if(val >= tags.length)
        val = 0;

    doFocus(val);
}

function onFocus() {
    doFocus(0);
}
</script>

<template>
    <ul ref="list"
        class="tag-list"
        tabindex="0"
        @focus="onFocus"
        @keydown.down.prevent="selectNext"
        @keydown.up.prevent="selectPrev"
        >
        <template v-for="tag, i in tags" :key="tag.name">
            <li class="list-item" :class="tag.type" :title="tag.name" tabindex="-1">
                <span class="name">{{ tag.name }}</span>
                <span class="count">{{ tag.count }}</span>
            </li>
        </template>
    </ul>
</template>

<style scoped>
.tag-list {
    margin: 0;
    padding: 0;

    /* Renders the list above the focus outline of the search box */
    position: relative;
    z-index: 1;
}

.list-item {
    background-color: #252525;
    border: 1px solid #555;
    color: #000;
    padding: 4px 8px;
    font-size: 16px;
    list-style: none;
    margin-top: -1px;
    display: flex;
    justify-content: space-between;
    cursor: pointer;
}

.list-item:hover, .list-item:focus {
    background-color: #303030;
}

.name {
    text-overflow: ellipsis;
    overflow: hidden;
}

.count {
    color: #aaa;
}

.tag {
    color: hsl(208, 56%, 75%);
}

.artist {
    color: #A00;
}

.copyright {
    color: #A0A;
}

.character {
    color: #0A0;
}

.metadata {
    color: #F80;
}

.deprecated, .unknown {
    color: #666;
}
</style>
