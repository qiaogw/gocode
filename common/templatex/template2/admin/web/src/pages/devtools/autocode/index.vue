<template>
  <q-page class="q-pa-xs">
    <div class="shadow-2 q-pa-xs">
      <q-card>
        <div v-if="dbVisable">
          <q-bar>选择数据库表</q-bar>
          <q-card-section class="row q-col-gutter-x-md">
            <q-select
              class="col-6 q-pb-md"
              outlined
              dense
              map-options
              option-value="id"
              option-label="name"
              v-model="db"
              :options="sourceList"
              label="数据源"
              @update:model-value="listTables"
            />
            <q-select
              class="col-6 q-pb-md"
              outlined
              dense
              :disable="tableList.length < 1"
              map-options
              option-value="name"
              option-label="tableComment"
              v-model="table"
              :options="tableList"
              label="表"
              @update:model-value="detaiTable"
            />
          </q-card-section>
        </div>
        <q-bar>表信息</q-bar>
        <q-card-section>
          <q-form class="row q-col-gutter-x-md">
            <q-input
              class="col-6 q-pb-md"
              outlined
              dense
              v-model="tableDetail.db"
              label="所属服务名称"
              :rules="[requiredRule]"
            />
            <q-input
              class="col-6 q-pb-md"
              outlined
              dense
              v-model="tableDetail.tableComment"
              label="业务菜单中文名"
              :rules="[requiredRule]"
            />
            <q-input
              class="col-6 q-pb-md"
              outlined
              dense
              v-model="tableDetail.packageName"
              label="包名"
              :rules="[requiredRule]"
            />
            <q-input
              class="col-6 q-pb-md"
              outlined
              dense
              disable
              v-model="tableDetail.name"
              label="表名"
              :rules="[requiredRule]"
            />
            <q-input
              class="col-6 q-pb-md"
              outlined
              dense
              v-model="tableDetail.table"
              label="结构体模型名"
              :rules="[requiredRule]"
            />
            <q-input
              class="col-6 q-pb-md"
              outlined
              dense
              v-model="tableDetail.tableUrl"
              label="文件名"
              :rules="[requiredRule]"
            />
            <q-input
              class="col-6 q-pb-md"
              outlined
              dense
              disable
              v-model="tableDetail.dir"
              label="文件目录"
            />
            <q-input
              class="col-6 q-pb-md"
              outlined
              dense
              v-model="tableDetail.author"
              label="作者"
            />
            <q-toggle
              label="需要用户认证"
              color="green"
              v-model="tableDetail.isAuth"
            />
            <q-toggle
              label="是否导入导出"
              color="green"
              v-model="tableDetail.isImport"
            />
          </q-form>
        </q-card-section>
      </q-card>
      <q-table
        dense
        flat
        bordered
        separator="cell"
        :columns="columns"
        v-model:rows="tableDetail.columns"
        row-key="id"
        :pagination="pagination"
        binary-state-sort
      >
        <template v-slot:top="table">
          <div class="row no-wrap full-width">
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
        <template v-slot:body-cell-require="props">
          <q-td :props="props">
            <div class="q-pa-md">
              <q-checkbox v-model="props.row.require" />
            </div>
          </q-td>
        </template>
        <template v-slot:body-cell-isList="props">
          <q-td :props="props">
            <div class="q-pa-md">
              <q-checkbox v-model="props.row.isList" />
            </div>
          </q-td>
        </template>
        <template v-slot:body-cell-isEdit="props">
          <q-td :props="props">
            <div class="q-pa-md">
              <q-checkbox v-model="props.row.isEdit" />
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
                icon="arrow_upward"
                :disabled="props.rowIndex === 0"
                @click="moveUpField(props.rowIndex)"
              >
                <q-tooltip>上移</q-tooltip></q-btn
              >
              <q-btn
                flat
                round
                dense
                color="primary"
                icon="arrow_downward"
                :disabled="props.rowIndex + 1 === tableDetail.columns.length"
                @click="moveDownField(props.rowIndex)"
              >
                <q-tooltip>下移</q-tooltip></q-btn
              >
              <com-edit
                label="字段"
                @confirm="edit(props.row, props.rowIndex)"
              />
              <com-del label="字段" @confirm="del(props.rowIndex)" />
            </div>
          </q-td>
        </template>
        <template v-slot:bottom>
          <div class="row no-wrap full-width">
            <q-space />
            <q-btn label="预览并生成代码" rounded color="primary" @click.stop>
              <q-tooltip>预览并生成代码</q-tooltip>
              <q-menu auto-close anchor="center left" self="center right">
                <div class="row no-wrap items-center q-pa-sm bg-cyan-2">
                  <span class="text-no-wrap"
                    >确认生成代码 <b>{{ tableDetail.tableComment }}</b> ？</span
                  >
                  <q-btn
                    flat
                    round
                    dense
                    color="primary"
                    icon="done"
                    @click="preview"
                  >
                    <q-tooltip>确认生成代码</q-tooltip></q-btn
                  >
                </div>
              </q-menu>
            </q-btn>
            <q-btn
              label="生成单服务代码"
              rounded
              color="deep-orange"
              @click="getCode"
            />
            <q-btn label="覆盖本机代码" rounded color="negative" @click.stop>
              <q-tooltip>覆盖本机代码</q-tooltip>
              <q-menu auto-close anchor="center left" self="center right">
                <div class="row no-wrap items-center q-pa-sm bg-cyan-2">
                  <span class="text-no-wrap"
                    >确认覆盖本机代码 <b>{{ label }}</b> ？</span
                  >
                  <q-btn
                    flat
                    round
                    dense
                    color="deep-orange"
                    icon="done"
                    @click="coverCode"
                  >
                    <q-tooltip>确认覆盖本机代码</q-tooltip></q-btn
                  >
                </div>
              </q-menu>
            </q-btn>
            <q-btn
              label="生成菜单"
              rounded
              color="deep-orange"
              @click="genMenu"
            />
          </div>
        </template>
      </q-table>
    </div>
    <q-dialog v-model="dialogVisible" persistent>
      <q-card style="width: 800px; min-width: 80vw">
        <q-bar class="bg-primary text-white">
          <span> {{ formType }}字段 </span>
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
                :disable="formType === '新建' ? false : true"
                v-model="form.gormName"
                label="字段名称"
                :rules="[requiredRule]"
              />
              <q-input
                dense
                outlined
                disable
                class="col-6"
                v-model="form.fieldName"
                label="go字段名称"
                :rules="[requiredRule]"
              />
              <q-input
                dense
                outlined
                disable
                class="col-6"
                v-model="form.fieldJson"
                label="json字段名称"
                :rules="[requiredRule]"
              />
              <q-input
                dense
                outlined
                disable
                class="col-6"
                v-model="form.tablename"
                label="表名"
                :rules="[requiredRule]"
              />
              <q-input
                class="col-6 q-pb-md"
                outlined
                dense
                v-model="form.columnComment"
                label="中文名称"
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
                v-model="form.dataType"
                :options="fieldOptions"
                label="go数据类型"
                :rules="[requiredRule]"
              />
              <q-select
                class="col-6 q-pb-md"
                outlined
                dense
                clearable
                map-options
                emit-value
                option-value="value"
                option-label="label"
                v-model="form.htmlType"
                :options="htmlOptions"
                label="HTML输入数据类型"
              />
              <q-select
                class="col-6 q-pb-md"
                outlined
                dense
                emit-value
                map-options
                clearable
                option-value="value"
                option-label="label"
                v-model="form.dictType"
                :options="dictList"
                label="关联字典"
              />
              <q-select
                class="col-6 q-pb-md"
                outlined
                dense
                map-options
                :disable="form.dictType.length > 1 ? true : false"
                option-value="name"
                option-label="tableComment"
                v-model="form.fkTable"
                :options="tableList"
                label="关系表"
                @update:model-value="detaiFkTable"
              />
              <q-select
                class="col-6 q-pb-md"
                outlined
                dense
                map-options
                emit-value
                :disable="fieldList.length < 1"
                option-value="name"
                option-label="columnComment"
                v-model="form.fkLabelId"
                :options="fieldList"
                label="关系表显示"
              />
              <q-select
                class="col-6 q-pb-md"
                outlined
                dense
                map-options
                emit-value
                :disable="fieldList.length < 1"
                option-value="name"
                option-label="columnComment"
                v-model="form.fkLabelName"
                :options="fieldList"
                label="关系表值"
              />
            </div>
            <div class="row justify-center q-pa-md">
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
    <q-dialog v-model="codeVisible" persistent>
      <q-card style="max-width: 100vw; min-width: 80vw">
        <q-bar class="bg-primary text-white">
          <span> {{ tableDetail.tableComment }} </span>
          <q-space />
          <q-btn dense flat icon="close" v-close-popup>
            <q-tooltip class="bg-white text-primary">Close</q-tooltip>
          </q-btn>
        </q-bar>
        <q-card-section>
          <q-splitter v-model="splitterModel">
            <template v-slot:before>
              <q-tabs
                v-model="tag"
                dense
                vertical
                class="text-teal"
                active-color="primary"
                indicator-color="primary"
              >
                <q-tab
                  v-for="item in code"
                  :key="item.name"
                  :name="item.name"
                  :label="item.name"
                />
              </q-tabs>
            </template>
            <template v-slot:after>
              <q-tab-panels
                animated
                swipeable
                vertical
                v-model="tag"
                transition-prev="jump-up"
                transition-next="jump-up"
              >
                <q-tab-panel
                  v-for="item in code"
                  :key="item.name"
                  :name="item.name"
                >
                  <div>
                    <com-code :code="item.code" />
                  </div>
                </q-tab-panel>
              </q-tab-panels>
            </template>
          </q-splitter>
        </q-card-section>
      </q-card>
    </q-dialog>
  </q-page>
