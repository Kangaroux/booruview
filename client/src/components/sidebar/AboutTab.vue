<script setup lang="ts">
import { COMMIT_SHA } from "@/config";
</script>

<template>
    <p>
        Booruview is
        <a href="https://github.com/Kangaroux/booru-viewer" target="_blank"
            >open source</a
        >
        and development is ongoing.
    </p>
    <p>
        Feedback and suggestions are welcome. You can use the Github issue
        tracker or send me an
        <a href="mailto:2302541+Kangaroux@users.noreply.github.com">email</a>.
    </p>
    <p>
        This site does not use tracking or cookies. Searches are cached briefly
        and entirely anonymous.
    </p>
    <p class="version" v-if="COMMIT_SHA !== undefined">ver. {{ COMMIT_SHA }}</p>
</template>

<style scoped>
.version {
    font-family: "Courier New", Courier, monospace;
    font-size: 12px;
    text-align: right;
}
</style>
