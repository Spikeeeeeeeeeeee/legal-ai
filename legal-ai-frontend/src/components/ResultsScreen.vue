<template>
  <div class="flex">
    <!-- Sidebar -->
    <Sidebar />

    <!-- Main Content -->
    <main class="flex-1 p-10 bg-gray-50 min-h-screen">
      <h2 class="text-3xl font-bold mb-8 text-gray-800">Executive Summary</h2>

      <div class="bg-white border rounded-xl shadow-md p-6 space-y-6">
        <!-- Original Text -->
        <section>
          <h3 class="text-lg font-semibold mb-2 text-gray-800">Original Text</h3>
          <p class="text-sm text-gray-700 bg-red-50 border border-red-200 p-4 rounded-md leading-relaxed">
            Tenant hereby agrees and covenants that Tenant shall not, without prior written consent of
            Landlord, which consent may be withheld in Landlord’s sole and absolute discretion, make any
            alterations, additions, or improvements to the Premises...
          </p>
        </section>

        <!-- Plain English Summary -->
        <section>
          <h3 class="text-lg font-semibold mb-2 text-gray-800">Plain English Summary</h3>
          <p class="text-sm text-gray-700 bg-green-50 border border-green-200 p-4 rounded-md leading-relaxed">
            You cannot make changes to your rental unit (like painting, adding fixtures, or renovating)
            without getting written permission from your landlord first...
          </p>
        </section>
      </div>

      <!-- Download Buttons -->
      <div class="mt-8 flex space-x-4">
        <button class="bg-teal-600 text-white px-6 py-3 rounded-lg shadow hover:bg-teal-700">
          Download Summary (PDF)
        </button>
        <button class="bg-gray-200 px-6 py-3 rounded-lg shadow hover:bg-gray-300">
          Share Results
        </button>
      </div>

      <!-- Footer -->
      <footer class="mt-16 border-t pt-6 text-sm text-gray-600">
        <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
          <div>
            <p class="font-semibold">LegalClear AI</p>
            <p>Making legal documents accessible to everyone</p>
          </div>
          <div>
            <p class="font-semibold">Privacy & Security</p>
            <p>Your documents are processed securely and never stored</p>
            <p>End-to-end encryption • GDPR compliant • SOC 2 certified</p>
          </div>
          <div>
            <p class="font-semibold">Support</p>
            <p>Privacy Policy</p>
            <p>Terms of Service</p>
            <p>Contact</p>
          </div>
        </div>
        <p class="mt-6 text-center text-xs">© 2025 LegalClear AI. All rights reserved.</p>
      </footer>
    </main>
  </div>
</template>

<script setup>
import Sidebar from "./Sidebar.vue"
</script>
