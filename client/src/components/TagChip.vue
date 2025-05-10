<script setup lang="ts">
import { onMounted, ref } from "vue";

const { jiggle, tag } = defineProps<{ jiggle?: boolean; tag: Tag }>();
const hasJiggled = ref(false);

onMounted(() => {
    if (jiggle) {
        // Prevents the jiggle animation from playing when the sidebar
        // is opened (display:none triggers animations)
        setTimeout(() => (hasJiggled.value = true), 1000);
    }
});
</script>

<template>
    <div
        class="chip"
        :class="{ [tag.type]: true, jiggle: jiggle && !hasJiggled }"
    >
        {{ tag.name }}
    </div>
</template>

<style scoped>
.chip {
    padding: 8px;
    margin-right: 4px;
    border: none;
    border-radius: 8px;
    display: inline-block;
    font-size: 16px;
}

@keyframes jiggle-anim {
    0% {
        transform: scale(0.9);
    }

    33% {
        transform: scale(1.1);
    }

    66% {
        transform: scale(0.95);
    }

    100% {
        transform: scale(1);
    }
}

.jiggle {
    animation: 300ms linear 0s jiggle-anim;
}

.tag {
    background-color: #303030;
    color: hsl(208, 56%, 75%);
}

.artist {
    background-color: #a00;
    color: #fff;
}

.copyright {
    background-color: #a0a;
    color: #fff;
}

.character {
    background-color: #0a0;
}

.metadata {
    background-color: #f80;
}

.deprecated,
.unknown {
    background-color: #6275ae;
    color: #0b1227;
}
</style>
