<template>
    <div class="text-center p-6 bg-gray-50 rounded-lg" role="status" aria-live="polite">
        <div class="flex flex-col items-center">
            <div class="rounded-full bg-gray-100 p-3">
                <svg class="h-12 w-12 text-gray-400" fill="none" viewBox="0 0 24 24" stroke="currentColor"
                    aria-hidden="true">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                        d="M20 13V6a2 2 0 00-2-2H6a2 2 0 00-2 2v7m16 0v5a2 2 0 01-2 2H6a2 2 0 01-2-2v-5m16 0h-2.586a1 1 0 00-.707.293l-2.414 2.414a1 1 0 01-.707.293h-3.172a1 1 0 01-.707-.293l-2.414-2.414A1 1 0 006.586 13H4" />
                </svg>
            </div>
            <h3 class="mt-4 text-sm font-medium text-gray-900">
                {{ title }}
            </h3>
            <p class="mt-1 text-sm text-gray-500">
                {{ description }}
            </p>
            <div v-if="$slots.action" class="mt-6">
                <slot name="action" />
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
withDefaults(defineProps<{
    title?: string
    description?: string
}>(), {
    title: 'No hay datos disponibles',
    description: 'No se encontraron registros que coincidan con los criterios de búsqueda.'
})
</script>
