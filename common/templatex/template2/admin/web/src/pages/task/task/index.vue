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
        v-model:selected="selected"
      >
        <template v-slot:top="table">
          <div class="row no-wrap full-width">
            <com-search @query="onRequest" v-model:filter="searchKey" />
            <q-space />
            <q-btn-group class="q-gutter-xs">
              <q-btn
                v-permission="'task:add'"
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
        <template v-slot:body-cell-level="props">
          <q-td :props="props">
            <q-chip
              dense
              text-color="white"
              :color="props.value ? 'positive' : 'grey'"
              >{{ formattaskLevel(props.value) }}
            </q-chip>
          </q-td>
        </template>
        <template v-slot:body-cell-dependencyStatus="props">
          <q-td :props="props">
            <q-chip
              dense
              text-color="white"
              :color="props.value ? 'positive' : 'grey'"
              >{{ formattaskDependency(props.value) }}
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
        <template v-slot:body-cell-spec="props">
          <q-td :props="props">
            <q-chip
              dense
              text-color="white"
              :color="props.value ? 'positive' : 'grey'"
              >{{ crontabFormat(props.value) }}
            </q-chip>
          </q-td>
        </template>
        <template v-slot:body-cell-multi="props">
          <q-td :props="props">
            <q-chip
              dense
              text-color="white"
              :color="props.value ? 'positive' : 'grey'"
              >{{ formatsys_yes_no(props.value) }}
            </q-chip>
          </q-td>
        </template>
        <template v-slot:body-cell-notifyStatus="props">
          <q-td :props="props">
            <q-chip
              dense
              text-color="white"
              :color="props.value ? 'positive' : 'grey'"
              >{{ formatsys_yes_no(props.value) }}
            </q-chip>
          </q-td>
        </template>
        <template v-slot:body-cell-notifyType="props">
          <q-td :props="props">
            <q-chip
              dense
              text-color="white"
              :color="props.value ? 'positive' : 'grey'"
              >{{ formattaskHookType(props.value) }}
            </q-chip>
          </q-td>
        </template>
        <template v-slot:body-cell-status="props">
          <q-td :props="props">
            <q-popup-edit
              v-model="props.row.status"
              title="更新状态"
              buttons
              @save="updateStatus(props.row)"
              v-slot="scope"
            >
              <q-toggle
                label="状态"
                color="green"
                dense
                autofocus
                v-model="scope.value"
                @update:model-value="updateStatus2"
              />
            </q-popup-edit>
            <q-chip
              dense
              text-color="white"
              :color="props.value ? 'positive' : 'grey'"
              >{{ formatsys_enabled(props.value) }}
            </q-chip>
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
                icon="play_circle_outline"
                @click.stop
              >
                <q-tooltip>手动执行一次</q-tooltip>
                <q-menu auto-close anchor="center left" self="center right">
                  <div
                    class="row no-wrap items-center q-pa-sm bg-warning text-white"
                  >
                    <span class="text-no-wrap"
                      >确认手动执行一次 <b>{{ props.row.name }}</b> ?</span
                    >
                    <q-btn
                      flat
                      round
                      dense
                      color="primary"
                      glossy
                      icon="play_circle"
                      @click="runTask(props.row)"
                    >
                      <q-tooltip>确认</q-tooltip></q-btn
                    >
                  </div>
                </q-menu></q-btn
              >
              <q-btn
                flat
                round
                dense
                color="secondary"
                glossy
                icon="list_alt"
                @click="logTask(props.row)"
              >
                <q-tooltip>查看任务结果</q-tooltip></q-btn
              >
              <com-edit
                v-permission="'task:edit'"
                label="任务信息"
                @confirm="edit(props.row)"
              />
              <com-del
                v-permission="'task:del'"
                label="任务信息"
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
          <span> {{ formType }}任务信息 </span>
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
                v-model="form.level"
                :options="taskLevelOptions"
                label="任务等级"
                :rules="[requiredRule]"
              />
              <q-select
                v-show="form.level === '1'"
                class="col-6 q-pb-md"
                outlined
                dense
                map-options
                emit-value
                option-value="value"
                option-label="label"
                v-model="form.dependencyStatus"
                :options="taskDependencyOptions"
                label="依赖关系"
              />
              <q-select
                v-show="form.level === '1'"
                class="col-6 q-pb-md"
                outlined
                dense
                filled
                map-options
                use-chips
                multiple
                option-value="id"
                option-label="name"
                v-model="form.dependencyTaskId"
                :options="taskDependencyTaskOptions"
                label="依赖任务"
              />

              <q-input
                v-show="form.level === '1'"
                outlined
                dense
                class="col-6 q-pb-md"
                v-model="form.spec"
                label="任务定时器"
                :rules="[requiredRule]"
              >
                <template v-slot:append>
                  <q-icon name="event" class="cursor-pointer">
                    <q-popup-proxy cover :breakpoint="600">
                      <ComCron v-model:cronStr="form.spec" />
                    </q-popup-proxy>
                  </q-icon>
                </template>
                <!-- -->
              </q-input>
              <q-chip
                v-show="form.level === '1'"
                class="col-5 q-pa-sm q-ml-md q-mb-md"
                color="teal"
                square
                icon="alarm"
                :label="crontabFormat(form.spec)"
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
                :rules="[requiredRule]"
              />
              <q-select
                v-show="form.protocol === '4' || form.protocol === '5'"
                class="col-6 q-pb-md"
                outlined
                dense
                filled
                map-options
                use-chips
                multiple
                option-value="id"
                option-label="name"
                v-model="form.nodes"
                :options="nodeOpions"
                label="节点"
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
                class="col-6 q-pb-md"
                v-model="form.args"
                label="命令参数"
              />

              <q-select
                v-show="form.protocol === '1'"
                class="col-6 q-pb-md"
                outlined
                dense
                map-options
                emit-value
                option-value="value"
                option-label="label"
                v-model="form.httpMethod"
                :options="http_methodsOptions"
                label="http请求方法"
              />
              <q-input
                outlined
                dense
                type="number"
                class="col-6 q-pb-md"
                v-model.number="form.timeout"
                label="超时时间(秒)"
                :rules="[requiredRule]"
              />

              <q-input
                outlined
                dense
                type="number"
                class="col-6 q-pb-md"
                v-model.number="form.retryTimes"
                label="重试次数"
                :rules="[requiredRule]"
              />
              <q-input
                outlined
                dense
                type="number"
                class="col-6 q-pb-md"
                v-model.number="form.retryInterval"
                label="重试间隔时间(秒)"
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
                v-model="form.notifyStatus"
                :options="taskNotifyTypeOptions"
                label="任务执行结束是否通知"
                :rules="[requiredRule]"
              />
              <q-select
                v-show="form.notifyStatus !== '0'"
                class="col-6 q-pb-md"
                outlined
                dense
                map-options
                emit-value
                option-value="value"
                option-label="label"
                v-model="form.notifyType"
                :options="taskHookTypeOptions"
                label="通知类型"
              />
              <q-select
                v-show="form.notifyStatus !== '0' && form.notifyType === 'mail'"
                class="col-6 q-pb-md"
                outlined
                dense
                map-options
                emit-value
                option-value="id"
                option-label="username"
                v-model="form.notifyReceiverId"
                :options="mailUsers"
                label="通知接受者ID"
              />
              <q-input
                v-show="form.notifyStatus === '3'"
                outlined
                dense
                class="col-6 q-pb-md"
                v-model="form.notifyKeyword"
                label="任务执行输出关键字，任务执行输出中包含此关键字将触发通知"
              />
              <q-input
                outlined
                dense
                class="col-6 q-pb-md"
                v-model="form.remark"
                label="备注"
              />
              <q-toggle
                class="col-6 q-pb-md"
                label="是否允许多实例运行"
                color="green"
                false-value="N"
                true-value="Y"
                v-model="form.multi"
              />
              <q-toggle
                class="col-6 q-pb-md"
                label="状态"
                color="green"
                v-model="form.status"
              />
            </div>
            <q-separator class="q-ma-sm" />
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
  listTask,
  createTask,
  updateTask,
  deleteTask,
  getTask,
  getChild,
  run,
} from 'src/api/task/task'
import { listNode } from 'src/api/task/node'
import { getTaskMailServer } from 'src/api/task/setting'

