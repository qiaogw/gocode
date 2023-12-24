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
                v-permission="'log:add'"
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
                v-permission="'log:del'"
                flat
                glossy
                label="任务日志"
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
        <template v-slot:body-cell-taskId="props">
          <q-td :props="props">
            <q-chip
              dense
              text-color="white"
              :color="props.value ? 'positive' : 'grey'"
              >{{ formatTaskTask(props.value) }}
            </q-chip>
          </q-td>
        </template>
        <template v-slot:body-cell-protocol="props">
          <q-td :props="props">
            <q-chip
              dense
              text-color="white"
              :color="props.value ? 'positive' : 'grey'"
              >{{ formattaskProtocol(props.value) }}
            </q-chip>
          </q-td>
        </template>
        <template v-slot:body-cell-status="props">
          <q-td :props="props">
            <q-chip
              v-if="props.value != '2'"
              dense
              text-color="white"
              :color="taskStatusColor(props.value)"
              >{{ formattaskStatus(props.value) }}
            </q-chip>
            <q-btn
              v-if="props.value === '2'"
              rounded
              size="md"
              color="primary"
              @click="stop(props.row)"
              >{{ formattaskStatus(props.value) }}
              <q-tooltip>点击结束任务</q-tooltip>
            </q-btn>
          </q-td>
        </template>

        <template v-slot:body-cell-actions="props">
          <q-td :props="props">
            <div class="q-gutter-xs">
              <q-btn
                dense
                color="positive"
                glossy
                size="sm"
                icon="list_alt"
                label="查看结果"
                @click="showTaskResult(props.row)"
              >
                <q-tooltip>查看任务结果</q-tooltip></q-btn
              >
              <com-edit
                v-permission="'log:edit'"
                label="任务日志"
                @confirm="edit(props.row)"
              />
              <com-del
                v-permission="'log:del'"
                label="任务日志"
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
          <span> {{ formType }}任务日志 </span>
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
                v-model="form.taskId"
                :options="taskTaskOptions"
                label="任务"
              />
              <q-input
                outlined
                dense
                class="col-6 q-pb-md"
                v-model="form.name"
                label="名称"
              />
              <q-input
                outlined
                dense
                class="col-6 q-pb-md"
                v-model="form.spec"
                label="任务crontab"
              />
              <q-select
                class="col-6 q-pb-md"
                outlined
                dense
                map-options
                emit-value
                option-value="value"
                option-label="label"
                v-model="form.protocol"
                :options="taskProtocolOptions"
                label="协议类型"
              />
              <q-input
                outlined
                dense
                class="col-6 q-pb-md"
                v-model="form.command"
                label="任务命令"
              />
              <q-input
                outlined
                dense
                type="number"
                class="col-6 q-pb-md"
                v-model.number="form.timeout"
                label="超时时间"
              />
              <q-input
                outlined
                dense
                type="number"
                class="col-6 q-pb-md"
                v-model.number="form.retryTimes"
                label="重试次数"
              />
              <q-input
                outlined
                dense
                class="col-6 q-pb-md"
                v-model="form.hostname"
                label="主机名称"
              />
              <q-input
                outlined
                dense
                class="col-6 q-pb-md"
                v-model="form.startTime"
                label="开始时间"
              />
              <q-input
                outlined
                dense
                class="col-6 q-pb-md"
                v-model="form.endTime"
                label="完成时间"
              />
              <q-select
                class="col-6 q-pb-md"
                outlined
                dense
                map-options
                emit-value
                option-value="value"
                option-label="label"
                v-model="form.status"
                :options="taskStatusOptions"
                label="状态"
              />
              <q-input
                outlined
                dense
                class="col-6 q-pb-md"
                v-model="form.result"
                label="执行结果"
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

    <q-dialog v-model="taskResultVisible" persistent>
      <q-card style="width: 800px; min-width: 40vw">
        <q-bar class="bg-primary text-white">
          <span> 任务结果详情 </span>
          <q-space />
          <q-btn dense flat icon="close" v-close-popup>
            <q-tooltip class="bg-white text-primary">关闭</q-tooltip>
          </q-btn>
        </q-bar>
        <q-card-section>
          <div>
            <q-chip size="xs" icon="bookmark"> 任务信息 </q-chip>
            <pre>{{ currentTaskResult.command }}</pre>
          </div>
          <div class="fit wrap items-start">
            <q-chip size="xs" icon="bookmark"> 执行结果 </q-chip>
            <pre>{{ currentTaskResult.result }}</pre>
          </div>
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
  listLog,
  createLog,
  updateLog,
  deleteLog,
  deleteLogList,
  stopTask,
} from 'src/api/task/log'
import { listTask } from 'src/api/task/task'

import { useQuasar } from 'quasar'
import { requiredRule } from 'src/utils/inputRule'
import { getIds } from 'src/utils/arrayOrObject'
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

