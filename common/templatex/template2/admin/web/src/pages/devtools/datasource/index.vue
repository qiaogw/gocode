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
            <q-btn-group flat>
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
                glossy
                dense
                @upload="uploadFn"
                title="导入"
                :uploadUrl="uploadUrl"
                fileType=".xlsx,.xls"
              />
              <q-btn
                color="primary"
                flat
                glossy
                dense
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
            <q-tooltip class="bg-white text-primary">Close</q-tooltip>
          </q-btn>
        </q-bar>
        <q-card-section>
          <q-form @submit="submit">
            <div class="row q-col-gutter-x-md dialog_form q-pt-md">
              <q-input
                dense
                outlined
                class="col-6"
                v-model="form.name"
                label="数据源名称"
                :rules="[requiredRule]"
              />
              <q-select
                class="col-6 q-pb-md"
                outlined
                dense
                filled
                use-chips
                emit-value
                map-options
                option-value="value"
                option-label="label"
                v-model="form.driver"
                :options="driverOptins"
                label="数据库类型"
                @update:model-value="listDriver"
              />
              <q-input
                class="col-6 q-pb-md"
                outlined
                dense
                v-model="form.host"
                label="主机地址"
                :rules="[requiredRule]"
              />
              <q-input
                class="col-6 q-pb-md"
                outlined
                dense
                v-model.number="form.port"
                type="number"
                label="端口"
                :rules="[requiredRule]"
              />
              <q-input
                class="col-6 q-pb-md"
                outlined
                dense
                v-model="form.user"
                label="用户名"
                :rules="[requiredRule]"
              />
              <q-input
                dense
                outlined
                class="col-6"
                label="密码"
                v-model="form.password"
                :rules="[requiredRule]"
              />
              <q-input
                class="col-6 q-pb-md"
                outlined
                dense
                v-model="form.dbname"
                label="数据库"
                :rules="[requiredRule]"
              />
              <q-input
                dense
                outlined
                class="col-6"
                label="前缀"
                v-model="form.tablePrefix"
              />
            </div>
            <q-input
              dense
              outlined
              class="col-6"
              label="配置(时区、编码、ssl等)"
              v-model="form.config"
              :rules="[requiredRule]"
            />
            <q-input
              dense
              outlined
              class="col-6"
              label="描述"
              v-model="form.remark"
            />
            <div class="row justify-between q-pa-md">
              <q-btn
                outline
                color="positive"
                icon="link_off"
                label="测试连接"
                @click="testLink"
              />
              <div>
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
  testDatasource,
  excelExport,
  excelImport,
  getTables,
} from 'src/api/devtools/datasource'
import { useQuasar } from 'quasar'
import { requiredRule } from 'src/utils/inputRule'
import { DictOptions, getDict } from 'src/utils/dict'
import { downloadAction } from 'src/api/manage'
import { getIds } from 'src/utils/arrayOrObject'
// import ComUpload from "components/comUpload";

const $q = useQuasar()
let { proxy } = getCurrentInstance()
const dialogVisible = ref(false)
const dataList = ref([])
const formType = ref('')
const dictId = ref(0)
const form = ref({})
const searchKey = ref('')
const uploadUrl = process.env.BASE_URL + '/gencode/gensource/import'
const exportUrl = '/gencode/gensource/export'
const exportTemplateUrl = '/gencode/gensource/exportTemplate'
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
      align: 'center',
      label: '名称',
      field: 'name',
      sortable: true,
    },
    {
      name: 'driver',
      align: 'left',
      label: '类型',
      field: 'driver',
      sortable: true,
    },
    {
      name: 'host',
      align: 'left',
      label: '地址',
      field: 'host',
    },
    {
      name: 'port',
      align: 'left',
      label: '端口',
      field: 'port',
    },
    {
      name: 'dbname',
      align: 'left',
      label: '数据库',
      field: 'dbname',
    },
    { name: 'actions', align: 'center', label: '操作', field: 'actions' },
  ]
})
const dictOptions = ref({})
const driverOptins = ref([])

onMounted(async () => {
  // dictOptions.value = await DictOptions()
  // driverOptins.value = dictOptions.value.data_driver
  driverOptins.value = await getDict('data_driver')
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
  formType.value = '新建'
  reset()
  dialogVisible.value = true
}

const listDriver = async (val) => {
  if (val === 'mysql') {
    form.value.config = 'charset=utf8&parseTime=True&loc=Local&timeout=1000ms'
  }
  if (val === 'postgres') {
    form.value.config = 'sslmode=disable TimeZone=Asia/Shanghai'
  }
}

const edit = (p) => {
  form.value = {
    ...p,
  }
  formType.value = '编辑'
  dialogVisible.value = true
}

const del = async (p) => {
  let req = {
    ids: [p],
  }
  let res = await deleteDatasource(req)
  onRequest()
}
const delList = async () => {
  // let res = await deleteDatasource()
  // onRequest()
  let req = {
    ids: getIds(selected.value),
  }
  let res = await deleteDatasource(req)
  onRequest()
}

const submit = async () => {
  // const res = undefined;
  if (formType.value === '编辑') {
    let res = await updateDatasource(form.value)
  } else if (formType.value === '新建') {
    let res = await createDatasource(form.value)
  } else {
    proxy.$error('路径错误')
  }
  dialogVisible.value = false
  reset()
  onRequest()
}
const testLink = async () => {
  let res = await testDatasource(form.value)
  // onRequest()
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
  downloadAction(exportUrl, '数据源.xlsx', queryReq)
}

const handleExportTemplate = () => {
  downloadAction(exportTemplateUrl, '数据源模板.xlsx')
}

const uploadFn = async (val) => {
  reset()
  onRequest()
}
</script>
