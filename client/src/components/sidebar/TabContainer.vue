<script setup lang="ts">
import { computed, ref } from "vue";
import AboutTab from "./AboutTab.vue";
import HelpTab from "./HelpTab.vue";
import SettingsTab from "./SettingsTab.vue";
import store from "@/store";

type Tab = "about" | "help" | "settings";
const currentTab = ref<Tab>("help");

// CBA writing the full setting each time
const closed = computed({
    get: () => store.settings.sidebarTabsHidden,
    set: (val: boolean) => (store.settings.sidebarTabsHidden = val),
});

function switchTab(tab: Tab) {
    currentTab.value = tab;
    closed.value = false;
    store.settings.save();
}

function toggleClose() {
    closed.value = !closed.value;
    store.settings.save();
}
</script>

<template>
    <div class="tab-container" :class="{ closed }">
        <header class="tabs">
            <button
                class="tab-btn"
                :class="{ active: currentTab === 'about' && !closed }"
                @click="switchTab('about')"
            >
                about
            </button>
            <button
                class="tab-btn"
                :class="{ active: currentTab === 'help' && !closed }"
                @click="switchTab('help')"
            >
                help
            </button>
            <button
                class="tab-btn"
                :class="{ active: currentTab === 'settings' && !closed }"
                @click="switchTab('settings')"
            >
                settings
            </button>

            <button
                class="tab-btn close-btn"
                :class="{ active: closed }"
                @click="toggleClose"
            >
                <i
                    class="bi"
                    :class="{
                        'bi-chevron-down': !closed,
                        'bi-chevron-up': closed,
                    }"
                ></i>
            </button>
        </header>

        <div class="tab-content" v-if="!closed">
            <KeepAlive>
                <HelpTab v-if="currentTab === 'help'" />
                <AboutTab v-else-if="currentTab === 'about'" />
                <SettingsTab v-else-if="currentTab === 'settings'" />
            </KeepAlive>
        </div>
    </div>
</template>

<!-- NOTE: no `scoped` here -->
<style lang="scss">
.tab-container {
    display: flex;
    flex-direction: column;

    &:not(.closed) {
        .tabs {
            border-bottom: 1px solid #555;
        }
    }
}

.tabs {
    display: flex;
    padding: 0 10px;
    gap: 5px;
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

    &.active {
        border-color: #695675;
        color: #bb9fce;
        background-color: #342b3a;
    }
}

.close-btn {
    margin-left: auto;
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

select {
    background-color: #222;
    border: 1px solid #555;
    color: #999;
    padding: 5px;
    border-radius: 5px;
}
</style>
