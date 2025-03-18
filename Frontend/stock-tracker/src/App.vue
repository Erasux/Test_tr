<script setup>
import StockList from "./components/StockList.vue";
import StockRecommendations from "./components/StockRecommendations.vue";
</script>

<template>
  <div class="container mt-4">
    <!-- Tabla de Todas las Acciones -->
    <StockList />

    <!-- Espaciado entre tablas -->
    <div class="my-5"></div>

    <!-- Tabla de Recomendaciones -->
    <StockRecommendations />
  </div>
</template>
