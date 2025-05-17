<script setup lang="ts">
import store from "@/store";

function onChangeColCount(e: Event) {
    store.settings.columnCount = parseInt((e.target as HTMLInputElement).value);
    store.settings.save();
}

function onChangeColSizing(e: Event) {
    store.settings.columnSizing = (e.target as HTMLInputElement).value as any;
    store.settings.save();
}

function onChangeColWidth(e: Event) {
    store.settings.columnWidth = parseInt((e.target as HTMLInputElement).value);
    store.settings.save();
}

function onChangeCloseSidebarOnSearch(e: Event) {
    store.settings.closeSidebarOnSearch = (
        e.target as HTMLInputElement
    ).checked;
    store.settings.save();
}
</script>

<template>
    <div class="settings-container">
        <div class="input-group">
            <label>column sizing</label>
            <div class="input">
                <select
                    :value="store.settings.columnSizing"
                    @change="onChangeColSizing"
                >
                    <option>fixed</option>
                    <option>dynamic</option>
                </select>
            </div>
        </div>

        <div class="input-group">
            <label>column count</label>
            <div class="input">
                <input
                    type="range"
                    min="1"
                    max="20"
                    step="1"
                    :value="store.settings.columnCount"
                    :disabled="store.settings.columnSizing !== 'fixed'"
                    @input="onChangeColCount"
                />
                <span class="value">{{ store.settings.columnCount }}</span>
            </div>
        </div>

        <div class="input-group">
            <label>column width</label>
            <div class="input">
                <input
                    type="range"
                    min="100"
                    max="1000"
                    step="10"
                    :value="store.settings.columnWidth"
                    :disabled="store.settings.columnSizing !== 'dynamic'"
                    @input="onChangeColWidth"
                />
                <span class="value">{{ store.settings.columnWidth }}</span>
            </div>
        </div>

        <div class="input-group">
            <label>
                <input
                    type="checkbox"
                    :checked="store.settings.closeSidebarOnSearch"
                    @change="onChangeCloseSidebarOnSearch"
                />
                searching closes sidebar</label
            >
        </div>
    </div>
</template>

<style lang="scss" scoped>
.settings-container {
}

.input {
    display: flex;

    input,
    select {
        flex: 1;
    }

    .value {
        min-width: 60px;
        text-align: center;
    }
}

.input-group {
    margin-top: 20px;
}
</style>
