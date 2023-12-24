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
                v-permission="'user:add'"
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
                v-permission="'user:del'"
                flat
                glossy
                label="用户"
                @confirm="delList"
              />
              <q-btn
                :disable="selected.length < 1"
                v-permission="'user:edit'"
                flat
                dense
                glossy
                icon="lock_reset"
                no-wrap
                color="primary"
                round
                @click.stop
              >
                <q-tooltip>重置密码</q-tooltip>
                <q-menu auto-close anchor="center left" self="center right">
                  <div class="row no-wrap items-center q-pa-sm bg-cyan-2">
                    <span class="text-no-wrap">确认 <b>重置密码</b> ？</span>
                    <q-btn
                      flat
                      round
                      dense
                      color="primary"
                      icon="mode_edit"
                      @click="resetPwd"
                    >
                      <q-tooltip>确认修改</q-tooltip></q-btn
                    >
                  </div>
                </q-menu>
              </q-btn>
              <q-btn
                v-permission="'user:export'"
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
                v-permission="'user:export'"
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
                v-permission="'user:import'"
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
        <template v-slot:body-cell-deptId="props">
          <q-td :props="props">
            <q-chip
              dense
              text-color="white"
              :color="props.value ? 'positive' : 'grey'"
              >{{ formatAdminDept(props.value) }}
            </q-chip>
          </q-td>
        </template>
        <template v-slot:body-cell-postId="props">
          <q-td :props="props">
            <q-chip
              dense
              text-color="white"
              :color="props.value ? 'positive' : 'grey'"
              >{{ formatAdminPost(props.value) }}
            </q-chip>
          </q-td>
        </template>
        <template v-slot:body-cell-status="props">
          <q-td :props="props">
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
              <com-edit
                v-permission="'user:edit'"
                label="用户"
                @confirm="edit(props.row)"
              />
              <com-del
                v-permission="'user:del'"
                label="用户"
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
          <span> 用户 </span>
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
                :disable="formType === '新建' ? false : true"
                v-model="form.username"
                label="用户名"
                :rules="[requiredRule]"
              />
              <q-input
                outlined
                dense
                class="col-6 q-pb-md"
                v-model="form.nickName"
                label="昵称"
                :rules="[requiredRule]"
              />
              <q-select
                class="col-6 q-pb-md"
                outlined
                dense
                filled
                map-options
                use-chips
                multiple
                option-value="id"
                option-label="name"
                v-model="form.roles"
                :options="adminRoleOptions"
                label="角色"
              />
              <q-select
                :disable="form.roles.length > 1 ? false : true"
                class="col-6 q-pb-md"
                outlined
                dense
                map-options
                emit-value
                option-value="id"
                option-label="name"
                v-model="form.roleId"
                :options="form.roles"
                label="默认角色"
              />
              <q-select
                class="col-6 q-pb-md"
                outlined
                dense
                map-options
                emit-value
                option-value="value"
                option-label="label"
                v-model="form.deptId"
                :options="adminDeptOptions"
                label="部门"
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
                v-model="form.postId"
                :options="adminPostOptions"
                label="职务"
                :rules="[requiredRule]"
              />
              <q-input
                outlined
                dense
                class="col-6 q-pb-md"
                v-model="form.mobile"
                label="手机号"
                :rules="[requiredRule]"
              />
              <q-input
                outlined
                dense
                class="col-6 q-pb-md"
                v-model="form.avatar"
                label="头像"
              />
              <q-input
                outlined
                dense
                class="col-6 q-pb-md"
                v-model="form.gender"
                label="性别"
              />
              <q-input
                outlined
                dense
                class="col-6 q-pb-md"
                v-model="form.email"
                label="邮箱"
              />
              <q-input
                outlined
                dense
                class="col-6 q-pb-md"
                v-model="form.baseColor"
                label="基础颜色"
              />
              <q-input
                outlined
                dense
                class="col-6 q-pb-md"
                v-model="form.activeColor"
                label="活跃颜色"
              />
              <q-input
                outlined
                dense
                type="number"
                class="col-6 q-pb-md"
                v-model.number="form.sort"
                label="排序"
                :rules="[requiredRule]"
              />
              <q-input
                outlined
                dense
                class="col-6 q-pb-md"
                v-model="form.remark"
                label="备注"
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
                :options="sys_enabledOptions"
                label="状态"
                :rules="[requiredRule]"
              />
              <q-input
                outlined
                dense
                class="col-6 q-pb-md"
                v-model="form.sideMode"
                label="用户主题"
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
  listUser,
  createUser,
  updateUser,
  deleteUser,
  getUser,
  resetPassword,
} from 'src/api/admin/user'
import { listDept } from 'src/api/admin/dept'
import { listPost } from 'src/api/admin/post'
import { listRole } from 'src/api/admin/role'

