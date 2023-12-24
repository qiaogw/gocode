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
        @request="onRequest"
        :grid="$q.screen.xs"
        v-model:filter="searchKey"
        v-model:pagination="pagination"
        binary-state-sort
      >
        <template v-slot:top="table">
          <div class="row no-wrap full-width">
            <com-search @query="onRequest" v-model:filter="searchKey" />
            <q-space />
            <div class="q-gutter-xs">
              <q-btn
                v-permission="'role:add'"
                icon="add"
                no-wrap
                color="primary"
                @click="create"
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
            </div>
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
                v-permission="'role:edit'"
                label="角色"
                @confirm="edit(props.row)"
              />
              <q-btn
                v-permission="'role:assign'"
                flat
                round
                dense
                color="positive"
                icon="add_link"
                @click.stop
              >
                <q-tooltip>关联用户</q-tooltip>
                <q-menu auto-close anchor="center left" self="center right">
                  <div
                    class="row no-wrap items-center q-pa-sm bg-warning text-white"
                  >
                    <span class="text-no-wrap"
                      >确认关联 <b>{{ props.row.name }}</b> ?</span
                    >
                    <q-btn
                      flat
                      round
                      dense
                      color="negative"
                      icon="add_link"
                      @click="relation(props.row)"
                    >
                      <q-tooltip>确认关联</q-tooltip></q-btn
                    >
                  </div>
                </q-menu>
              </q-btn>
              <q-btn
                v-permission="'role:assign'"
                flat
                round
                dense
                color="warning"
                icon="engineering"
                @click.stop
              >
                <q-tooltip>授权</q-tooltip>
                <q-menu auto-close anchor="center left" self="center right">
                  <div
                    class="row no-wrap items-center q-pa-sm bg-warning text-white"
                  >
                    <span class="text-no-wrap"
                      >确认授权 <b>{{ props.row.name }}</b> ?</span
                    >
                    <q-btn
                      flat
                      round
                      dense
                      color="negative"
                      icon="engineering"
                      @click="permission(props.row)"
                    >
                      <q-tooltip>确认授权</q-tooltip></q-btn
                    >
                  </div>
                </q-menu>
              </q-btn>
              <com-del
                v-permission="'role:del'"
                label="角色"
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
          <span> {{ formType }}角色 </span>
          <q-space />
          <q-btn dense flat icon="close" v-close-popup>
            <q-tooltip class="bg-white text-primary">Close</q-tooltip>
          </q-btn>
        </q-bar>
        <q-card-section>
          <q-form @submit="submit">
            <div>
              <q-input
                dense
                outlined
                class="col-6"
                v-model="form.name"
                label="角色名称"
                :rules="[requiredRule]"
              />
              <q-input
                outlined
                v-model="form.code"
                label="角色编码"
                :rules="[requiredRule]"
              />
              <q-input
                dense
                outlined
                class="col-6"
                label="排序"
                type="number"
                v-model.number="form.sort"
                :rules="[requiredRule]"
              />
              <q-input
                dense
                outlined
                class="col-6"
                label="描述"
                v-model="form.remark"
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
    <RolePermission ref="rolePermissionDialog" />
    <userRelation ref="userRelationDialog" />
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
  listRole,
  createRole,
  updateRole,
  deleteRole,
  getRole,
} from 'src/api/admin/role'
import { useQuasar } from 'quasar'
import { requiredRule } from 'src/utils/inputRule'
import RolePermission from './item/permission.vue'
import userRelation from './item/usertree.vue'
import { useRoute } from 'vue-router'

const $q = useQuasar()
let { proxy } = getCurrentInstance()
const dialogVisible = ref(false)

const route = useRoute()

const dataList = ref([])
const formType = ref('')

const form = ref({})
const searchKey = ref('')
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
      align: 'center',
      label: '序号',
      field: 'sort',
      sortable: true,
    },
    {
      name: 'name',
      align: 'center',
      label: '名称',
      field: 'name',
      sortable: true,
    },
    {
      name: 'code',
      align: 'left',
      label: '代码',
      field: 'code',
      sortable: true,
    },
    {
      name: 'remark',
      align: 'left',
      label: '描述',
      field: 'remark',
      classes: 'ellipsis',
      style: 'max-width: 100px',
    },
    { name: 'actions', align: 'center', label: '操作', field: 'actions' },
  ]
})
onMounted(async () => {
  onRequest()
  // console.log(route.meta)
  // console.log(router)
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

  let table = await listRole(queryReq)

  pagination.value = val.pagination

  pagination.value.rowsNumber = table.count
  dataList.value = table.list
}

const create = () => {
  formType.value = '新建'
  reset()
  dialogVisible.value = true
}
const edit = (p) => {
  form.value = {
    ...p,
  }
  formType.value = '编辑'
  dialogVisible.value = true
}

const del = async (p) => {
  let res = await deleteRole(p)
  onRequest()
}

const submit = async () => {
  // const res = undefined;
  if (formType.value === '编辑') {
    let res = await updateRole(form.value)
  } else if (formType.value === '新建') {
    let res = await createRole(form.value)
  } else {
    proxy.$error('路径错误')
  }
  dialogVisible.value = false
  reset()
  onRequest()
}
const rolePermissionDialog = ref(null)
const permission = async (p) => {
  rolePermissionDialog.value.show(p)
}
const userRelationDialog = ref(null)
const relation = async (p) => {
  userRelationDialog.value.show(p)
}
</script>
