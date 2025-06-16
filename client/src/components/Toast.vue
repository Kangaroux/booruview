<script setup lang="ts">
type ToastKind = "danger";
const { kind } = defineProps<{ kind: ToastKind }>();
defineEmits(["dismiss"]);
</script>

<template>
    <div
        class="toast"
        :class="{ [`toast-${kind}`]: true }"
        @click="$emit('dismiss')"
    >
        <slot></slot>
    </div>
</template>

<style lang="scss" scoped>
.toast {
    position: absolute;
    bottom: 20px;
    left: 50%;
    transform: translateX(-50%);
    z-index: 999;
    border-radius: 10px;
    padding: 10px 15px;
    font-size: 16px;
    cursor: pointer;
}

.toast-danger {
    background-color: #700;
}
</style>
