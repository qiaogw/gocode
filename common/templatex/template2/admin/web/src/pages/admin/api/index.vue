<template>
  <q-page class="q-pa-xs">
    <div class="shadow-2 q-pa-xs">
      <q-table
        dense
        flat
        bordered
        separator="cell"
        :columns="columns"
        :rows="dataList"
        row-key="id"
        v-model:filter="searchKey"
        @request="onRequest"
        :grid="$q.screen.xs"
        v-model:pagination="pagination"
        binary-state-sort
        :selected-rows-label="getSelectedString"
        selection="multiple"
        v-model:selected="selected"
      >
        <template v-slot:top="table">
          <div class="row no-wrap full-width">
            <com-search @query="onRequest" v-model:filter="searchKey" />
            <q-space />
            <q-btn-group class="q-gutter-xs">
              <q-btn
                v-permission="'api:add'"
                flat
                dense
                glossy
                icon="add"
                no-wrap
                color="primary"
                @click="create"
                ><q-tooltip>新建</q-tooltip></q-btn
              >
              <q-btn
                v-permission="'api:export'"
                flat
                dense
                glossy
                icon="download"
                no-wrap
                color="primary"
                @click="handleExport"
                ><q-tooltip>导出</q-tooltip></q-btn
              >
              <q-btn
                v-permission="'api:export'"
                flat
                dense
                glossy
                icon="downloading"
                no-wrap
                color="primary"
                @click="handleExportTemplate"
                ><q-tooltip>导出模板</q-tooltip></q-btn
              >
              <com-upload
                v-permission="'api:import'"
                flat
                dense
                glossy
                @upload="uploadFn"
                title="导入"
                :uploadUrl="uploadUrl"
                fileType=".xlsx,.xls"
              />
              <q-btn
                color="primary"
                flat
                dense
                glossy
                no-wrap
                v-show="$q.screen.gt.md"
                @click="table.toggleFullscreen"
                :icon="table.inFullscreen ? 'fullscreen_exit' : 'fullscreen'"
              >
                <q-tooltip>切换全屏</q-tooltip>
              </q-btn>
            </q-btn-group>
          </div>
        </template>
        <template v-slot:body-cell-remark="props">
          <q-td key="remark" :props="props">
            <div>
              <q-tooltip anchor="bottom middle" self="top middle">
                {{ props.value }}
              </q-tooltip>
              <div class="ellipsis">
                {{ props.value }}
              </div>
            </div>
          </q-td>
        </template>
        <template v-slot:body-cell-method="props">
          <q-td :props="props">
            <q-chip
              dense
              text-color="white"
              :color="props.value ? 'positive' : 'grey'"
              >{{ formathttp_methods(props.value) }}
            </q-chip>
          </q-td>
        </template>

        <template v-slot:body-cell-actions="props">
          <q-td :props="props">
            <div class="q-gutter-xs">
              <com-edit
                v-permission="'api:edit'"
                label="api"
                @confirm="edit(props.row)"
              />
              <com-del
                v-permission="'api:del'"
                label="api"
                @confirm="del(props.row)"
              />
            </div>
          </q-td>
        </template>
      </q-table>
    </div>
    <q-dialog v-model="dialogVisible" persistent>
      <q-card style="width: 800px; min-width: 40vw">
        <q-bar class="bg-primary text-white">
          <span> {{ formType }}api </span>
          <q-space />
          <q-btn dense flat icon="close" v-close-popup>
            <q-tooltip class="bg-white text-primary">关闭</q-tooltip>
          </q-btn>
        </q-bar>
        <q-card-section>
          <q-form @submit="submit">
            <div class="row q-col-gutter-x-md dialog_form q-pt-md">
              <q-input
                outlined
                dense
                class="col-6 q-pb-md"
                v-model="form.title"
                label="标题"
                :rules="[requiredRule]"
              />
              <q-input
                outlined
                dense
                class="col-6 q-pb-md"
                v-model="form.path"
                label="地址"
                :rules="[requiredRule]"
              />
              <q-select
                class="col-6 q-pb-md"
                outlined
                dense
                map-options
                emit-value
                option-value="value"
                option-label="label"
                v-model="form.method"
                :options="http_methodsOptions"
                label="请求类型"
              />
              <q-input
                outlined
                dense
                class="col-6 q-pb-md"
                v-model="form.module"
                label="api组"
                :rules="[requiredRule]"
              />
              <q-input
                outlined
                dense
                class="col-6 q-pb-md"
                v-model="form.remark"
                label="说明"
              />
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
import { listApi, createApi, updateApi, deleteApi } from 'src/api/admin/api'
import { useQuasar } from 'quasar'
import { requiredRule } from 'src/utils/inputRule'
import { DictOptions, getOptionsByList, getDictLabel } from 'src/utils/dict'
import { downloadAction } from 'src/api/manage'
import { getIds } from 'src/utils/arrayOrObject'

