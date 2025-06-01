<script setup lang="ts">
import store from "@/store";
</script>

<template>
    <footer class="page-nav">
        <button
            class="btn-primary btn-rounded"
            @click="store.prevPage()"
            v-if="store.currentPage > 1"
        >
            <i class="bi bi-arrow-left"></i> prev
        </button>
        <button
            class="btn-primary btn-rounded"
            @click="store.nextPage()"
            v-if="store.currentPage < store.maxPage()"
        >
            next <i class="bi bi-arrow-right"></i>
        </button>
        <p>
            page {{ store.currentPage }} of {{ store.maxPage() }} ({{
                store.totalPostCount
            }}
            results)
        </p>
    </footer>
</template>

<style lang="scss" scoped>
@import "@/assets/buttons";

.page-nav {
    margin-top: 40px;
    text-align: center;

    .btn-primary:nth-of-type(2) {
        margin-left: 10px;
    }
}
</style>