const taskResultVisible = ref(false)
const currentTaskResult = reactive({
  command: '',
  result: '',
})

const dataList = ref([])
const formType = ref('')
const dictId = ref(0)
const form = ref({})
const dictOptions = ref({})
const taskTaskOptions = ref([])
const formatTaskTask = (prop) => {
  return getDictLabel(taskTaskOptions.value, prop)
}
const taskProtocolOptions = ref([])
const formattaskProtocol = (prop) => {
  if (!prop) {
    prop = false
  }
  return getDictLabel(dictOptions.value.taskProtocol, prop)
}
const taskStatusOptions = ref([])
const taskStatusColor = (prop) => {
  let tcolor = 'positive'
  if (!prop) {
    return tcolor
  }
  switch (prop) {
    case '-1':
      tcolor = 'grey'
      break
    case '0':
      tcolor = 'negative'
      break
    case '1':
      tcolor = 'secondary'
      break
    case '2':
      tcolor = 'info'
      break
    case '3':
      tcolor = 'positive'
      break
    case '4':
      tcolor = 'blue-grey'
      break
    default:
      tcolor = 'positive'
  }
  return tcolor
}
const formattaskStatus = (prop) => {
  if (!prop) {
    prop = false
  }
  return getDictLabel(dictOptions.value.taskStatus, prop)
}

const searchKey = ref('')

const selected = ref([])
const pagination = ref({
  sortBy: 'start_time',
  descending: true,
  page: 1,
  rowsPerPage: 10,
  rowsNumber: 0,
})
const columns = computed(() => {
  return [
    {
      name: 'taskId',
      align: 'left',
      label: '任务',
      field: 'taskId',
      sortable: true,
      classes: 'ellipsis',
    },
    {
      name: 'spec',
      align: 'left',
      label: '任务crontab',
      field: 'spec',
      sortable: true,
      classes: 'ellipsis',
    },
    {
      name: 'protocol',
      align: 'left',
      label: '协议类型',
      field: 'protocol',
      sortable: true,
      classes: 'ellipsis',
    },
    {
      name: 'startTime',
      align: 'left',
      label: '开始时间',
      field: 'startTime',
      sortable: true,
      classes: 'ellipsis',
    },
    {
      name: 'endTime',
      align: 'left',
      label: '完成时间',
      field: 'endTime',
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

    { name: 'actions', align: 'center', label: '操作', field: 'actions' },
  ]
})

onMounted(async () => {
  dictOptions.value = await DictOptions()
  const queryReqTaskTask = {
    pageIndex: 0,
    pageSize: 9999,
  }
  let resTaskTask = await listTask(queryReqTaskTask)
  taskTaskOptions.value = getOptionsByList(resTaskTask.list, 'name', 'id')
  taskProtocolOptions.value = await getDict('taskProtocol')
  taskStatusOptions.value = await getDict('taskStatus')

  if (!route.query.id) {
    onRequest()
  } else {
    let val = {
      pagination: {
        sortBy: 'start_time',
        descending: true,
        page: 1,
        rowsPerPage: 10,
        rowsNumber: 0,
        taskId: route.query.id,
      },
    }

    onRequest(val)
  }
})

const reset = () => {
  pagination.value = {
    sortBy: 'start_time',
    descending: true,
    page: 1,
    rowsPerPage: 10,
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
  const { page, rowsPerPage, sortBy, descending, taskId } = val.pagination
  let queryReq = {}

  queryReq.pageSize = rowsPerPage
  queryReq.pageIndex = page
  queryReq.sortBy = sortBy
  queryReq.descending = descending
  queryReq.searchKey = val.filter
  queryReq.taskId = taskId

  // console.log(val.pagination)
  // console.log('queryReq', queryReq)

  let table = await listLog(queryReq)

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
  let res = await deleteLog(p)
  onRequest()
}

const delList = async () => {
  // let res = await deleteDatasource()
  // onRequest()
  // console.log('dellist')
  let req = {
    ids: getIds(selected.value),
  }
  let res = await deleteLogList(req)
  onRequest()
}
const submit = async () => {
  // const res = undefined
  if (formType.value === '编辑') {
    let res = await updateLog(form.value)
  } else if (formType.value === '新建') {
    let res = await createLog(form.value)
  } else {
    proxy.$error('请求错误')
  }
  dialogVisible.value = false
  reset()
  onRequest()
}

const showTaskResult = (item) => {
  // const res = undefined
  taskResultVisible.value = true
  currentTaskResult.command = item.command
  currentTaskResult.result = item.result
}
const stop = async (item) => {
  // const res = undefined
  // console.log(item)
  let req = {
    id: item.id,
  }
  await stopTask(req)
}
</script>
<style scoped>
pre {
  white-space: pre-wrap;
  word-wrap: break-word;
  padding-left: 10px;
  padding-top: none;
  margin: 10px;
  background-color: #4c4c4c;
  color: white;
}
</style>
