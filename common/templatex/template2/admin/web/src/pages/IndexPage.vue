<template>
  <q-page class="flex flex-center">
    <q-card>
      <!-- Option 1 -->
      <div>Direct store</div>
      <!-- Read the state value directly -->
      <div>{{ store.counter }}</div>
      <!-- Use getter directly -->
      <div>{{ store.doubleCount }}</div>

      <!-- Manipulate state directly -->
      <q-btn @click="store.counter--">-</q-btn>
      <!-- Use an action -->
      <q-btn @click="store.increment()">+</q-btn>
    </q-card>

    <q-card>
      <!-- Option 2 -->
      <div>Indirect store</div>
      <!-- Use the computed state -->
      <div>{{ count }}</div>
      <!-- Use the computed getter -->
      <div>{{ doubleCountValue }}</div>

      <!-- Use the exposed function -->
      <q-btn @click="decrementCount()">-</q-btn>
      <!-- Use the exposed function -->
      <q-btn @click="incrementCount()">+</q-btn>
    </q-card>

    <q-card>
      <!-- Option 3 -->
      <div>Destructured store</div>
      <!-- Use the destructured state -->
      <div>{{ counter }}</div>
      <!-- Use the destructured getter -->
      <div>{{ doubleCount }}</div>

      <!-- Manipulate state directly-->
      <q-btn @click="counter--">-</q-btn>
      <!-- Use an action -->
      <q-btn @click="increment()">+</q-btn>
    </q-card>
    <div>
      <q-btn round size="sm" color="accent" @click="apiClick">api</q-btn>
      <q-btn
        v-permission="['function_edit']"
        round
        size="sm"
        color="accent"
        @click="menuClick"
        >menu</q-btn
      >
    </div>
    <q-list dense bordered separator>
      <q-item clickable v-close-popup v-ripple>
        <q-item-section>
          <q-chip> Vue: {{ $quasarVersion }} </q-chip>
        </q-item-section>
      </q-item>
    </q-list>
  </q-page>
</template>

<script>
import { computed, getCurrentInstance } from 'vue'
import { useCounterStore } from 'stores/example-store'
import { useUserStore } from 'stores/user'
import { storeToRefs } from 'pinia'
import { useQuasar } from 'quasar'
import { getMenu } from 'src/api/admin/auth'
import { getUser, listUser } from 'src/api/admin/user'
export default {
  name: 'IndexPage',
  setup() {
    // const { info } = info;
    const $q = useQuasar()
    let { proxy } = getCurrentInstance()
    const store = useCounterStore()
    let userStore = useUserStore()

    // Option 2: use computed and functions to use the store
    const count = computed(() => store.counter)
    const doubleCountValue = computed(() => store.doubleCount)
    const incrementCount = () => store.increment() // use action
    const decrementCount = () => store.counter-- // manipulate directly

    // Option 3: use destructuring to use the store in the template
    const { counter, doubleCount } = storeToRefs(store) // state and getters need "storeToRefs"
    const { increment } = store // actions can be destructured directly
    const apiClick = () => {
      $q.dialog({
        title: 'Alert',
        message: 'Some message',
      })
        .onOk(() => {
          // console.log('OK')
        })
        .onCancel(() => {
          // console.log('Cancel')
        })
        .onDismiss(() => {
          // console.log('I am triggered on both OK and Cancel')
        })
      listUser().then((r) => {
        // console.log(r)
        proxy.$info(r.picPath)
      })
    }
    const menuClick = () => {
      // console.log(info);
      getMenu().then((res) => {
        // console.log(res)
      })
    }
    return {
      // Option 1: return the store directly and couple it in the template
      store,

      // Option 2: use the store in functions and compute the state to use in the template
      count,
      doubleCountValue,
      incrementCount,
      decrementCount,

      // Option 3: pass the destructed state, getters and actions to the template
      counter,
      increment,
      doubleCount,
      apiClick,
      menuClick,
    }
  },
}
</script>