import { useQuasar } from 'quasar'
import { requiredRule } from 'src/utils/inputRule'
import {
  DictOptions,
  getOptionsByList,
  getDictLabel,
  getDict,
} from 'src/utils/dict'
import { downloadAction } from 'src/api/manage'
import { getIds } from 'src/utils/arrayOrObject'

const $q = useQuasar()
let { proxy } = getCurrentInstance()
const dialogVisible = ref(false)

const dataList = ref([])
const formType = ref('')
const dictId = ref(0)
const form = ref({})

const enableDictOptions = ref({})
const adminDeptOptions = ref([])
const formatAdminDept = (prop) => {
  return getDictLabel(adminDeptOptions.value, prop)
}
const adminPostOptions = ref([])
const formatAdminPost = (prop) => {
  return getDictLabel(adminPostOptions.value, prop)
}
const sys_enabledOptions = ref([])
const formatsys_enabled = (prop) => {
  if (!prop) {
    prop = false
  }
  return getDictLabel(enableDictOptions.value, prop)
}
const adminRoleOptions = ref([])
const formatAdminRole = (prop) => {
  return getDictLabel(adminRoleOptions.value, prop)
}
const searchKey = ref('')
const uploadUrl = process.env.BASE_URL + '/admin/user/import'
const exportUrl = '/admin/user/export'
const exportTemplateUrl = '/admin/user/exportTemplate'
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
      name: 'sort',
      align: 'left',
      label: '序号',
      field: 'sort',
      sortable: true,
      classes: 'ellipsis',
    },
    {
      name: 'username',
      align: 'left',
      label: '用户名',
      field: 'username',
      sortable: true,
      classes: 'ellipsis',
    },
    {
      name: 'nickName',
      align: 'left',
      label: '昵称',
      field: 'nickName',
      sortable: true,
      classes: 'ellipsis',
    },
    {
      name: 'mobile',
      align: 'left',
      label: '手机号',
      field: 'mobile',
      sortable: true,
      classes: 'ellipsis',
    },
    {
      name: 'email',
      align: 'left',
      label: '邮箱',
      field: 'email',
      sortable: true,
      classes: 'ellipsis',
    },
    {
      name: 'deptId',
      align: 'left',
      label: '部门',
      field: 'deptId',
      sortable: true,
      classes: 'ellipsis',
    },
    {
      name: 'postId',
      align: 'left',
      label: '职务',
      field: 'postId',
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
  enableDictOptions.value = await getDict('sys_enabled')
  const queryReqAdminDept = {
    pageIndex: 0,
    pageSize: 9999,
  }
  let resAdminDept = await listDept(queryReqAdminDept)
  adminDeptOptions.value = getOptionsByList(resAdminDept.list, 'name', 'id')
  const queryReqAdminPost = {
    pageIndex: 0,
    pageSize: 9999,
  }
  let resAdminPost = await listPost(queryReqAdminPost)
  adminPostOptions.value = getOptionsByList(resAdminPost.list, 'name', 'id')
  sys_enabledOptions.value = await getDict('sys_enabled')
  let resAdminRole = await listRole(queryReqAdminPost)
  adminRoleOptions.value = resAdminRole.list
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
    roles: [],
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

  let table = await listUser(queryReq)

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
const edit = async (p) => {
  reset()
  // form.value = {
  //   ...p,
  // }
  let req = {
    id: p.id,
  }
  let res = await getUser(req)
  form.value = res
  formType.value = '编辑'
  dialogVisible.value = true
}

const del = async (p) => {
  let res = await deleteUser(p)
  onRequest()
}

const delList = async () => {
  let req = {
    ids: getIds(selected.value),
  }
  let res = await deleteUser(req)
  onRequest()
}
const resetPwd = async () => {
  let req = {
    ids: getIds(selected.value),
  }
  await resetPassword(req)
}

const submit = async () => {
  // const res = undefined;
  // form.value.postId = form.value.post.id
  // form.value.deptId = form.value.dept.id
  if (formType.value === '编辑') {
    let res = await updateUser(form.value)
  } else if (formType.value === '新建') {
    let res = await createUser(form.value)
  } else {
    proxy.$error('请求错误')
  }
  dialogVisible.value = false
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
  downloadAction(exportUrl, '用户-导出.xlsx', queryReq)
}

const handleExportTemplate = () => {
  downloadAction(exportTemplateUrl, '用户模板.xlsx')
}

const uploadFn = async (val) => {
  reset()
  onRequest()
}
</script>
