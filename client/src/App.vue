<script setup lang="ts">
import { ref } from "vue";
import store from "@/store";
import PostContainer from "./components/PostContainer.vue";
import Sidebar from "./components/Sidebar.vue";
import SearchHelp from "./components/SearchHelp.vue";
import FullscreenView from "./components/fullscreen-view/FullscreenView.vue";

const showHelp = ref(localStorage.getItem("hide-help") === null);

function onCloseHelp() {
    showHelp.value = false;
    localStorage.setItem("hide-help", "1");
}
</script>

<template>
    <div
        class="app"
        :class="{
            'sidebar-closed': store.sidebarClosed,
            'sidebar-open': !store.sidebarClosed,
        }"
    >
        <Sidebar
            :closed="store.sidebarClosed"
            @toggle="store.sidebarClosed = !store.sidebarClosed"
        />
        <main>
            <SearchHelp
                v-if="showHelp && !store.hasResults()"
                @on-close="onCloseHelp"
            />
            <FullscreenView v-if="store.fullscreenPost !== null" />
            <div class="main-content" v-if="store.hasResults()">
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
            </div>
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
}

main {
    flex: 1;

    @media (max-width: 600px) {
        .sidebar-open & {
            display: none;
        }
    }
}

.main-content {
    overflow-y: scroll;
    height: 100%;
}
</style>
