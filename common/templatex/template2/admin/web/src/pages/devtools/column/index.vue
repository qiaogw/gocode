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
                v-permission="'datasource:add'"
                flat
                dense
                glossy
                icon="add"
                no-wrap
                color="primary"
                @click="create"
                ><q-tooltip>新建</q-tooltip></q-btn
              >
              <com-del
                :disable="selected.length < 1"
                v-permission="'datasource:del'"
                flat
                glossy
                label="数据源"
                @confirm="delList"
              />
              <q-btn
                v-permission="'datasource:export'"
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
                v-permission="'datasource:export'"
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
                v-permission="'datasource:import'"
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
        <template v-slot:body-cell-driver="props">
          <q-td :props="props">
            <q-chip
              dense
              text-color="white"
              :color="props.value ? 'positive' : 'grey'"
              >{{ formatdata_driver(props.value) }}
            </q-chip>
          </q-td>
        </template>

        <template v-slot:body-cell-actions="props">
          <q-td :props="props">
            <div class="q-gutter-xs">
              <com-edit
                v-permission="'datasource:edit'"
                label="数据源"
                @confirm="edit(props.row)"
              />
              <com-del
                v-permission="'datasource:del'"
                label="数据源"
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
          <span> 数据源 </span>
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
                v-model="form.name"
                label="名称"
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
                v-model="form.driver"
                :options="data_driverOptions"
                label="数据库类型"
              />
              <q-input
                outlined
                dense
                class="col-6 q-pb-md"
                v-model="form.host"
                label="主机地址"
                :rules="[requiredRule]"
              />
              <q-input
                outlined
                dense
                class="col-6 q-pb-md"
                v-model="form.port"
                label="端口"
                :rules="[requiredRule]"
              />
              <q-input
                outlined
                dense
                class="col-6 q-pb-md"
                v-model="form.user"
                label="用户"
                :rules="[requiredRule]"
              />
              <q-input
                outlined
                dense
                class="col-6 q-pb-md"
                v-model="form.password"
                label="密码"
                :rules="[requiredRule]"
              />
              <q-input
                outlined
                dense
                class="col-6 q-pb-md"
                v-model="form.dbname"
                label="数据库"
                :rules="[requiredRule]"
              />
              <q-input
                outlined
                dense
                class="col-6 q-pb-md"
                v-model="form.config"
                label="配置"
                :rules="[requiredRule]"
              />
              <q-input
                outlined
                dense
                class="col-6 q-pb-md"
                v-model="form.tablePrefix"
                label="表前缀"
                :rules="[requiredRule]"
              />
              <q-input
                outlined
                dense
                class="col-6 q-pb-md"
                v-model="form.remark"
                label="remark"
                :rules="[requiredRule]"
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
import {
  listDatasource,
  createDatasource,
  updateDatasource,
  deleteDatasource,
} from 'src/api/devtools/datasource'
import { useQuasar } from 'quasar'
import { requiredRule } from 'src/utils/inputRule'
import { DictOptions, getOptionsByList, getDictLabel } from 'src/utils/dict'
import { downloadAction } from 'src/api/manage'

const $q = useQuasar()
let { proxy } = getCurrentInstance()
const dialogVisible = ref(false)
const dataList = ref([])
const formType = ref('')
const dictId = ref(0)
const form = ref({})
const dictOptions = ref({})
const data_driverOptions = ref([])
const formatdata_driver = (prop) => {
  if (!prop) {
    prop = false
  }
  return getDictLabel(dictOptions.value.data_driver, prop)
}

const searchKey = ref('')
const uploadUrl = process.env.BASE_URL + '/gencode/gensource/import'
const exportUrl = process.env.BASE_URL + '/gencode/gensource/export'
const exportTemplateUrl =
  process.env.BASE_URL + '/gencode/gensource/exportTemplate'
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
      name: 'name',
      align: 'left',
      label: '名称',
      field: 'name',
      sortable: true,
      classes: 'ellipsis',
    },
    {
      name: 'driver',
      align: 'left',
      label: '数据库类型',
      field: 'driver',
      sortable: true,
      classes: 'ellipsis',
    },
    {
      name: 'host',
      align: 'left',
      label: '主机地址',
      field: 'host',
      sortable: true,
      classes: 'ellipsis',
    },
    {
      name: 'port',
      align: 'left',
      label: '端口',
      field: 'port',
      sortable: true,
      classes: 'ellipsis',
    },
    {
      name: 'dbname',
      align: 'left',
      label: '数据库',
      field: 'dbname',
      sortable: true,
      classes: 'ellipsis',
    },
    { name: 'actions', align: 'center', label: '操作', field: 'actions' },
  ]
})

onMounted(async () => {
  dictOptions.value = await DictOptions()
  data_driverOptions.value = dictOptions.value.data_driver
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

  let table = await listDatasource(queryReq)

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
  let res = await deleteDatasource(p)
  onRequest()
}

const submit = async () => {
  // const res = undefined;
  if (formType.value === '编辑') {
    let res = await updateDatasource(form.value)
  } else if (formType.value === '新建') {
    let res = await createDatasource(form.value)
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
  downloadAction(exportUrl, '数据源-导出.xlsx', queryReq)
}

const handleExportTemplate = () => {
  downloadAction(exportTemplateUrl, '数据源模板.xlsx')
}

const uploadFn = async (val) => {
  reset()
  onRequest()
}
</script>
