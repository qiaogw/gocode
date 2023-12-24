<template>
  <q-dialog v-model="linkVisible" persistent>
    <q-card style="width: 1080px; min-width: 80vw">
      <q-bar class="bg-primary text-white">
        <span> 角色关联用户 </span>
        <q-space />
        <q-btn dense flat icon="close" v-close-popup>
          <q-tooltip class="bg-white text-primary">Close</q-tooltip>
        </q-btn>
      </q-bar>
      <q-card-section>
        <q-form>
          <div>
            <q-table
              dense
              flat
              bordered
              separator="cell"
              :columns="columns"
              :rows="role.users"
              row-key="id"
              :filter="filter"
              :grid="$q.screen.xs"
              binary-state-sort
              selection="multiple"
              v-model:selected="selected"
            >
              <template v-slot:top="table">
                <div class="row no-wrap full-width">
                  <q-input
                    borderless
                    dense
                    outlined
                    debounce="300"
                    color="primary"
                    v-model="filter"
                  >
                    <template v-slot:append>
                      <q-icon name="search" />
                    </template>
                  </q-input>
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
                      color="primary"
                      flat
                      dense
                      glossy
                      no-wrap
                      v-show="$q.screen.gt.md"
                      @click="table.toggleFullscreen"
                      :icon="
                        table.inFullscreen ? 'fullscreen_exit' : 'fullscreen'
                      "
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
              <template v-slot:body-cell-actions="props">
                <q-td :props="props">
                  <div class="q-gutter-xs">
                    <com-del
                      v-permission="'user:del'"
                      label="用户"
                      @confirm="del(props.rowIndex)"
                    />
                  </div>
                </q-td>
              </template>
            </q-table>
          </div>
          <div class="row justify-center q-pa-md">
            <q-btn
              outline
              color="primary"
              icon="mdi-close-thick"
              label="关闭"
              v-close-popup
            />
          </div>
        </q-form>
      </q-card-section>
    </q-card>
    <userList ref="userListDialog" :role="role" />
  </q-dialog>
</template>

<script setup>
import { ref, computed } from 'vue'
import { getRole, updateRole } from 'src/api/admin/role'
import { listDept } from 'src/api/admin/dept'
import { listPost } from 'src/api/admin/post'
import userList from './users.vue'

import { getOptionsByList, getDictLabel, getDict } from 'src/utils/dict'
const linkVisible = ref(false)

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
const selected = ref([])
const role = ref({})
const filter = ref('')
const queryReq = {
  pageIndex: 0,
  pageSize: 9999,
}
const adminDeptOptions = ref([])
const adminPostOptions = ref([])
const enableDictOptions = ref([])

const show = async (record) => {
  role.value = await getRole(record)
  linkVisible.value = true

  let resAdminDept = await listDept(queryReq)
  adminDeptOptions.value = getOptionsByList(resAdminDept.list, 'name', 'id')
  let resAdminPost = await listPost(queryReq)
  adminPostOptions.value = getOptionsByList(resAdminPost.list, 'name', 'id')
  enableDictOptions.value = await getDict('sys_enabled')
}

const formatAdminDept = (prop) => {
  return getDictLabel(adminDeptOptions.value, prop)
}
const formatAdminPost = (prop) => {
  return getDictLabel(adminPostOptions.value, prop)
}

const formatsys_enabled = (prop) => {
  if (!prop) {
    prop = false
  }
  return getDictLabel(enableDictOptions.value, prop)
}

// const emit = defineEmits(['showUsers'])
const userListDialog = ref(null)
const create = () => {
  userListDialog.value.show()
  // emit('showUsers')
}

const del = async (index) => {
  // let req = {
  //   ids: [p],
  // }
  role.value.users.splice(index, 1)
  // console.log(role.value.users)
  let res = await updateRole(role.value)
  // onRequest()
}

const delList = async () => {
  // let res = await deleteDatasource()
  // onRequest()
  let req = {
    ids: getIds(selected.value),
  }
  let res = await deleteApi(req)
}

defineExpose({
  show,
})
</script>
