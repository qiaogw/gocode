<template>
  <q-btn flat round dense color="primary" icon="mode_edit" @click.stop>
    <q-tooltip>修改</q-tooltip>
    <q-menu auto-close anchor="center left" self="center right">
      <div class="row no-wrap items-center q-pa-sm bg-cyan-2">
        <span class="text-no-wrap"
          >确认修改 <b>{{ label }}</b> ？</span
        >
        <q-btn
          flat
          round
          dense
          color="primary"
          icon="mode_edit"
          @click="$emit('confirm')"
        >
          <q-tooltip>确认修改</q-tooltip></q-btn
        >
      </div>
    </q-menu>
  </q-btn>
</template>

<script>
export default {
  name: 'BtnEdit',
  emits: ['confirm'],
}
</script>
<script setup>
const props = defineProps({
  label: String,
})
</script>
