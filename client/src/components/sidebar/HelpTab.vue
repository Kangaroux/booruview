<script setup lang="ts">
const emit = defineEmits(["on-close"]);
</script>

<template>
    <p>Quick reference:</p>
    <ul>
        <li>Searching without tags shows the most recent posts</li>
        <li><code>blue_sky</code> → search for blue_sky</li>
        <li><code>-blue_sky</code> → exclude blue_sky</li>
        <li><code>blue*</code> → tags that start with blue</li>
        <li><code>sort:score</code> → sort by most popular</li>
        <li><code>rating:general</code> → SFW only</li>
    </ul>
    <p>
        <a
            href="https://gelbooru.com/index.php?page=wiki&s=&s=view&id=26263"
            target="_blank"
            >Gelbooru search cheatsheet</a
        >
    </p>
</template>

<style lang="scss" scoped>
code {
    font-family: "Courier New", Courier, monospace;
    background-color: rgba(121, 121, 121, 0.2);
    padding: 0 3px;
    border: 1px solid #4c4c4c;
    border-radius: 2px;
}

ul {
    padding-left: 25px;
    margin: 16px 0;
}
</style>
