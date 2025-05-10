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
        <Sidebar />
        <main>
            <SearchHelp
                v-if="showHelp && store.hasResults()"
                @on-close="onCloseHelp"
            />
            <template v-if="store.hasResults()">
                <PostContainer :posts="store.postsForCurrentPage() || []" />
                <footer>
                    <p>
                        page {{ store.currentPage }} of
                        {{ store.maxPage() }} ({{ store.totalPostCount }}
                        results)
                    </p>
                    <button
                        @click="store.prevPage()"
                        v-if="store.currentPage > 1"
                    >
                        &lt;&lt; prev page
                    </button>
                    <button
                        @click="store.nextPage()"
                        v-if="store.currentPage < store.maxPage()"
                    >
                        next page &gt;&gt;
                    </button>
                </footer>
            </template>
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

main {
    flex: 1;
    overflow-y: scroll;
}
</style>