</template>

<script>
export default {
  name: 'autocode',
}
</script>
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
  getTables,
  getTable,
  getColumns,
  previewCode,
  genCode,
  gendb,
  genCoverCode,
} from 'src/api/devtools/datasource'
import { getGenTable } from 'src/api/devtools/table'
import { listDict } from 'src/api/admin/dictType'
import { useQuasar, format } from 'quasar'
import { requiredRule } from 'src/utils/inputRule'
import {
  DictOptions,
  getDictLabel,
  getOptionsByList,
  getDict,
} from 'src/utils/dict'
import { camelCase, titleUpperCase } from 'src/utils/str'
import { useRoute, useRouter } from 'vue-router'

const router = useRouter()

const $q = useQuasar()
const tag = ref('')
let { proxy } = getCurrentInstance()

const route = useRoute()
const splitterModel = ref(20)

const dialogVisible = ref(false)
const codeVisible = ref(false)

const formType = ref('')
const dictId = ref(0)
const db = ref(null)
const table = ref(null)
const tableDetail = ref({})

const sourceList = ref([])
const tableList = ref([])
const fieldList = ref([])

const code = ref([])

const form = ref({})

const pagination = ref({
  sortBy: 'id',
  descending: false,
  page: 1,
  rowsPerPage: 9999,
})
const columns = computed(() => {
  return [
    {
      name: 'sort',
      align: 'left',
      label: '序列',
      field: 'sort',
    },
    {
      name: 'gormName',
      align: 'center',
      label: '名称',
      field: 'gormName',
      sortable: true,
    },
    {
      name: 'columnComment',
      align: 'left',
      label: '中文名',
      field: 'columnComment',
      sortable: true,
      classes: 'ellipsis',
      style: 'max-width: 100px',
    },
    {
      name: 'fieldName',
      align: 'left',
      label: 'go字段名',
      field: 'fieldName',
    },
    {
      name: 'dbType',
      align: 'left',
      label: '数据库类型',
      field: 'dbType',
    },
    {
      name: 'dataType',
      align: 'left',
      label: 'go数据类型',
      field: 'dataType',
    },
    {
      name: 'require',
      align: 'left',
      label: '必填',
      field: 'require',
      style: 'max-width: 30px',
    },
    {
      name: 'isList',
      align: 'left',
      label: '显示',
      field: 'isList',
      style: 'max-width: 30px',
    },
    {
      name: 'isEdit',
      align: 'left',
      label: '编辑',
      field: 'isEdit',
      style: 'max-width: 30px',
    },
    { name: 'actions', align: 'center', label: '操作', field: 'actions' },
  ]
})

