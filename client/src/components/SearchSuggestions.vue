<script setup lang="ts">
import { useTemplateRef, watchPostEffect } from "vue";

const emit = defineEmits<{ onClick: [index: number] }>();

const props = defineProps<{
    selectedIndex: number;
    tags: Tag[];
}>();

const listRef = useTemplateRef("list");

watchPostEffect(() => {
    if (!listRef.value || props.tags.length === 0 || props.selectedIndex < 0) {
        return;
    }

    const item = listRef.value.children.item(props.selectedIndex);

    if (!item) {
        console.error("shouldn't happen");
    } else {
        (item as HTMLLIElement).focus();
    }
});
</script>

<template>
    <ul ref="list" class="tag-list" tabindex="0">
        <template v-for="(tag, i) in tags" :key="tag.name">
            <li
                class="list-item"
                :class="tag.type"
                :title="tag.name"
                tabindex="-1"
                @click="$emit('onClick', i)"
            >
                <span class="name">{{ tag.name }}</span>
                <span class="count" v-if="tag.type !== 'unknown'">{{ tag.count }}</span>
            </li>
        </template>
    </ul>
</template>

<style scoped lang="scss">
.tag-list {
    margin: 0;
    padding: 0;
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

    &:hover,
    &:focus {
        background-color: #303030;
    }
}

.name {
    text-overflow: ellipsis;
    overflow: hidden;
}

.count {
    color: #aaa;
}

.tag {
    color: #9cc2e3;
}

.artist {
    color: #a00;
}

.copyright {
    color: #a0a;
}

.character {
    color: #0a0;
}

.metadata {
    color: #f80;
}

.deprecated,
.unknown {
    color: #6275ae;
}
</style>
