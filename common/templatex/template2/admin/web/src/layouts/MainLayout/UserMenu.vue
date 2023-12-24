<template>
  <q-btn dense flat>
    {{ t('Welcome') }}
    <comShowName showMyName style="margin: 0 5px" />
    <comAvatar loginUser size="28px" />
    <q-menu class="row no-wrap items-center justify-around q-pa-md">
      <div class="column">
        <q-list class="q-py-md" dense style="min-width: 100px">
          <q-item clickable v-close-popup>
            <q-item-section @click="showProfile">个人中心</q-item-section>
          </q-item>
          <q-separator />
          <q-item clickable v-if="user.roles.length > 1">
            <q-item-section>切换角色</q-item-section>
            <q-item-section side>
              <q-icon name="keyboard_arrow_right" />
            </q-item-section>

            <q-menu anchor="top end" self="top start">
              <q-list>
                <q-item v-for="n in user.roles" :key="n.id" dense clickable>
                  <q-item-section @click="setRole(n.id)">{{
                    n.name
                  }}</q-item-section>
                </q-item>
              </q-list>
            </q-menu>
          </q-item>
          <q-separator />
          <q-item clickable v-close-popup>
            <q-item-section>其他</q-item-section>
          </q-item>
        </q-list>
      </div>

      <q-separator vertical inset class="q-mx-lg" />

      <div class="column items-center">
        <comAvatar loginUser size="88px" />

        <div class="text-subtitle1 q-mt-md q-mb-md">
          <comShowName showMyName />
        </div>
        <div class="text-h7 q-mt-md q-mb-md">
          <span> 当前角色：{{ user.role.name }} </span>
        </div>

        <div class="row no-wrap q-gutter-md">
          <q-btn
            icon="logout"
            color="primary"
            :label="$t('Logout')"
            push
            size="sm"
            v-close-popup
            @click="logout"
          />
        </div>
      </div>
    </q-menu>
  </q-btn>
</template>

<script setup>
import { useUserStore } from 'src/stores/user'
import comShowName from 'src/components/comShowName/index.vue'
import comAvatar from 'src/components/comAvatar/index.vue'
import { useQuasar } from 'quasar'
import { useI18n } from 'vue-i18n'
import { useRouter } from 'vue-router'
import { computed, onMounted, ref } from 'vue'
import { setMeRole } from 'src/api/admin/user'

const userStore = useUserStore()
const $q = useQuasar()
const { t } = useI18n()
const router = useRouter()
// const user = ref({})

const logout = () => {
  $q.dialog({
    title: t('Logout'),
    message: t('Confirm') + t('Logout') + '?',
    cancel: true,
    persistent: true,
  }).onOk(() => {
    userStore.HandleLogout()
    router.push({ path: '/login' })
  })
}

const user = computed(() => {
  const user = userStore.GetInfo()
  if (!user) {
    return
  }
  return user
})

const showProfile = async () => {
  router.push({ name: 'person' })
}

const setRole = async (roleId) => {
  let data = {
    roleId: roleId,
  }
  let res = await setMeRole(data)
  userStore.setToken(res.accessToken)
  userStore.setInfo(res.user)
  setTimeout(() => {
    window.location.reload()
  }, 50)
  // userStore.GetInfo()
}
</script>
