<template>
  <q-dialog v-model="userVisible" persistent>
    <q-card style="width: 1080px; min-width: 80vw">
      <q-bar class="bg-primary text-white">
        <span> 角色添加用户 </span>
        <q-space />
        <q-btn dense flat icon="close" v-close-popup>
          <q-tooltip class="bg-white text-primary">Close</q-tooltip>
        </q-btn>
      </q-bar>
      <q-card-section>
        <q-form @submit="submit">
          <div>
            <q-table
              dense
              flat
              bordered
              separator="cell"
              :columns="columns"
              :rows="userList"
              row-key="id"
              v-model:filter="searchKey"
              :grid="$q.screen.xs"
              binary-state-sort
              selection="multiple"
              v-model:selected="selected"
            >
              <template v-slot:top="table">
                <div class="row no-wrap full-width">
                  <com-search @query="onRequest" v-model:filter="searchKey" />
                  <q-space />
                  <q-btn-group class="q-gutter-xs">
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
  </q-dialog>
</template>

<script setup>
import { ref, toRefs, computed } from 'vue'
import { getRole, setRoleUser } from 'src/api/admin/role'
import { listDept } from 'src/api/admin/dept'
import { listPost } from 'src/api/admin/post'
import XEUtils from 'xe-utils'

import { getOptionsByList, getDictLabel, getDict } from 'src/utils/dict'
import { listNoUser } from 'src/api/admin/user'

const props = defineProps({
  role: Object,
})
const { role } = toRefs(props)

const userVisible = ref(false)
const userList = ref([])
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
  ]
})
const queryReq = {
  pageIndex: 0,
  pageSize: 9999,
}
const adminDeptOptions = ref([])
const adminPostOptions = ref([])
const enableDictOptions = ref([])

const selected = ref([])

const searchKey = ref('')
const pagination = ref({
  sortBy: 'id',
  descending: false,
  page: 1,
  rowsPerPage: 10,
  rowsNumber: 0,
})

const show = async () => {
  let resAdminDept = await listDept(queryReq)
  adminDeptOptions.value = getOptionsByList(resAdminDept.list, 'name', 'id')
  let resAdminPost = await listPost(queryReq)
  adminPostOptions.value = getOptionsByList(resAdminPost.list, 'name', 'id')
  enableDictOptions.value = await getDict('sys_enabled')
  userVisible.value = true
  onRequest()
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
  queryReq.id = role.value.id

  let table = await listNoUser(queryReq)

  pagination.value = val.pagination

  pagination.value.rowsNumber = table.count
  userList.value = table.list
  let tt = XEUtils.filter(table.list, (item) => {
    return XEUtils.filter(item.roles, (it) => it.id !== role.value.id)
  })
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

const submit = async () => {
  userVisible.value = false
}

defineExpose({
  show,
})
</script>
