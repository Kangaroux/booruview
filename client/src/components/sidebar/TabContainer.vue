<script setup lang="ts">
import { ref } from "vue";
import InfoTab from "./AboutTab.vue";

type Tab = "about" | "settings";
const currentTab = ref<Tab>("about");
</script>

<template>
    <div class="tab-container">
        <header class="tabs">
            <button
                class="tab-btn"
                :class="{ active: currentTab === 'about' }"
                @click="currentTab = 'about'"
            >
                about
            </button>
            <button
                class="tab-btn"
                :class="{ active: currentTab === 'settings' }"
                @click="currentTab = 'settings'"
            >
                settings
            </button>
        </header>

        <div class="tab-content">
            <KeepAlive>
                <InfoTab v-if="currentTab === 'about'" />
            </KeepAlive>
        </div>
    </div>
</template>

<!-- NOTE: no `scoped` here -->
<style lang="scss">
.tab-container {
    margin-top: 40px;
    flex: 1;
    display: flex;
    flex-direction: column;
    min-height: 0;
}

.tabs {
    border-bottom: 1px solid #555;
    padding: 0 10px;
}

.tab-btn {
    border: 1px solid #555;
    border-bottom: none;
    padding: 5px 10px;
    border-radius: 4px 4px 0 0;
    background-color: #1e1e1e;
    color: #999;
    cursor: pointer;
    font-size: 16px;
    margin-right: 5px;

    &.active {
        border-color: #695675;
        color: #bb9fce;
        background-color: #342b3a;
    }
}

.tab-content {
    padding: 0 10px;
    flex: 1;
    min-height: 0;
    overflow-y: scroll;
}

p {
    font-size: 16px;
    color: #999;

    a,
    a:visited {
        color: #bb9fce;
    }
}
</style>
