<template>
  <q-btn flat round dense color="negative" icon="delete" @click.stop>
    <q-tooltip>删除</q-tooltip>
    <q-menu auto-close anchor="center left" self="center right">
      <div class="row no-wrap items-center q-pa-sm bg-warning text-white">
        <span class="text-no-wrap"
          >确认删除 <b>{{ label }}</b> ?</span
        >
        <q-btn
          flat
          round
          dense
          color="negative"
          icon="mdi-delete-empty"
          @click="$emit('confirm')"
        >
          <q-tooltip>确认删除</q-tooltip></q-btn
        >
      </div>
    </q-menu>
  </q-btn>
</template>
<script>
export default {
  name: 'comDict',
  emits: ['confirm'],
}
</script>
<script setup>
const props = defineProps({
  label: String,
})
</script>
