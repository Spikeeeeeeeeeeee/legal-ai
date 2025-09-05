<template>
  <div class="w-full max-w-4xl mx-auto text-center py-12">
    <!-- Heading -->
    <h2 class="text-3xl font-bold text-gray-800 mb-4">
      Understand Your Legal Documents in Plain English
    </h2>
    <p class="text-gray-600 max-w-2xl mx-auto mb-8">
      AI-powered analysis that transforms complex legal jargon into clear, accessible language
    </p>

    <!-- Step Logic -->
    <div>
      <UploadForm v-if="step === 'upload'" @next="goToStep" />
      <ResultsScreen v-if="step === 'results'" />
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import UploadForm from '../components/UploadForm.vue'
import ResultsScreen from '../components/ResultsScreen.vue'

const step = ref('upload')

function goToStep(nextStep) {
  step.value = nextStep
}
</script>
