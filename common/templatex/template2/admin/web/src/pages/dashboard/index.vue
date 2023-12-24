<template>
  <q-page class="flex flex-left">
    <div class="q-pa-md">
      <div class="row justify-left q-gutter-sm">
        <q-intersection
          v-for="(item, index) in dataList"
          :key="index"
          transition="scale"
          class="example-item"
        >
          <q-card class="q-ma-sm">
            <img src="images/meta.png" />
            <q-separator color="orange" inset />
            <q-card-section>
              <div class="text-h6">{{ item.name }}</div>
              <div class="text-subtitle2">￥{{ item.price }}</div>
              <div class="row justify-end">
                <q-btn color="secondary" label="订购" />
              </div>
            </q-card-section>
          </q-card>
        </q-intersection>
      </div>
    </div>
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
  listBundle,
  createBundle,
  updateBundle,
  deleteBundle,
  getBundle,
} from 'src/api/drug/bundle'

import { listGoods } from 'src/api/drug/goods'

import { useQuasar } from 'quasar'
import { requiredRule } from 'src/utils/inputRule'
import { FormatDate } from 'src/utils/date'
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

const dataList = ref([])
const dictOptions = ref({})
const goodsTypeOptions = ref([])
const goodsOptions = ref([])

const searchKey = ref('')
const pagination = ref({
  sortBy: 'id',
  descending: false,
  page: 1,
  rowsPerPage: 10,
  rowsNumber: 0,
})

onMounted(async () => {
  dictOptions.value = await DictOptions()
  goodsTypeOptions.value = await getDict('goodsType')
  const queryPost = {
    pageIndex: 0,
    pageSize: 9999,
  }
  let gs = await listGoods(queryPost)
  goodsOptions.value = gs.list
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

  let table = await listBundle(queryReq)

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
</script>
<style lang="sass" scoped>
.example-item
  height: 290px
  width: 290px
</style>
