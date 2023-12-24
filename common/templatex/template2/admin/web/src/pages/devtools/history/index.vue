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
                v-permission="'gencode_history:export'"
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

        <template v-slot:body-cell-actions="props">
          <q-td :props="props">
            <div class="q-gutter-xs">
              <q-btn
                flat
                round
                dense
                color="primary"
                icon="auto_fix_high"
                @click="autoFix(props.row)"
              >
                <q-tooltip>代码生成</q-tooltip></q-btn
              >
              <q-btn
                flat
                round
                dense
                color="primary"
                icon="format_shapes"
                @click="formbuild(props.row)"
              >
                <q-tooltip>表单生成</q-tooltip></q-btn
              >
              <com-del
                v-permission="'gencode_history:del'"
                label="表模块"
                @confirm="del(props.row)"
              />
            </div>
          </q-td>
        </template>
      </q-table>
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
  listGenTable,
  createGenTable,
  updateGenTable,
  deleteGenTable,
} from 'src/api/devtools/table'
import { useQuasar } from 'quasar'
import { requiredRule } from 'src/utils/inputRule'
import { DictOptions, getOptionsByList, getDictLabel } from 'src/utils/dict'
import { downloadAction } from 'src/api/manage'
import { useRouter, useRoute } from 'vue-router'

const router = useRouter()
const $q = useQuasar()
let { proxy } = getCurrentInstance()
const route = useRoute()
const dialogVisible = ref(false)
const dataList = ref([])
const formType = ref('')
const dictId = ref(0)
const form = ref({})
const dictOptions = ref({})

const queryReq = ref({})
const searchKey = ref('')
const uploadUrl = process.env.BASE_URL + '/gencode/gentable/import'
const exportUrl = process.env.BASE_URL + '/gencode/gentable/export'
const exportTemplateUrl =
  process.env.BASE_URL + '/gencode/gentable/exportTemplate'
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
      name: 'tableComment',
      align: 'left',
      label: '中文名',
      field: 'tableComment',
      sortable: true,
      classes: 'ellipsis',
    },
    {
      name: 'packageName',
      align: 'left',
      label: '包名',
      field: 'packageName',
      sortable: true,
      classes: 'ellipsis',
    },
    {
      name: 'tableUrl',
      align: 'left',
      label: '文件名',
      field: 'tableUrl',
      sortable: true,
      classes: 'ellipsis',
    },
    {
      name: 'pkg',
      align: 'left',
      label: '根目录',
      field: 'pkg',
      sortable: true,
      classes: 'ellipsis',
    },
    {
      name: 'parentPkg',
      align: 'left',
      label: '项目路径',
      field: 'parentPkg',
      sortable: true,
      classes: 'ellipsis',
    },

    { name: 'actions', align: 'center', label: '操作', field: 'actions' },
  ]
})

onMounted(async () => {
  if (route.query.sourceId) {
    queryReq.value.sourceId = Number(route.query.sourceId)
  }
  dictOptions.value = await DictOptions()
  await onRequest()
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
  queryReq.value = {}
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

  queryReq.value.pageSize = rowsPerPage
  queryReq.value.pageIndex = page
  queryReq.value.sortBy = sortBy
  queryReq.value.descending = descending
  queryReq.value.searchKey = val.filter

  let table = await listGenTable(queryReq.value)

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

const del = async (p) => {
  let res = await deleteGenTable(p)
  onRequest()
}

const submit = async () => {
  // const res = undefined;
  if (formType.value === '编辑') {
    let res = await updateGenTable(form.value)
  } else if (formType.value === '新建') {
    let res = await createGenTable(form.value)
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
  downloadAction(exportUrl, '表模块-导出.xlsx', queryReq)
}

const handleExportTemplate = () => {
  downloadAction(exportTemplateUrl, '表模块模板.xlsx')
}

const uploadFn = async (val) => {
  reset()
  onRequest()
}
const autoFix = async (prop) => {
  if (prop) {
    router.push({
      name: 'autoCodeEdit',
      query: {
        id: prop.id,
        tablePrefix: prop.source.tablePrefix,
      },
    })
  } else {
    router.push({ name: 'autocode' })
  }
}
const formbuild = async (prop) => {
  if (prop) {
    router.push({
      name: 'formbuild',
      query: {
        id: prop.id,
        tablePrefix: prop.source.tablePrefix,
      },
    })
  }
}
</script>
