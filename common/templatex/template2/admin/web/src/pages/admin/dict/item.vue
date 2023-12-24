<template>
  <div>
    <q-table
      dense
      flat
      bordered
      separator="cell"
      :columns="columns"
      :rows="dataList"
      row-key="id"
      @request="onRequest"
      v-model:filter="searchKey"
      :grid="$q.screen.xs"
      v-model:selected="selected"
      v-model:pagination="pagination"
      binary-state-sort
    >
      <template v-slot:top="table">
        <div v-show="dictId" class="row no-wrap full-width">
          <com-search @query="onRequest" v-model:filter="searchKey" />
          <q-btn icon="autorenew" no-wrap color="primary" @click="reset"
            ><q-tooltip>重置</q-tooltip></q-btn
          >
          <q-space />
          <div class="q-gutter-xs">
            <q-btn icon="add" no-wrap color="primary" @click="create"
              ><q-tooltip>新建</q-tooltip></q-btn
            >
            <q-btn
              color="primary"
              no-wrap
              v-show="$q.screen.gt.md"
              @click="table.toggleFullscreen"
              :icon="table.inFullscreen ? 'fullscreen_exit' : 'fullscreen'"
              ><q-tooltip>切换全屏</q-tooltip></q-btn
            >
          </div>
        </div>
      </template>
      <template v-slot:body-cell-enabled="props">
        <q-td :props="props">
          <q-chip
            dense
            text-color="white"
            :color="props.value ? 'positive' : 'grey'"
            >{{ showDict(props.value) }}</q-chip
          >
        </q-td>
      </template>
      <template v-slot:body-cell-actions="props">
        <q-td :props="props">
          <div class="q-gutter-xs">
            <com-edit label="字典详值" @confirm="edit(props.row)" />
            <com-del label="字典详值" @confirm="del(props.row)" />
          </div>
        </q-td>
      </template>
    </q-table>
    <q-dialog v-model="dialogVisible" persistent>
      <q-card style="width: 800px; min-width: 40vw">
        <q-bar class="bg-primary text-white">
          <span> {{ formType }}字典详值 </span>
          <q-space />
          <q-btn dense flat icon="close" v-close-popup>
            <q-tooltip class="bg-white text-primary">Close</q-tooltip>
          </q-btn>
        </q-bar>
        <q-card-section>
          <q-form @submit="submit">
            <div>
              <q-input
                outlined
                v-model="form.label"
                label="字典标签"
                :rules="[requiredRule]"
              />
              <q-input
                outlined
                v-model="form.value"
                label="字典值"
                :rules="[requiredRule]"
              />
              <q-input
                outlined
                label="排序"
                type="number"
                v-model.number="form.sort"
                :rules="[requiredRule]"
              />
              <q-toggle label="是否生效" color="green" v-model="form.enabled" />
            </div>
            <div class="row justify-center q-pa-md">
              <q-btn
                outline
                color="primary"
                icon="mdi-close-thick"
                label="关闭"
                v-close-popup
              />
              <q-btn
                class="q-mx-sm"
                color="primary"
                icon="mdi-check-bold"
                label="提交"
                type="submit"
              />
            </div>
          </q-form>
        </q-card-section>
      </q-card>
    </q-dialog>
  </div>
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
  listDict,
  createDict,
  updateDict,
  deleteDict,
} from 'src/api/admin/dictData'

import { requiredRule } from 'src/utils/inputRule'
import { DictOptions, getDictLabel } from 'src/utils/dict'
import ComSearch from 'src/components/comSearch/index.vue'
import { useQuasar } from 'quasar'

const $q = useQuasar()
let { proxy } = getCurrentInstance()
const props = defineProps({
  dictId: String,
})
const dictId = ref(props.dictId)
const dialogVisible = ref(false)
const dataList = ref([])
const formType = ref('')
const form = ref({})
const searchKey = ref('')
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
      name: 'label',
      align: 'center',
      label: '字典标签',
      field: 'label',
      sortable: true,
    },
    {
      name: 'value',
      align: 'center',
      label: '字典值',
      field: 'value',
      sortable: true,
    },
    {
      name: 'enabled',
      align: 'left',
      label: '状态',
      field: 'enabled',
    },
    { name: 'actions', align: 'center', label: '操作', field: 'actions' },
  ]
})
const selected = ref([])
const dictOptions = ref({})

const showDict = (prop) => {
  if (!prop) {
    prop = false
  }

  return getDictLabel(dictOptions.value.sys_enabled, prop)
}

onMounted(async () => {
  // tableRef.value.requestServerInteraction();
  dictOptions.value = await DictOptions()
})
const reset = () => {
  pagination.value = {
    sortBy: 'sort',
    descending: false,
    page: 1,
    rowsPerPage: 10,
    rowsNumber: 0,
  }
  form.value = {
    enabled: true,
    sort: 10,
    dictTypeId: dictId,
  }
}
const getSelectedString = () => {
  return selected.value.length === 0
    ? ''
    : `已选择${selected.value.length}
      行`
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
  queryReq.dictTypeId = dictId.value ? dictId.value : '0000'
  let table = await listDict(queryReq)

  pagination.value = val.pagination

  pagination.value.rowsNumber = table.count
  dataList.value = table.list
}

const create = () => {
  reset()
  formType.value = '新建'
  dialogVisible.value = true
}
const edit = (p) => {
  form.value = {
    ...p,
  }
  formType.value = '编辑'
  dialogVisible.value = true
}

const del = async (p) => {
  let res = await deleteDict(p)
  onRequest()
}

const submit = async () => {
  // const res = undefined;
  if (formType.value === '编辑') {
    let res = await updateDict(form.value)
  } else if (formType.value === '新建') {
    let res = await createDict(form.value)
  } else {
    proxy.$error('路径错误')
  }

  dialogVisible.value = false
  reset()
  onRequest()
}
watch(props, () => {
  dictId.value = props.dictId
  reset()
  onRequest()
})
</script>
