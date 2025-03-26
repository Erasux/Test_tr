<template>
    <div class="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded relative" role="alert"
        aria-live="polite">
        <div class="flex items-center">
            <div class="flex-shrink-0">
                <svg class="h-5 w-5 text-red-400" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
                    <path fill-rule="evenodd"
                        d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.28 7.22a.75.75 0 00-1.06 1.06L8.94 10l-1.72 1.72a.75.75 0 101.06 1.06L10 11.06l1.72 1.72a.75.75 0 101.06-1.06L11.06 10l1.72-1.72a.75.75 0 00-1.06-1.06L10 8.94 8.28 7.22z"
                        clip-rule="evenodd" />
                </svg>
            </div>
            <div class="ml-3">
                <h3 class="text-sm font-medium">
                    {{ title || 'Error' }}
                </h3>
                <div class="mt-1 text-sm">
                    {{ message }}
                </div>
            </div>
            <div class="ml-auto pl-3">
                <button v-if="dismissible" type="button"
                    class="inline-flex rounded-md p-1.5 text-red-500 hover:bg-red-200 focus:outline-none focus:ring-2 focus:ring-red-600 focus:ring-offset-2"
                    @click="$emit('dismiss')" aria-label="Cerrar mensaje de error">
                    <svg class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                        <path
                            d="M6.28 5.22a.75.75 0 00-1.06 1.06L8.94 10l-3.72 3.72a.75.75 0 101.06 1.06L10 11.06l3.72 3.72a.75.75 0 101.06-1.06L11.06 10l3.72-3.72a.75.75 0 00-1.06-1.06L10 8.94 6.28 5.22z" />
                    </svg>
                </button>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
interface Props {
    message: string
    title?: string
    dismissible?: boolean
}

defineProps<Props>()
defineEmits<{
    (e: 'dismiss'): void
}>()
</script>