import { useQuasar } from 'quasar'
import { requiredRule } from 'src/utils/inputRule'
import {
  DictOptions,
  getOptionsByList,
  getDictLabel,
  getDict,
} from 'src/utils/dict'
import { downloadAction } from 'src/api/manage'
import { useRouter, useRoute } from 'vue-router'

const router = useRouter()

import ComCron from 'src/components/comCron/index.vue'
import cronstrue from 'cronstrue/i18n'

const route = useRoute()

const $q = useQuasar()
let { proxy } = getCurrentInstance()
const dialogVisible = ref(false)
const dataList = ref([])
const formType = ref('')
// const dictId = ref(0)
const form = ref({})
const dictOptions = ref({})
const mailUsers = ref([])
const taskLevelOptions = ref([])

const taskDependencyTaskOptions = ref([])

const taskStatus = ref(null)
const formattaskLevel = (prop) => {
  if (!prop) {
    prop = false
  }
  return getDictLabel(dictOptions.value.taskLevel, prop)
}
const taskDependencyOptions = ref([])
const formattaskDependency = (prop) => {
  if (!prop) {
    prop = false
  }
  return getDictLabel(dictOptions.value.taskDependency, prop)
}
const taskProtocolOptions = ref([])
const formattaskProtocol = (prop) => {
  if (!prop) {
    prop = false
  }
  return getDictLabel(dictOptions.value.taskProtocol, prop)
}
const http_methodsOptions = ref([])
const formathttp_methods = (prop) => {
  if (!prop) {
    prop = false
  }
  return getDictLabel(dictOptions.value.http_methods, prop)
}
const nodeOpions = ref([])

