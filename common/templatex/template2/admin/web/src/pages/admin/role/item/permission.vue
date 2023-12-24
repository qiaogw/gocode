<template>
  <q-dialog v-model="permissionVisible" persistent maximized position="right">
    <q-card style="max-width: 80vw; min-width: 60vw">
      <q-bar class="bg-primary text-white">
        <span> {{ row.name }}权限 </span>
        <q-space />
        <q-btn dense flat icon="close" v-close-popup>
          <q-tooltip class="bg-white text-primary">Close</q-tooltip>
        </q-btn>
      </q-bar>
      <q-tabs
        v-model="tab"
        dense
        class="text-grey"
        active-color="primary"
        indicator-color="primary"
        narrow-indicator
      >
        <q-tab name="menu" label="菜单权限" />
        <q-tab name="api" label="api权限" />
        <q-tab name="data" label="数据权限" />
      </q-tabs>
      <q-separator />

      <q-tab-panels v-model="tab" animated>
        <q-tab-panel name="menu">
          <role-menu v-model:row="row" v-if="tab === 'menu'" @submit="submit" />
        </q-tab-panel>

        <q-tab-panel name="api">
          <role-api v-model:row="row" v-if="tab === 'api'" @submit="submit" />
        </q-tab-panel>

        <q-tab-panel name="data">
          <role-data v-model:row="row" v-if="tab === 'data'" />
        </q-tab-panel>
      </q-tab-panels>
    </q-card>
  </q-dialog>
</template>

<script setup>
import { ref, watch } from 'vue'
import RoleMenu from './menu.vue'
import RoleApi from './api.vue'
import RoleData from './data.vue'
import { getRole } from 'src/api/admin/role'

const permissionVisible = ref(false)
const row = ref({})
const tab = ref('menu')
const show = async (record) => {
  row.value = await getRole(record)
  permissionVisible.value = true
}
const submit = async () => {
  row.value = []
  permissionVisible.value = false
}

defineExpose({
  show,
})
</script>
