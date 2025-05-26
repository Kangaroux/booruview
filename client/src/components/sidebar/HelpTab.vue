<script setup lang="ts">
const emit = defineEmits(["on-close"]);
</script>

<template>
    <p>
        Booruview lets you search and browse
        <a href="https://gelbooru.com" target="_blank">Gelbooru</a>
        posts.
    </p>
    <p>
        Add tags to your search one at a time. To exclude a tag, use
        <code>-tag</code>
        . Results must match every tag in your search (AND).
    </p>
    <p>SFW Example:</p>
    <p>
        Type
        <code>rating:general</code>
        and confirm the tag. Type
        <code>blue_sky</code>
        and confirm. Click search.
    </p>
    <p>
        <a
            href="https://gelbooru.com/index.php?page=wiki&s=&s=view&id=26263"
            target="_blank"
        >
            Gelbooru Search Help
        </a>
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
</style>
