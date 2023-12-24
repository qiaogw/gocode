<template>
  <q-page class="q-pa-xs">
    <div class="shadow-2 q-pa-xs">
      <q-card>
        <q-tabs
          v-model="tab"
          inline-label
          align="justify"
          narrow-indicator
          class="bg-purple text-white shadow-2"
        >
          <q-tab name="mails" icon="mail" label="Mails" />
          <q-tab name="webhook" icon="login" label="Webhook" />
        </q-tabs>
        <q-separator />

        <q-tab-panels v-model="tab" animated>
          <q-tab-panel name="mails">
            <taskMail />
          </q-tab-panel>

          <q-tab-panel name="webhook">
            <div class="text-h6">暂时不可用！</div>
            暂时不可用！
          </q-tab-panel>
        </q-tab-panels>
      </q-card>
    </div>
  </q-page>
</template>

<script setup>
import {
  computed,
  onMounted,
  ref,
  reactive,
  watch,
  getCurrentInstance,
} from 'vue'
import {
  listSetting,
  createSetting,
  updateSetting,
  deleteSetting,
} from 'src/api/task/setting'

import { useQuasar } from 'quasar'
import { requiredRule } from 'src/utils/inputRule'
import {
  DictOptions,
  getOptionsByList,
  getDictLabel,
  getDict,
} from 'src/utils/dict'
import { downloadAction } from 'src/api/manage'
import { useRoute } from 'vue-router'

import taskMail from './mail.vue'

const route = useRoute()

const $q = useQuasar()
let { proxy } = getCurrentInstance()

const tab = ref('mails')

const dialogVisible = ref(false)
const dataList = ref([])
const formType = ref('')
const dictId = ref(0)
const form = ref({})
const dictOptions = ref({})
const taskHookTypeOptions = ref([])
const formattaskHookType = (prop) => {
  if (!prop) {
    prop = false
  }
  return getDictLabel(dictOptions.value.taskHookType, prop)
}
const sys_enabledOptions = ref([])
const formatsys_enabled = (prop) => {
  if (!prop) {
    prop = false
  }
  return getDictLabel(dictOptions.value.sys_enabled, prop)
}

const searchKey = ref('')

const selected = ref([])
const pagination = ref({
  sortBy: 'id',
  descending: false,
  page: 1,
  rowsPerPage: 10,
  rowsNumber: 0,
})
const columns = computed(() => {
  return [
    {
      name: 'code',
      align: 'left',
      label: '回掉类型',
      field: 'code',
      sortable: true,
      classes: 'ellipsis',
    },
    {
      name: 'key',
      align: 'left',
      label: '回掉模版key',
      field: 'key',
      sortable: true,
      classes: 'ellipsis',
    },
    {
      name: 'status',
      align: 'left',
      label: '状态',
      field: 'status',
      sortable: true,
      classes: 'ellipsis',
    },
    {
      name: 'value',
      align: 'left',
      label: '回掉模版',
      field: 'value',
      sortable: true,
      classes: 'ellipsis',
    },
    {
      name: 'remark',
      align: 'left',
      label: '备注',
      field: 'remark',
      sortable: true,
      classes: 'ellipsis',
    },

    { name: 'actions', align: 'center', label: '操作', field: 'actions' },
  ]
})

onMounted(async () => {
  dictOptions.value = await DictOptions()
  taskHookTypeOptions.value = await getDict('taskHookType')
  sys_enabledOptions.value = await getDict('sys_enabled')
  onRequest()
})

const reset = () => {
  pagination.value = {
    sortBy: 'id',
    descending: false,
    page: 1,
    rowsPerPage: 10,
    rowsNumber: 0,
  }
  form.value = {
    status: true,
  }
  dictId.value = 0
}

const onRequest = async (val) => {
  if (!val) {
    val = { pagination: pagination.value }
  }
  if (!val.pagination) {
    val.pagination = pagination.value
  }
  if (!val.filter) {
    val.filter = searchKey.value
  }
  const { page, rowsPerPage, sortBy, descending } = val.pagination
  let queryReq = {}
  queryReq.pageSize = rowsPerPage
  queryReq.pageIndex = page
  queryReq.sortBy = sortBy
  queryReq.descending = descending
  queryReq.searchKey = val.filter

  let table = await listSetting(queryReq)

  pagination.value = val.pagination

  pagination.value.rowsNumber = table.count
  if (table.list) {
    dataList.value = table.list
  }
}
const getSelectedString = () => {
  return selected.value.length === 0
    ? ''
    : `${selected.value.length} record${
        selected.value.length > 1 ? 's' : ''
      } selected of ${dataList.value.length}`
}

const create = () => {
  reset()
  formType.value = '新建'
  dialogVisible.value = true
}
const edit = (p) => {
  reset()
  form.value = {
    ...p,
  }
  formType.value = '编辑'
  dialogVisible.value = true
}

const del = async (p) => {
  let res = await deleteSetting(p)
  onRequest()
}

const delList = async () => {
  // let res = await deleteDatasource()
  // onRequest()
  let req = {
    ids: getIds(selected.value),
  }
  let res = await deleteSetting(req)
  onRequest()
}
const submit = async () => {
  // const res = undefined
  if (formType.value === '编辑') {
    let res = await updateSetting(form.value)
  } else if (formType.value === '新建') {
    let res = await createSetting(form.value)
  } else {
    proxy.$error('请求错误')
  }
  dialogVisible.value = false
  reset()
  onRequest()
}
</script>
