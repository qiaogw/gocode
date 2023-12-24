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
                v-permission="'pkg:add'"
                flat
                dense
                glossy
                icon="add"
                no-wrap
                color="primary"
                @click="create"
              >
                <q-tooltip>新建</q-tooltip>
              </q-btn>
              <com-del
                :disable="selected.length < 1"
                v-permission="'pkg:del'"
                flat
                glossy
                label="包"
                @confirm="delList"
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
        <template v-slot:body-cell-sourceId="props">
          <q-td :props="props">
            <q-chip
              dense
              text-color="white"
              :color="props.value ? 'positive' : 'grey'"
              >{{ formatGenSource(props.value) }}
            </q-chip>
          </q-td>
        </template>
        <template v-slot:body-cell-mode="props">
          <q-td :props="props">
            <q-chip
              dense
              text-color="white"
              :color="props.value ? 'positive' : 'grey'"
              >{{ formatgenMode(props.value) }}
            </q-chip>
          </q-td>
        </template>
        <template v-slot:body-cell-actions="props">
          <q-td :props="props">
            <div class="q-gutter-xs">
              <q-btn
                v-permission="'pkg:edit'"
                label="生成代码"
                flat
                round
                dense
                color="positive"
                icon="qr_code"
                @click="autoFix(props.row)"
              >
                <q-tooltip>生成代码</q-tooltip></q-btn
              >
              <q-btn
                v-permission="'pkg:edit'"
                flat
                round
                dense
                color="primary"
                icon="auto_fix_high"
                label="查看记录"
                @click="history(props.row)"
              >
                <q-tooltip>查看记录</q-tooltip></q-btn
              >
              <com-edit
                v-permission="'pkg:edit'"
                label="包"
                @confirm="edit(props.row)"
              />
              <com-del
                v-permission="'pkg:del'"
                label="包"
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
          <span> {{ formType }}包 </span>
          <q-space />
          <q-btn dense flat icon="close" v-close-popup>
            <q-tooltip class="bg-white text-primary">关闭</q-tooltip>
          </q-btn>
        </q-bar>
        <q-card-section>
          <q-form @submit="submit">
            <div class="row q-col-gutter-x-md dialog_form q-pt-md">
              <q-select
                class="col-6 q-pb-md"
                outlined
                dense
                map-options
                emit-value
                option-value="value"
                option-label="label"
                v-model="form.sourceId"
                :options="genSourceOptions"
                label="数据源"
                :rules="[requiredRule]"
              />
              <q-input
                outlined
                dense
                class="col-6 q-pb-md"
                v-model="form.name"
                label="名称"
                :rules="[requiredRule]"
              />
              <q-input
                outlined
                dense
                class="col-6 q-pb-md"
                v-model="form.label"
                label="标题"
              />
              <q-select
                class="col-6 q-pb-md"
                outlined
                dense
                map-options
                emit-value
                option-value="value"
                option-label="label"
                v-model="form.mode"
                :options="genModeOptions"
                label="模式"
              />
              <q-input
                outlined
                dense
                type="number"
                class="col-6 q-pb-md"
                v-model.number="form.sort"
                label="排序"
              />
              <q-input
                outlined
                dense
                class="col-6 q-pb-md"
                v-model="form.remark"
                label="remark"
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
import { computed, getCurrentInstance, onMounted, ref } from 'vue'
import {
  createPkg,
  deletePkg,
  genPkgCode,
  listPkg,
  updatePkg,
} from 'src/api/devtools/pkg'
import { listDatasource, previewCode } from 'src/api/devtools/datasource'
import { requiredRule } from 'src/utils/inputRule'
import { useQuasar } from 'quasar'
import {
  DictOptions,
  getDictLabel,
  getOptionsByList,
  getDict,
} from 'src/utils/dict'
import { useRouter, useRoute } from 'vue-router'
const router = useRouter()

const $q = useQuasar()
let { proxy } = getCurrentInstance()
const dialogVisible = ref(false)
const dataList = ref([])
const formType = ref('')
const dictId = ref(0)
const form = ref({})
const dictOptions = ref({})
const genSourceOptions = ref([])
const tableDetail = ref({})
const genModeOptions = ref([])
const formatgenMode = (prop) => {
  if (!prop) {
    prop = false
  }
  return getDictLabel(dictOptions.value.genMode, prop)
}
const formatGenSource = (prop) => {
  return getDictLabel(genSourceOptions.value, prop)
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
      name: 'sourceId',
      align: 'left',
      label: '数据源',
      field: 'sourceId',
      sortable: true,
      classes: 'ellipsis',
    },
    {
      name: 'name',
      align: 'left',
      label: '名称',
      field: 'name',
      sortable: true,
      classes: 'ellipsis',
    },
    {
      name: 'label',
      align: 'left',
      label: '标题',
      field: 'label',
      sortable: true,
      classes: 'ellipsis',
    },
    {
      name: 'mode',
      align: 'left',
      label: '模式',
      field: 'mode',
      sortable: true,
      classes: 'ellipsis',
    },

    { name: 'actions', align: 'center', label: '操作', field: 'actions' },
  ]
})

onMounted(async () => {
  dictOptions.value = await DictOptions()
  genModeOptions.value = await getDict('genMode')
  const queryReqGenSource = {
    pageIndex: 0,
    pageSize: 9999,
  }
  let resGenSource = await listDatasource(queryReqGenSource)
  genSourceOptions.value = getOptionsByList(resGenSource.list, 'name', 'id')
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

  let table = await listPkg(queryReq)

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
  let res = await deletePkg(p)
  await onRequest()
}

const delList = async () => {
  // let res = await deleteDatasource()
  // onRequest()
  let req = {
    ids: getIds(selected.value),
  }
  let res = await deletePkg(req)
  onRequest()
}
const submit = async () => {
  // const res = undefined;
  if (formType.value === '编辑') {
    let res = await updatePkg(form.value)
  } else if (formType.value === '新建') {
    let res = await createPkg(form.value)
  } else {
    proxy.$error('请求错误')
  }
  dialogVisible.value = false
  reset()
  await onRequest()
}

const autoFix = async (prop) => {
  tableDetail.value.id = prop.id
  let res = await genPkgCode(tableDetail.value)
}

const history = (prop) => {
  if (prop) {
    router.push({
      name: 'gencode_history',
      query: {
        sourceId: prop.sourceId,
      },
    })
  } else {
    router.push({ name: 'gencode_history' })
  }
}
</script>
