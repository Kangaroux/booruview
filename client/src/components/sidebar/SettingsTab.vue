<script setup lang="ts">
import store from "@/store";

function onChangeColWidth(e: Event) {
    store.settings.columnWidth = parseInt((e.target as HTMLInputElement).value);
    store.settings.save();
}
</script>

<template>
    <div class="settings-container">
        <div class="input-group">
            <label>column width</label>
            <div class="input">
                <input
                    type="range"
                    min="100"
                    max="1000"
                    :value="store.settings.columnWidth"
                    @input="onChangeColWidth"
                />
                <span class="value">{{ store.settings.columnWidth }}</span>
            </div>
        </div>
    </div>
</template>

<style lang="scss" scoped>
.settings-container {
}

.input {
    display: flex;

    input {
        flex: 1;
    }

    .value {
        margin-left: 10px;
    }
}
</style>
