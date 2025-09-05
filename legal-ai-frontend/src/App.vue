<template>
  <div id="app" class="min-h-screen flex flex-col bg-white">
    <Navbar />
    <main class="flex-1 flex justify-center items-start">
      <router-view />
    </main>
  </div>
</template>

<script setup>
import Navbar from './components/Navbar.vue'
</script>
