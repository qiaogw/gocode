<template>
  <div>
    <q-btn
      dense
      round
      flat
      :icon="`${!right ? 'settings' : 'close'}`"
      @click="right = !right"
    />
    <q-dialog v-model="right" position="right">
      <q-card style="height: 100%; width: 400px" class="text-center">
        <q-card-section>
          <div class="text-h6">{{ $t('System') + $t('Setting') }}</div>
        </q-card-section>
        <q-card-section>
          <DarkTheme />

          <q-separator />
        </q-card-section>
      </q-card>
    </q-dialog>
  </div>
</template>

<script setup>
// import comTheme from "src/components/comTheme";
import DarkTheme from 'src/components/comTheme/DarkTheme.vue'
import { ref } from 'vue'

const right = ref(false)
</script>