const $q = useQuasar()
let { proxy } = getCurrentInstance()
const dialogVisible = ref(false)
const dataList = ref([])
const formType = ref('')
const dictId = ref(0)
const form = ref({})
const dictOptions = ref({})
const http_methodsOptions = ref([])
const formathttp_methods = (prop) => {
  if (!prop) {
    prop = false
  }
  return getDictLabel(dictOptions.value.http_methods, prop)
}

const searchKey = ref('')
const uploadUrl = process.env.BASE_URL + '/admin/api/import'
const exportUrl = process.env.BASE_URL + '/admin/api/export'
const exportTemplateUrl = process.env.BASE_URL + '/admin/api/exportTemplate'
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
      name: 'title',
      align: 'left',
      label: '标题',
      field: 'title',
      sortable: true,
      classes: 'ellipsis',
    },
    {
      name: 'path',
      align: 'left',
      label: '地址',
      field: 'path',
      sortable: true,
      classes: 'ellipsis',
    },
    {
      name: 'method',
      align: 'left',
      label: '请求类型',
      field: 'method',
      sortable: true,
      classes: 'ellipsis',
    },
    {
      name: 'module',
      align: 'left',
      label: 'api组',
      field: 'module',
      sortable: true,
      classes: 'ellipsis',
    },
    { name: 'actions', align: 'center', label: '操作', field: 'actions' },
  ]
})

onMounted(async () => {
  dictOptions.value = await DictOptions()
  http_methodsOptions.value = dictOptions.value.http_methods
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

  let table = await listApi(queryReq)

  pagination.value = val.pagination

  pagination.value.rowsNumber = table.count
  dataList.value = table.list
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
  let res = await deleteApi(p)
  onRequest()
}
const delList = async () => {
  // let res = await deleteDatasource()
  // onRequest()
  let req = {
    ids: getIds(selected.value),
  }
  let res = await deleteApi(req)
  onRequest()
}

const submit = async () => {
  // const res = undefined;
  if (formType.value === '编辑') {
    let res = await updateApi(form.value)
  } else if (formType.value === '新建') {
    let res = await createApi(form.value)
  } else {
    proxy.$error('请求错误')
  }
  dialogVisible.value = false
  reset()
  onRequest()
}

const handleExport = () => {
  let queryReq = {}
  let val = {}
  val.pagination = pagination.value
  queryReq.pageSize = val.pagination.rowsPerPage
  queryReq.pageIndex = val.pagination.page
  queryReq.sortBy = val.pagination.sortBy
  queryReq.descending = val.pagination.descending
  queryReq.searchKey = searchKey.value
  downloadAction(exportUrl, 'api-导出.xlsx', queryReq)
}

const handleExportTemplate = () => {
  downloadAction(exportTemplateUrl, 'api模板.xlsx')
}

const uploadFn = async (val) => {
  reset()
  onRequest()
}
</script>
