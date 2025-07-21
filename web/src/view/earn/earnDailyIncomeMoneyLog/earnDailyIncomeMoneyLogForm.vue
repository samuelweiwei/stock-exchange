
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="id字段:" prop="id">
          <el-input v-model.number="formData.id" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="用户id:" prop="uid">
          <el-input v-model.number="formData.uid" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="理财产品ID:" prop="productId">
          <el-input v-model.number="formData.productId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="下单ID:" prop="subscribeId">
          <el-input v-model.number="formData.subscribeId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="当前收益:" prop="earnings">
          <el-input-number v-model="formData.earnings" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="当天利率:" prop="interestRates">
          <el-input-number v-model="formData.interestRates" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="买入数量:" prop="boughtNum">
          <el-input-number v-model="formData.boughtNum" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="模拟|||默认不是模拟:" prop="isMoni">
          <el-switch v-model="formData.isMoni" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="到帐时间:" prop="offeredAt">
          <el-input v-model.number="formData.offeredAt" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="createdAt字段:" prop="createdAt">
          <el-input v-model.number="formData.createdAt" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="updatedAt字段:" prop="updatedAt">
          <el-date-picker v-model="formData.updatedAt" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
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
  createEarnDailyIncomeMoneyLog,
  updateEarnDailyIncomeMoneyLog,
  findEarnDailyIncomeMoneyLog
} from '@/api/earn/earnDailyIncomeMoneyLog'

defineOptions({
    name: 'EarnDailyIncomeMoneyLogForm'
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
            id: undefined,
            uid: undefined,
            productId: undefined,
            subscribeId: undefined,
            earnings: 0,
            interestRates: 0,
            boughtNum: 0,
            isMoni: false,
            offeredAt: undefined,
            createdAt: undefined,
            updatedAt: new Date(),
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findEarnDailyIncomeMoneyLog({ ID: route.query.id })
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
               res = await createEarnDailyIncomeMoneyLog(formData.value)
               break
             case 'update':
               res = await updateEarnDailyIncomeMoneyLog(formData.value)
               break
             default:
               res = await createEarnDailyIncomeMoneyLog(formData.value)
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
