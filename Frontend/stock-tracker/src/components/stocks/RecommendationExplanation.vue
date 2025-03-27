<template>
  <div class="bg-white p-4 rounded-lg shadow-md mb-6">
    <h2 class="text-xl font-semibold mb-4 text-gray-800">¿Cómo se calcula la recomendación?</h2>
    <div class="space-y-4">
      <div class="bg-gray-50 p-4 rounded-lg">
        <h3 class="font-medium text-gray-700 mb-2">Puntuación Total = Impacto de Precio + Impacto de Rating + Puntuación del Broker</h3>
        <div class="space-y-2 text-sm text-gray-600">
          <p><span class="font-medium">Impacto de Precio:</span></p>
          <ul class="list-disc pl-5">
            <li>+3 puntos si hay aumento de precio objetivo</li>
            <li>+2 puntos adicionales si el aumento es mayor a $10</li>
          </ul>
        </div>
      </div>

      <div class="bg-gray-50 p-4 rounded-lg">
        <p class="font-medium text-gray-700 mb-2">Impacto de Rating:</p>
        <ul class="list-disc pl-5 text-sm text-gray-600">
          <li>+3 puntos: Cambios positivos (ej: Sell → Buy, Market Perform → Buy)</li>
          <li>+2 puntos: Cambios moderados (ej: Neutral → Buy, Underweight → Market Perform)</li>
          <li>+1 punto: Cambios menores (ej: Sell → Neutral, Market Perform → Sell)</li>
        </ul>
      </div>

      <div class="bg-gray-50 p-4 rounded-lg">
        <p class="font-medium text-gray-700 mb-2">Recomendación Final:</p>
        <ul class="list-disc pl-5 text-sm text-gray-600">
          <li><span class="text-green-600 font-medium">Strong Buy:</span> Puntuación ≥ 7</li>
          <li><span class="text-green-500 font-medium">Buy:</span> Puntuación ≥ 5</li>
          <li><span class="text-gray-500 font-medium">Hold:</span> Puntuación ≥ 3</li>
          <li><span class="text-red-600 font-medium">Sell:</span> Puntuación < 3</li>
        </ul>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
// No se requiere lógica adicional
</script>
