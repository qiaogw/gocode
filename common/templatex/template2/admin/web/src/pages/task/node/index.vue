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
                v-permission="'node:add'"
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
                v-permission="'node:del'"
                flat
                glossy
                label="主机节点"
                @confirm="delList"
              />
              <q-btn
                v-permission="'node:export'"
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
                v-permission="'node:export'"
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
                v-permission="'node:import'"
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
              <q-icon name="help_outline" class="text-purple cursor-pointer">
                <q-popup-proxy :offset="[10, 10]">
                  <q-banner class="bg-purple text-white">
                    <template v-slot:avatar>
                      <q-icon name="help" />
                    </template>
                    {{ route.meta.remark }}
                  </q-banner>
                </q-popup-proxy>
              </q-icon>
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

        <template v-slot:body-cell-actions="props">
          <q-td :props="props">
            <div class="q-gutter-xs">
              <q-btn
                flat
                round
                dense
                color="primary"
                icon="insert_link"
                @click="ping(props.row)"
              >
                <q-tooltip>测试节点连接</q-tooltip></q-btn
              >
              <com-edit
                v-permission="'node:edit'"
                label="主机节点"
                @confirm="edit(props.row)"
              />
              <com-del
                v-permission="'node:del'"
                label="主机节点"
                @confirm="del(props.row)"
              />
            </div>
          </q-td>
        </template>
        <template v-slot:body-cell-enable="props">
          <q-td :props="props">
            <q-chip
              dense
              text-color="white"
              :color="props.value ? 'positive' : 'grey'"
              >{{ formatsys_enabled(props.value) }}
            </q-chip>
          </q-td>
        </template>
        <template v-slot:body-cell-linked="props">
          <q-td :props="props">
            <q-chip
              dense
              text-color="white"
              :color="props.value ? 'positive' : 'grey'"
              >{{ format_linked(props.value) }}
            </q-chip>
          </q-td>
        </template>
      </q-table>
    </div>
    <q-dialog v-model="dialogVisible" persistent>
      <q-card style="width: 800px; min-width: 40vw">
        <q-bar class="bg-primary text-white">
          <span> {{ formType }}主机节点 </span>
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
                label="节点名称"
              />
              <q-input
                outlined
                dense
                class="col-6 q-pb-md"
                v-model="form.app"
                label="节点认证名称"
              />
              <q-input
                outlined
                dense
                class="col-6 q-pb-md"
                v-model="form.host"
                label="节点地址"
              />
              <q-input
                outlined
                dense
                type="number"
                class="col-6 q-pb-md"
                v-model.number="form.port"
                label="节点端口"
              />
              <q-input
                outlined
                dense
                class="col-6 q-pb-md"
                v-model="form.token"
                label="认证令牌"
              />

              <q-input
                outlined
                dense
                class="col-6 q-pb-md"
                v-model="form.remark"
                label="备注"
              />
              <q-toggle label="可用状态" color="green" v-model="form.enable" />
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
  listNode,
  createNode,
  updateNode,
  deleteNode,
  pingNode,
} from 'src/api/task/node'

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
const route = useRoute()

const $q = useQuasar()
let { proxy } = getCurrentInstance()
const dialogVisible = ref(false)
const dataList = ref([])
const formType = ref('')
const dictId = ref(0)
const form = ref({})
const dictOptions = ref({})

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
      name: 'name',
      align: 'left',
      label: '节点名称',
      field: 'name',
      sortable: true,
      classes: 'ellipsis',
    },
    {
      name: 'app',
      align: 'left',
      label: '节点认证名称',
      field: 'app',
      sortable: true,
      classes: 'ellipsis',
    },
    {
      name: 'host',
      align: 'left',
      label: '节点地址',
      field: 'host',
      sortable: true,
      classes: 'ellipsis',
    },
    {
      name: 'port',
      align: 'left',
      label: '主机端口',
      field: 'port',
      sortable: true,
      classes: 'ellipsis',
    },
    {
      name: 'token',
      align: 'left',
      label: '认证令牌',
      field: 'token',
      sortable: true,
      classes: 'ellipsis',
    },
    {
      name: 'enable',
      align: 'left',
      label: '是否启用',
      field: 'enable',
      sortable: true,
      classes: 'ellipsis',
    },
    {
      name: 'linked',
      align: 'left',
      label: '连接状态',
      field: 'linked',
      sortable: true,
      classes: 'ellipsis',
    },
    { name: 'actions', align: 'center', label: '操作', field: 'actions' },
  ]
})

const formatsys_enabled = (prop) => {
  if (!prop) {
    prop = false
  }
  return getDictLabel(dictOptions.value.sys_enabled, prop)
}

const format_linked = (prop) => {
  if (!prop) {
    prop = false
  }
  return prop ? '已连接' : '断开'
}
onMounted(async () => {
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

  let table = await listNode(queryReq)

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
  let res = await deleteNode(p)
  onRequest()
}

const delList = async () => {
  // let res = await deleteDatasource()
  // onRequest()
  let req = {
    ids: getIds(selected.value),
  }
  let res = await deleteNode(req)
  onRequest()
}
const submit = async () => {
  // const res = undefined
  if (formType.value === '编辑') {
    let res = await updateNode(form.value)
  } else if (formType.value === '新建') {
    let res = await createNode(form.value)
  } else {
    proxy.$error('请求错误')
  }
  dialogVisible.value = false
  reset()
  onRequest()
}
const uploadUrl = process.env.BASE_URL + '/task/node/import'
const exportUrl = '/task/node/export'
const exportTemplateUrl = '/task/node/exportTemplate'
const handleExport = () => {
  let queryReq = {}
  let val = {}
  val.pagination = pagination.value
  queryReq.pageSize = val.pagination.rowsPerPage
  queryReq.pageIndex = val.pagination.page
  queryReq.sortBy = val.pagination.sortBy
  queryReq.descending = val.pagination.descending
  queryReq.searchKey = searchKey.value
  downloadAction(exportUrl, '主机节点-导出.xlsx', queryReq)
}

const handleExportTemplate = () => {
  downloadAction(exportTemplateUrl, '主机节点模板.xlsx')
}

const uploadFn = async (val) => {
  reset()
  onRequest()
}
const ping = async (val) => {
  let res = await pingNode(val)
}
</script>
