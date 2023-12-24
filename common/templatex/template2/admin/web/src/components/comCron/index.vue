<template>
  <div class="q-pa-none row items-start q-gutter-md">
    <q-card class="my-card">
      <q-card-section class="bg-teal text-white">
        <div class="text-h6">选择定时器计划</div>
      </q-card-section>
      <q-separator />
      <q-card-section>
        <q-list bordered>
          <q-item v-ripple>
            <q-item-section>
              <q-item-label caption>定时器计划</q-item-label>
              <q-item-label>{{ cronFormat }}</q-item-label>
            </q-item-section>
          </q-item>
          <q-item v-ripple>
            <q-item-section>
              <q-item-label caption>定时器计划描述</q-item-label>
              <q-item-label> {{ timeCronStr }}!</q-item-label>
            </q-item-section>
          </q-item>
        </q-list>
        <!-- <div class="text-overline text-orange-9">定时器计划</div>
        <div class="text-h5 q-mt-sm q-mb-xs">定时器计划描述</div>
        <div class="text-caption text-grey">
          {{ timeCronStr }}
        </div> -->
      </q-card-section>
      <q-separator />
      <q-card-section>
        <div class="form">
          <div class="row q-col-gutter-x-md dialog_form q-pt-md">
            <q-select
              class="col-12 q-pb-md"
              outlined
              dense
              map-options
              emit-value
              v-model="cronType"
              :options="cronTypes"
              @change="typeChange"
              label="任务计划"
            />
            <q-select
              v-show="cronType == '每周'"
              v-model="cronData.week"
              class="col-12 q-pb-md"
              outlined
              dense
              map-options
              emit-value
              option-value="value"
              option-label="title"
              :options="weekOption"
              label="请选择星期"
            />
            <q-select
              v-show="cronType == '每月'"
              v-model="cronData.month"
              class="col-12 q-pb-md"
              outlined
              dense
              map-options
              emit-value
              option-value="value"
              option-label="title"
              :options="monthOption"
              label="请选择日期"
            />
          </div>
          <div class="row justify-center q-pa-none">
            <q-time
              class="col-8"
              v-show="
                cronType == '每天' || cronType == '每周' || cronType == '每月'
              "
              color="orange"
              v-model="cronData.time"
              with-seconds
            />
          </div>
        </div>
      </q-card-section>
    </q-card>
  </div>
</template>

<script setup>
import { onMounted, watch, toRefs, ref, getCurrentInstance } from 'vue'
import cronstrue from 'cronstrue/i18n'
import { CronToStruct } from 'src/utils/date'

import { useQuasar } from 'quasar'

const $q = useQuasar()

let { proxy } = getCurrentInstance()

const props = defineProps({
  cronStr: {
    type: String,
    required: false,
    default: '0 0 12 * * ?',
  },
})

const emit = defineEmits(['update:cronStr'])

const { cronStr } = toRefs(props)
const visible = ref(false)
const timeCronStr = ref('')
const cronType = ref('') // 天\周\月
const cronFormat = ref('')
// const week = ref(null) // 星期几
// const month = ref(null) // 几号
// const time = ref('') // 时间
const cronData = ref({
  week: '',
  month: '',
  time: '',
})
const cronTypes = ref([
  '每30秒',
  '每分钟',
  '每五分钟',
  '每十分钟',
  '每半小时',
  '每小时',
  '每天',
  '每周',
  '每月',
])

const cronOptions = ref([
  {
    label: '每30秒',
    value: 'Every30s',
  },
  {
    label: '每分钟',
    value: 'EveryMinute',
  },
  {
    label: '每五分钟',
    value: 'EveryFiveMinute',
  },
  {
    label: '每十分钟',
    value: 'EveryTenMinute',
  },
  {
    label: '每半小时',
    value: 'EveryHalfHour',
  },
  {
    label: '每小时',
    value: 'EveryHour',
  },
  {
    label: '每天',
    value: 'EveryDay',
  },
  {
    label: '每周',
    value: 'EveryWeek',
  },
  {
    label: '每月',
    value: 'EveryMonth',
  },
])
const weekOption = ref([
  {
    title: '星期一',
    value: '1',
    cron: 1,
  },
  {
    title: '星期二',
    value: '2',
    cron: 2,
  },
  {
    title: '星期三',
    value: '3',
    cron: 3,
  },
  {
    title: '星期四',
    value: '4',
    cron: 4,
  },
  {
    title: '星期五',
    value: '5',
    cron: 5,
  },
  {
    title: '星期六',
    value: '6',
    cron: 6,
  },
  {
    title: '星期日',
    value: '0',
    cron: 0,
  },
])

const monthOption = ref([])

onMounted(async () => {
  initDate(cronStr.value)
})

