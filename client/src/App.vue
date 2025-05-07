<script setup lang="ts">
import { ref } from "vue";
import store from "@/store";
import PostContainer from "./components/PostContainer.vue";
import Sidebar from "./components/Sidebar.vue";
import SearchHelp from "./components/SearchHelp.vue";

const showHelp = ref(localStorage.getItem("hide-help") === null);

function onCloseHelp() {
    showHelp.value = false;
    localStorage.setItem("hide-help", "1");
}
</script>

<template>
    <div class="app">
        <header class="sidebar-container">
            <Sidebar />
        </header>
        <main>
            <SearchHelp
                v-if="showHelp && store.posts.length === 0"
                @on-close="onCloseHelp"
            />
            <PostContainer v-if="store.posts.length > 0" :posts="store.posts" />
        </main>
    </div>
</template>

<style scoped lang="scss">
@import "assets/colors";

.app {
    display: flex;
    flex-direction: row;
    width: 100%;
    height: 100%;
    gap: 10px;
}

header {
    position: relative;
}

main {
    flex: 1;
    overflow-y: scroll;
}
</style>