const sys_yes_noOptions = ref([])
const formatsys_yes_no = (prop) => {
  if (!prop) {
    prop = false
  }
  return getDictLabel(dictOptions.value.sys_yes_no, prop)
}
const taskHookTypeOptions = ref([])
const formattaskHookType = (prop) => {
  if (!prop) {
    prop = false
  }
  return getDictLabel(dictOptions.value.taskHookType, prop)
}
const sys_enabledOptions = ref([])
const formatsys_enabled = (prop) => {
  if (!prop) {
    prop = false
  }
  return getDictLabel(dictOptions.value.sys_enabled, prop)
}
const taskNotifyTypeOptions = ref([])
const formattaskNotifyType = (prop) => {
  if (!prop) {
    prop = false
  }
  return getDictLabel(dictOptions.value.taskNotifyType, prop)
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
      name: 'name',
      align: 'left',
      label: '名称',
      field: 'name',
      sortable: true,
      classes: 'ellipsis',
    },
    {
      name: 'level',
      align: 'left',
      label: '任务等级',
      field: 'level',
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
  let queryReq = {}
  let res = await getTaskMailServer(queryReq)
  mailUsers.value = res.mailUsers
  res = await listNode(queryReq)
  nodeOpions.value = res.list
  res = await getChild(queryReq)
  taskDependencyTaskOptions.value = res.list
  // console.log(res)
  taskLevelOptions.value = await getDict('taskLevel')
  taskDependencyOptions.value = await getDict('taskDependency')
  taskProtocolOptions.value = await getDict('taskProtocol')
  http_methodsOptions.value = await getDict('http_methods')
  sys_yes_noOptions.value = await getDict('sys_yes_no')
  sys_yes_noOptions.value = await getDict('sys_yes_no')
  taskHookTypeOptions.value = await getDict('taskHookType')
  sys_enabledOptions.value = await getDict('sys_enabled')
  taskNotifyTypeOptions.value = await getDict('taskNotifyType')
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
    notifyStatus: '0',
    level: '1',
    timeout: 0,
    retryTimes: 3,
    retryInterval: 60,
  }
  // dictId.value = 0
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

  let table = await listTask(queryReq)

  pagination.value = val.pagination

  pagination.value.rowsNumber = table.count
  if (table.list) {
    dataList.value = table.list
  }
}

const create = () => {
  reset()

  formType.value = '新建'
  dialogVisible.value = true
}
const edit = async (p) => {
  reset()

  let res = await getTask(p)

  form.value = res

  formType.value = '编辑'
  dialogVisible.value = true
}

const del = async (p) => {
  let res = await deleteTask(p)
  onRequest()
}

const submit = async () => {
  // const res = undefined
  if (formType.value === '编辑') {
    let res = await updateTask(form.value)
  } else if (formType.value === '新建') {
    let res = await createTask(form.value)
  } else {
    proxy.$error('请求错误')
  }
  dialogVisible.value = false
  reset()
  onRequest()
}

const crontabFormat = (row) => {
  // console.log(row.spec)
  if (!row) {
    return ''
  }
  return cronstrue.toString(row, { locale: 'zh_CN' })
}
const runTask = async (row) => {
  let res = await run(row)
  // console.log(res)
}
const logTask = (row) => {
  if (row) {
    router.push({
      name: 'task_log',
      query: {
        id: row.id,
      },
    })
  } else {
    router.push({ name: 'task_log' })
  }
}
const updateStatus = async (val) => {
  // console.log('updateStatus:', val.status)
  val.status = taskStatus.value
  let res = await updateTask(val)
  onRequest()
  // console.log(res)
}
const updateStatus2 = async (val, evt) => {
  // console.log('updateStatus2:', val, evt)
  taskStatus.value = val
}
</script>
