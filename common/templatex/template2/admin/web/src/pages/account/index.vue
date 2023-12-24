<template>
  <q-page class="q-pa-xs">
    <q-card flat square>
      <q-card-section class="q-pa-none q-py-md row">
        <div class="column col-xs-2 items-center">
          <div>
            <div class="q-pa-none q-py-md">
              <comAvatar loginUser size="88px" />
              <div class="text-subtitle1 q-mt-md q-mb-md">
                <comShowName showMyName />
              </div>
              <div class="text-h7 q-mt-md q-mb-md">
                <span> 角色：{{ roleName }} </span>
              </div>
            </div>
            <q-tabs
              v-model="tab"
              align="left"
              dense
              class="text-grey"
              active-color="primary"
              indicator-color="primary"
              vertical
            >
              <q-tab
                name="basicSettings"
                label="基本设置"
                style="justify-content: left"
                content-class="q-pl-md"
              />
              <q-tab
                name="safeSettings"
                label="安全设置"
                style="justify-content: left"
                content-class="q-pl-md"
              />
              <q-tab
                name="accountBind"
                label="账号绑定"
                disable
                style="justify-content: left"
                content-class="q-pl-md"
              />
              <q-tab
                disable
                name="newMsg"
                label="新消息通知"
                style="justify-content: left"
                content-class="q-pl-md"
              />
              <!-- <q-tab name="api" label="api权限" />
      <q-tab name="data" label="数据权限" /> -->
            </q-tabs>
          </div>
        </div>
        <q-separator :vertical="$q.screen.gt.xs" v-show="$q.screen.gt.xs" />
        <div class="col-sm col-xs-12 q-px-md q-pt-none">
          <q-tab-panels
            v-model="tab"
            animated
            transition-prev="fade"
            transition-next="fade"
          >
            <q-tab-panel name="basicSettings" class="row q-pt-sm">
              <div class="text-h5 col-12 q-mb-md">基本设置</div>
              <div class="lt-sm col-xs-12 q-mb-md">
                <span class="text-center block">
                  <q-img
                    src="/img/user/woman.jpg"
                    width="180px"
                    :ratio="10 / 10"
                  />
                </span>
                <span class="text-center block">
                  <q-btn
                    unelevated
                    color="primary"
                    label="更换头像"
                    icon="unarchive"
                  />
                </span>
              </div>
              <div
                class="col-md-6 col-sm-5 col-xs-12 q-gutter-y-md q-pt-md q-pl-md q-pb-md"
              >
                <q-input
                  outlined
                  dense
                  square
                  label="移动电话"
                  v-model="user.mobile"
                />
                <q-input
                  outlined
                  dense
                  square
                  label="邮箱"
                  v-model="user.email"
                />
                <q-input
                  outlined
                  dense
                  square
                  label="昵称"
                  v-model="user.nickName"
                />

                <q-select
                  outlined
                  dense
                  square
                  behavior="menu"
                  label="国家\地区"
                  options-dense
                  :options="['中国', '韩国']"
                  v-model="user.country"
                />
                <span class="row q-gutter-x-sm">
                  <q-select
                    class="col"
                    outlined
                    dense
                    square
                    behavior="menu"
                    label="所在省份"
                    options-dense
                    :options="['湖北省', '广东省']"
                    v-model="user.province"
                  />
                  <q-select
                    class="col"
                    outlined
                    dense
                    square
                    behavior="menu"
                    label="所在城市"
                    options-dense
                    :options="['深圳市', '佛山市']"
                    v-model="user.city"
                  />
                </span>
                <q-input
                  type="text"
                  outlined
                  dense
                  square
                  label="详细地址"
                  v-model="user.address"
                />
                <q-input
                  type="text"
                  outlined
                  dense
                  square
                  label="公司"
                  v-model="user.company"
                />

                <span class="row q-gutter-x-sm">
                  <q-select
                    class="col-3"
                    outlined
                    dense
                    square
                    behavior="menu"
                    label="前缀"
                    options-dense
                    :options="['+86', '+87']"
                    v-model="user.phonePrefix"
                  />
                  <q-input
                    class="col"
                    outlined
                    dense
                    square
                    label="联系电话"
                    v-model="user.phone"
                  />
                </span>
                <q-btn label="更新基本信息" color="primary" unelevated />
              </div>
            </q-tab-panel>
            <q-tab-panel name="safeSettings" class="q-pt-sm">
              <div class="text-h5 col-12 q-mb-md">安全设置</div>
              <q-list class="text-body2">
                <q-item>
                  <q-item-section>
                    <q-item-label>账户密码</q-item-label>
                  </q-item-section>
                  <q-item-section avatar>
                    <q-btn
                      flat
                      unelevated
                      color="primary"
                      label="修改"
                      @click="editPwd"
                    />
                  </q-item-section>
                </q-item>
                <q-separator inset="" spaced="10px" />
                <q-item>
                  <q-item-section>
                    <q-item-label>绑定手机</q-item-label>
                    <q-item-label class="text-grey-6">
                      绑定手机：{{ user.mobile }}
                    </q-item-label>
                  </q-item-section>
                  <q-item-section avatar>
                    <q-btn flat unelevated color="primary" label="修改" />
                  </q-item-section>
                </q-item>
              </q-list>
            </q-tab-panel>
            <q-tab-panel name="accountBind" class="q-pt-sm">
              <div class="text-h5 col-12 q-mb-md">账号绑定</div>
              <q-list class="text-body2">
                <q-item>
                  <q-item-section avatar>
                    <q-icon
                      size="xl"
                      color="warning"
                      style="cursor: pointer"
                      class="iconfont icontaobao q-ml-sm"
                    />
                  </q-item-section>
                  <q-item-section>
                    <q-item-label>绑定淘宝</q-item-label>
                    <q-item-label class="text-grey-6"
                      >{{ accountBindData.bindTaoBaoNo }}
                    </q-item-label>
                  </q-item-section>
                  <q-item-section avatar>
                    <q-btn flat unelevated color="primary" label="绑定" />
                  </q-item-section>
                </q-item>
                <q-separator inset="" spaced="10px" />
                <q-item>
                  <q-item-section avatar>
                    <q-icon
                      size="xl"
                      color="primary"
                      style="cursor: pointer"
                      class="iconfont iconzhifubao q-ml-sm"
                    />
                  </q-item-section>
                  <q-item-section>
                    <q-item-label>绑定支付宝</q-item-label>
                    <q-item-label class="text-grey-6">
                      {{ accountBindData.bindZfbNo }}
                    </q-item-label>
                  </q-item-section>
                  <q-item-section avatar>
                    <q-btn flat unelevated color="primary" label="绑定" />
                  </q-item-section>
                </q-item>
                <q-separator inset="" spaced="10px" />
                <q-item>
                  <q-item-section avatar>
                    <q-icon
                      size="xl"
                      color="info"
                      style="cursor: pointer"
                      class="iconfont iconweixin q-ml-sm"
                    />
                  </q-item-section>
                  <q-item-section>
                    <q-item-label>绑定微信</q-item-label>
                    <q-item-label class="text-grey-6"
                      >{{ accountBindData.bindWechatNo }}
                    </q-item-label>
                  </q-item-section>
                  <q-item-section avatar>
                    <q-btn flat unelevated color="primary" label="绑定" />
                  </q-item-section>
                </q-item>
                <q-separator inset="" spaced="10px" />
              </q-list>
            </q-tab-panel>
            <q-tab-panel name="newMsg">
              <div class="text-h5 col-12 q-mb-md">新消息通知</div>
              <q-list class="text-body2">
                <q-item>
                  <q-item-section>
                    <q-item-label>账户密码</q-item-label>
                    <q-item-label class="text-grey-6"
                      >其他用户的消息将以站内信的形式通知
                    </q-item-label>
                  </q-item-section>
                  <q-item-section avatar>
                    <q-toggle
                      v-model="newMsgData.passwordMsg"
                      checked-icon="check"
                      color="primary"
                      unchecked-icon="clear"
                    />
                  </q-item-section>
                </q-item>
                <q-separator inset="" spaced="10px" />
                <q-item>
                  <q-item-section>
                    <q-item-label>系统消息</q-item-label>
                    <q-item-label class="text-grey-6">
                      系统消息将以站内信的形式通知
                    </q-item-label>
                  </q-item-section>
                  <q-item-section avatar>
                    <q-toggle
                      v-model="newMsgData.systemMsg"
                      checked-icon="check"
                      color="primary"
                      unchecked-icon="clear"
                    />
                  </q-item-section>
                </q-item>
                <q-separator inset="" spaced="10px" />
                <q-item>
                  <q-item-section>
                    <q-item-label>待办任务</q-item-label>
                    <q-item-label class="text-grey-6"
                      >待办任务将以站内信的形式通知
                    </q-item-label>
                  </q-item-section>
                  <q-item-section avatar>
                    <q-toggle
                      v-model="newMsgData.waitTaskMsg"
                      checked-icon="check"
                      color="red"
                      unchecked-icon="clear"
                    />
                  </q-item-section>
                </q-item>
                <q-separator inset="" spaced="10px" />
              </q-list>
            </q-tab-panel>
          </q-tab-panels>
        </div>
      </q-card-section>
    </q-card>
    <q-dialog v-model="dialogVisiblePwd" persistent>
      <q-card style="width: 800px; min-width: 40vw">
        <q-bar class="bg-primary text-white">
          <span> 职务 </span>
          <q-space />
          <q-btn dense flat icon="close" v-close-popup>
            <q-tooltip class="bg-white text-primary">关闭</q-tooltip>
          </q-btn>
        </q-bar>
        <q-card-section>
          <q-form @submit="submitPwd">
            <div class="q-col-gutter-x-md dialog_form q-pt-md">
              <q-input
                outlined
                dense
                :type="isPwd ? 'password' : 'text'"
                class="q-pb-md"
                v-model="formPwd.oldPassword"
                label="原密码"
                :rules="[requiredRule]"
              >
                <template v-slot:append>
                  <q-icon
                    :name="isPwd ? 'visibility_off' : 'visibility'"
                    class="cursor-pointer"
                    @click="isPwd = !isPwd"
                  /> </template
              ></q-input>
              <q-input
                outlined
                dense
                :type="isPwd ? 'password' : 'text'"
                class="q-pb-md"
                v-model="formPwd.password"
                label="新密码"
                :rules="[requiredRule]"
              >
                <template v-slot:append>
                  <q-icon
                    :name="isPwd ? 'visibility_off' : 'visibility'"
                    class="cursor-pointer"
                    @click="isPwd = !isPwd"
                  /> </template
              ></q-input>
              <q-input
                outlined
                dense
                :type="isPwd ? 'password' : 'text'"
                class="q-pb-md"
                v-model="formPwd.rePassword"
                label="新密码重复"
                :rules="[requiredRule]"
              >
                <template v-slot:append>
                  <q-icon
                    :name="isPwd ? 'visibility_off' : 'visibility'"
                    class="cursor-pointer"
                    @click="isPwd = !isPwd"
                  /> </template
              ></q-input>
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
import { useUserStore } from 'src/stores/user'
import comShowName from 'src/components/comShowName/index.vue'
import comAvatar from 'src/components/comAvatar/index.vue'
import { useQuasar } from 'quasar'
import { useI18n } from 'vue-i18n'
import { useRouter } from 'vue-router'
import { computed, onMounted, getCurrentInstance, ref } from 'vue'
import { setMeRole } from 'src/api/admin/user'
import { setPassword } from 'src/api/admin/user'
import { requiredRule } from 'src/utils/inputRule'

const tab = ref('basicSettings')
const userStore = useUserStore()
const $q = useQuasar()
let { proxy } = getCurrentInstance()
const { t } = useI18n()
const formPwd = ref({})
const dialogVisiblePwd = ref(false)
const isPwd = ref(true)

// const user = computed(() => {
//   const user = userStore.GetInfo()
//   console.log(user)
//   return user
// })

const user = ref({})
const roleName = ref('')
onMounted(async () => {
  user.value = await userStore.GetInfo()
  roleName.value = user.value.role.name
})

const editPwd = () => {
  dialogVisiblePwd.value = true
}
const submitPwd = async () => {
  if (formPwd.value.password !== formPwd.value.rePassword) {
    proxy.$error('重复密码不一致')
    return
  }
  await setPassword(formPwd.value)
  dialogVisiblePwd.value = false
}
</script>
