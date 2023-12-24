<template>
  <div>
    <div class="q-pa-md q-ml-xs">
      <q-chip size="18px" icon="bookmark"> Mails通知模板支持的变量 </q-chip>
      <div class="text-h6 q-ml-md">
        <q-badge color="primary">TaskId</q-badge>
        任务ID
      </div>
      <div class="text-h6 q-ml-md">
        <q-badge color="primary">TaskName</q-badge>
        任务名称
      </div>
      <div class="text-h6 q-ml-md">
        <q-badge color="primary">Status</q-badge>
        任务执行结果状态
      </div>
      <div class="text-h6 q-ml-md">
        <q-badge color="primary">Result</q-badge>
        任务执行输出
      </div>
    </div>
    <q-separator />

    <q-form @submit="submit" class="q-pa-md">
      <q-chip size="18px" icon="bookmark"> 邮件服务器配置 </q-chip>
      <div class="row q-col-gutter-x-md dialog_form q-pt-md">
        <q-input
          outlined
          dense
          class="col-6 q-pb-md"
          v-model="form.host"
          label="SMTP服务器"
        />
        <q-input
          outlined
          dense
          type="number"
          class="col-6 q-pb-md"
          v-model.number="form.port"
          label="SMTP服务器端口"
        />
        <q-input
          outlined
          dense
          class="col-6 q-pb-md"
          v-model="form.user"
          label="邮箱用户名"
        />
        <q-input
          outlined
          dense
          class="col-6 q-pb-md"
          v-model="form.password"
          label="密码"
        />
      </div>
      <q-chip size="18px" icon="bookmark"> 邮件通知模板支持html </q-chip>
      <div class="row q-col-gutter-x-md dialog_form q-pt-md">
        <q-input
          outlined
          dense
          type="textarea"
          class="col-12 q-pb-md"
          v-model="form.template"
          label="邮件模版"
        />
      </div>
      <q-separator class="q-ma-sm" />
      <q-table
        dense
        flat
        bordered
        separator="cell"
        :columns="columns"
        :rows="form.mailUsers"
        row-key="id"
        @request="onRequest"
        :grid="$q.screen.xs"
        v-model:pagination="pagination"
        binary-state-sort
      >
        <template v-slot:top="table">
          <div class="row no-wrap full-width">
            <q-space />
            <q-btn-group class="q-gutter-xs">
              <q-btn
                v-permission="'appsns:add'"
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
            </q-btn-group>
          </div>
        </template>
        <template v-slot:body-cell-actions="props">
          <q-td :props="props">
            <div class="q-gutter-xs">
              <com-del
                v-permission="'appsns:del'"
                label="通知用户邮件配置"
                @confirm="del(props.row)"
              />
            </div>
          </q-td>
        </template>
      </q-table>
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
    <q-dialog v-model="dialogVisible" persistent>
      <q-card style="width: 800px; min-width: 40vw">
        <q-bar class="bg-primary text-white">
          <span> 通知用户邮件配置 </span>
          <q-space />
          <q-btn dense flat icon="close" v-close-popup>
            <q-tooltip class="bg-white text-primary">关闭</q-tooltip>
          </q-btn>
        </q-bar>
        <q-card-section>
          <q-form @submit="mailSubmit">
            <div class="row q-col-gutter-x-md dialog_form q-pt-md">
              <q-input
                outlined
                dense
                class="col-6 q-pb-md"
                v-model="mailForm.username"
                label="用户名称"
                :rules="[requiredRule]"
              />
              <q-input
                outlined
                dense
                class="col-6 q-pb-md"
                v-model="mailForm.email"
                label="邮件地址"
                :rules="[emailRule]"
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
  </div>
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
  listSetting,
  createMailUser,
  updateMail,
  deleteSetting,
  getTaskMailServer,
  deleteMailUser,
} from 'src/api/task/setting'

import { useQuasar } from 'quasar'
import { requiredRule, emailRule } from 'src/utils/inputRule'
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
const form = ref({})
const formType = ref('')
const mailForm = ref({})
const dialogVisible = ref(false)
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
      field: 'username',
      sortable: true,
      classes: 'ellipsis',
    },
    {
      name: 'email',
      align: 'left',
      label: '邮件地址',
      field: 'email',
      sortable: true,
      classes: 'ellipsis',
    },
    { name: 'actions', align: 'center', label: '操作', field: 'actions' },
  ]
})

const getSelectedString = () => {
  return selected.value.length === 0
    ? ''
    : `${selected.value.length} record${
        selected.value.length > 1 ? 's' : ''
      } selected of ${dataList.value.length}`
}
onMounted(async () => {
  onRequest()
})
const onRequest = async (val) => {
  let queryReq = {}
  let res = await getTaskMailServer(queryReq)
  form.value = res
}
const reset = () => {
  pagination.value = {
    sortBy: 'id',
    descending: false,
    page: 1,
    rowsPerPage: 10,
    rowsNumber: 0,
  }
  mailForm.value = {
    username: '',
    email: '',
  }
}
const create = () => {
  reset()
  dialogVisible.value = true
}

const del = async (p) => {
  let res = await deleteMailUser(p)
  onRequest()
}
const mailSubmit = async () => {
  await createMailUser(mailForm.value)
  dialogVisible.value = false
  onRequest()
}
const submit = async () => {
  // const res = undefined
  await updateMail(form.value)
  dialogVisible.value = false
}
</script>
