<template>
  <q-page class="q-pa-sm row">
    <div class="col-md-6 col-xs-12 shadow-2 q-pa-xs">
      <q-table
        dense
        flat
        bordered
        separator="cell"
        :columns="columns"
        :rows="dataList"
        row-key="id"
        @request="onRequest"
        @row-click="clickDict"
        @selection="clickDict"
        :grid="$q.screen.xs"
        v-model:filter="searchKey"
        v-model:selected="selected"
        v-model:pagination="pagination"
        binary-state-sort
      >
        <template v-slot:top="table">
          <div class="row no-wrap full-width">
            <com-search @query="onRequest" v-model:filter="searchKey" />
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
        <template v-slot:body-cell-remark="props">
          <q-td key="remark" :props="props">
            <div>
              <q-tooltip anchor="bottom middle" self="top middle">{{
                props.value
              }}</q-tooltip>
              <div class="ellipsis">
                {{ props.value }}
              </div>
            </div>
          </q-td>
        </template>
        <template v-slot:body-cell-actions="props">
          <q-td :props="props">
            <div class="q-gutter-xs">
              <com-edit label="字典" @confirm="edit(props.row)" />
              <com-del label="字典" @confirm="del(props.row)" />
            </div>
          </q-td>
        </template>
      </q-table>
    </div>
    <div class="col-md-6 col-xs-12 shadow-2 q-pa-xs">
      <dictItem :dictId="dictId" />
    </div>

    <q-dialog v-model="dialogVisible" persistent>
      <q-card style="width: 800px; min-width: 40vw">
        <q-bar class="bg-primary text-white">
          <span> {{ formType }}字典列表 </span>
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
                v-model="form.name"
                label="字典名称"
                :rules="[requiredRule]"
              />
              <q-input
                outlined
                v-model="form.type"
                label="字典编码"
                :rules="[requiredRule]"
              />
              <q-input
                outlined
                label="描述"
                v-model="form.remark"
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
  listDict,
  createDict,
  updateDict,
  deleteDict,
} from 'src/api/admin/dictType'

import { requiredRule } from 'src/utils/inputRule'
import { DictOptions } from 'src/utils/dict'
import ComSearch from 'src/components/comSearch/index.vue'
import dictItem from './item.vue'
import { useQuasar } from 'quasar'

const $q = useQuasar()
let { proxy } = getCurrentInstance()
const dialogVisible = ref(false)
const dataList = ref([])
const formType = ref('')
const dictId = ref(0)
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
      name: 'name',
      align: 'center',
      label: '名称',
      field: 'name',
      sortable: true,
    },
    {
      name: 'type',
      align: 'left',
      label: '字典代码',
      field: 'type',
      sortable: true,
    },
    {
      name: 'remark',
      align: 'left',
      label: '描述',
      field: 'remark',
      classes: 'ellipsis',
      style: 'max-width: 100px',
    },
    { name: 'actions', align: 'center', label: '操作', field: 'actions' },
  ]
})
const selected = ref([])
const dictOptions = ref({})

onMounted(async () => {
  // tableRef.value.requestServerInteraction();
  dictOptions.value = await DictOptions()
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
    enabled: true,
  }
  dictId.value = 0
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

  let table = await listDict(queryReq)

  pagination.value = val.pagination

  pagination.value.rowsNumber = table.count
  dataList.value = table.list
}

const create = () => {
  formType.value = '新建'
  reset()
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
const clickDict = (evt, row, index) => {
  if (dictId.value === row.id) {
    dictId.value = 0
    selected.value.pop()
  } else {
    dictId.value = row.id
    selected.value.pop()
    selected.value.push(row)
  }
}
</script>
<style lang="scss" scoped>
.selected {
  background-color: $secondary;
}
</style>
