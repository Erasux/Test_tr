<template>
  <div>
    <NavigationBar />
    <router-view></router-view>
  </div>
</template>

<script setup lang="ts">
import NavigationBar from "@/components/Navbar.vue";
</script>