const fieldOptions = ref([])
const htmlOptions = ref([])

const dictList = ref([])
const columnIndex = ref(0)
const dbVisable = ref(true)

const init = async () => {
  htmlOptions.value = await getDict('htmlType')
  fieldOptions.value = await getDict('filed_type')

  const queryReq = {
    pageIndex: 0,
    pageSize: 9999,
  }
  let res = await listDict(queryReq)
  dictList.value = getOptionsByList(res.list, 'name', 'type')

  // onRequest()
  listsource()
  let id = route.query.id
  let tablePrefix = route.query.tablePrefix
  if (id) {
    dbVisable.value = false
    getTableById(id, tablePrefix)
  }
}

onMounted(() => {
  if (route.name === 'autoCodeEdit') {
    if (!route.query.id) {
      router.push({ name: 'gencode_history' })
    }
  }
  init()
})

watch(
  () => route.query.id,
  (id) => {
    if (route.name === 'autoCodeEdit') {
      init()
    } else {
      reset()
      dbVisable.value = true
      tableDetail.value = {}
    }
  }
)

watch(
  () => form.value.gormName,
  (newValue, oldValue) => {
    form.value.fieldJson = camelCase(form.value.gormName)
    form.value.fieldName = titleUpperCase(form.value.fieldJson)
  }
)