const initDate = (val) => {
  // cronType.value = '每天'
  const arr = []
  for (let i = 1; i < 32; i++) {
    arr.push({
      title: i + '号',
      value: i + '',
      cron: i,
    })
  }
  monthOption.value = arr
  formatCronData(val)
  const valueArr = val.split(' ')
  const clockCornArr = valueArr.slice(0, 3).reverse()
  cronData.value.time = clockCornArr.join(':')
  switch (val) {
    case '*/30 * * * * *':
      cronType.value = '每30秒'
      break
    case '0 */1 * * * *':
      cronType.value = '每分钟'
      break
    case '0 */5 * * * *':
      cronType.value = '每五分钟'
      break
    case '0 */10 * * * *':
      cronType.value = '每十分钟'
      break
    case '0 0,30 * * * *':
      cronType.value = '每半小时'
      break
    case '0 0 * * * *':
      cronType.value = '每小时'
      break
    default:
  }
  const parsedCron = CronToStruct(val)

  if (
    parsedCron.month.length > 2 &&
    parsedCron.dayOfMonth.length > 2 &&
    parsedCron.dayOfWeek.length > 2 &&
    parsedCron.hour.length === 1 &&
    parsedCron.minute.length === 1 &&
    parsedCron.second.length === 1
  ) {
    cronType.value = '每天'
  } else if (
    parsedCron.month.length > 2 &&
    parsedCron.dayOfWeek.length > 2 &&
    parsedCron.dayOfMonth.length === 1 &&
    parsedCron.hour.length === 1 &&
    parsedCron.minute.length === 1 &&
    parsedCron.second.length === 1
  ) {
    cronType.value = '每月'
    cronData.value.month = parsedCron.dayOfMonth[0]
  } else if (
    parsedCron.month.length > 2 &&
    parsedCron.dayOfMonth.length > 2 &&
    parsedCron.dayOfWeek.length === 1 &&
    parsedCron.hour.length === 1 &&
    parsedCron.minute.length === 1 &&
    parsedCron.second.length === 1
  ) {
    cronType.value = '每周'
    cronData.value.week = parsedCron.dayOfWeek[0]
  }

  //   // handleSummit()
  // }
}

const typeChange = (t) => {
  // if (t === '每周' && !cronData.value.week) {
  //   cronData.value.week = weekOption.value[0]
  // }
  // if (t === '每月' && !cronData.value.month) {
  //   cronData.value.month = monthOption.value[0]
  // }
}

watch(cronData.value, (newValue, oldValue) => {
  formatCronData(newValue)
})
watch(cronType, (newValue, oldValue) => {
  formatCronType(newValue)
})

const formatCronType = (val) => {
  switch (cronType.value) {
    case '每30秒':
      cronFormat.value = '*/30 * * * * *'
      break
    case '每分钟':
      cronFormat.value = '0 */1 * * * *'
      break
    case '每五分钟':
      cronFormat.value = '0 */5 * * * *'
      break
    case '每十分钟':
      cronFormat.value = '0 */10 * * * *'
      break
    case '每半小时':
      cronFormat.value = '0 0,30 * * * *'
      break
    case '每小时':
      cronFormat.value = '0 0 * * * *'
      break
    default:
      cronFormat.value = cronStr.value
  }
  const parsedCron = CronToStruct(cronFormat.value)
  timeCronStr.value = cronstrue.toString(cronFormat.value, {
    locale: 'zh_CN',
    use24HourTimeFormat: true,
    verbose: true,
  })
  emit('update:cronStr', cronFormat.value)
}
const formatCronData = (val) => {
  let clockCornArr = cronData.value.time.split(':').reverse()
  switch (cronType.value) {
    case '每30秒':
      cronFormat.value = '*/30 * * * * *'
      break
    case '每分钟':
      cronFormat.value = '0 */1 * * * *'
      break
    case '每五分钟':
      cronFormat.value = '0 */5 * * * *'
      break
    case '每十分钟':
      cronFormat.value = '0 */10 * * * *'
      break
    case '每半小时':
      cronFormat.value = '0 0,30 * * * *'
      break
    case '每小时':
      cronFormat.value = '0 0 * * * *'
      break
    case '每天':
      clockCornArr = cronData.value.time.split(':').reverse()
      cronFormat.value = clockCornArr.join(' ') + ' * * ?'
      // console.log('每天:', cronFormat.value)
      break
    case '每周':
      // console.log('cronData.value.time:', cronData.value.time)
      clockCornArr = cronData.value.time.split(':').reverse()
      cronFormat.value = clockCornArr.join(' ') + ' ? * ' + cronData.value.week
      // console.log('每周end:', cronFormat.value)
      // console.log('每周end每周clockCornArr:', clockCornArr)
      break
    case '每月':
      // console.log('每月:', cronFormat.value)
      clockCornArr = cronData.value.time.split(':').reverse()
      // console.log('clockCornArr:', clockCornArr)
      cronFormat.value =
        clockCornArr.join(' ') + ' ' + cronData.value.month + ' * ?'
      break
    default:
      // console.log('default', val)
      cronFormat.value = cronStr.value
  }

  // console.log('cronFormat', cronFormat.value)
  const parsedCron = CronToStruct(cronFormat.value)
  // console.log('watch parsedCron', parsedCron)
  timeCronStr.value = cronstrue.toString(cronFormat.value, {
    locale: 'zh_CN',
    use24HourTimeFormat: true,
    verbose: true,
  })
  // console.log(timeCronStr.value)
  emit('update:cronStr', cronFormat.value)
  // emit('change', cronFormat.value) // 每月,1号,14:52:36 和 36 52 14 1 * ?
}
</script>
<style lang="sass" scoped>
.my-card
  width: 100%
  width: 500px
</style>
