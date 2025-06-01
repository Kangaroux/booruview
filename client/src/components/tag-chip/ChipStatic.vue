<script setup lang="ts">
import { computed, onMounted, ref } from "vue";

const {
    jiggle = false,
    state,
    tag,
} = defineProps<{ jiggle?: boolean; state: TagState; tag: Tag }>();
const hasJiggled = ref(false);

const cls = computed(() => ({
    [`state-${state}`]: true,
    [tag.type]: true,
    jiggle: jiggle && !hasJiggled,
}));

onMounted(() => {
    if (jiggle) {
        // Prevents the jiggle animation from playing when the sidebar
        // is opened (display:none triggers animations)
        setTimeout(() => (hasJiggled.value = true), 1000);
    }
});
</script>

<template>
    <div class="chip" :class="cls">
        <i class="bi bi-check-lg" v-if="state === 'include'"></i>
        {{ tag.name
        }}<span class="dep-warning" v-if="tag.type === 'deprecated'">
            (deprecated)</span
        >
    </div>
</template>

<style lang="scss" scoped>
@import "@/assets/colors";

.chip {
    padding: 8px;
    margin: 0 4px 4px 0;
    border: none;
    border-radius: 8px;
    display: inline-block;
    font-size: 16px;
    word-break: break-all;
    cursor: pointer;

    &.state-exclude {
        filter: brightness(0.8);
        text-decoration: line-through;
    }

    .dep-warning {
        color: #f44;
    }
}

.chip-options {
    button {
        border: 1px solid white;
        background: $color-lightgray;
        color: white;
    }
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

.deprecated,
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

.unknown {
    background-color: #6275ae;
    color: #0b1227;
}
</style>
