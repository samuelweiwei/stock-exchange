
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="币种，例如BTC:" prop="currency">
          <el-input v-model="formData.currency" :clearable="true"  placeholder="请输入币种，例如BTC" />
       </el-form-item>
        <el-form-item label="第三方厂商名称：aa pay:" prop="thirdRechargeChannel">
          <el-input v-model="formData.thirdRechargeChannel" :clearable="true"  placeholder="请输入第三方厂商名称：aa pay" />
       </el-form-item>
        <el-form-item label="thirdCurrencyCode字段:" prop="thirdCurrencyCode">
          <el-input v-model="formData.thirdCurrencyCode" :clearable="true"  placeholder="请输入thirdCurrencyCode字段" />
       </el-form-item>
        <el-form-item label="thirdCoinCode字段:" prop="thirdCoinCode">
          <el-input v-model="formData.thirdCoinCode" :clearable="true"  placeholder="请输入thirdCoinCode字段" />
       </el-form-item>
        <el-form-item label="是否需要查询第三方汇率:" prop="requireExchangeRate">
          <el-switch v-model="formData.requireExchangeRate" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="价格，以USDT计价:" prop="priceUsdt">
          <el-input-number v-model="formData.priceUsdt" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="充值方式：系统充值，快捷充值:" prop="rachargeType">
          <el-input v-model="formData.rachargeType" :clearable="true"  placeholder="请输入充值方式：系统充值，快捷充值" />
       </el-form-item>
        <el-form-item label="渠道1 地址  2 银行卡 3 其他:" prop="channel">
          <el-input v-model="formData.channel" :clearable="true"  placeholder="请输入渠道1 地址  2 银行卡 3 其他" />
       </el-form-item>
        <el-form-item label="对应渠道的地址:" prop="address">
          <el-input v-model="formData.address" :clearable="true"  placeholder="请输入对应渠道的地址" />
       </el-form-item>
        <el-form-item label="类型，数字货币，法币:" prop="coinType">
          <el-input v-model="formData.coinType" :clearable="true"  placeholder="请输入类型，数字货币，法币" />
       </el-form-item>
        <el-form-item label="排序权重，默认为0:" prop="sortOrder">
          <el-input v-model.number="formData.sortOrder" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="状态，开启或关闭:" prop="STATUS">
          <el-input v-model="formData.STATUS" :clearable="true"  placeholder="请输入状态，开启或关闭" />
       </el-form-item>
        <el-form-item label="创建者:" prop="createdBy">
          <el-input v-model.number="formData.createdBy" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="更新者:" prop="updatedBy">
          <el-input v-model.number="formData.updatedBy" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="删除者:" prop="deletedBy">
          <el-input v-model.number="formData.deletedBy" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createRechargeChannels,
  updateRechargeChannels,
  findRechargeChannels
} from '@/api/userfund/rechargeChannels'

defineOptions({
    name: 'RechargeChannelsForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'


const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            currency: '',
            thirdRechargeChannel: '',
            thirdCurrencyCode: '',
            thirdCoinCode: '',
            requireExchangeRate: false,
            priceUsdt: 0,
            rachargeType: '',
            channel: '',
            address: '',
            coinType: '',
            sortOrder: undefined,
            STATUS: '',
            createdBy: undefined,
            updatedBy: undefined,
            deletedBy: undefined,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findRechargeChannels({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createRechargeChannels(formData.value)
               break
             case 'update':
               res = await updateRechargeChannels(formData.value)
               break
             default:
               res = await createRechargeChannels(formData.value)
               break
           }
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
