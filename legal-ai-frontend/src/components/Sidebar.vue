<template>
  <aside class="w-72 min-h-screen bg-gray-50 border-r shadow-sm flex flex-col">
    <!-- Header -->
    <div class="p-6 border-b">
      <h2 class="font-bold text-lg text-teal-700 leading-snug">
        📑 LegalClear AI
      </h2>
      <p class="text-sm text-gray-500">
        Lease Agreement - Clause 12
      </p>
    </div>

    <!-- Menu -->
    <nav class="flex-1 p-4 space-y-2">
      <button
        v-for="item in menu"
        :key="item"
        class="w-full text-left px-4 py-3 rounded-lg transition
              hover:bg-teal-50 hover:text-teal-700
              focus:bg-teal-600 focus:text-white
              active:bg-teal-700"
        :class="{
          'bg-teal-600 text-white font-semibold': active === item,
          'text-gray-700': active !== item
        }"
        @click="active = item"
      >
        {{ item }}
      </button>
    </nav>

    <!-- Footer -->
    <div class="p-4 border-t text-xs text-gray-500">
      © 2025 LegalClear AI
    </div>
  </aside>
</template>

<script setup>
import { ref } from "vue"

const active = ref("Executive Summary")

const menu = [
  "Analyze New Document",
  "Executive Summary",
  "Key Terms",
  "Potential Concerns",
  "Ask Questions",
]
</script>