const reset = () => {
  pagination.value = {
    sortBy: 'sort',
    descending: false,
    page: 1,
    rowsPerPage: 9999,
    rowsNumber: 0,
  }
  form.value = {
    dictType: '',
    tablename: tableDetail.value.name,
  }
  fieldList.value = []
  dictId.value = 0
  columnIndex.value = 0
}

const create = () => {
  formType.value = '新建'
  reset()
  dialogVisible.value = true
}
const edit = (p, index) => {
  reset()
  form.value = {
    ...p,
  }

  columnIndex.value = index
  formType.value = '编辑'
  dialogVisible.value = true
}

const del = async (index) => {
  tableDetail.value.columns.splice(index, 1)
}

const moveUpField = (index) => {
  if (index === 0) {
    return
  }
  tableDetail.value.columns[index].sort--
  tableDetail.value.columns[index - 1].sort++

  const oldUpField = tableDetail.value.columns[index - 1]
  oldUpField.sort = index
  tableDetail.value.columns.splice(index - 1, 1)
  tableDetail.value.columns.splice(index, 0, oldUpField)
}
const moveDownField = (index) => {
  const fCount = tableDetail.value.columns.length
  if (index === fCount - 1) {
    return
  }
  tableDetail.value.columns[index + 1].sort--
  tableDetail.value.columns[index].sort++
  const oldDownField = tableDetail.value.columns[index + 1]
  tableDetail.value.columns.splice(index + 1, 1)
  tableDetail.value.columns.splice(index, 0, oldDownField)
}

const submit = async () => {
  // const res = undefined;
  if (formType.value === '编辑') {
    tableDetail.value.columns[columnIndex.value] = form.value

    // let res = await updateDatasource(form.value);
  } else if (formType.value === '新建') {
    tableDetail.value.columns.push(form.value)
  } else {
    proxy.$error('路径错误')
  }
  dialogVisible.value = false
  reset()
  // onRequest()
}
const listsource = async () => {
  let res = await listDatasource()
  sourceList.value = res.list
}
const listTables = async (val) => {
  let data = {
    id: val.id,
  }
  let res = await getTables(data)
  tableList.value = res.list
}

const detaiTable = async (val) => {
  tableDetail.value = {}
  let data = {
    id: db.value.id,
    name: val.name,
    comment: val.tableComment,
  }

  let res = await getTable(data)
  tableDetail.value = res
  tableDetail.value.tablePrefix = db.value.tablePrefix
  tableDetail.value.columns.forEach((row, index) => {
    row.sort = index
  })
}

const getTableById = async (id, tablePrefix) => {
  tableDetail.value = {}
  let req = {
    id: id,
  }
  let res = await getGenTable(req)
  tableDetail.value = res
  tableDetail.value.tablePrefix = tablePrefix
  tableDetail.value.columns.forEach((row, index) => {
    row.sort = index
  })
  db.value = {}
  db.value.id = res.sourceId
  let reqTables = {
    id: res.sourceId,
  }
  await listTables(reqTables)
}
const detaiFkTable = async (val) => {
  fieldList.value = {}
  let data = {
    id: db.value.id,
    name: val.name,
    comment: val.tableComment,
  }
  form.value.fkTable = val.name
  let res = await getColumns(data)
  fieldList.value = res.list
}
const preview = async (prop) => {
  tableDetail.value.id = db.value.id
  tableDetail.value.sourceId = db.value.id
  let res = await previewCode(tableDetail.value)
  code.value = res.list
  codeVisible.value = true
  tag.value = res.list[0].name
}

const getCode = async (prop) => {
  tableDetail.value.id = db.value.id
  tableDetail.value.sourceId = db.value.id
  let res = await genCode(tableDetail.value)
}
const coverCode = async (prop) => {
  tableDetail.value.id = db.value.id
  tableDetail.value.sourceId = db.value.id
  let res = await genCoverCode(tableDetail.value)
}
const genMenu = async (prop) => {
  tableDetail.value.id = db.value.id
  tableDetail.value.sourceId = db.value.id
  let res = await gendb(tableDetail.value)
}
</script>
