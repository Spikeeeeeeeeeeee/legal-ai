<template>
  <div class="border-2 border-dashed border-gray-300 rounded-lg p-8 text-center bg-gray-50">
    <div class="flex flex-col items-center">
      <!-- Upload Icon -->
      <div class="text-3xl text-blue-500">⬇</div>
      <p class="mt-2 text-gray-600">Drop your legal document here</p>
      <p class="text-sm text-gray-500">Supports PDF, DOC, DOCX, TXT (up to 10MB)</p>

      <!-- File Input -->
      <input type="file" class="hidden" id="fileUpload" @change="handleFileUpload" />
      <label
        for="fileUpload"
        class="mt-4 px-6 py-2 bg-teal-600 text-white rounded-lg cursor-pointer hover:bg-teal-700"
      >
        Choose File
      </label>
    </div>

    <!-- Show uploaded file name -->
    <div v-if="fileName" class="mt-4 text-gray-700">
      Uploaded: <span class="font-medium">{{ fileName }}</span>
    </div>

    <!-- Demo Button -->
    <div class="text-center mt-6">
      <button
        @click="showResults"
        class="px-6 py-2 border rounded-lg hover:bg-gray-100"
      >
        Try Demo Document
      </button>
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue"

const fileName = ref("")
const emit = defineEmits(['next'])

function handleFileUpload(event) {
  const file = event.target.files[0]
  if (file) {
    fileName.value = file.name
    showResults()
  }
}

function showResults() {
  emit('next', 'results')
}
</script>
