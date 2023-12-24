<template>
  <q-card-section>
    <q-option-group
      v-model="row.dataScope"
      :options="dataScopeOptions"
      color="green"
    />
    <q-input
      dense
      outlined
      class="col-6"
      v-model="row.dataFilter"
      label="自定数据权限过滤条件"
      v-show="row.dataScope === '5'"
    />
  </q-card-section>
  <div class="justify-center row q-gutter-x-xs" style="width: 100%">
    <q-btn
      outline
      color="primary"
      icon="mdi-close-thick"
      label="关闭"
      v-close-popup
    />
    <q-btn color="primary" @click="handleSave">
      {{ $t('Save') }}
    </q-btn>
  </div>
</template>

<script setup>
import {
  computed,
  onMounted,
  onBeforeMount,
  ref,
  toRefs,
  watch,
  getCurrentInstance,
} from 'vue'
import { updateRole } from 'src/api/admin/role'
import {
  DictOptions,
  getOptionsByList,
  getDictLabel,
  getDict,
} from 'src/utils/dict'
import { getIds } from 'src/utils/arrayOrObject'
const dataScopeOptions = ref([])
const formatDataScope = (prop) => {
  return getDictLabel(dataScopeOptions.value, prop)
}

const props = defineProps({
  row: {
    type: Object,
    required: true,
  },
})
const { row } = toRefs(props)
onMounted(async () => {
  dataScopeOptions.value = await getDict('data_scope')
})
const handleSave = async () => {
  await updateRole(row.value)
}
</script>
