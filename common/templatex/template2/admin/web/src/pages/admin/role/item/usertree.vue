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
        <q-scroll-area
          :thumb-style="thumbStyle"
          style="width: 100%; height: 70vh"
        >
          <q-tree
            :nodes="userTree"
            node-key="uuid"
            label-key="title"
            tick-strategy="leaf"
            v-model:selected="selected"
            v-model:ticked="ticked"
            v-model:expanded="expanded"
          >
            <template v-slot:default-header="prop">
              <div class="row q-gutter-xs no-wrap full-width q-pr-lg">
                <q-icon :name="prop.node.icon" />
                <q-chip>{{ prop.node.nickName }}</q-chip>
                <q-space />
                <q-badge
                  v-if="prop.node.button"
                  transparent
                  color="warning"
                  class="q-mr-md"
                  ><q-icon name="radio_button_checked" />按钮:</q-badge
                >
                <q-checkbox
                  v-for="item in prop.node.button"
                  :key="item.uuid"
                  dense
                  v-model="ticked"
                  :label="item.title"
                  :value="item.uuid"
                  :val="item.uuid"
                  color="cyan"
                />
              </div>
            </template>
          </q-tree>
        </q-scroll-area>
      </q-card-section>
      <div class="justify-center row q-gutter-x-xs" style="width: 100%">
        <q-btn color="primary" @click="handleClear">
          {{ $t('Clear') + $t('All') }}
        </q-btn>
        <q-btn color="primary" @click="handleAll">
          {{ $t('Select') + $t('All') }}
        </q-btn>
        <q-btn
          outline
          color="primary"
          icon="mdi-close-thick"
          label="关闭"
          v-close-popup
        />
        <q-btn color="primary" @click="handleSave">
          {{ $t('Save') }}
        </q-btn>
      </div>
    </q-card>
  </q-dialog>
</template>

<script setup>
import {
  computed,
  onMounted,
  ref,
  toRefs,
  watch,
  getCurrentInstance,
} from 'vue'
import { treeUser } from 'src/api/admin/user'
import XEUtils from 'xe-utils'
import { setRoleUser, getRole } from 'src/api/admin/role'
import { ArrayToTree } from 'src/utils/arrayAndTree'
const thumbStyle = ref({
  right: '2px',
  borderRadius: '5px',
  backgroundColor: '#027be3',
  width: '5px',
  opacity: 0.75,
})

const role = ref({})

const selected = ref('')
const ticked = ref([])
const expanded = ref([])
const userTree = ref([])
const userList = ref([])
const linkVisible = ref(false)

const show = async (record) => {
  role.value = await getRole(record)
  linkVisible.value = true
  let queryReq = {
    sortBy: 'sort',
    descending: false,
    pageIndex: 1,
    pageSize: 99999,
  }

  let table = await treeUser(queryReq)
  expanded.value = []
  userList.value = table.list
  if (table.list) {
    table.list.forEach((item) => {
      expanded.value.push(item.uuid)
    })
  }
  userTree.value = XEUtils.toArrayTree(table.list, {
    key: 'uuid',
    parentKey: 'parentId',
    strict: false,
  })

  onRequest()
}

const onRequest = async () => {
  ticked.value = []
  if (role.value.users) {
    role.value.users.forEach((item) => {
      if (!item.uuid) {
        return
      }
      ticked.value.push(item.uuid)
    })
  }
}
const handleAll = () => {
  ticked.value = []
  userList.value.forEach((item) => {
    ticked.value.push(item.uuid)
  })
}
const handleClear = () => {
  ticked.value = []
}
const emit = defineEmits(['submit', 'update:row'])

const handleSave = async () => {
  // 过滤出在 ids 数组中的对象，然后取出 uuid 属性组成新的数组
  const ids = userList.value
    .filter((item) => ticked.value.includes(item.uuid))
    .map((item) => ({ id: item.id }))
  // console.log(ticked.value)
  role.value.ids = ids
  await setRoleUser(role.value)
  linkVisible.value = false
}

defineExpose({
  show,
})
</script>
